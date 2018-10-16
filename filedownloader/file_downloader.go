package filedownloader

import (
	"fmt"
	"net/http"
)

type FileDownloader struct {
	httpClient IHttpClient
	fileHandler IFileOperationHandler
}

func (fd FileDownloader) DownloadFile(filepath string, url string) (err error) {

	if fd.httpClient == nil {
		fd.httpClient = http.DefaultClient
	}

	if fd.fileHandler == nil {
		fd.fileHandler = &FileOperationHandler{}
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("AUTH_TOKEN", "My AuthToken Key")

	// Create the file
	out, err := fd.fileHandler.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, httpErr := fd.httpClient.Do(req)

	if httpErr != nil {
		return httpErr
	}

	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = fd.fileHandler.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
