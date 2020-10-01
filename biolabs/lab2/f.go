

package main

import (
"fmt"
"io"
"os"
"strconv"
)

const (
	chunk       = 512
	outputShare = 100
)

const (
	start = iota
	title
	body
)

func parse(path string, num int, check func(uint8)) []string {
	file, err := os.Open(path)
	if err != nil {
		panic("couldn't open file to read (" + err.Error() + ")")
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic("couldn't close file to read (" + err.Error() + ")")
		}
	}()

	buff := make([]byte, chunk)
	res := make([]string, 1)
	state := start
	n, err := file.Read(buff)
	for err == nil {
		for i := 0; i < n; i++ {
			switch state {
			case start:
				if buff[i] == '>' {
					state = title
				}
			case title:
				if buff[i] == '\n' {
					state = body
				}
			case body:
				if buff[i] == '>' {
					if num == len(res) {
						return res
					}
					state = title
					res = append(res, "")
				} else if buff[i] != '\n' {
					res[len(res)-1] += toStr(buff[i])
				} else {
					check(buff[i])
				}
			default:
				panic("wrong state")
			}
		}
		n, err = file.Read(buff)
	}
	if err != io.EOF {
		panic("couldn't read file (" + err.Error() + ")")
	}

	if num > 0 && num != len(res) {
		panic("too little data")
	}

	return res
}

func writeSeqs(w io.Writer, seqs ...string) {
	cur := 0
	lens := make([]int, len(seqs))
	for i := range seqs {
		lens[i] = len(seqs[i])
	}
	all := max(lens[0], lens[1:]...)
	for cur < all {
		for i := range seqs {
			if _, err := fmt.Fprint(w, "seq"+strconv.Itoa(i+1)+": "); err != nil {
				panic("couldn't write (" + err.Error() + ")")
			}
			if cur < len(seqs[i]) {
				if _, err := fmt.Fprintln(w, seqs[i][cur:min(cur+outputShare, len(seqs[i]))]); err != nil {
					panic("couldn't write (" + err.Error() + ")")
				}
			}
		}
		if _, err := fmt.Fprint(w, "\n"); err != nil {
			panic("couldn't write (" + err.Error() + ")")
		}
		cur += outputShare
	}
}

func writeScore(w io.Writer, score int) {
	if _, err := fmt.Fprintln(w, "Score: "+strconv.Itoa(score)); err != nil {
		panic("couldn't write (" + err.Error() + ")")
	}
}

func toStr(c uint8) string {
	return fmt.Sprintf("%c", c)
}

func max(a int, as ...int) int {
	max := a
	for _, i := range as {
		if i > max {
			max = i
		}
	}
	return max
}

func min(a int, as ...int) int {
	min := a
	for _, i := range as {
		if i < min {
			min = i
		}
	}
	return min
}

func printMatrix(m [][]int) {
	for _, i := range m {
		for _, j := range i {
			fmt.Print(j)
			fmt.Print("\t")
		}
		fmt.Println()
	}
}

