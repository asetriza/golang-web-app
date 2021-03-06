package rest

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
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
			name:  "OK",
			body:  `{"name":"name","description":"description","notifyDate":1}`,
			input: todo,
			mockBehavior: func(r *mock_service.MockTodo, todo common.Todo) {
				r.EXPECT().Create(context.Background(), todo).Return(0, errors.New("internal server error"))
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"internal server error"}`,
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

func TestHandler_getTodos(t *testing.T) {
	type mockBehavior func(r *mock_service.MockTodo, userID int, pagination common.Pagination)

	testTable := []struct {
		name         string
		body         string
		userID       int
		mockBehavior mockBehavior
		setUserID    gin.HandlerFunc
		pagination   common.Pagination
		statusCode   int
		responseBody string
	}{
		{
			name:   "OK",
			body:   ``,
			userID: 1,
			mockBehavior: func(r *mock_service.MockTodo, userID int, pagination common.Pagination) {
				todo := common.Todo{
					ID:          0,
					UserID:      0,
					Name:        "",
					Description: "",
					NotifyDate:  0,
					Done:        false,
				}
				r.EXPECT().GetAll(context.Background(), userID, pagination).Return([]common.Todo{todo}, nil)
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			pagination: common.Pagination{
				CurrentPage:  1,
				ItemsPerPage: 2,
			},
			statusCode:   http.StatusOK,
			responseBody: `{"todos":[{"id":0,"userId":0,"name":"","description":"","notifyDate":0,"done":false}]}`,
		},
		{
			name:   "GetAll todos error",
			body:   ``,
			userID: 1,
			mockBehavior: func(r *mock_service.MockTodo, userID int, pagination common.Pagination) {
				r.EXPECT().GetAll(context.Background(), userID, pagination).Return([]common.Todo{}, errors.New("Internal server error"))
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			pagination: common.Pagination{
				CurrentPage:  1,
				ItemsPerPage: 2,
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"Internal server error"}`,
		},
		{
			name:         "user id error",
			body:         ``,
			userID:       1,
			mockBehavior: func(r *mock_service.MockTodo, userID int, pagination common.Pagination) {},
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
			tc.mockBehavior(mockServ, tc.userID, tc.pagination)

			service := service.Service{Todo: mockServ}
			rest := REST{Service: &service}

			// Init Endpoint
			r := gin.New()
			r.GET("/api/todo", tc.setUserID, rest.getTodos)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/todo?currentPage=1&itemsPerPage=2", bytes.NewBufferString(tc.body))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, w.Body.String(), tc.responseBody)
		})
	}
}

func TestHandler_getTodo(t *testing.T) {
	type mockBehavior func(r *mock_service.MockTodo, userID int, todoID interface{})

	testTable := []struct {
		name         string
		body         string
		userID       int
		todoID       interface{}
		mockBehavior mockBehavior
		setUserID    gin.HandlerFunc
		statusCode   int
		responseBody string
	}{
		{
			name:   "OK",
			body:   ``,
			userID: 1,
			todoID: 1,
			mockBehavior: func(r *mock_service.MockTodo, userID int, todoID interface{}) {
				r.EXPECT().Get(context.Background(), userID, todoID).Return(common.Todo{}, nil)
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusOK,
			responseBody: `{"todo":{"id":0,"userId":0,"name":"","description":"","notifyDate":0,"done":false}}`,
		},
		{
			name:   "Get todo error",
			body:   ``,
			userID: 1,
			todoID: 1,
			mockBehavior: func(r *mock_service.MockTodo, userID int, todoID interface{}) {
				r.EXPECT().Get(context.Background(), userID, todoID).Return(common.Todo{}, sql.ErrNoRows)
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusNotFound,
			responseBody: `{"message":"sql: no rows in result set"}`,
		},
		{
			name:         "user id error",
			body:         ``,
			userID:       1,
			todoID:       1,
			mockBehavior: func(r *mock_service.MockTodo, userID int, todoID interface{}) {},
			setUserID: func(c *gin.Context) {
				c.Set("afd", 1)
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"userCtx not found"}`,
		},
		{
			name:         "todo id param string error",
			body:         ``,
			userID:       1,
			todoID:       "asd",
			mockBehavior: func(r *mock_service.MockTodo, userID int, todoID interface{}) {},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusBadRequest,
			responseBody: `{"message":"id param must be int"}`,
		},
		{
			name:   "internal server error",
			body:   ``,
			userID: 1,
			todoID: 1,
			mockBehavior: func(r *mock_service.MockTodo, userID int, todoID interface{}) {
				r.EXPECT().Get(context.Background(), userID, todoID).Return(common.Todo{}, errors.New("internal server error"))
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"internal server error"}`,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockServ := mock_service.NewMockTodo(c)
			tc.mockBehavior(mockServ, tc.userID, tc.todoID)

			service := service.Service{Todo: mockServ}
			rest := REST{Service: &service}

			// Init Endpoint
			r := gin.New()
			r.GET("/api/todo/:id", tc.setUserID, rest.getTodo)

			// Create Request
			w := httptest.NewRecorder()
			url := fmt.Sprintf("/api/todo/%v", tc.todoID)
			fmt.Println(url)
			req := httptest.NewRequest("GET", url, bytes.NewBufferString(tc.body))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, w.Body.String(), tc.responseBody)
		})
	}
}

