package parmnemonics

import (
	"sync"
)

var mnemonics = map[string]string{
	"2": "ABC",
	"3": "DEF",
	"4": "GHI",
	"5": "JKL",
	"6": "MNO",
	"7": "PQRS",
	"8": "TUV",
	"9": "WXYZ",
}

func generateMnemonics(digits string, index int, current string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	// If all digits have been processed, send the current combination
	if index == len(digits) {
		results <- current
		return
	}
	// Get the letters corresponding to the current digit
	letters := mnemonics[string(digits[index])]
	for _, letter := range letters {
		// For each letter, call the function recursively with the next digit
		wg.Add(1)
		go generateMnemonics(digits, index+1, current+string(letter), results, wg)
	}
}

func Run(phoneNumber string) {
	results := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)

	go generateMnemonics(phoneNumber, 0, "", results, &wg)

	go func() {
		wg.Wait()
		close(results)
	}()
}
