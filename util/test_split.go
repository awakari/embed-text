package util

func TextSplitWithOverlaps(in string, chunkSize, overlapSize int) []string {

	var chunks []string
	l := len(in)

	for i := 0; i < l; {
		// Ensure we don't go beyond text length
		end := i + chunkSize
		if end > l {
			end = l
		}

		// Extract chunk
		chunk := in[i:end]
		chunks = append(chunks, chunk)

		// Move index forward, keeping the overlap
		i += chunkSize - overlapSize
	}

	return chunks
}
