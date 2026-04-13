package internal

import (
	"fmt"

	tiktoken "github.com/pkoukk/tiktoken-go"
)

// this file contains the code for splitting the text into chunks based on the token count, and the overlap between chunks. 
// The main function is TokenChunker, which takes in the content, chunk size, overlap, and model name, and returns a slice of chunks.
// arguments:
// - content: the text to be split into chunks
// - chunkSize: the maximum number of tokens in each chunk
// - overlap: the number of tokens that should overlap between consecutive chunks
// - model: the name of the model to determine the tokenizer to use
// returns:
// - a slice of strings, where each string is a chunk of the original content, and an error if any occurs during the tokenization process
// The function works as follows:
// the function first checks if the content is empty, and if so, returns an empty slice. 
// It then checks if the chunk size is greater than the overlap, and if not, returns an error. 
// Next, it loads the tokenizer for the specified model and encodes the content into tokens. 
// If there are no tokens, it returns an empty slice. 
// Finally, it iterates through the tokens and creates chunks based on the specified chunk size and overlap, and returns the resulting chunks.

func TokenChunker(content string, chunkSize int, overlap int, model string) ([]string, error) {
	if content == "" {
		return []string{}, nil
	}

	if chunkSize <= overlap {
		return nil, fmt.Errorf("chunkSize must be > overlap")
	}

	enc, err := tiktoken.EncodingForModel(model)
	if err != nil {
		return nil, fmt.Errorf("failed to load tokenizer: %w", err)
	}

	tokens := enc.Encode(content, nil, nil)

	if len(tokens) == 0 {
		return []string{}, nil
	}

	step := chunkSize - overlap
	chunks := make([]string, 0, len(tokens)/step+1)

	for i := 0; i < len(tokens); {
		end := i + chunkSize

		if end > len(tokens) {
			end = len(tokens)
			i = end - chunkSize
			if i < 0 {
				i = 0
			}
		}

		chunkTokens := tokens[i:end]
		chunkText := enc.Decode(chunkTokens)

		chunks = append(chunks, chunkText)

		if end == len(tokens) {
			break
		}

		i += step
	}

	return chunks, nil
}
