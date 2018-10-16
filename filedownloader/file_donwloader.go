package filedownloader

type IFileDownloader interface {
	DownloadFile(filepath string, url string) (err error)
}