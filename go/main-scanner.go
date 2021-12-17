package main

import (
	"bufio"
	//"fmt"
	"log"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	var s string
	for scanner.Scan() {
		s = scanner.Text()
		w.Write([]byte(s))
		w.Write([]byte("\n"))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	w.Flush()
}
