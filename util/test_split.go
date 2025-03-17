package util

func TextSplitWithOverlaps(in string, chunkSize, overlapSize int) []string {

	var chunks []string
	l := len(in)

	// Edge case when the input text length is exactly the chunk size
	if l <= chunkSize {
		return []string{in}
	}

	for i := 0; i < l; {
		// Ensure we don't go beyond text length
		end := i + chunkSize
		if end > l {
			end = l
		}

		// Extract chunk
		chunk := in[i:end]
		if i == 0 || len(chunk) > overlapSize {
			chunks = append(chunks, chunk)
		}

		// Move index forward, keeping the overlap
		i += chunkSize - overlapSize

		// Avoid adding a chunk if we've reached the end of the text
		if i >= l {
			break
		}
	}

	return chunks
}
