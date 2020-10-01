package main

import (
	"flag"
	"io"
	"os"
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	var gapArg int
	var input, output, scoring string
	flag.StringVar(&input, "i", "", "input files")
	flag.IntVar(&gapArg, "gap", 0, "custom gap")
	flag.IntVar(&gapArg, "g", 0, "custom gap")
	flag.StringVar(&output, "o", "", "output file")
	flag.StringVar(&scoring, "s", "default", "scoring (blossum|dna|default)")
	flag.Parse()

	var check func(uint8)
	var scoreF func(uint8, uint8) int
	var gap int

	switch scoring {
	case "blossum":
		check = blossumAlph
		scoreF = blossum
		gap = blossumGap
	case "dna":
		check = dnaAlph
		scoreF = dna
		gap = dnaGap
	case "default":
		check = defAlph
		scoreF = def
		gap = defGap
	}

	if isFlagPassed("gap") || isFlagPassed("g") {
		gap = gapArg
	}

	var seq1, seq2 string
	if !isFlagPassed("i") {
		panic("undefined input")
	}
	if len(flag.Args()) > 0 && flag.Args()[0][0] != '-' {
		seq1, seq2 = parse(input, 1, check)[0], parse(flag.Args()[0], 1, check)[0]
	} else {
		seqs := parse(input, 2, check)
		seq1, seq2 = seqs[0], seqs[1]
	}

	nwd := &SWData{
		Scores: scoreF,
		Gap:    gap,
	}

	a, b, score := nwd.Align(seq1, seq2)

	var w io.Writer

	if isFlagPassed("o") {
		file, err := os.Create(output)
		if err != nil {
			panic("couldn't open file to write (" + err.Error() + ")")
		}
		defer func() {
			if err := file.Close(); err != nil {
				panic("couldn't close file to write (" + err.Error() + ")")
			}
		}()
		w = file
	} else {
		w = os.Stdout
	}

	writeSeqs(w, a, b)
	writeScore(w, score)
}
