// https://codingchallenges.fyi/challenges/challenge-wc

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

type WcMode uint8

const (
	Invalid WcMode = 0
	Default WcMode = 1
	Bytes   WcMode = 2
	Lines   WcMode = 3
	Words   WcMode = 4
	Chars   WcMode = 5
)

type WcArgs struct {
	path *string
	mode WcMode
}

func getMode(arg string) (WcMode, error) {
	switch arg {
	case "-c":
		return Bytes, nil
	case "-l":
		return Lines, nil
	case "-w":
		return Words, nil
	case "-m":
		return Chars, nil
	default:
		return Invalid, fmt.Errorf("invalid argument %s", arg)
	}
}

func parseArgs(args []string) (*WcArgs, error) {
	switch len(args) {
	case 0:
		{
			wcArgs := WcArgs{
				path: nil,
				mode: Default,
			}

			return &wcArgs, nil
		}
	case 1:
		{
			var wcArgs WcArgs
			mode, modeErr := getMode(args[0])

			if modeErr == nil {
				// only mode passed
				wcArgs = WcArgs{
					path: nil,
					mode: mode,
				}
			} else {
				// only path passed
				wcArgs = WcArgs{
					path: &args[0],
					mode: Default,
				}
			}

			return &wcArgs, nil
		}
	case 2:
		{
			mode := Default

			mode, modeErr := getMode(args[0])
			if modeErr != nil {
				return nil, modeErr
			}

			wcArgs := WcArgs{
				path: &args[1],
				mode: mode,
			}

			return &wcArgs, nil
		}
	default:
		return nil, fmt.Errorf("invalid arguments")
	}
}

func eachRune(input io.Reader, fn func(rune)) {
	reader := bufio.NewReader(input)

	for {
		currentRune, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		fn(currentRune)
	}
}

func countBytes(next *func(rune)) (*int, func(rune)) {
	byteCount := 0

	fn := func(currentRune rune) {
		byteCount += len(string(currentRune))

		if next != nil {
			(*next)(currentRune)
		}
	}

	return &byteCount, fn
}

func countLines(next *func(rune)) (*int, func(rune)) {
	linesCount := 0

	fmt.Println("countLines")

	fn := func(currentRune rune) {
		if currentRune == '\n' {
			linesCount += 1
		}

		if next != nil {
			(*next)(currentRune)
		}
	}

	return &linesCount, fn
}

func countWords(next *func(rune)) (*int, func(rune)) {
	wordsCount := 0
	currentWordLength := 0

	fmt.Println("countWords")

	fn := func(currentRune rune) {
		if unicode.IsSpace(currentRune) {
			if currentWordLength > 0 {
				wordsCount += 1
			}

			currentWordLength = 0
		} else {
			currentWordLength += 1
		}

		if next != nil {
			(*next)(currentRune)
		}
	}

	return &wordsCount, fn
}

func countChars(next *func(rune)) (*int, func(rune)) {
	charsCount := 0

	fmt.Println("countChars")

	fn := func(currentRune rune) {
		charsCount += 1

		if next != nil {
			(*next)(currentRune)
		}
	}

	return &charsCount, fn
}

func printResult(count int, path *string) {
	if path == nil {
		fmt.Printf("%12d", count)
	} else {
		fmt.Printf("%12d %s", count, *path)
	}
}

type CalcFn = func(*func(rune)) (*int, func(rune))

func runPipeline(input io.Reader, calcFns [](CalcFn)) *[]*int {
	var counters []*int
	var prevFn *func(rune) = nil

	for _, calcFn := range calcFns {
		counter, nextFn := calcFn(prevFn)
		prevFn = &nextFn
		counters = append(counters, counter)
	}

	eachRune(input, *prevFn)

	return &counters
}

func runOne(input io.Reader, fn CalcFn) *int {
	counters := *runPipeline(input, [](CalcFn){fn})
	return counters[0]
}

func main() {
	wcArgs, parseArgsError := parseArgs(os.Args[1:])
	if parseArgsError != nil {
		log.Fatalln(parseArgsError)
	}

	var input *os.File

	if wcArgs.path == nil {
		input = os.Stdin
	} else {
		var inputErr error
		input, inputErr = os.Open(*wcArgs.path)
		if inputErr != nil {
			log.Fatalln(inputErr)
		}
		defer input.Close()
	}

	switch wcArgs.mode {
	case Bytes:
		counter := *runOne(input, countBytes)
		printResult(counter, wcArgs.path)
	case Lines:
		counter := *runOne(input, countLines)
		printResult(counter, wcArgs.path)
	case Words:
		counter := *runOne(input, countWords)
		printResult(counter, wcArgs.path)
	case Chars:
		counter := *runOne(input, countChars)
		printResult(counter, wcArgs.path)
	case Default:
		counters := *runPipeline(input, [](CalcFn){countChars, countLines, countWords})

		if wcArgs.path == nil {
			fmt.Printf("%12d %12d %12d", *counters[0], *counters[1], *counters[2])
		} else {
			fmt.Printf("%12d %12d %12d %s", *counters[0], *counters[1], *counters[2], *wcArgs.path)
		}
	}
}
