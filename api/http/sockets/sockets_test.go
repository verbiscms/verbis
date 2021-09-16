package sockets

import (
	"bytes"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/logger"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

// SocketsTestSuite defines the helper used for websocket
// testing.
type SocketsTestSuite struct {
	suite.Suite
	LogWriter bytes.Buffer
}

// TestSockets asserts testing has begun.
func TestSockets(t *testing.T) {
	suite.Run(t, &SocketsTestSuite{})
}

// SetupSuite assign the logger to a buffer.
func (t *SocketsTestSuite) SetupSuite() {
	b := bytes.Buffer{}
	t.LogWriter = b
	logger.SetOutput(&t.LogWriter)
	logger.SetLevel(logrus.TraceLevel)
}

// Reset the log writer.
func (t *SocketsTestSuite) Reset() {
	t.LogWriter.Reset()
}

// Setup is a helper to obtain a a new test server handler
// for websocket testing.
func (t *SocketsTestSuite) Setup(handler http.Handler) (*websocket.Conn, func()) {
	// Create test server with the echo handler.
	s := httptest.NewServer(handler)

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fail("error dialing websocket", err)
	}

	return ws, func() {
		s.Close()
		ws.Close()
	}
}

func (t *SocketsTestSuite) TestClose() {
	t.Setup(Handler())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		Close()
	}()
	wg.Wait()
}
