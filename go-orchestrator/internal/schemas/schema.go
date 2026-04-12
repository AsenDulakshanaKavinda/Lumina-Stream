package pipeline

type FileJob struct {
    Path string
}

type ChunkJob struct {
    FilePath string
    ChunkID  int
    Content  string
}