package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	stdout, err := runCommand("java", "-jar", "../logisim-evolution.jar", "../main.circ", "-load", "program", "-tty", "table")
	if err != nil {
		os.Exit(1)
	}

	expected := []string{"1335", "EECD", "1234", "ffff", "ffff", "ffff", "ffff", "ffff", "ffff", "8888"}

	scanner := bufio.NewScanner(stdout)
	scanner.Scan() // ignores first print
	var i = 0
	for scanner.Scan() {
		line := scanner.Text()

		binaryStr := strings.ReplaceAll(line, " ", "")[0:16]
		gotNum, errBinary := binaryStringToNumber(binaryStr)

		expectedNum, errHex := hexStringToNumber(expected[i])

		if errBinary != nil || errHex != nil {
			fmt.Printf("Parsing error. ErrBinary: '%t', ErrHex: '%t'.", errBinary, errHex)
		}
		if gotNum != expectedNum {
			fmt.Printf("Got different then expected. Expected: %d, but got: %d.\n", expectedNum, gotNum)
		}

		i++
		if len(expected) == i {
			if scanner.Scan() {
				fmt.Printf("Size is different")
			}
			fmt.Print("breaked")
			break
		}
	}

	fmt.Print("closed")
	stdout.Close()
}

func binaryStringToNumber(s string) (uint64, error) {
	ui, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return 0, err
	}
	return ui, nil
}

func hexStringToNumber(s string) (uint64, error) {
	ui, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0, err
	}
	return ui, nil
}

func runCommand(name string, arg ...string) (io.ReadCloser, error) {
	cmd := exec.Command(name, arg...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return stdout, nil
}