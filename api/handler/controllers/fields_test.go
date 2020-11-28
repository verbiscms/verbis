package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getFieldsMock is a helper to obtain a mock fields controller
// for testing.
func getFieldsMock(m models.FieldsRepository) *AuthController {
	return &AuthController{
		store: &models.Store{
			Fields: m,
		},
	}
}

// Test_NewAuth - Test construct
func Test_NewFields(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &AuthController{
		store:  &store,
		config: config,
	}
	got := newFields(&store, config)
	assert.Equal(t, got, want)
}

// TestAuthController_Get - Test Get route
func TestFieldController_Get(t *testing.T) {


	tt := map[string]struct {
		name    string
		want    string
		status  int
		message string
		input   interface{}
		mock    func(m *mocks.FieldsRepository)
	}{
		"Success": {
			want:    `{"biography":null,"created_at":"0001-01-01T00:00:00Z","email":"","email_verified_at":null,"facebook":null,"first_name":"","id":0,"instagram":null,"last_name":"","linked_in":null,"profile_picture_id":null,"role":{"description":"","id":0,"name":""},"twitter":null,"updated_at":"0001-01-01T00:00:00Z","uuid":"00000000-0000-0000-0000-000000000000"}`,
			status:  200,
			message: "Successfully logged in & session started",
			input:   login,
			mock: func(m *mocks.FieldsRepository) {
				m.On("GetLayout").Return(&[]domain.FieldGroup{}, nil)
			},
		},
	}

	for name, test := range tt {

		t.Run(name, func(t *testing.T) {
			rr := newTestSuite(t)
			mock := &mocks.FieldsRepository{}
			test.mock(mock)

			body, err := json.Marshal(test.input)
			if err != nil {
				t.Fatal(err)
			}

			rr.RequestAndServe("GET", "/fields", "/fields", bytes.NewBuffer(body), func(g *gin.Context) {
				getFieldsMock(mock).Login(g)
			})

			rr.Run(test.want, test.status, test.message)
		})
	}
}
}