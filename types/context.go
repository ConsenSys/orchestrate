package types

import (
	log "github.com/sirupsen/logrus"
)

// Logger ...
type Logger struct {
	DefaultFields map[string]interface{}
}

// HandlerFunc is base type for a function processing a Trace
type HandlerFunc func(ctx *Context)

// Context allows us to transmit information through middlewares
type Context struct {
	// T stores information about transaction lifecycle in high level types
	T *Trace

	// Message that triggered Context execution (typically a sarama.ConsumerMessage)
	Msg interface{}

	// Keys is a key/value pair
	Keys map[string]interface{}

	// Handlers to be executed on context
	handlers []HandlerFunc

	// Handler being executed
	index int

	// Logger
	Logger Logger
}

// NewContext creates a new context
func NewContext() *Context {
	t := NewTrace()
	return &Context{
		T:     t,
		Keys:  make(map[string]interface{}),
		index: -1,
	}
}

// Reset re-initialize context
func (ctx *Context) Reset() {
	ctx.Msg = nil
	ctx.T.Reset()
	ctx.Keys = make(map[string]interface{})
	ctx.handlers = nil
	ctx.index = -1
}

// Next should be used in middleware
// It executes pending handlers
func (ctx *Context) Next() {
	ctx.index++
	for s := len(ctx.handlers); ctx.index < s; ctx.index++ {
		ctx.handlers[ctx.index](ctx)
	}
}

// Error attaches an error to context.
func (ctx *Context) Error(err error) *Error {
	if err == nil {
		panic("err is nil")
	}

	e, ok := err.(*Error)
	if !ok {
		e = &Error{
			Err:  err,
			Type: ErrorTypeUnknown,
		}
	}
	ctx.T.Errors = append(ctx.T.Errors, e)

	return e
}

// Abort prevents pending handlers to be executed
func (ctx *Context) Abort() {
	ctx.index = len(ctx.handlers)
}

// AbortWithError calls `Abort()` and `Error()``
func (ctx *Context) AbortWithError(err error) *Error {
	ctx.Abort()
	return ctx.Error(err)
}

// Prepare re-initializes context, set handlers and set message
func (ctx *Context) Prepare(handlers []HandlerFunc, msg interface{}) {
	ctx.Reset()
	ctx.handlers = handlers
	ctx.Msg = msg
}

// AddDefaultFields creates a new context
func (l *Logger) AddDefaultFields(fields map[string]interface{}) *log.Entry {
	for k, v := range fields {
		l.DefaultFields[k] = v
	}
	return log.WithFields(l.DefaultFields)
}

// DelDefaultFields creates a new context
func (l *Logger) DelDefaultFields(fields []string) *log.Entry {
	for _, v := range fields {
		delete(l.DefaultFields, v)
	}
	return log.WithFields(l.DefaultFields)
}

// WithFields creates a new context
func (l *Logger) WithFields(fields map[string]interface{}) *log.Entry {
	ctxFields := l.DefaultFields

	for k, v := range fields {
		ctxFields[k] = v
	}
	return log.WithFields(ctxFields)
}
