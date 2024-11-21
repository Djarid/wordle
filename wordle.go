package wordle

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Djarid/wordle/words"
)

func main() {
	// secret := words.GetWord()
	// ws := newWordleState(words.GetWord())
	for i := 0; i < 6; i++ {
		guess := userGuess()
		if validateGuess(guess) {

		}
	}

	// fmt.Println(secret)
}

func userGuess() string {
	var guess string
	fmt.Println("Please enter guess:")
	fmt.Scanln(&guess)
	return strings.ToUpper(guess)
}

func validateGuess(g string) bool {
	if len(g) != 5 {
		fmt.Println("Input must be exactly 5 characters")
		return false
	}

	alphabet := []rune("abcdefghijklmonpqrstuvwxyz")
	for _, v := range g {
		if !slices.Contains(alphabet, v) {
			fmt.Println("Only letters with from the 26 character English alphabet are valid")
		}
	}

	if !words.IsWord(g) {
		fmt.Println("Not a valid word")
		return false
	}

	return true

}

func newWordleState(word string) wordleState {
	w := wordleState{}
	copy(w.word[:], word)
	return w
}

const (
	maxGuesses = 6
	wordSize   = 5
)

type wordleState struct {
	word      [wordSize]byte
	guesses   [maxGuesses]string
	currGuess int
}

type letterStatus int

const (
	none letterStatus = iota
	absent
	present
	correct
)

type letter struct {
	char   byte
	status letterStatus
}

type guess [wordSize]letter

func newLetter(character byte) letter {
	return letter{char: character, status: none}
}

func newGuess(word string) guess {
	g := guess{}
	for i, v := range word {
		g[i] = newLetter(byte(v))
	}
	return g
}

func (g *guess) string() string {
	var s [len(g)]byte
	for i, v := range g {
		s[i] = byte(v.char)
	}
	return string(s[:])
}

func (g *guess) updateLettersWithWord(word [wordSize]byte) {
	for i, v := range g {

		switch {
		case v.char == word[i]:
			g[i].status = correct
		case slices.Contains([]byte(word[:]), byte(v.char)):
			g[i].status = present
		default:
			g[i].status = absent
		}
	}
}

func (w *wordleState) appendGuess(g guess) error {
	w.currGuess--

	return nil
}
