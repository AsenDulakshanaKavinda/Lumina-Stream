package internal

import (
	"context"
	pipeline "orchestrator/internal/schemas"
	"orchestrator/utils"
	"os"
)



func StartSplitterPool(
	ctx context.Context,
	numWorkers int,
	fileJobs <-chan pipeline.FileJob,
	chunkJobs chan<- pipeline.ChunkJob,
) {
	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			for {
				select {
				case <-ctx.Done():
					return
				
				case job, ok := <-fileJobs:
					if !ok {
						return
					}

					data, err := os.ReadFile(job.Path)
					if err != nil {
						utils.Log.Error().Err(err).Msg("Error while reading file content.")
						continue
					}

					chunks, err := TokenChunker(string(data), 100, 20, "model") // todo - update using config
					if err != nil {
						utils.Log.Error().Err(err).Msg("Error while splitting the content.")
						continue
					}

					for i, chunk := range chunks {
						select {
							case chunkJobs <- pipeline.ChunkJob{
								FilePath: job.Path,
								ChunkID: i,
								Content: chunk,
							}:
							case <-ctx.Done():
								return
						}
					}
				}
			}
		}(i)
	}
}








