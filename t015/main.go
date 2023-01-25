package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'findRange' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER num as parameter.
 */

func big_small(arr []int) (int, int, int) {
	small := arr[0]
	small2 := arr[0]
	big := arr[0]

	for i := range arr {
		if small > arr[i] && arr[i] != 0 {
			small = arr[i]
		}
		if small2 > arr[i] {
			small2 = arr[i]
		}

		if big < arr[i] && arr[i] != 0 {
			big = arr[i]
		}
	}

	return big, small, small2
}

func findRange(num int32) int64 {
	var str = strconv.Itoa(int(num))
	var i_arr = make([]int, len(str))

	for i := range str {
		i_arr[i], _ = strconv.Atoi(string(str[i]))
	}

	var big_result, small_result int

	var big, _, small2 = big_small(i_arr)

	big_result, _ = strconv.Atoi(strings.Replace(str, strconv.Itoa(small2), "9", -1))
	small_result, _ = strconv.Atoi(strings.Replace(str, strconv.Itoa(big), "1", -1))

	fmt.Printf("%v", int64(big_result-small_result))
	//r, _ := strconv.Atoi(str)
	return int64(big_result - small_result)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	numTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	num := int32(numTemp)

	result := findRange(num)

	fmt.Fprintf(writer, "%d\n", result)

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
