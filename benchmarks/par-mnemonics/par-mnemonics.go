package parmnemonics

import (
	"fmt"
	"strings"
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

/*
 * A map from digit strings to the words that represent them,
 * e.g. "5282" -> Set("Java", "Kata", "Lava", ...)
 */
func wordsForNumParallel(digits string, index int, current string, results chan<- string, wg *sync.WaitGroup) {
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
		go wordsForNumParallel(digits, index+1, current+string(letter), results, wg)
	}
}

func wordsForNum(phoneNumber string) map[string][]string {
	results := make(chan string)
	var nums []string
	var wg sync.WaitGroup

	wg.Add(1)

	go wordsForNumParallel(phoneNumber, 0, "", results, &wg)

	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		nums = append(nums, result)
	}
	wordsMap := make(map[string][]string)
	wordsMap[phoneNumber] = nums
	return wordsMap
}

/* Maps a word to the digit string it can represent, e.g. "Java" -> "5282" */
func wordCode(input string) string {
	var result string = ""
	for _, char := range strings.ToUpper(input) {
		for key, mnemonic := range mnemonics {
			for _, character := range mnemonic {
				if character == char {
					result += key
				}
			}
		}
	}
	return result
}

func encodeParallel(word string, currentResult []string, resultChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	number := wordCode(word)
	wordMap := wordsForNum(number)

	for numCode, wordList := range wordMap {
		if strings.HasSuffix(number, numCode) {
			for _, word := range wordList {
				remainingLength := len(number) - len(numCode)

				if remainingLength == 0 {
					resultChan <- word
				} else {
					var innerWg sync.WaitGroup
					intermediateResults := encode(number[:remainingLength], word, currentResult)

					for _, intermediate := range intermediateResults {
						innerWg.Add(1)
						go func(intermediate, word string) {
							defer innerWg.Done()
							resultChan <- intermediate + " " + word
						}(intermediate, word)
					}
					innerWg.Wait()
				}
			}
		}
	}
}

func encode(number string, words string, currentResult []string) []string {
	wordMap := wordsForNum(words)
	result := []string{}

	for numCode, wordList := range wordMap {
		if strings.HasSuffix(number, numCode) {
			for _, word := range wordList {
				remainingLength := len(number) - len(numCode)
				if remainingLength == 0 {
					result = append(result, word)
				} else {
					intermediateResults := encode(number[:remainingLength], words, currentResult)
					for _, intermediate := range intermediateResults {
						result = append(result, intermediate+" "+word)
						fmt.Println(intermediate + " " + word)
					}
				}
			}
		}
	}
	return result
}

func Run() {
	words := []string{
		"Scala",
		"rocks",
		"Pack",
		"brocks",
		"GWT",
		"implicit",
		"nice",
		"ScalaGWT",
		"cat",
		"EFPL",
		"Lausanne",
		"sCala",
		"ROcks",
		"pAck",
		"Java",
		"Apple",
		"Google",
		"Rochester",
		"Utah",
		"Rice",
		"wyr",
		"lxm",
		"Scala",
		"rocks",
		"Pack",
		"brocks",
		"GWT",
		"implicit",
		"nice",
		"ScalaGWT",
		"cat",
		"EFPL",
		"Lausanne",
		"sCala",
		"ROcks",
		"pAck",
		"Java",
		"Apple",
		"Google",
		"Rochester",
		"Utah",
		"Rice",
		"wyr",
		"lxm",
	}

	currentResult := []string{}

	resultChan := make(chan string)
	var wg sync.WaitGroup

	for x := 0; x < 30; x++ {
		for _, word := range words {
			wg.Add(1)
			go encodeParallel(word, currentResult, resultChan, &wg)
		}
		go func() {
			wg.Wait()
			close(resultChan)
		}()
	}
}
