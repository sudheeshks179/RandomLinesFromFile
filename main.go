package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func random(min, max int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}

func parseFlags() (file string, lines, lineLength int64) {
	avgLineLength := flag.Int64("avgLineSize", 20, "avg size of a line")
	fileName := flag.String("filename", "words_alpha.txt", "Filename with path.Mandatory")
	nrOfLines := flag.Int64("numberOfLines", 1, "Number of lines to print")
	flag.Parse()

	file = *fileName
	lines = *nrOfLines
	lineLength = *avgLineLength
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "This is not helpful.\n")
	}
	if len(file) == 0 {
		fmt.Println("file name is mandatory")
		fmt.Println("./RandomLinesFromFile -filename /tmp/file.txt -numberOfLines 100")
		file = "words_alpha.txt"
	}
	if lines <= 0 {
		fmt.Println("numberOfLines should be greater than 0.")
		lines = 1
	}
	if lineLength <= 0 {
		fmt.Println("avgLineLength should be greater than 0.")
		lineLength = 20
	}
	return
}
func main() {
	file := "words_alpha.txt"
	var lines, lineLength int64
	if len(os.Args) == 2 {
		lineLength = 20
		var err error
		lines, err = strconv.ParseInt(os.Args[1], 10, 0)
		if err != nil || lines <= 0 {
			lines = 1
		}

	} else {
		file, lines, lineLength = parseFlags()
	}
	//fmt.Println("Filename: ", file, " No of lines to print: ", lines,
	//	" Avg Line Size: ", lineLength)

	usingScanner(file, lines, lineLength)
}

func usingScanner(fileName string, nrOfLines, lineLength int64) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist.")
			return
		}
	}
	//fmt.Println("Size in bytes:", fileInfo.Size())
	rand.Seed(time.Now().Unix())
	arr := rand.Perm(int(fileInfo.Size() / lineLength))
	if int(nrOfLines+1) > len(arr) {
		fmt.Println("Error while processing. Reduce no of lines to print",
			"or add more lines to file and try again.")
		return
	}
	myArr := arr[:nrOfLines+1]
	sort.Ints(myArr)
	//fmt.Println("myArr :", myArr)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File does not exist. Error:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// Scan for next token.
	var j, i int
	for scanner.Scan() {
		if i == myArr[j] {
			//fmt.Println("Line no: ", i, " Line[", j+1, "]: ", scanner.Text())
			fmt.Println(scanner.Text())
			j++
			if j == int(nrOfLines) || j == len(myArr) {
				break
			}
		}
		i++
	}
}