func TestHandler_updateTodo(t *testing.T) {
	type mockBehavior func(r *mock_service.MockTodo, todo common.Todo)

	testTable := []struct {
		name         string
		body         string
		todoID       int
		todo         common.Todo
		mockBehavior mockBehavior
		setUserID    gin.HandlerFunc
		statusCode   int
		responseBody string
	}{
		{
			name:   "OK",
			body:   `{"name":"name","description":"description","notifyDate":1,"done":false}`,
			todoID: 1,
			todo: common.Todo{
				ID:          1,
				UserID:      1,
				Name:        "name",
				Description: "description",
				NotifyDate:  1,
				Done:        false,
			},
			mockBehavior: func(r *mock_service.MockTodo, todo common.Todo) {
				r.EXPECT().Update(context.Background(), todo).Return(nil)
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusOK,
			responseBody: `{"id":1}`,
		},
		{
			name:   "OK",
			body:   `{"name":"name","notifyDate":1,"done":false}`,
			todoID: 1,
			todo: common.Todo{
				ID:          1,
				UserID:      1,
				Name:        "name",
				Description: "description",
				NotifyDate:  1,
				Done:        false,
			},
			mockBehavior: func(r *mock_service.MockTodo, todo common.Todo) {},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusBadRequest,
			responseBody: `{"message":"Key: 'Todo.Description' Error:Field validation for 'Description' failed on the 'required' tag"}`,
		},
		{
			name:   "OK",
			body:   `{"name":"name","description":"description","notifyDate":1,"done":false}`,
			todoID: 1,
			todo: common.Todo{
				ID:          1,
				UserID:      1,
				Name:        "name",
				Description: "description",
				NotifyDate:  1,
				Done:        false,
			},
			mockBehavior: func(r *mock_service.MockTodo, todo common.Todo) {
				r.EXPECT().Update(context.Background(), todo).Return(errors.New("internal server error"))
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"internal server error"}`,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockServ := mock_service.NewMockTodo(c)
			tc.mockBehavior(mockServ, tc.todo)

			service := service.Service{Todo: mockServ}
			rest := REST{Service: &service}

			// Init Endpoint
			r := gin.New()
			r.PUT("/api/todo/:id", tc.setUserID, rest.updateTodo)

			// Create Request
			w := httptest.NewRecorder()
			url := fmt.Sprintf("/api/todo/%d", tc.todoID)
			req := httptest.NewRequest("PUT", url, bytes.NewBufferString(tc.body))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, w.Body.String(), tc.responseBody)
		})
	}
}

func TestHandler_deleteTodo(t *testing.T) {
	type mockBehavior func(r *mock_service.MockTodo, userID, todoID int)

	testTable := []struct {
		name         string
		body         string
		todoID       int
		userID       int
		mockBehavior mockBehavior
		setUserID    gin.HandlerFunc
		statusCode   int
		responseBody string
	}{
		{
			name:   "OK",
			body:   ``,
			todoID: 1,
			userID: 1,
			mockBehavior: func(r *mock_service.MockTodo, userID, todoID int) {
				r.EXPECT().Delete(context.Background(), userID, todoID).Return(nil)
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusOK,
			responseBody: `{"id":1}`,
		},
		{
			name:   "OK",
			body:   ``,
			todoID: 1,
			userID: 1,
			mockBehavior: func(r *mock_service.MockTodo, userID, todoID int) {
				r.EXPECT().Delete(context.Background(), userID, todoID).Return(errors.New("internal server error"))
			},
			setUserID: func(c *gin.Context) {
				c.Set(userCtx, 1)
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"internal server error"}`,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockServ := mock_service.NewMockTodo(c)
			tc.mockBehavior(mockServ, tc.userID, tc.todoID)

			service := service.Service{Todo: mockServ}
			rest := REST{Service: &service}

			// Init Endpoint
			r := gin.New()
			r.DELETE("/api/todo/:id", tc.setUserID, rest.deleteTodo)

			// Create Request
			w := httptest.NewRecorder()
			url := fmt.Sprintf("/api/todo/%d", tc.todoID)
			req := httptest.NewRequest("DELETE", url, bytes.NewBufferString(tc.body))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, w.Body.String(), tc.responseBody)
		})
	}
}
