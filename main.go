package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

func loadBankData(path string) ([]Bank, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var banksData []Bank
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ",")
		if len(lineParts) != 3 {
			return nil, fmt.Errorf("invalid line form: %s", line)
		}
		binFrom, err := strconv.Atoi(lineParts[1])
		if err != nil {
			return nil, err
		}
		binTo, err := strconv.Atoi(lineParts[2])
		if err != nil {
			return nil, err
		}
		record := Bank{
			Name:    lineParts[0],
			BinFrom: binFrom,
			BinTo:   binTo,
		}
		banksData = append(banksData, record)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return banksData, nil
}

func extractBIN(cardNumber string) int {
	var cleaned []rune
	for _, r := range cardNumber {
		if r >= '0' && r <= '9' {
			cleaned = append(cleaned, r)
		}
	}
	if len(cleaned) < 6 {
		return 0
	}
	binStr := string(cleaned[:6])
	binInt, err := strconv.Atoi(binStr)
	if err != nil {
		return 0
	}
	return binInt
}

func identifyBank(bin int, banks []Bank) string {

	for _, bank := range banks {
		if bin >= bank.BinFrom && bin <= bank.BinTo {
			return bank.Name
		}
	}
	return "Неизвестный банк"
}

func main() {
	// banksData, err := loadBankData("banks.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println(banksData)
	// }
	banks, err := loadBankData("banks.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(identifyBank(400000, banks))
}
