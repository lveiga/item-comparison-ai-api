package database

import (
	"os"
)

// FileStore defines the interface for file operations
type FileStore interface {
	Read(filename string) ([]byte, error)
	Write(filename string, data []byte, perm os.FileMode) error
	CheckLiveness(filename string) error
}

func NewClient(fileStore FileStore) *Database {
	return &Database{}
}

// Database is an implementation of FileStore that uses os package for file operations
type Database struct{}

func (fs *Database) Read(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (fs *Database) Write(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

func (conn *Database) CheckLiveness(filename string) error {
	_, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return nil
}
