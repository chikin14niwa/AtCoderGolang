package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Io struct {
	reader    *bufio.Reader
	writer    *bufio.Writer
	tokens    []string
	nextToken int
}

func NewIO() *Io {
	return &Io{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (io *Io) Flush() {
	err := io.writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *Io) NextLine() string {
	var buffer []byte
	for {
		line, isPrefix, err := io.reader.ReadLine()
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, line...)
		if !isPrefix {
			break
		}
	}
	return string(buffer)
}

func (io *Io) Next() string {
	for io.nextToken >= len(io.tokens) {
		line := io.NextLine()
		io.tokens = strings.Fields(line)
		io.nextToken = 0
	}
	r := io.tokens[io.nextToken]
	io.nextToken++
	return r
}

func (io *Io) NextInt() int {
	i, err := strconv.Atoi(io.Next())
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) NextFloat() float64 {
	i, err := strconv.ParseFloat(io.Next(), 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) PrintLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *Io) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.writer, format, a...)
}

func (io *Io) PrintIntLn(a []int) {
	b := []interface{}{}
	for _, x := range a {
		b = append(b, x)
	}
	io.PrintLn(b...)
}

func (io *Io) PrintStringLn(a []string) {
	b := []interface{}{}
	for _, x := range a {
		b = append(b, x)
	}
	io.PrintLn(b...)
}

type Point struct {
	P int
	C int
}

func main() {
	io := NewIO()
	defer io.Flush()

	N := io.NextInt()
	points := make([]Point, 2*N)
	for i := 0; i < N; i++ {
		A, B := io.NextInt(), io.NextInt()
		points[2*i] = Point{P: A, C: 1}
		points[2*i+1] = Point{P: A + B, C: -1}
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].P < points[j].P
	})
	ans := make([]int, N+1)
	var cnt int
	for i := 0; i < len(points)-1; i++ {
		cnt += points[i].C
		ans[cnt] += points[i+1].P - points[i].P
	}
	for i := 1; i <= N; i++ {
		io.Printf("%d ", ans[i])
	}
	io.PrintLn()
}
