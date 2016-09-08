package main

import (
	"bufio"
	"fmt"
	"github.com/yanyiwu/gojieba"
	"io"
	"os"
	"strings"
)

func main() {
	x := gojieba.NewJieba()
	defer x.Free()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		s, err := inputReader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		s = strings.TrimSpace(s)

		if s != "" {
			words := x.CutForSearch(s, true)
			fmt.Println(strings.Join(words, " "))
		} else {
			fmt.Println("get empty \n")
		}
	}
}
