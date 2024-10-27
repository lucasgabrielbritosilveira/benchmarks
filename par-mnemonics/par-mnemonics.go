package parmnemonics

import (
	"unicode"
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

func WordCode(word string) string {
	/* TODO Invert the mnemonics map to give a map from chars 'A' ... 'Z' to '2' ... '9'
	 * e.g. Map(E -> 3, X -> 9, N -> 6, T -> 8, Y -> 9,...)  */
	charCode := make(map[rune]string)

	for key, value := range mnemonics {
		for _, char := range value {
			charCode[char] = key
		}
	}

	word_code := ""

	//	Give a word and return the number

	for _, character := range word {
		character = unicode.ToUpper(character)
		word_code += charCode[character]
	}

	return word_code

}

/** A map from digit strings to the words that represent them,
*  e.g. "5282" -> Set("Java", "Kata", "Lava", ...) */
func wordsForNum(word_code string) map[string][]string {
	//list := []string{}
	ret := make(map[string][]string)
	return ret
}

// WordForNum parallel
func wordsForNumParallel() {

}
