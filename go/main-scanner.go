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
	defer w.Flush()

	scanner := bufio.NewScanner(os.Stdin)

	var b []byte
	for scanner.Scan() {
		b = scanner.Bytes()
		w.Write(b)
		w.WriteByte('\n')
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
