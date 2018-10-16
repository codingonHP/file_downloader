package main

import (
	"FileDownloader/filedownloader"
	"fmt"
)

func main() {
  fd := filedownloader.FileDownloader{}
  err := fd.DownloadFile("./index.html", "https://www.google.com")
	if err != nil {
		fmt.Println(err)
		return
	}

  fmt.Println("file downloaded successfully")
}