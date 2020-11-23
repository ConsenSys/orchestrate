package tcp

import (
	"context"
	"io"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	traefikstatic "github.com/containous/traefik/v2/pkg/config/static"
	"github.com/containous/traefik/v2/pkg/log"
	"github.com/sirupsen/logrus"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/tcp/metrics"
)

// EntryPoint is a TCP server
type EntryPoint struct {
	name string
	addr string

	handler Handler

	lis *atomic.Value

	timeouts  *traefikstatic.RespondingTimeouts
	lifecycle *traefikstatic.LifeCycle

	metrics metrics.TPCMetrics

	doneOnce sync.Once
	done     chan struct{}
}

type listenerValue struct {
	l net.Listener
}

// NewEntryPoint creates a new EntryPoint
func NewEntryPoint(name string, config *traefikstatic.EntryPoint, handler Handler, reg metrics.TPCMetrics) *EntryPoint {
	return &EntryPoint{
		name:      name,
		addr:      config.Address,
		handler:   handler,
		timeouts:  config.Transport.RespondingTimeouts,
		lifecycle: config.Transport.LifeCycle,
		lis:       &atomic.Value{},
		metrics:   reg,
		done:      make(chan struct{}),
	}
}

func (e *EntryPoint) Name() string {
	return e.name
}

func (e *EntryPoint) Addr() string {
	lis := e.listener()
	if lis != nil {
		return lis.Addr().String()
	}
	return ""
}

func (e *EntryPoint) listener() net.Listener {
	v, ok := e.lis.Load().(*listenerValue)
	if ok {
		return v.l
	}
	return nil
}

func (e *EntryPoint) with(ctx context.Context) context.Context {
	return log.With(ctx, log.Str("entrypoint", e.name))
}

func (e *EntryPoint) Serve(ctx context.Context, l net.Listener) error {
	logger := log.FromContext(e.with(ctx))
	logger.WithField("address", l.Addr()).Infof("start serving tcp entrypoint")

	e.lis.Store(&listenerValue{l})

	for {
		conn, err := l.Accept()
		if err != nil {
			select {
			case <-e.done:
				return http.ErrServerClosed
			default:
			}

			logger.Error(err)
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				continue
			}

			return err
		}

		// TODO: switch this to Trace when updating logger interface
		logger.WithFields(logrus.Fields{
			"destination.address": conn.RemoteAddr(),
			"destination.source":  conn.LocalAddr(),
		}).Debugf("accepted tcp connection %v", conn)

		writeCloser, err := writeCloser(conn)
		if err != nil {
			panic(err)
		}

		go func() {
			// Enforce read/write deadlines at the connection level,
			// because when we're peeking the first byte to determine whether we are doing TLS,
			// the deadlines at the server level are not taken into account.
			if e.timeouts.ReadTimeout > 0 {
				err = writeCloser.SetReadDeadline(time.Now().Add(time.Duration(e.timeouts.ReadTimeout)))
				if err != nil {
					logger.WithError(err).Errorf("could not set read deadline")
				}
			}

			if e.timeouts.WriteTimeout > 0 {
				err = writeCloser.SetWriteDeadline(time.Now().Add(time.Duration(e.timeouts.WriteTimeout)))
				if err != nil {
					logger.WithError(err).Errorf("could not set write deadline")
				}
			}

			e.handler.ServeTCP(writeCloser)
		}()
	}
}

// Serve handler provided on entrypoint
func (e *EntryPoint) ListenAndServe(ctx context.Context) error {
	listener, err := Listen(
		"tcp", e.addr,
		KeepAliveOpt(3*time.Minute),
		MetricsOpt(e.name, e.metrics),
	)
	if err != nil {
		return err
	}

	return e.Serve(ctx, listener)
}

// Shutdown stops the TCP connections
func (e *EntryPoint) Shutdown(ctx context.Context) error {
	e.doneOnce.Do(func() {
		close(e.done)
	})

	logger := log.FromContext(e.with(ctx))
	logger.Infof("shutting down tcp entrypoint...")

	reqAcceptGraceTimeOut := time.Duration(e.lifecycle.RequestAcceptGraceTimeout)
	if reqAcceptGraceTimeOut > 0 {
		logger.Infof("waiting %s for incoming requests to cease...", reqAcceptGraceTimeOut)
		time.Sleep(reqAcceptGraceTimeOut)
	}

	// Stop accepting new connection
	lis := e.listener()
	if lis != nil {
		return lis.Close()
	}

	graceTimeOut := time.Duration(e.lifecycle.GraceTimeOut)
	if graceTimeOut > 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, graceTimeOut)
		defer cancel()
		logger.Infof("waiting %s seconds before killing connections...", graceTimeOut)
	}

	if handler, ok := e.handler.(Shutdownable); ok {
		err := Shutdown(ctx, handler)
		if err != nil {
			logger.WithError(err).Errorf("error while shutting down tcp entrypoint")
		}
		return err
	}

	logger.Infof("tcp entrypoint shutted down")

	return nil
}

func (e *EntryPoint) Close() error {
	var err error
	if handler, ok := e.handler.(io.Closer); ok {
		err = Close(handler)
	}
	return err
}
