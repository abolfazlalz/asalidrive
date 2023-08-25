package storage

type IStorage interface {
  ShouldPath(path ...string) string
  Path(path ...string) (string, error)
  Copy(src string, dst string) error
  Move(src string, dst string) error
  Write(fileName string, b []byte) error
  Read(fileName string) ([]byte, error)
}
