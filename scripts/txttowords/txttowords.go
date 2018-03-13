package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := ""
	fmt.Scanln(&filename)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wf, err := os.Create("words.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := bufio.NewScanner(f)
	w := bufio.NewWriter(wf)
	n := 0
	for s.Scan() {
		t := s.Text()
		i := strings.Index(t, "=")
		if i > -1 {
			n++
			fmt.Fprintf(w, "%s=", t[:i-1])
		}
	}
	w.Flush()
	fmt.Printf("%d entries written\n", n)
}
