package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func useNewReader(filename string) {
	var count int = 0

	fin, error := os.OpenFile(filename, os.O_RDONLY, 0)
	if error != nil {
		panic(error)
	}
	defer fin.Close()
	/*create a Reader*/
	rd := bufio.NewReader(fin)
	/*read the file and stop when meet err or EOF*/
	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		count++
		/*for each line, process it.
		  if you want it ouput format in command-line, you need clean the '\f'*/
		line = strings.Replace(line, "\f", "", -1)
		fmt.Printf("the line %d: %s", count, line)
	}
}

func useNewScanner(filename string) {
	var count int = 0

	fin, error := os.OpenFile(filename, os.O_RDONLY, 0)
	if error != nil {
		panic(error)
	}
	defer fin.Close()

	sc := bufio.NewScanner(fin)
	/*default split the file use '\n'*/
	for sc.Scan() {
		count++
		fmt.Printf("the line %d: %s\n", count, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Prinfln("An error has hippened")
	}
}

var LineSplit = func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	/*read some*/
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	/*find the index of the byte '\n'
	  and find another line begin i+1
	  default token doesn't include '\n'*/
	if i := bytes.IndexByte(data, '\n'); i > 0 {
		return i + 1, dropCR(data[0:i]), nil
	}

	/*at EOF, we have a final, non-terminal line*/
	if atEOF {
		return len(data), dropCR(data), nil
	}

	/*read some more*/
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	/*drop the '\f'
	  if you don't need, you can delete it*/
	if i := bytes.IndexByte(data, '\f'); i >= 0 {
		tmp := [][]byte{data[0:i], data[(i + 1):]}
		sep := []byte("")
		data = bytes.Join(tmp, sep)
	}
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func useSplit(filename string) {
	var count int = 0

	fin, error := os.OpenFile(filename, os.O_RDONLY, 0)
	if error != nil {
		panic(error)
	}
	defer fin.Close()

	sc := bufio.NewScanner(fin)
	/*Specifies the matching function, default read by lines*/
	sc.Split(LineSplit)
	/*begin scan*/
	for sc.Scan() {
		count++
		fmt.Printf("the line %d: %s\n", count, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Prinfln("An error has hippened")
	}
}

func main() {
	filename := "test.txt"
	fmt.Println("From func useNewReader()")
	useNewReader(filename)

	fmt.Println("From func useNewScanner()")
	useNewScanner(filename)

	fmt.Println("From func useNewSplit()")
	useSplit(filename)
}
