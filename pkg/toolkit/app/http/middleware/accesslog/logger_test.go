// nolint // reason
package accesslog

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/traefik/traefik/v2/pkg/types"
	traefiktypes "github.com/traefik/paerser/types"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	logFileNameSuffix       = "/traefik/logger/test.log"
	testContent             = "Hello, World"
	testServiceName         = "http://127.0.0.1/testService"
	testRouterName          = "testRouter"
	testStatus              = 123
	testContentSize   int64 = 12
	testHostname            = "TestHost"
	testUsername            = "TestUser"
	testPath                = "testpath"
	testPort                = 8181
	testProto               = "HTTP/0.0"
	testScheme              = "http"
	testMethod              = http.MethodPost
	testReferer             = "testReferer"
	testUserAgent           = "testUserAgent"
	testRetryAttempts       = 2
	testStart               = time.Now()
)

func TestLogRotation(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "traefik_")
	if err != nil {
		t.Fatalf("Error setting up temporary directory: %s", err)
	}
	defer os.RemoveAll(tempDir)

	fileName := filepath.Join(tempDir, "traefik.log")
	rotatedFileName := fileName + ".rotated"

	config := &types.AccessLog{FilePath: fileName, Format: CommonFormat}
	logHandler, err := NewHandler(config)
	if err != nil {
		t.Fatalf("Error creating new log handler: %s", err)
	}
	defer logHandler.Close()

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	next := func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}

	iterations := 20
	halfDone := make(chan bool)
	writeDone := make(chan bool)
	go func() {
		for i := 0; i < iterations; i++ {
			logHandler.ServeHTTP(recorder, req, http.HandlerFunc(next))
			if i == iterations/2 {
				halfDone <- true
			}
		}
		writeDone <- true
	}()

	<-halfDone
	err = os.Rename(fileName, rotatedFileName)
	if err != nil {
		t.Fatalf("Error renaming file: %s", err)
	}

	err = logHandler.Rotate()
	if err != nil {
		t.Fatalf("Error rotating file: %s", err)
	}

	select {
	case <-writeDone:
		gotLineCount := lineCount(t, fileName) + lineCount(t, rotatedFileName)
		if iterations != gotLineCount {
			t.Errorf("Wanted %d written log lines, got %d", iterations, gotLineCount)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatalf("test timed out")
	}

	close(halfDone)
	close(writeDone)
}

func lineCount(t *testing.T, fileName string) int {
	t.Helper()
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Error reading from file %s: %s", fileName, err)
	}

	count := 0
	for _, line := range strings.Split(string(fileContents), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		count++
	}

	return count
}

func TestLoggerCLF(t *testing.T) {
	logrus.SetFormatter(new(CommonLogFormatter))

	tmpDir := createTempDir(t, CommonFormat)
	defer os.RemoveAll(tmpDir)

	logFilePath := filepath.Join(tmpDir, logFileNameSuffix)
	config := &types.AccessLog{FilePath: logFilePath, Format: CommonFormat}
	doLogging(t, config)

	logData, err := ioutil.ReadFile(logFilePath)
	require.NoError(t, err)

	expectedLog := ` TestHost - TestUser [13/Apr/2016:07:14:19 -0700] "POST testpath HTTP/0.0" 123 12 "testReferer" "testUserAgent" 1 "testRouter" "http://127.0.0.1/testService" 1ms`
	assertValidLogData(t, expectedLog, logData)
}

func TestAsyncLoggerCLF(t *testing.T) {
	logrus.SetFormatter(new(CommonLogFormatter))

	tmpDir := createTempDir(t, CommonFormat)
	defer os.RemoveAll(tmpDir)

	logFilePath := filepath.Join(tmpDir, logFileNameSuffix)
	config := &types.AccessLog{FilePath: logFilePath, Format: CommonFormat, BufferingSize: 1024}
	doLogging(t, config)

	logData, err := ioutil.ReadFile(logFilePath)
	require.NoError(t, err)

	expectedLog := ` TestHost - TestUser [13/Apr/2016:07:14:19 -0700] "POST testpath HTTP/0.0" 123 12 "testReferer" "testUserAgent" 1 "testRouter" "http://127.0.0.1/testService" 1ms`
	assertValidLogData(t, expectedLog, logData)
}

func assertString(exp string) func(t *testing.T, actual interface{}) {
	return func(t *testing.T, actual interface{}) {
		t.Helper()

		assert.Equal(t, exp, actual)
	}
}

