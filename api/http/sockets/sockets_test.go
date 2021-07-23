package sockets

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/common/paths"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// SocketsTestSuite defines the helper used for websocket
// testing.
type SocketsTestSuite struct {
	suite.Suite
	logger  bytes.Buffer
	apiPath string
}

// TestSockets
//
// Assert testing has begun.
func TestSockets(t *testing.T) {
	suite.Run(t, &SocketsTestSuite{})
}

// BeforeTest
//
// Assign the logger to a buffer.
func (t *SocketsTestSuite) BeforeTest(suiteName, testName string) {
	b := bytes.Buffer{}
	t.logger = b
	logger.SetOutput(&t.logger)
}

// SetupSuite
//
// Reassign API path for testing.
func (t *SocketsTestSuite) SetupSuite() {
	logger.SetOutput(ioutil.Discard)
	wd, err := os.Getwd()
	t.NoError(err)
	t.apiPath = filepath.Join(filepath.Dir(wd), "../")
}

// Setup
//
// A helper to obtain a a new test server handler
// for testing.
func (t *SocketsTestSuite) Setup() (*websocket.Conn, func()) {
	//log := &bytes.Buffer{}
	//logger.Init(&environment.Env{})
	//logger.SetOutput(log)

	d := &deps.Deps{
		Store:  nil,
		Config: nil,
		Options: &domain.Options{
			ActiveTheme: "verbis",
		},
		Paths: paths.Paths{
			Base: t.apiPath + "/test/testdata",
		},
		Running: false,
	}

	// Create test server with the echo handler.
	s := httptest.NewServer(http.Handler(Admin(d)))

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fail("error dialing websocket", err) //nolint
	}

	return ws, func() {
		s.Close()
		ws.Close()
	}
}
