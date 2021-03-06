package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(test int) {
	// implement my answer
	out(fmt.Sprintf("Case #%v: my answer\n", test))
}

func main() {
	fin, _ := os.OpenFile("input.txt", os.O_RDONLY, 0644)
	scanner = bufio.NewScanner(fin)
	scanner.Split(bufio.ScanWords)
	defer fin.Close()

	fout, _ := os.Create("output.txt")
	writer = bufio.NewWriter(fout)
	defer fout.Close()

	testNum := nextInt()
	for test := 1; test <= testNum; test++ {
		solve(test)
	}
}

// input
var scanner *bufio.Scanner

func next() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	scanner.Scan()
	ret, _ := strconv.Atoi(scanner.Text())
	return ret
}

func nextInt64() int64 {
	scanner.Scan()
	ret, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return ret
}

var writer *bufio.Writer

// output
func out(str string) {
	fmt.Print(str)
	fmt.Fprint(writer, str)
	writer.Flush()
}
