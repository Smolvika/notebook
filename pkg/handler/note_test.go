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

func TestHandler_createNote(t *testing.T) {
	type mockBehavior func(r *mock_service.MockNote, userId int, note notebook.Note)

	testTable := []struct {
		name                 string
		userInfo             notebook.UsersNote
		inputBody            string
		inputUser            notebook.Note
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			userInfo:  notebook.UsersNote{UserId: 1},
			inputBody: `{"date":"20.02","description":"test"}`,
			inputUser: notebook.Note{
				Date:        "20.02",
				Description: "test",
			},
			mockBehavior: func(r *mock_service.MockNote, userId int, note notebook.Note) {
				r.EXPECT().Create(userId, note).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:      "userId Is Missing",
			userInfo:  notebook.UsersNote{UserId: -10},
			inputBody: `{"date":"20.02","description":"test"}`,
			inputUser: notebook.Note{
				Date:        "20.02",
				Description: "test",
			},
			mockBehavior:         func(r *mock_service.MockNote, userId int, note notebook.Note) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"user id not found"}`,
		},
		{
			name:      "userId Invalid Type",
			inputBody: `{"date":"20.02","description":"test"}`,
			inputUser: notebook.Note{
				Date:        "20.02",
				Description: "test",
			},
			mockBehavior:         func(r *mock_service.MockNote, userId int, note notebook.Note) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"user id is of invalid type"}`,
		},
		{
			name:      "Some data is missing",
			userInfo:  notebook.UsersNote{UserId: 1},
			inputBody: `{"data":"20.02"}`,
			inputUser: notebook.Note{
				Date: "20.02",
			},
			mockBehavior:         func(r *mock_service.MockNote, userId int, note notebook.Note) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Key: 'Note.Description' Error:Field validation for 'Description' failed on the 'required' tag"}`,
		},
		{
			name:      "Service Error",
			userInfo:  notebook.UsersNote{UserId: 1},
			inputBody: `{"date":"20.02","description":"test"}`,
			inputUser: notebook.Note{
				Date:        "20.02",
				Description: "test",
			},
			mockBehavior: func(r *mock_service.MockNote, userId int, note notebook.Note) {
				r.EXPECT().Create(userId, note).Return(0, errors.New("service error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"service error"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			note := mock_service.NewMockNote(c)
			testCase.mockBehavior(note, testCase.userInfo.UserId, testCase.inputUser)

			services := &service.Service{Note: note}
			handler := NewHandler(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/api/notes", func(c *gin.Context) {
				if testCase.userInfo.UserId == 0 {
					c.Set(userCtx, "some wrong info")
				} else if testCase.userInfo.UserId < 0 {
					c.Set("err", "some wrong info")
				} else {
					c.Set(userCtx, testCase.userInfo.UserId)
				}
			},
				handler.createNote)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/notes",
				bytes.NewBufferString(testCase.inputBody))

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())

		})
	}

}

func TestHandler_getAllNotes(t *testing.T) {
	type mockBehavior func(r *mock_service.MockNote, userId int)
	testTable := []struct {
		name                 string
		userInfo             notebook.UsersNote
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			userInfo: notebook.UsersNote{UserId: 1},
			mockBehavior: func(r *mock_service.MockNote, userId int) {
				r.EXPECT().GetAll(userId).Return([]notebook.Note{{Id: 1, Date: "26.02", Description: "купить корм"}}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"id":1,"date":"26.02","description":"купить корм"}]`,
		},
		{
			name:                 "userId Is Missing",
			userInfo:             notebook.UsersNote{UserId: -1},
			mockBehavior:         func(r *mock_service.MockNote, userId int) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"user id not found"}`,
		},
		{
			name:                 "userId Invalid Type",
			mockBehavior:         func(r *mock_service.MockNote, userId int) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"user id is of invalid type"}`,
		},
		{
			name:     "Service Error",
			userInfo: notebook.UsersNote{UserId: 1},
			mockBehavior: func(r *mock_service.MockNote, userId int) {
				r.EXPECT().GetAll(userId).Return(nil, errors.New("service error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"service error"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			note := mock_service.NewMockNote(c)
			testCase.mockBehavior(note, testCase.userInfo.UserId)

			services := &service.Service{Note: note}
			handler := NewHandler(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/api/notes", func(c *gin.Context) {
				if testCase.userInfo.UserId == 0 {
					c.Set(userCtx, "some wrong info")
				} else if testCase.userInfo.UserId < 0 {
					c.Set("err", "some wrong info")
				} else {
					c.Set(userCtx, testCase.userInfo.UserId)
				}
			}, handler.getAllNotes)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/notes", nil)

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())

		})
	}

}

func TestHandler_getNoteById(t *testing.T) {
	type mockBehavior func(note *mock_service.MockNote, userId, noteId int)

	testTable := []struct {
		name                 string
		userInfo             notebook.UsersNote
		Id                   string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			userInfo: notebook.UsersNote{UserId: 1, NotesId: 1},
			Id:       "1",
			mockBehavior: func(r *mock_service.MockNote, userId, noteId int) {
				r.EXPECT().GetById(userId, noteId).Return(notebook.Note{Id: 1, Date: "26.02", Description: "купить корм"}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"date":"26.02","description":"купить корм"}`,
		},
		{
			name:                 "userId Is Missing",
			userInfo:             notebook.UsersNote{UserId: -1},
			Id:                   "-1",
			mockBehavior:         func(r *mock_service.MockNote, userId, noteId int) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"user id not found"}`,
		},
		{
			name:                 "userId Invalid Type",
			userInfo:             notebook.UsersNote{UserId: 0},
			Id:                   "0",
			mockBehavior:         func(r *mock_service.MockNote, userId, noteId int) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"user id is of invalid type"}`,
		},
		{
			name:                 "Invalid id Param",
			userInfo:             notebook.UsersNote{UserId: 1, NotesId: 1},
			Id:                   "f",
			mockBehavior:         func(r *mock_service.MockNote, userId, noteId int) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid id param"}`,
		},
		{
			name:     "Service Error",
			userInfo: notebook.UsersNote{UserId: 1, NotesId: 1},
			Id:       "1",
			mockBehavior: func(r *mock_service.MockNote, userId, noteId int) {
				r.EXPECT().GetById(userId, noteId).Return(notebook.Note{}, errors.New("service error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"service error"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			note := mock_service.NewMockNote(c)
			testCase.mockBehavior(note, testCase.userInfo.UserId, testCase.userInfo.NotesId)

			services := &service.Service{Note: note}
			handler := NewHandler(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/api/notes/:id", func(c *gin.Context) {
				if testCase.userInfo.UserId == 0 {
					c.Set(userCtx, "some wrong info")
				} else if testCase.userInfo.UserId < 0 {
					c.Set("err", testCase.userInfo.UserId)
				} else {
					c.Set(userCtx, testCase.userInfo.UserId)
				}
			}, handler.getNoteById)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/notes/"+testCase.Id, nil)

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())

		})
	}

}

