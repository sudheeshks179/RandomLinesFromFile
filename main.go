package main

import (
	"bufio"
	"flag"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

func random(min, max int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}

var avgLineLength = flag.Int("avgLinesize", 20, "avg size of a line")
var filename = flag.String("filename", "", "filename with path.Mandatory")
var nrOfLines = flag.Int("numberOfLines", 1, "line number")

func main() {

	flag.Parse()
	log.Println("Filename: ", *filename, " No of lines to print: ", *nrOfLines)
	if len(*filename) == 0 {
		log.Fatal("file name is mandatory")
		log.Fatal("./RandomLinesFromFile -filename /tmp/file.txt -numberOfLines 100")
		return
	}
	usingScanner(*filename)
}

func usingScanner(filename string) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
			return
		}
	}
	log.Println("Size in bytes:", fileInfo.Size())
	rand.Seed(time.Now().Unix())
	arr := rand.Perm(int(fileInfo.Size()) / *avgLineLength)
	myArr := arr[:*nrOfLines+1]
	sort.Ints(myArr)
	//log.Println("myArr :", myArr)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("File does not exist. Error:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// Scan for next token.
	var j, i int
	for scanner.Scan() {
		if i == myArr[j] {
			log.Println("Line no: ", i, " Line[", j+1, "]: ", scanner.Text())
			j++
			if j == *nrOfLines {
				break
			}
		}
		i++
	}
}
