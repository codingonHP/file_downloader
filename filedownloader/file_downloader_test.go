package filedownloader

import (
	"io"
	"net/http"
	"os"
	"testing"
)

type FakeFileOperationHandler struct { }
func (fo *FakeFileOperationHandler) Create(name string) (*os.File, error) {
	return &os.File{}, nil
}

func (fo *FakeFileOperationHandler) Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	return 0, nil
}

type FakeReadCloser struct {}
func (fr FakeReadCloser) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (fr FakeReadCloser) Close() error {
	return nil
}

type ClientMock struct {}
func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	if req.Header.Get("AUTH_TOKEN") == "" {
		panic("AUTH_HEADER MISSING")
	}
	return &http.Response{ StatusCode:http.StatusOK, Body: FakeReadCloser{} }, nil
}

func TestFileDownloader_DownloadFile(t *testing.T) {
	// Arrange
	fakeReader := &FakeFileOperationHandler{}
	fakeHttp := &ClientMock{}
	testFileDownloader := FileDownloader{
		fileHandler:fakeReader,
		httpClient:fakeHttp,
	}

	//Act
	err := testFileDownloader.DownloadFile("temp-path", "http://fakeurl.com")

	//Assert
	if err != nil {
		t.Fail()
	}
}
