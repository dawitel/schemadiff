package storage

import (
	"os"
)

type FileStorage struct{}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (s *FileStorage) SaveMigration(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func (s *FileStorage) EnsureDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}