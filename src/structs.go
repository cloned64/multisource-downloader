package ms_downloader


type FileConfig struct {
	Filename string
	Size int64
	Etag string
}

type ChunkInfo struct {
	Start int64
	Length int64
	Index int64
}

type Chunk struct {
	Info ChunkInfo
	Data []byte
}

type Settings struct {
	Retries int
	ChunkSize int
}
