package storage

import (
  "path/filepath"
  "os"
)

type LocalStorage struct {
  path string
}

func (ls LocalStorage) Path(path ...string) (string, error) {
  absPath, err := filepath.Abs(ls.path)
  if err != nil {
    return "", err
  }
  if len(path) > 0 {
    absPath = absPath + "/" + path[0]
  }
  return absPath, nil
}

func (ls LocalStorage) ShouldPath(path ...string) string {
  absPath, _ := ls.Path(path...)
  return absPath
}

func (ls LocalStorage) Copy(src string, dst string) error {
  data, err := ls.Read(src)
  if err != nil {
    return err
  }
  return ls.Write(dst, data)
}

func (ls LocalStorage) Read(path string) ([]byte, error) {
  filePath := ls.ShouldPath(path)
  return os.ReadFile(filePath)
}

func (ls LocalStorage) Move(src string, dst string) error {
  return os.Rename(ls.ShouldPath(src), ls.ShouldPath(dst))
}

func (ls LocalStorage) Write(fileName string, b []byte) error {
  return os.WriteFile(ls.ShouldPath(fileName), b, 0644)
}

func (ls LocalStorage) Delete(fileName string) error {
  return os.Remove(ls.ShouldPath(fileName))
}

func NewLocal() IStorage {
  return &LocalStorage {
    path: "uploads",
  }
}

