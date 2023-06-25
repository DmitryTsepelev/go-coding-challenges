package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func printPrompt() {
	fmt.Print("ccsh> ")
}

const LS = "ls"
const PWD = "pwd"
const EXIT = "exit"
const CAT = "cat"
const CD = "cd"
const UNIQ = "uniq"

func cd(in *io.Reader, out io.Writer, bashState *BashState, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("expected argument")
	}

	targetPath := args[0]

	if targetPath[0] == '/' {
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			return fmt.Errorf("no such file or directory: " + args[0])
		}
		bashState.currentDir = targetPath
		return nil
	}

	targetPathCmp := strings.Split(targetPath, "/")
	currentPathCmp := strings.Split(bashState.currentDir, "/")

	for _, currentTargetDir := range targetPathCmp {
		switch currentTargetDir {
		case ".":
			continue
		case "..":
			currentPathCmp = currentPathCmp[0 : len(currentPathCmp)-1]
		default:
			{
				currentPathCmp = append(currentPathCmp, currentTargetDir)
				candidateDir := strings.Join(currentPathCmp, "/")
				if _, err := os.Stat(candidateDir); os.IsNotExist(err) {
					return fmt.Errorf("no such file or directory: " + args[0])
				}
			}
		}
	}

	bashState.currentDir = strings.Join(currentPathCmp, "/")

	return nil
}

func ls(in *io.Reader, out io.Writer, bashState *BashState, args []string) error {
	entries, err := os.ReadDir(bashState.currentDir)
	if err != nil {
		return err
	}

	var output string
	for _, e := range entries {
		output += fmt.Sprintf("%-16s", e.Name())
	}
	output += "\n"

	out.Write([]byte(output))

	return nil
}

func pwd(in *io.Reader, out io.Writer, bashState *BashState, args []string) error {
	out.Write([]byte(bashState.currentDir + "\n"))
	return nil
}

func cat(in *io.Reader, out io.Writer, bashState *BashState, args []string) error {
	var content []byte
	var err error

	if in == nil {
		if len(args) == 0 {
			return fmt.Errorf("expected argument")
		}

		path := bashState.currentDir + "/" + args[0]

		content, err = os.ReadFile(path)
	} else {
		content, err = ioutil.ReadAll(*in)
	}

	if err != nil {
		return err
	}

	output := string(content) + "\n"
	out.Write([]byte(output))
	return nil
}

func exit(in *io.Reader, out io.Writer, bashState *BashState, args []string) error {
	bashState.exited = true
	return nil
}

func uniq(in *io.Reader, out io.Writer, bashState *BashState, args []string) error {
	var scanner *bufio.Scanner

	if in == nil {
		if len(args) == 0 {
			return fmt.Errorf("expected argument")
		}

		path := bashState.currentDir + "/" + args[0]

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(*in)
	}

	uniqLines := []string{}
	prevLines := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if !prevLines[line] {
			prevLines[line] = true
			uniqLines = append(uniqLines, line)
		}
	}

	output := strings.Join(uniqLines, "\n")
	out.Write([]byte(output + "\n"))

	return nil
}

func defaultCommand(*io.Reader, io.Writer, *BashState, []string) error {
	return fmt.Errorf("No such file or directory")
}

type BashState struct {
	exited     bool
	currentDir string
}

type Command func(*io.Reader, io.Writer, *BashState, []string) error

var commands = map[string]Command{
	LS:   ls,
	PWD:  pwd,
	CAT:  cat,
	CD:   cd,
	EXIT: exit,
	UNIQ: uniq,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	state := &BashState{
		exited:     false,
		currentDir: currentDir,
	}

	for {
		printPrompt()

		input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		input = strings.Replace(input, "\n", "", -1)

		subcommands := strings.Split(input, "|")

		var prevBuffer bytes.Buffer

		for idx, subcommand := range subcommands {
			// TODO: extract to parseCommand
			cmp := strings.Fields(subcommand)
			var command string
			args := make([]string, 0)

			switch len(cmp) {
			case 0:
				continue
			case 1:
				command = cmp[0]
			default:
				{
					command = cmp[0]
					args = cmp[1:]
				}
			}
			// end extract

			commandFn := commands[command]
			if commandFn == nil {
				commandFn = defaultCommand
			}

			var commandErr error
			var reader io.Reader = &prevBuffer
			var writer io.Writer = &prevBuffer

			if len(subcommands) == 1 {
				commandErr = commandFn(nil, os.Stdout, state, args)
			} else if idx == (len(subcommands) - 1) {
				commandErr = commandFn(&reader, os.Stdout, state, args)
			} else if idx == 0 {
				commandErr = commandFn(nil, writer, state, args)
			} else {
				var nextBuffer bytes.Buffer
				commandErr = commandFn(&reader, &nextBuffer, state, args)
				prevBuffer = nextBuffer
			}

			if commandErr != nil {
				fmt.Println(commandErr)
			}

			if state.exited {
				break
			}
		}

		if state.exited {
			break
		}
	}
}
