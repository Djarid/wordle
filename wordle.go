package wordle

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/Djarid/wordle/words"
)

func main() {
	// secret := words.GetWord()
	// ws := newWordleState(words.GetWord())
	// for i := 0; i < 6; i++ {
	// 	guess := userGuess()
	// }

	// fmt.Println(secret)
}

func userGuess() string {
	var guess string
	fmt.Println("Please enter guess:")
	fmt.Scanln(&guess)
	return strings.ToUpper(guess)
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
	guesses   [maxGuesses]guess
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
	// fmt.Printf("newGuess: %s\n", g.string())
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
	// iterate through g, (int32) test the value 65 > v.char < 90
	if len(g) != wordSize {
		return errors.New("input must be exactly 5 characters")
	}

	alphabet := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for _, v := range g {
		if !slices.Contains(alphabet, v.char) {
			return errors.New("only letters with from the 26 character English alphabet are valid")
		}
	}

	if !words.IsWord(g.string()) {
		return errors.New("not a valid word")
	}

	if w.currGuess > maxGuesses {
		return errors.New("too many guesses")
	}

	w.guesses[w.currGuess] = g
	w.currGuess++

	return nil
}

// func (w *wordleState) isWordGuessed(g guess) bool {
func (w *wordleState) isWordGuessed() bool {
	result := true
	for _, v := range w.guesses[w.currGuess-1] {
		if v.status != correct {
			result = false
		}
	}
	return result
}

func (w *wordleState) shouldEndGame() bool {
	//check if word correct.

	if w.isWordGuessed() || w.currGuess >= maxGuesses {
		//fmt.Printf("You've guessed correct!")
		return true
	}
	return false
}
