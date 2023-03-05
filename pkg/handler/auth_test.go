package handler

import (
	"bytes"
	"errors"
	"github.com/Smolvika/notebook.git"
	"github.com/Smolvika/notebook.git/pkg/service"
	mock_service "github.com/Smolvika/notebook.git/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, user notebook.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            notebook.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"name":"Test","username":"test", "password": "ghjghj"}`,
			inputUser: notebook.User{
				Name:     "Test",
				Username: "test",
				Password: "ghjghj",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user notebook.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:      "Some data is missing",
			inputBody: `{"name":"Test","password": "ghjghj"}`,
			inputUser: notebook.User{
				Name:     "Test",
				Password: "ghjghj",
			},
			mockBehavior:         func(s *mock_service.MockAuthorization, user notebook.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag"}`,
		},
		{
			name:      "Problems on the service",
			inputBody: `{"name":"Test","username":"test","password":"ghjghj"}`,
			inputUser: notebook.User{
				Name:     "Test",
				Username: "test",
				Password: "ghjghj",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user notebook.User) {
				s.EXPECT().CreateUser(user).Return(0, errors.New("wrong password"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"wrong password"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := New(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())

		})
	}

}

func TestHandler_signIn(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user notebook.User)
	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            notebook.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"test","password":"ghjghj"}`,
			inputUser: notebook.User{
				Username: "test",
				Password: "ghjghj",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user notebook.User) {
				s.EXPECT().GenerateToken(user.Username, user.Password).Return("token", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"token":"token"}`,
		},
		{
			name:      "bad request",
			inputBody: `{"password":"ghjghj"}`,
			inputUser: notebook.User{
				Password: "ghjghj",
			},
			mockBehavior:         func(s *mock_service.MockAuthorization, user notebook.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Key: 'singInInput.Username' Error:Field validation for 'Username' failed on the 'required' tag"}`,
		},
		{
			name:      "wrong data",
			inputBody: `{"username":"test","password":"ghjghj"}`,
			inputUser: notebook.User{
				Username: "test",
				Password: "ghjghj",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user notebook.User) {
				s.EXPECT().GenerateToken(user.Username, user.Password).Return("token", errors.New("wrong data"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"wrong data"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := New(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/sign-in", handler.signIn)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-in",
				bytes.NewBufferString(testCase.inputBody))

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)

		})
	}
}