func assertNotEmpty() func(t *testing.T, actual interface{}) {
	return func(t *testing.T, actual interface{}) {
		t.Helper()

		assert.NotEqual(t, "", actual)
	}
}

func assertFloat64(exp float64) func(t *testing.T, actual interface{}) {
	return func(t *testing.T, actual interface{}) {
		t.Helper()

		assert.Equal(t, exp, actual)
	}
}

func assertFloat64NotZero() func(t *testing.T, actual interface{}) {
	return func(t *testing.T, actual interface{}) {
		t.Helper()

		assert.NotZero(t, actual)
	}
}

func TestLoggerJSON(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	testCases := []struct {
		desc     string
		config   *types.AccessLog
		tls      bool
		expected map[string]func(t *testing.T, value interface{})
	}{
		{
			desc: "default config",
			config: &types.AccessLog{
				FilePath: "",
				Format:   JSONFormat,
			},
			expected: map[string]func(t *testing.T, value interface{}){
				RequestContentSize:        assertFloat64(0),
				RequestHost:               assertString(testHostname),
				RequestAddr:               assertString(testHostname),
				RequestMethod:             assertString(testMethod),
				RequestPath:               assertString(testPath),
				RequestProtocol:           assertString(testProto),
				RequestScheme:             assertString(testScheme),
				RequestPort:               assertString("-"),
				DownstreamStatus:          assertFloat64(float64(testStatus)),
				DownstreamContentSize:     assertFloat64(float64(len(testContent))),
				OriginContentSize:         assertFloat64(float64(len(testContent))),
				OriginStatus:              assertFloat64(float64(testStatus)),
				RequestRefererHeader:      assertString(testReferer),
				RequestUserAgentHeader:    assertString(testUserAgent),
				RouterName:                assertString(testRouterName),
				ServiceURL:                assertString(testServiceName),
				ClientUsername:            assertString(testUsername),
				ClientHost:                assertString(testHostname),
				ClientPort:                assertString(fmt.Sprintf("%d", testPort)),
				ClientAddr:                assertString(fmt.Sprintf("%s:%d", testHostname, testPort)),
				"level":                   assertString("info"),
				"type":                    assertString("accesslog"),
				"msg":                     assertString("OK"),
				"downstream_Content-Type": assertString("text/plain; charset=utf-8"),
				RequestCount:              assertFloat64NotZero(),
				Duration:                  assertFloat64NotZero(),
				Overhead:                  assertFloat64NotZero(),
				RetryAttempts:             assertFloat64(float64(testRetryAttempts)),
				"time":                    assertNotEmpty(),
				StartLocal:                assertNotEmpty(),
				StartUTC:                  assertNotEmpty(),
			},
		},
		{
			desc: "default config, with TLS request",
			config: &types.AccessLog{
				FilePath: "",
				Format:   JSONFormat,
			},
			tls: true,
			expected: map[string]func(t *testing.T, value interface{}){
				RequestContentSize:        assertFloat64(0),
				RequestHost:               assertString(testHostname),
				RequestAddr:               assertString(testHostname),
				RequestMethod:             assertString(testMethod),
				RequestPath:               assertString(testPath),
				RequestProtocol:           assertString(testProto),
				RequestScheme:             assertString("https"),
				RequestPort:               assertString("-"),
				DownstreamStatus:          assertFloat64(float64(testStatus)),
				DownstreamContentSize:     assertFloat64(float64(len(testContent))),
				OriginContentSize:         assertFloat64(float64(len(testContent))),
				OriginStatus:              assertFloat64(float64(testStatus)),
				RequestRefererHeader:      assertString(testReferer),
				RequestUserAgentHeader:    assertString(testUserAgent),
				RouterName:                assertString(testRouterName),
				ServiceURL:                assertString(testServiceName),
				ClientUsername:            assertString(testUsername),
				ClientHost:                assertString(testHostname),
				ClientPort:                assertString(fmt.Sprintf("%d", testPort)),
				ClientAddr:                assertString(fmt.Sprintf("%s:%d", testHostname, testPort)),
				"level":                   assertString("info"),
				"type":                    assertString("accesslog"),
				"msg":                     assertString("OK"),
				"downstream_Content-Type": assertString("text/plain; charset=utf-8"),
				RequestCount:              assertFloat64NotZero(),
				Duration:                  assertFloat64NotZero(),
				Overhead:                  assertFloat64NotZero(),
				RetryAttempts:             assertFloat64(float64(testRetryAttempts)),
				"time":                    assertNotEmpty(),
				StartLocal:                assertNotEmpty(),
				StartUTC:                  assertNotEmpty(),
			},
		},
		{
			desc: "default config drop all fields",
			config: &types.AccessLog{
				FilePath: "",
				Format:   JSONFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "drop",
				},
			},
			expected: map[string]func(t *testing.T, value interface{}){
				"level":                   assertString("info"),
				"type":                    assertString("accesslog"),
				"msg":                     assertString("OK"),
				"time":                    assertNotEmpty(),
				"downstream_Content-Type": assertString("text/plain; charset=utf-8"),
				RequestRefererHeader:      assertString(testReferer),
				RequestUserAgentHeader:    assertString(testUserAgent),
			},
		},
		{
			desc: "default config drop all fields and headers",
			config: &types.AccessLog{
				FilePath: "",
				Format:   JSONFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "drop",
					Headers: &types.FieldHeaders{
						DefaultMode: "drop",
					},
				},
			},
			expected: map[string]func(t *testing.T, value interface{}){
				"level": assertString("info"),
				"type":  assertString("accesslog"),
				"msg":   assertString("OK"),
				"time":  assertNotEmpty(),
			},
		},
		{
			desc: "default config drop all fields and redact headers",
			config: &types.AccessLog{
				FilePath: "",
				Format:   JSONFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "drop",
					Headers: &types.FieldHeaders{
						DefaultMode: "redact",
					},
				},
			},
			expected: map[string]func(t *testing.T, value interface{}){
				"level":                   assertString("info"),
				"type":                    assertString("accesslog"),
				"msg":                     assertString("OK"),
				"time":                    assertNotEmpty(),
				"downstream_Content-Type": assertString("REDACTED"),
				RequestRefererHeader:      assertString("REDACTED"),
				RequestUserAgentHeader:    assertString("REDACTED"),
			},
		},
		{
			desc: "default config drop all fields and headers but kept someone",
			config: &types.AccessLog{
				FilePath: "",
				Format:   JSONFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "drop",
					Names: map[string]string{
						RequestHost: "keep",
					},
					Headers: &types.FieldHeaders{
						DefaultMode: "drop",
						Names: map[string]string{
							"Referer": "keep",
						},
					},
				},
			},
			expected: map[string]func(t *testing.T, value interface{}){
				RequestHost:          assertString(testHostname),
				"level":              assertString("info"),
				"type":               assertString("accesslog"),
				"msg":                assertString("OK"),
				"time":               assertNotEmpty(),
				RequestRefererHeader: assertString(testReferer),
			},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			tmpDir := createTempDir(t, JSONFormat)
			defer os.RemoveAll(tmpDir)

			logFilePath := filepath.Join(tmpDir, logFileNameSuffix)

			test.config.FilePath = logFilePath
			if test.tls {
				doLoggingTLS(t, test.config)
			} else {
				doLogging(t, test.config)
			}

			logData, err := ioutil.ReadFile(logFilePath)
			require.NoError(t, err)

			jsonData := make(map[string]interface{})
			err = json.Unmarshal(logData, &jsonData)
			require.NoError(t, err)

			assert.Equal(t, len(test.expected), len(jsonData))

			for field, assertion := range test.expected {
				assertion(t, jsonData[field])
			}
		})
	}
}