func TestHandler_updateNote(t *testing.T) {
	type mackBehavior func(r *mock_service.MockNote, userId, noteId int, input notebook.UpdateNoteInput)
	da, de := "20.03", "не покупать корм"
	testTable := []struct {
		name                 string
		userInfo             notebook.UsersNote
		Id                   string
		inputBody            string
		inputUser            notebook.UpdateNoteInput
		mockBehavior         mackBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{

		{name: "Ok",
			userInfo:  notebook.UsersNote{UserId: 1, NotesId: 1},
			Id:        "1",
			inputBody: `{"date":"20.03","description":"не покупать корм"}`,
			inputUser: notebook.UpdateNoteInput{
				Date:        &da,
				Description: &de,
			},
			mockBehavior: func(r *mock_service.MockNote, userId, noteId int, input notebook.UpdateNoteInput) {
				r.EXPECT().Update(userId, noteId, input).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"status":"ok"}`,
		},

		{name: "userId Is Missing",
			userInfo:  notebook.UsersNote{UserId: -1, NotesId: 1},
			Id:        "1",
			inputBody: `{"date":"20.03","description":"не покупать корм"}`,
			inputUser: notebook.UpdateNoteInput{
				Date:        &da,
				Description: &de,
			},
			mockBehavior:         func(r *mock_service.MockNote, userId, noteId int, input notebook.UpdateNoteInput) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"user id not found"}`,
		},
		{
			name:      "userId Invalid Type",
			userInfo:  notebook.UsersNote{UserId: 0, NotesId: 1},
			Id:        "0",
			inputBody: `{"date":"20.03","description":"не покупать корм"}`,
			inputUser: notebook.UpdateNoteInput{
				Date:        &da,
				Description: &de,
			},
			mockBehavior:         func(r *mock_service.MockNote, userId, noteId int, input notebook.UpdateNoteInput) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"user id is of invalid type"}`,
		},
		{
			name:                 "Invalid id Param",
			userInfo:             notebook.UsersNote{UserId: 1, NotesId: 1},
			Id:                   "rrr",
			mockBehavior:         func(r *mock_service.MockNote, userId, noteId int, input notebook.UpdateNoteInput) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid id param"}`,
		},
		{name: "Service Error",
			userInfo:  notebook.UsersNote{UserId: 1, NotesId: 1},
			Id:        "1",
			inputBody: `{"date":"20.03","description":"не покупать корм"}`,
			inputUser: notebook.UpdateNoteInput{
				Date:        &da,
				Description: &de,
			},
			mockBehavior: func(r *mock_service.MockNote, userId, noteId int, input notebook.UpdateNoteInput) {
				r.EXPECT().Update(userId, noteId, input).Return(errors.New("service error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"service error"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			note := mock_service.NewMockNote(c)
			testCase.mockBehavior(note, testCase.userInfo.UserId, testCase.userInfo.NotesId, testCase.inputUser)

			services := &service.Service{Note: note}
			handler := NewHandler(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.PUT("/api/notes/:id", func(c *gin.Context) {
				if testCase.userInfo.UserId == 0 {
					c.Set(userCtx, "some wrong info")
				} else if testCase.userInfo.UserId < 0 {
					c.Set("err", testCase.userInfo.UserId)
				} else {
					c.Set(userCtx, testCase.userInfo.UserId)
				}
			}, handler.updateNote)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/api/notes/"+testCase.Id, bytes.NewBufferString(testCase.inputBody))

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())

		})
	}

}
