package rest

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asetriza/golang-web-app/internal/common"
	"github.com/asetriza/golang-web-app/internal/service"
	mock_service "github.com/asetriza/golang-web-app/internal/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_singUp(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, input common.User, clientIP string)

	user := common.User{
		Name:     "name",
		Username: "username",
		Email:    "email@gmail.com",
		Password: "password",
	}

	testTable := []struct {
		name         string
		body         string
		clientIP     string
		input        common.User
		mockBehavior mockBehavior
		statusCode   int
		responseBody string
	}{
		{
			name:     "OK",
			body:     `{"name":"name","email":"email@gmail.com","username":"username","password":"password"}`,
			clientIP: "192.0.2.1",
			input:    user,
			mockBehavior: func(r *mock_service.MockAuthorization, input common.User, clientIP string) {
				r.EXPECT().CreateUser(context.Background(), input, clientIP).Return(service.Credentials{}, nil)
			},
			statusCode:   http.StatusOK,
			responseBody: `{"credentials":{"token":"","refreshToken":""}}`,
		},
		{
			name:         "Empty name",
			body:         `{"name":"","email":"email@gmail.com","username":"username","password":"password"}`,
			clientIP:     "192.0.2.1",
			input:        user,
			mockBehavior: func(r *mock_service.MockAuthorization, input common.User, clientIP string) {},
			statusCode:   http.StatusBadRequest,
			responseBody: `{"message":"Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
		},
		{
			name:         "Incorrect email",
			body:         `{"name":"name","email":"email","username":"username","password":"password"}`,
			clientIP:     "192.0.2.1",
			input:        user,
			mockBehavior: func(r *mock_service.MockAuthorization, input common.User, clientIP string) {},
			statusCode:   http.StatusBadRequest,
			responseBody: `{"message":"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag"}`,
		},
		{
			name:     "Internal server error",
			body:     `{"name":"name","email":"email@gmail.com","username":"username","password":"password"}`,
			clientIP: "192.0.2.1",
			input:    user,
			mockBehavior: func(r *mock_service.MockAuthorization, input common.User, clientIP string) {
				r.EXPECT().CreateUser(context.Background(), input, clientIP).Return(service.Credentials{}, errors.New("Internal server error"))
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"Internal server error"}`,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockServ := mock_service.NewMockAuthorization(c)
			tc.mockBehavior(mockServ, tc.input, tc.clientIP)

			service := service.Service{Authorization: mockServ}
			rest := REST{Service: &service}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-up", rest.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(tc.body))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, w.Body.String(), tc.responseBody)
		})
	}
}

func TestHandler_singIn(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, username, password, clientIP string)

	user := signInInput{
		Username: "username",
		Password: "password",
	}

	testTable := []struct {
		name         string
		body         string
		clientIP     string
		input        signInInput
		mockBehavior mockBehavior
		statusCode   int
		responseBody string
	}{
		{
			name:     "OK",
			body:     `{"username":"username","password":"password"}`,
			clientIP: "192.0.2.1",
			input:    user,
			mockBehavior: func(r *mock_service.MockAuthorization, username, password, clientIP string) {
				r.EXPECT().CreateCredentials(context.Background(), username, password, clientIP).Return(service.Credentials{}, nil)
			},
			statusCode:   http.StatusOK,
			responseBody: `{"credentials":{"token":"","refreshToken":""}}`,
		},
		{
			name:         "Empty field username",
			body:         `{"username":"","password":"password"}`,
			clientIP:     "192.0.2.1",
			input:        user,
			mockBehavior: func(r *mock_service.MockAuthorization, username, password, clientIP string) {},
			statusCode:   http.StatusBadRequest,
			responseBody: `{"message":"Key: 'signInInput.Username' Error:Field validation for 'Username' failed on the 'required' tag"}`,
		},
		{
			name:     "Internal server error",
			body:     `{"username":"username","password":"password"}`,
			clientIP: "192.0.2.1",
			input:    user,
			mockBehavior: func(r *mock_service.MockAuthorization, username, password, clientIP string) {
				r.EXPECT().CreateCredentials(context.Background(), username, password, clientIP).Return(service.Credentials{}, errors.New("Internal server error"))
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"Internal server error"}`,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockServ := mock_service.NewMockAuthorization(c)
			tc.mockBehavior(mockServ, tc.input.Username, tc.input.Password, tc.clientIP)

			service := service.Service{Authorization: mockServ}
			rest := REST{Service: &service}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-in", rest.signIn)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(tc.body))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, w.Body.String(), tc.responseBody)
		})
	}
}