func TestNewLogHandlerOutputStdout(t *testing.T) {
	logrus.SetFormatter(new(CommonLogFormatter))

	testCases := []struct {
		desc        string
		config      *types.AccessLog
		expectedLog string
	}{
		{
			desc: "default config",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
			},
			expectedLog: `TestHost - TestUser [13/Apr/2016:07:14:19 -0700] "POST testpath HTTP/0.0" 123 12 "testReferer" "testUserAgent" 23 "testRouter" "http://127.0.0.1/testService" 1ms`,
		},
		{
			desc: "default config with empty filters",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Filters:  &types.AccessLogFilters{},
			},
			expectedLog: `TestHost - TestUser [13/Apr/2016:07:14:19 -0700] "POST testpath HTTP/0.0" 123 12 "testReferer" "testUserAgent" 23 "testRouter" "http://127.0.0.1/testService" 1ms`,
		},
		{
			desc: "Status code filter not matching",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Filters: &types.AccessLogFilters{
					StatusCodes: []string{"200"},
				},
			},
			expectedLog: ``,
		},
		{
			desc: "Status code filter matching",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Filters: &types.AccessLogFilters{
					StatusCodes: []string{"123"},
				},
			},
			expectedLog: `TestHost - TestUser [13/Apr/2016:07:14:19 -0700] "POST testpath HTTP/0.0" 123 12 "testReferer" "testUserAgent" 23 "testRouter" "http://127.0.0.1/testService" 1ms`,
		},
		{
			desc: "Duration filter not matching",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Filters: &types.AccessLogFilters{
					MinDuration: traefiktypes.Duration(1 * time.Hour),
				},
			},
			expectedLog: ``,
		},
		{
			desc: "Duration filter matching",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Filters: &types.AccessLogFilters{
					MinDuration: traefiktypes.Duration(1 * time.Millisecond),
				},
			},
			expectedLog: `TestHost - TestUser [13/Apr/2016:07:14:19 -0700] "POST testpath HTTP/0.0" 123 12 "testReferer" "testUserAgent" 23 "testRouter" "http://127.0.0.1/testService" 1ms`,
		},
		{
			desc: "Retry attempts filter matching",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Filters: &types.AccessLogFilters{
					RetryAttempts: true,
				},
			},
			expectedLog: `TestHost - TestUser [13/Apr/2016:07:14:19 -0700] "POST testpath HTTP/0.0" 123 12 "testReferer" "testUserAgent" 23 "testRouter" "http://127.0.0.1/testService" 1ms`,
		},
		{
			desc: "Default mode keep",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "keep",
				},
			},
			expectedLog: `TestHost - TestUser [13/Apr/2016:07:14:19 -0700] "POST testpath HTTP/0.0" 123 12 "testReferer" "testUserAgent" 23 "testRouter" "http://127.0.0.1/testService" 1ms`,
		},
		{
			desc: "Default mode keep with override",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "keep",
					Names: map[string]string{
						ClientHost: "drop",
					},
				},
			},
			expectedLog: `- - TestUser [13/Apr/2016:07:14:19 -0700] "POST testpath HTTP/0.0" 123 12 "testReferer" "testUserAgent" 23 "testRouter" "http://127.0.0.1/testService" 1ms`,
		},
		{
			desc: "Default mode drop",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "drop",
				},
			},
			expectedLog: `- - - [-] "- - -" - - "testReferer" "testUserAgent" - "-" "-" 0ms`,
		},
		{
			desc: "Default mode drop with override",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "drop",
					Names: map[string]string{
						ClientHost:     "drop",
						ClientUsername: "keep",
					},
				},
			},
			expectedLog: `- - TestUser [-] "- - -" - - "testReferer" "testUserAgent" - "-" "-" 0ms`,
		},
		{
			desc: "Default mode drop with header dropped",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "drop",
					Names: map[string]string{
						ClientHost:     "drop",
						ClientUsername: "keep",
					},
					Headers: &types.FieldHeaders{
						DefaultMode: "drop",
					},
				},
			},
			expectedLog: `- - TestUser [-] "- - -" - - "-" "-" - "-" "-" 0ms`,
		},
		{
			desc: "Default mode drop with header redacted",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "drop",
					Names: map[string]string{
						ClientHost:     "drop",
						ClientUsername: "keep",
					},
					Headers: &types.FieldHeaders{
						DefaultMode: "redact",
					},
				},
			},
			expectedLog: `- - TestUser [-] "- - -" - - "REDACTED" "REDACTED" - "-" "-" 0ms`,
		},
		{
			desc: "Default mode drop with header redacted",
			config: &types.AccessLog{
				FilePath: "",
				Format:   CommonFormat,
				Fields: &types.AccessLogFields{
					DefaultMode: "drop",
					Names: map[string]string{
						ClientHost:     "drop",
						ClientUsername: "keep",
					},
					Headers: &types.FieldHeaders{
						DefaultMode: "keep",
						Names: map[string]string{
							"Referer": "redact",
						},
					},
				},
			},
			expectedLog: `- - TestUser [-] "- - -" - - "REDACTED" "testUserAgent" - "-" "-" 0ms`,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			// NOTE: It is not possible to run these cases in parallel because we capture Stdout

			file, restoreStdout := captureStdout(t)
			defer restoreStdout()

			doLogging(t, test.config)

			written, err := ioutil.ReadFile(file.Name())
			require.NoError(t, err, "unable to read captured stdout from file")
			assertValidLogData(t, test.expectedLog, written)
		})
	}
}

