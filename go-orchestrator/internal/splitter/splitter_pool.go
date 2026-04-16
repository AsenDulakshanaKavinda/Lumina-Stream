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
				// 1. If the context is done, should stop the worker and return
				case <-ctx.Done():
					return
				
				// 2. If receive a job from the fileJobs channel, 
				// should read the file content, split the content into chunks, and send the chunks to the chunkJobs channel
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








