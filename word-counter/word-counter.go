package wordcounter

import (
	"bufio"
	"io"
)

func Count(input io.Reader, countLines, countBytes bool) int {
	//Read text from a reader (such as files)

	//Scanner is a convenient way of reading data delimitated by newlines or spaces
	scanner := bufio.NewScanner(input)

	//If the count lines flag is not set, we ant to count words so we define
	// the scanner to split type to words (default is split by lines)
	if !countLines {
		//By default a scanner reads lines of data,
		//but we want the scanner to read words instead
		//So we set the the split function of the scanner to scan words
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	//Variable to hold the count of the words
	wordCount := 0

	//Loop through each word or "token" with Scan and increment the counter
	for scanner.Scan() {
		wordCount++
	}

	//Return the total
	return wordCount

}