func assertValidLogData(t *testing.T, expected string, logData []byte) {
	if expected == "" {
		assert.Zero(t, len(logData))
		t.Log(string(logData))
		return
	}

	result, err := ParseAccessLog(string(logData))
	require.NoError(t, err)

	resultExpected, err := ParseAccessLog(expected)
	require.NoError(t, err)

	formatErrMessage := fmt.Sprintf(`
	Expected: %s
	Actual:   %s`, expected, string(logData))

	require.Equal(t, len(resultExpected), len(result), formatErrMessage)
	assert.Equal(t, resultExpected[ClientHost], result[ClientHost], formatErrMessage)
	assert.Equal(t, resultExpected[ClientUsername], result[ClientUsername], formatErrMessage)
	assert.Equal(t, resultExpected[RequestMethod], result[RequestMethod], formatErrMessage)
	assert.Equal(t, resultExpected[RequestPath], result[RequestPath], formatErrMessage)
	assert.Equal(t, resultExpected[RequestProtocol], result[RequestProtocol], formatErrMessage)
	assert.Equal(t, resultExpected[OriginStatus], result[OriginStatus], formatErrMessage)
	assert.Equal(t, resultExpected[OriginContentSize], result[OriginContentSize], formatErrMessage)
	assert.Equal(t, resultExpected[RequestRefererHeader], result[RequestRefererHeader], formatErrMessage)
	assert.Equal(t, resultExpected[RequestUserAgentHeader], result[RequestUserAgentHeader], formatErrMessage)
	assert.Regexp(t, regexp.MustCompile("[0-9]*"), result[RequestCount], formatErrMessage)
	assert.Equal(t, resultExpected[RouterName], result[RouterName], formatErrMessage)
	assert.Equal(t, resultExpected[ServiceURL], result[ServiceURL], formatErrMessage)
	assert.Regexp(t, regexp.MustCompile("[0-9]*ms"), result[Duration], formatErrMessage)
}

