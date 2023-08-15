package main

import (
	"fmt"
	dl "ms_downloader/src"
	"time"
	
)

// type FileInfo struct {
// 	filename string
// 	size int64
// }

func main() {

	fmt.Println("Hello, 世界")
	url_ := "https://releases.ubuntu.com/22.04.3/ubuntu-22.04.3-desktop-amd64.iso"
	// url_ := "https://mega.nz/linux/repo/xUbuntu_23.04/amd64/megasync-xUbuntu_23.04_amd64.deb"

	fucker, err := dl.GetHeader(url_)
	fmt.Println(fucker)
	fmt.Println(err)

	// testChunk := dl.Chunk{Start: 0, Length: 1024, Index: 0}

	// chunk, err := dl.DownloadChunk(url_, testChunk)
	// fmt.Println(err)
	// fmt.Print(chunk.Data)
	// fmt.Println(chunk)


	s_ := dl.Settings{
		Retries: 10,
		ChunkSize: 1024*1024,
		MaxWorkers: 100,
		OutputPath: "output" }

	urls := make([]string, 1)
	urls[0] = url_
	
	start := time.Now()
	filePath, err := dl.Runner(urls, s_)
	
	if err != nil {
		fmt.Println(err)
		fmt.Println("File could not be downloaded")
	}

	duration := time.Since(start)
	fmt.Printf("Download Time: %v", duration)

	hash, err := dl.Hash_file_md5(filePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("MD5 Hash sould not be calculated")
	}
	fmt.Printf("MD5 hash: \n\r\t%v", hash)

}
