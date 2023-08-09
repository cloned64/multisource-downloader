package main

import (
	"fmt"
)

import (
	dl "ms_downloader/src"
)

// type FileInfo struct {
// 	filename string
// 	size int64
// }


func main() {

	fmt.Println("Hello, 世界")
	url_ := "https://releases.ubuntu.com/22.04.2/ubuntu-22.04.2-desktop-amd64.iso"


	fucker, err := dl.GetHeader(url_)
	fmt.Println(fucker)
	fmt.Println(err)

	testChunk := dl.ChunkInfo{Start: 0, Length: 1024, Index: 0}

	chunk, err := dl.GetFileRange(url_, testChunk)
	fmt.Println(err)
	fmt.Print(chunk.Data)

}