func captureStdout(t *testing.T) (out *os.File, restoreStdout func()) {
	file, err := ioutil.TempFile("", "testlogger")
	require.NoError(t, err, "failed to create temp file")

	original := os.Stdout
	os.Stdout = file

	restoreStdout = func() {
		os.Stdout = original
		os.RemoveAll(file.Name())
	}

	return file, restoreStdout
}

func createTempDir(t *testing.T, prefix string) string {
	tmpDir, err := ioutil.TempDir("", prefix)
	require.NoError(t, err, "failed to create temp dir")

	return tmpDir
}

func doLoggingTLSOpt(t *testing.T, config *types.AccessLog, enableTLS bool) {
	logger, err := NewHandler(config)
	require.NoError(t, err)
	defer logger.Close()

	if config.FilePath != "" {
		_, err = os.Stat(config.FilePath)
		require.NoError(t, err, fmt.Sprintf("logger should create %s", config.FilePath))
	}

	req := &http.Request{
		Header: map[string][]string{
			"User-Agent": {testUserAgent},
			"Referer":    {testReferer},
		},
		Proto:      testProto,
		Host:       testHostname,
		Method:     testMethod,
		RemoteAddr: fmt.Sprintf("%s:%d", testHostname, testPort),
		URL: &url.URL{
			User: url.UserPassword(testUsername, ""),
			Path: testPath,
		},
	}
	if enableTLS {
		req.TLS = &tls.ConnectionState{}
	}

	logger.ServeHTTP(httptest.NewRecorder(), req, http.HandlerFunc(logWriterTestHandlerFunc))
}

func doLoggingTLS(t *testing.T, config *types.AccessLog) {
	doLoggingTLSOpt(t, config, true)
}

func doLogging(t *testing.T, config *types.AccessLog) {
	doLoggingTLSOpt(t, config, false)
}

func logWriterTestHandlerFunc(rw http.ResponseWriter, r *http.Request) {
	if _, err := rw.Write([]byte(testContent)); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	logData := GetLogData(r)
	if logData != nil {
		logData.Core[RouterName] = testRouterName
		logData.Core[ServiceURL] = testServiceName
		logData.Core[OriginStatus] = testStatus
		logData.Core[OriginContentSize] = testContentSize
		logData.Core[RetryAttempts] = testRetryAttempts
		logData.Core[StartUTC] = testStart.UTC()
		logData.Core[StartLocal] = testStart.Local()
	} else {
		http.Error(rw, "LogData is nil", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(testStatus)
}
