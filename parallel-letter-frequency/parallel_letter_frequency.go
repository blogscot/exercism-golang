package letter

// ConcurrentFrequency calculates the letter frequency of the given
// texts concurrently.
func ConcurrentFrequency(texts []string) FreqMap {
	freqChan := make(chan FreqMap)

	for _, text := range texts {
		go func(s string) { freqChan <- Frequency(s) }(text)
	}

	freqMap := FreqMap{}

	for range texts {
		for k, v := range <-freqChan {
			freqMap[k] += v
		}
	}
	return freqMap
}
