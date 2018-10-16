package filedownloader

import "os"
import "io"

//IFileOperationHandler Create file for a given name and path
type IFileOperationHandler interface {
	Create(name string) (*os.File, error)
    Copy(dst io.Writer, src io.Reader) (written int64, err error)
}
