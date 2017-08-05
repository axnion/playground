package main

import (
	"errors"
	"testing"

	git "gopkg.in/src-d/go-git.v4"

	"github.com/stretchr/testify/mock"
)

/*
----------MOCKS-------------------------------------------------------------------------------------
*/

type mockObj struct {
	mock.Mock
}

func (m *mockObj) Output() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (m *mockObj) Chdir(dir string) error {
	args := m.Called(dir)
	return args.Error(0)
}

func (m *mockObj) PlainClone(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
	args := m.Called(path, isBare, o)
	return args.Get(0).(*git.Repository), args.Error(1)
}

/*
---------TESTS--------------------------------------------------------------------------------------
*/

func TestExecuteCommandSuccessful(t *testing.T) {
	mock := new(mockObj)
	stacktrace := []byte("Pass")
	mock.On("Output").Return(stacktrace, nil)

	executeCommand(mock)

	mock.AssertExpectations(t)
}

func TestExecuteCommandFailure(t *testing.T) {
	mock := new(mockObj)
	stacktrace := []byte("Fail")
	err := errors.New("Something failed")
	mock.On("Output").Return(stacktrace, err)

	executeCommand(mock)

	mock.AssertExpectations(t)
}

func TestGetRepository(t *testing.T) {
	payload := new(payload)
	payload.After = "11111"
	payload.Repository.Clone_url = "https://bla.git"
	payload.Repository.Ssh_url = "https://bla.git"

	mock := new(mockObj)

	mock.On("PlainClone").Return(nil, errors.New("OMG!"))
	mock.On("Chdir").Return(errors.New("OMG!"))

	getRepository(payload, mock)

	mock.AssertExpectations(t)
}
