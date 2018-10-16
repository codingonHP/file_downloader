package filedownloader

import (
	"io"
	"os"
)

type FileOperationHandler struct { }

func (fo *FileOperationHandler) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func (fo *FileOperationHandler) Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	return io.Copy(dst, src)
}
