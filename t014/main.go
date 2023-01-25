package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'getFinalString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func getFinalString(s string) string {
	c := 0
	progress := true
	for progress {
		progress = strings.Contains(s, "AWS")
		s = strings.Replace(s, "AWS", "", -1)
		if progress {
			c++
		}
	}
	if len(s) == 0 {
		return "-1"
	}
	return s
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := getFinalString(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
