package mock

import (
	mockFile "github.com/stretchr/testify/mock"
)

type File struct {
	mockFile.Mock
}

func (mf *File) Read(p []byte) (n int, err error) {
	args := mf.Called(p)
	return args.Int(0), args.Error(1)
}

func (mf *File) Write(p []byte) (n int, err error) {
	args := mf.Called(p)
	return args.Int(0), args.Error(1)
}

func (mf *File) Close() error {
	return mf.Called().Error(0)
}
