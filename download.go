package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	//  https://www.learningcontainer.com/wp-content/uploads/2020/05/sample-mp4-file.mp4
	//  https://www.facebook.com/favicon.ico

	fileUrl := "https://www.facebook.com/favicon.ico"
	r, _ := http.NewRequest("GET", fileUrl, nil)
	err := DownloadFile(path.Base(r.URL.Path), fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("filename: " + path.Base(r.URL.Path))
	fmt.Println("Downloaded: " + fileUrl)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
