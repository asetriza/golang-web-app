package rest

import (
	"bytes"
	"context"
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

func TestHandler_createTodo(t *testing.T) {
	type mockBehavior func(r *mock_service.MockTodo, todo common.Todo)

	todo := common.Todo{
		UserID:      1,
		Name:        "name",
		Description: "description",
		NotifyDate:  1,
		Done:        false,
	}

	testTable := []struct {
		name         string
		body         string
		input        common.Todo
		mockBehavior mockBehavior
		setUserID    gin.HandlerFunc
		statusCode   int
		responseBody string
	}{
		{
			name:  "OK",
			body:  `{"name":"name","description":"description","notifyDate":1}`,
			input: todo,
			mockBehavior: func(r *mock_service.MockTodo, todo common.Todo) {
				r.EXPECT().Create(context.Background(), todo).Return(1, nil)
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusOK,
			responseBody: `{"id":1}`,
		},
		{
			name:         "Empty fields",
			body:         `{"name":"","description":"description","notifyDate":1}`,
			input:        todo,
			mockBehavior: func(r *mock_service.MockTodo, todo common.Todo) {},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusBadRequest,
			responseBody: `{"message":"Key: 'Todo.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
		},
		{
			name:         "Internal server error",
			body:         `{"name":"name","description":"description","notifyDate":1}`,
			input:        todo,
			mockBehavior: func(r *mock_service.MockTodo, todo common.Todo) {},
			setUserID: func(c *gin.Context) {
				c.Set("afd", 1)
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"userCtx not found"}`,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockServ := mock_service.NewMockTodo(c)
			tc.mockBehavior(mockServ, tc.input)

			service := service.Service{Todo: mockServ}
			rest := REST{Service: &service}

			// Init Endpoint
			r := gin.New()
			r.POST("/api/todo", tc.setUserID, rest.createTodo)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/todo", bytes.NewBufferString(tc.body))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, w.Body.String(), tc.responseBody)
		})
	}
}
