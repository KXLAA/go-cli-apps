package main

import (
	//read text
	"flag"
	"fmt" //print formatted output
	"os"  //os system resources

	counter "github.com/KXLAA/go-cli-apps/word-counter"
)

func main() {
	// CLI flags
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count Bytes")

	//Parse the flags provided by the user
	flag.Parse()

	//Word count function
	fmt.Println(counter.Count(os.Stdin, *lines, *bytes))
}
