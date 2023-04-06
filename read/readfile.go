package agv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFileData(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func GetFloatData(fileName string) (numbers [3]float64, err error) {
	open, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(open)
	i := 0
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	err = open.Close()
	if err != nil {
		return numbers, err
	}
	return numbers, nil
}
