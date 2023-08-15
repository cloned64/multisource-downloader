package ms_downloader

type File struct {
	Filename string
	Size     int64
	Etag     string
}

type Chunk struct {
	Start    int64
	Length   int64
	Index    int64
	Data     []byte
	Complete bool
	Trys     int
}

type Settings struct {
	Retries    int
	ChunkSize  int
	MaxWorkers int
	OutputPath string
	Resume     bool
}
