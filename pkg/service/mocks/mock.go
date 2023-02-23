package mock_service

import (
	"github.com/Smolvika/notebook.git"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

func (m *MockAuthorization) CreateUser(user notebook.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

type MockNote struct {
	ctrl     *gomock.Controller
	recorder *MockNoteMockRecorder
}

type MockNoteMockRecorder struct {
	mock *MockNote
}

func NewMockNote(ctrl *gomock.Controller) *MockNote {
	mock := &MockNote{ctrl: ctrl}
	mock.recorder = &MockNoteMockRecorder{mock}
	return mock
}

func (m *MockNote) EXPECT() *MockNoteMockRecorder {
	return m.recorder
}

func (m *MockNote) Create(userId int, note notebook.Note) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId, note)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNoteMockRecorder) Create(userId, note interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNote)(nil).Create), userId, note)
}

func (m *MockNote) GetAll(userId int) ([]notebook.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userId)
	ret0, _ := ret[0].([]notebook.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNoteMockRecorder) GetAll(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockNote)(nil).GetAll), userId)
}

func (m *MockNote) GetById(userId, noteId int) (notebook.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, noteId)
	ret0, _ := ret[0].(notebook.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNoteMockRecorder) GetById(userId, noteId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockNote)(nil).GetById), userId, noteId)
}

func (m *MockNote) Delete(userId, noteId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, noteId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockNoteMockRecorder) Delete(userId, noteId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNote)(nil).Delete), userId, noteId)
}

func (m *MockNote) Update(userId, noteId int, input notebook.UpdateNoteInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, noteId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockNoteMockRecorder) Update(userId, noteId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNote)(nil).Update), userId, noteId, input)
}
