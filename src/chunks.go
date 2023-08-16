package ms_downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func MakeChunks(f File, config Settings) []Chunk {

	count := f.Size / int64(config.ChunkSize)
	rem := f.Size % int64(config.ChunkSize)

	var data []Chunk

	start := int64(0)
	length := int64(config.ChunkSize)
	for n := int64(0); n < count; n++ {
		data = append(data, Chunk{Start: start, Length: length, Index: n})
		start += length
	}

	if rem > 0 {
		data = append(data, Chunk{Start: start, Length: rem, Index: int64(len(data))})
	}

	return data
}

// grabs a chunk from a url
func DownloadChunk(url string, info *Chunk) error {

	end := info.Start + info.Length - 1
	log.Println(fmt.Sprintf("fetching segment (%d, %d) for url %s", info.Start, end, url))
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", info.Start, end))

	log.Println(req.Header)

	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 && res.StatusCode != 206 {
		return fmt.Errorf("Error code %v received, check your url and connection", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	info.Data = body

	log.Printf("Got Chunk %v", info.Index)

	return nil
}

func WriteChunk(file *os.File, chunk Chunk) error {

	_, err := file.WriteAt(chunk.Data, chunk.Start)
	return err
}
