package main

import (
	"bufio"
	//"fmt"
	"log"
	"os"
)

func main() {
	err := solve()
	if err != nil {
		log.Fatalln(err)
	}
}

func solve() error {
	w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	var s string
	for scanner.Scan() {
		s = scanner.Text()
		w.WriteString(s)
		w.WriteByte('\n')
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}
