package ms_downloader

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"github.com/schollz/progressbar/v3"
)

func DownloadRoutine(inputChan chan Chunk, outputChan chan Chunk, retryChan chan Chunk, url string) {
	for chuInfo := range inputChan {
		err := DownloadChunk(url, &chuInfo)
		if err != nil {
			retryChan <- chuInfo
			chuInfo.Trys++
		} else {
			outputChan <- chuInfo
		}
	}
}

// This will retry a chunk if it fails to download or write to file
// Can be expanded to limit the number of retrys or to switch source in the future
func Retryer(retryChan chan Chunk, downloadChan chan Chunk) {
	for chu := range retryChan {
		downloadChan <- chu
	}

}


func ChunkWriter(chunkChannel chan Chunk, retryChan chan Chunk, done *sync.WaitGroup, bar *progressbar.ProgressBar,file *os.File) {

	for chu := range chunkChannel {
		err := WriteChunk(file, chu)
		if err != nil {
			retryChan <- chu
		} else {
			chu.Complete = true
			bar.Add(1)
			done.Done()
		}
	}
}


func Runner(urls []string, settings Settings) (string, error) {
	if urls == nil || len(urls) < 1 {
		return "", fmt.Errorf("Source urls not present")
	}

	// get file info for the first url
	fileDesc, err := GetHeader(urls[0])
	if err != nil {
		return "", err
	}

	chunkMeta := MakeChunks(fileDesc, settings)

	log.Println("CHUNKS")

	// create channels
	size := len(chunkMeta)
	downloadChan := make(chan Chunk, size)
	writingChan := make(chan Chunk, size)
	retryChan := make(chan Chunk, size)

	// setup the waitgroup
	var done sync.WaitGroup
	done.Add(len(chunkMeta))

	//create writer goroutine
	filePath := fmt.Sprintf("%s/%s", settings.OutputPath, fileDesc.Filename)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	bar := progressbar.Default(int64(len(chunkMeta)), "downloading")
	// bar := progressbar.Default(fileDesc.Size, "downloading")

	go Retryer(retryChan, downloadChan)
	go ChunkWriter(writingChan, retryChan, &done, bar, file)

	// create downloader goroutines
	for n := 0; n < settings.MaxWorkers; n++ {
		go DownloadRoutine(downloadChan, writingChan, retryChan, urls[0])

	}

	// put chunks onto channel to download
	for _, chunk := range chunkMeta {
		downloadChan <- chunk
	}
	// wait for progress checker to say it it done

	done.Wait()
	Echo("File is done!")

	return filePath, nil
}

// This is a functgion i found to calculate md5 hash
// https://mrwaggel.be/post/generate-md5-hash-of-a-file-in-golang
func Hash_file_md5(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string
	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	//Tell the program to call the following function when the current function returns
	defer file.Close()
	//Open a new hash interface to write to
	hash := md5.New()
	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}
