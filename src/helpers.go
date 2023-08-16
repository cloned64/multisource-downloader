package ms_downloader

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

// This should probably just be an io.MultiWriter object but I didn't know about that  until after I wrote most of the code
func Echo(msg any) {
	log.Println(msg)
	fmt.Println(msg)
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

func Hash_file_sha256(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var sha256string string
	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return sha256string, err
	}
	//Tell the program to call the following function when the current function returns
	defer file.Close()
	//Open a new hash interface to write to
	hash := sha256.New()
	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return sha256string, err
	}
	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:32]
	//Convert the bytes to a string
	sha256string = hex.EncodeToString(hashInBytes)
	return sha256string, nil
}
