package controllers

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

// getMediaMock is a helper to obtain a mock user controller
// for testing.
func getMediaMock(m models.MediaRepository) *MediaController {
	return &MediaController{
		store: &models.Store{
			Media: m,
		},
	}
}

// Test_NewMedia - Test construct
func Test_NewMedia(t *testing.T) {
	store := models.Store{}
	config := config.Configuration{}
	want := &MediaController{
		store:  &store,
		config: config,
	}
	got := newMedia(&store, config)
	assert.Equal(t, got, want)
}

// TestUserController_Get - Test Get route
//func TestMediaController_Get(t *testing.T) {
//
//	media := []domain.Media{
//		{
//			Id:          0,
//			UUID:        uuid.UUID{},
//			Url:         "",
//			Title:       nil,
//			Alt:         nil,
//			Description: nil,
//			FilePath:    "",
//			FileSize:    0,
//			FileName:    "",
//			Sizes:       nil,
//			Type:        "",
//			UserID:      0,
//			CreatedAt:   time.Time{},
//			UpdatedAt:   time.Time{},
//		},
//	}
//	pagination := http.Params{Page: 1, Limit: 15, OrderBy: "id", OrderDirection: "asc", Filters: nil}
//
//	tt := []struct {
//		name       string
//		want       string
//		status     int
//		message    string
//		mock func(u *modelMocks.MediaRepository)
//	}{
//		{
//			name:       "Success",
//			want:       `sdfdsfgsdfg`,
//			status:     200,
//			message:    "Successfully obtained users",
//			mock: func(u *modelMocks.MediaRepository) {
//				u.On("Get", pagination).Return(media, 1, nil)
//			},
//		},
//		{
//			name:       "Not Found",
//			want:       `{}`,
//			status:     200,
//			message:    "no users found",
//			mock: func(u *modelMocks.MediaRepository) {
//				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no users found"})
//			},
//		},
//		{
//			name:       "Conflict",
//			want:       `{}`,
//			status:     400,
//			message:    "conflict",
//			mock: func(u *modelMocks.MediaRepository) {
//				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.CONFLICT, Message: "conflict"})
//			},
//		},
//		{
//			name:       "Invalid",
//			want:       `{}`,
//			status:     400,
//			message:    "invalid",
//			mock: func(u *modelMocks.MediaRepository) {
//				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INVALID, Message: "invalid"})
//			},
//		},
//		{
//			name:       "Internal Error",
//			want:       `{}`,
//			status:     500,
//			message:    "internal",
//			mock: func(u *modelMocks.MediaRepository) {
//				u.On("Get", pagination).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal"})
//			},
//		},
//	}
//
//	for _, test := range tt {
//
//		t.Run(test.name, func(t *testing.T) {
//			rr := newResponseRecorder(t)
//			mock := &modelMocks.MediaRepository{}
//			test.mock(mock)
//
//			rr.RequestAndServe("GET", "/users", "/users", nil, func(g *gin.Context) {
//				getMediaMock(mock).Get(g)
//			})
//
//			assert.JSONEq(t, test.want, rr.Data())
//			assert.Equal(t, test.status, rr.recorder.Code)
//			assert.Equal(t, test.message, rr.Message())
//			assert.Equal(t, rr.recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
//		})
//	}
//}
