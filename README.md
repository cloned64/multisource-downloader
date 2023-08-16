# Multisource Downloader #

This is a simple tool to download large files concurrently

## How to run ##

Tested on Go 1.21 on a Mint Linux computer

You should be able to use the following command from the root folder
    go run main.go --url https://go.dev/dl/go1.21.0.linux-amd64.tar.gz


## Usage ##
```
usage: file downloader [-h|--help] [-r|--retries <integer>] [-c|--chunk-size
                       <integer>] [-t|--threads <integer>] [-o|--output
                       "<value>"] -f|--filename "<value>"

                       Downloads a file from multiple sources

Arguments:

  -h  --help        Print help information
  -r  --retries     Not implemented but, the number of retries before giving up. Default: 10
  -c  --chunk-size  Chunk size in bytes. Default: 1048576
  -t  --threads     Number of threads to use. Default: 100
  -o  --output      Folder to save file. Default: output
  -f  --filename    Folder to save file
```
## Some urls to try ##

* Recent Ubuntu release
    * 4.7 Gb
    * https://releases.ubuntu.com/22.04.3/ubuntu-22.04.3-desktop-amd64.iso
    * SHA256: a435f6f393dda581172490eda9f683c32e495158a780b5a1de422ee77d98e909
* Python 3.11.4
    * 26 Mb
    * https://www.python.org/ftp/python/3.11.4/Python-3.11.4.tgz
    * MD5: bf6ec50f2f3bfa6ffbdb385286f2c628
* Go 1.21.0 for Linux 2.6.32 or later
    * 63 Mb
    * https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
    * SHA256: d0398903a16ba2232b389fb31032ddf57cac34efda306a0eebac34f0965a0742

## Notes to improve ##
1. To make it truly multisource add a descriptor file or an arguement that takes an list of urls
2. Once you have a list of urls assign some of downloader routines to each
3. Utilize the Chunk.Trys count and Settings.Retries to switch downloader routines to new urls
4. the temp folder and Settings.Resume were intended to added the ability to resume a download if interrupted
5. Would have liked to give the program a hash to verify the download automatically
6. Read settings from file sort of like a torrent file
7. Add some unit tests, particularly for generating the file descriptor and chunks