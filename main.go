package main

import (
	"fmt"
	"log"
	dl "ms_downloader/src"
	"os"
	"time"

	"github.com/akamensky/argparse"
)

func main() {

	logfile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logfile)

	settings, url_ := Args()
	urls := make([]string, 1)
	urls[0] = url_

	start := time.Now()
	filePath, err := dl.Runner(urls, settings)

	if err != nil {
		dl.Echo(err)
		dl.Echo("File could not be downloaded")
	}

	duration := time.Since(start)
	fmt.Printf("Download Time: %v\n", duration)

	hash, err := dl.Hash_file_md5(filePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("MD5 Hash could not be calculated")
	}
	fmt.Printf("MD5 hash: \n\r\t%v\n", hash)

	hash_sha, err := dl.Hash_file_sha256(filePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("SHA256 Hash could not be calculated")
	}
	fmt.Printf("SHA256 hash: \n\r\t%v\n", hash_sha)

}

func Args() (dl.Settings, string) {
	parser := argparse.NewParser("file downloader", "Downloads a file from multiple sources")

	Retries := parser.Int("r", "retries",
		&argparse.Options{Required: false, Help: "Number of retries before giving up", Default: 10})
	ChunkSize := parser.Int("c", "chunk-size",
		&argparse.Options{Required: false, Help: "Chunk size in bytes", Default: 1024 * 1024})
	MaxWorkers := parser.Int("t", "threads",
		&argparse.Options{Required: false, Help: "Number of threads to use", Default: 100})
	OutputPath := parser.String("o", "output",
		&argparse.Options{Required: false, Help: "Folder to save file", Default: "output"})

	fmt.Println(os.Args)

	filename := parser.String("f", "filename",
		&argparse.Options{Required: false, Help: "Folder to save file", Default: "https://www.python.org/ftp/python/3.11.4/Python-3.11.4.tgz"})
	// filename := parser.String("f", "filename",
	// 	&argparse.Options{Required: false, Help: "Folder to save file", Default: "https://releases.ubuntu.com/22.04.3/ubuntu-22.04.3-desktop-amd64.iso"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	var s_ = dl.Settings{
		Retries:    *Retries,
		ChunkSize:  *ChunkSize,
		MaxWorkers: *MaxWorkers,
		OutputPath: *OutputPath,
	}

	return s_, *filename
}
