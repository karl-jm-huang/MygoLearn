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
		page, err := rd.ReadString('\f')
		if err != nil || err == io.EOF {
			/*if it has no '\f' behind the last line*/
			if err == io.EOF && len(page) != 0 {
				count++
				fmt.Printf("the page %d:\n%s\n", count, page)
			}
			break
		}
		count++
		/*for each line, process it.
		  if you want it ouput format in command-line, you need clean the '\f'*/
		page = strings.Replace(page, "\f", "", -1)
		fmt.Printf("the page %d:\n%s\n", count, page)
	}
}

var LineSplit = func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	/*read some*/
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	/*find the index of the byte '\f'
	  and find another line begin i+1
	  default token doesn't include '\n'*/
	if i := bytes.IndexByte(data, '\f'); i > 0 {
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
		fmt.Printf("the line %d:\n%s\n", count, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Prinfln("An error has hippened")
	}
}

func main() {
	filename := "test.txt"
	fmt.Println("From func useNewReader()")
	useNewReader(filename)

	fmt.Println("From func useNewSplit()")
	useSplit(filename)
}
