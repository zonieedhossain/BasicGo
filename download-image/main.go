package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName    string
	fullUrlFile string
)

func main() {

	fullUrlFile = "https://maps.googleapis.com/maps/api/staticmap?center=23.770474%2C90.362925&key=AIzaSyCeLaDxdBF1ipP-rADJZhitd15MOZoC-1I%0A%0A&maptype=roadmap&markers=color%3Agreen%7c23.770474%2C90.362925&size=2480x720&zoom=13"

	// Build fileName from fullPath
	buildFileName()

	// Create blank file
	file := createFile()

	// Put content on file
	putFile(file, httpClient())

}

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullUrlFile)

	checkError(err)

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	fmt.Println(size)
	defer file.Close()

	checkError(err)

	fmt.Printf("Just Downloaded a file %s with size %d", fileName, size)
}

func buildFileName() {
	fileUrl, err := url.Parse(fullUrlFile)
	checkError(err)

	path := fileUrl.Path
	segments := strings.Split(path, "/")

	fileName = segments[len(segments)-1]
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func createFile() *os.File {
	file, err := os.Create(fileName)

	checkError(err)
	return file
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
