package sockets

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// SocketsTestSuite defines the helper used for websocket
// testing.
type SocketsTestSuite struct {
	test.HandlerSuite
}

// TestSockets
//
// Assert testing has begun.
func TestSockets(t *testing.T) {
	suite.Run(t, &SocketsTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a a new test server handler
// for testing.
func (t *SocketsTestSuite) Setup() (*websocket.Conn, func()) {
	d := &deps.Deps{
		Store:  nil,
		Config: nil,
		Site:   domain.Site{},
		Options: &domain.Options{
			ActiveTheme: "test",
		},
		Paths:   paths.Paths{},
		Running: false,
	}

	// Create test server with the echo handler.
	s := httptest.NewServer(http.Handler(Admin(d)))

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
