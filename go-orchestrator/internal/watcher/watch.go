package watcher

import (
	"context"
	pipeline "orchestrator/internal/schemas"
	utils "orchestrator/utils"

	"github.com/fsnotify/fsnotify"
)

func WatchFiles(
	ctx context.Context, 
	filepath string, 
	fileJobs chan<- pipeline.FileJob, // send
) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		utils.Log.Error().Err(err).Msg("Error while creating watcher")
	}
	defer watcher.Close()

	if err := watcher.Add(filepath); err != nil {
		utils.Log.Err(err).Msg("Error while watching the filepath")
	}

	for {
		select {
		// 1. If the context is done, stop the watcher and return
		case <-ctx.Done():
			utils.Log.Info().Msg("Watcher shutting down")
			return

		// 2. If receive an event, should check if it's a write or create event, 
		// and if so, should send a job to the fileJobs channel
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op & fsnotify.Write == fsnotify.Write ||
				event.Op & fsnotify.Create == fsnotify.Create {
					select {
					// send a job to the fileJobs channel
					case fileJobs <- pipeline.FileJob{Path: event.Name}:
						utils.Log.Info().Str("file", event.Name).Msg("job dispatched")

					case <-ctx.Done():
						return
					}
				}

		// 3. If receive an error, should log the error
		case err := <-watcher.Errors:
			utils.Log.Error().Err(err).Msg("watcher error")
		}
	}
}