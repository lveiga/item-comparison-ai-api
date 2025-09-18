package database

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFileStore is a mock implementation of the FileStore interface for testing purposes

type MockFileStore struct {
	ReadFunc          func(filename string) ([]byte, error)
	WriteFunc         func(filename string, data []byte, perm os.FileMode) error
	CheckLivenessFunc func(filename string) error
}

func (m *MockFileStore) Read(filename string) ([]byte, error) {
	if m.ReadFunc != nil {
		return m.ReadFunc(filename)
	}
	return nil, errors.New("ReadFunc not implemented")
}

func (m *MockFileStore) Write(filename string, data []byte, perm os.FileMode) error {
	if m.WriteFunc != nil {
		return m.WriteFunc(filename, data, perm)
	}
	return errors.New("WriteFunc not implemented")
}

func (m *MockFileStore) CheckLiveness(filename string) error {
	if m.CheckLivenessFunc != nil {
		return m.CheckLivenessFunc(filename)
	}
	return nil
}

func TestNewClient(t *testing.T) {
	mockFileStore := &MockFileStore{}
	client := NewClient(mockFileStore)
	assert.NotNil(t, client)
}

func TestOSFileStore_Read(t *testing.T) {
	// Create a temporary directory for test files
	testDir, err := os.MkdirTemp("", "db_test")
	assert.NoError(t, err)
	defer os.RemoveAll(testDir) // Clean up the directory after tests

	// Create a temporary file with some content
	tempFile, err := os.CreateTemp(testDir, "test.txt")
	assert.NoError(t, err)
	defer tempFile.Close()

	content := []byte("hello world")
	_, err = tempFile.Write(content)
	assert.NoError(t, err)

	// Use an Database to read the file
	fs := &Database{}
	readContent, err := fs.Read(tempFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, content, readContent)
}

func TestOSFileStore_Write(t *testing.T) {
	// Create a temporary directory for test files
	testDir, err := os.MkdirTemp("", "db_test")
	assert.NoError(t, err)
	defer os.RemoveAll(testDir) // Clean up the directory after tests

	// Create a temporary file
	tempFile, err := os.CreateTemp(testDir, "test.txt")
	assert.NoError(t, err)
	defer tempFile.Close()

	// Use an Database to write to the file
	fs := &Database{}
	content := []byte("hello world")
	err = fs.Write(tempFile.Name(), content, 0644)
	assert.NoError(t, err)

	// Read the file and assert that the content is correct
	readContent, err := os.ReadFile(tempFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, content, readContent)
}
