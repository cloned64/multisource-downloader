package ms_downloader

import (
	"net/http"
	"fmt"
	"io"
)


func GetFileRange(url string, info ChunkInfo) (Chunk, error){

	end := info.Start + info.Length - 1
	fmt.Println(fmt.Sprintf("fetching segment (%d, %d) for url %s", info.Start, end, url))
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", info.Start , end))

	fmt.Println(req.Header)

	if err != nil {
		return Chunk{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Chunk{}, err
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	if res.StatusCode != 200 && res.StatusCode != 206{
		return Chunk{}, fmt.Errorf("Error code %v received, check your url and connection", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Chunk{}, err
	}
	

	fmt.Println(body)

	return Chunk{Info: info, Data: body}, nil
}