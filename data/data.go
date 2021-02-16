package data

import (
	"bufio"
	detailsPack "factory/details"
	"log"
	"os"
	"strconv"
	"strings"
)

func formDetailObjectFromText(text string, detailArr []detailsPack.Detals) {
	str := strings.Split(text, " ")
	machineNum := -1
	for i, el := range str {
		if i == 0 {
			machineNum, _ = strconv.Atoi(str[0])
			continue
		}
		detailNum, _ := strconv.Atoi(el)

		detailArr[detailNum-1].Vector[machineNum-1] = true
	}
}

func ReadDataAndFormPrimaryArr() ([]detailsPack.Detals, int, int) {
	file, err := os.Open("data\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	primaryStr := strings.Split(scanner.Text(), " ")

	machineSize, _ := strconv.Atoi(primaryStr[0])
	detailSize, _ := strconv.Atoi(primaryStr[1])

	detailArr := make([]detailsPack.Detals, detailSize)
	for i := 0; i < detailSize; i++ {
		detailArr[i].Vector = make([]bool, machineSize)
	}
	for scanner.Scan() {
		formDetailObjectFromText(scanner.Text(), detailArr)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return detailArr, detailSize, machineSize
}
