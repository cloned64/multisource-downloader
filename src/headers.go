package ms_downloader

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
)

func onError(err error) (File, error) {
	return File{}, err
}

func checkHeaderValue(header http.Header, key string) (string, error) {
	value, ok := header[key]
	if ok {
		fmt.Printf("%v = %v\n", key, value[0])
		return value[0], nil
	} else {
		return "", fmt.Errorf("%v not present", key)
	}
}

func GetHeader(url string) (File, error) {

	fmt.Println("fetching headers for url")
	req, err := http.NewRequest("HEAD", url, nil)

	if err != nil {
		return onError(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return onError(err)
	}
	defer res.Body.Close()

	// check status codes
	if res.StatusCode != 200 {
		return onError(fmt.Errorf("Error code %v received, invalid target url", res.StatusCode))
	}

	// checking for required header values
	fileSizeStr, err := checkHeaderValue(res.Header, "Content-Length")
	if err != nil {
		return onError(err)
	}

	fileSize, err := strconv.ParseInt(fileSizeStr, 0, 64)
	if err != nil {
		return onError(err)
	}

	acceptRange, err := checkHeaderValue(res.Header, "Accept-Ranges")
	if err != nil {
		return onError(err)
	}
	if acceptRange != "bytes" {
		return onError(fmt.Errorf("Accept-Ranges not bytes"))
	}

	etag, err := checkHeaderValue(res.Header, "Etag")
	if err != nil {
		fmt.Println("ETag not found, validation will not be possible")
	}

	filename := path.Base(req.URL.Path)
	fmt.Println(filename)

	return File{Size: fileSize, Filename: filename, Etag: etag}, nil

}
