package wordle

import (
	"fmt"
	"testing"

	"github.com/Djarid/wordle/words"
)

func TestNewWordleState(t *testing.T) {
	word := "HELLO"
	ws := newWordleState(word)
	wordleAsString := string(ws.word[:])

	t.Log("Created wordleState:")
	t.Logf("    word: %s", wordleAsString)
	t.Logf("    guesses: %v", ws.guesses)
	t.Logf("    currGuess: %v", ws.currGuess)
}

func statusToString(status letterStatus) string {
	switch status {
	case none:
		return "none"
	case correct:
		return "correct"
	case present:
		return "present"
	case absent:
		return "absent"
	default:
		return "unknown"
	}
}

func TestNewGuess(t *testing.T) {
	wordToGuess := "YEILD"
	guess := newGuess(wordToGuess)

	t.Logf("New guess: %s", guess.string())

	// Check that the letter and status are correct
	for i, l := range guess {
		t.Logf("    letter %d: %c, %s", i, l.char, statusToString(l.status))

		if l.char != wordToGuess[i] || l.status != none {
			t.Errorf(
				"letter[%d] = %c, %s; want %c, none",
				i,
				l.char,
				statusToString(l.status),
				wordToGuess[i],
			)
		}
	}
}

func TestUpdateLettersWithWord(t *testing.T) {
	wordToGuess := "YIELD"
	guess := newGuess(wordToGuess)
	t.Logf("New guess: %s", guess.string())

	var word [wordSize]byte
	copy(word[:], "HELLO")
	guess.updateLettersWithWord(word)

	statuses := []letterStatus{
		absent,
		absent,
		present,
		correct,
		absent,
	}

	for i, v := range guess {
		//  t.Logf("    letter %d: %c, %s", i, v.char, statusToString(v.status))

		if v.status != statuses[i] {
			t.Errorf(
				"ERR letter %d: %c, %s; want %c, %s",
				i,
				v.char,
				statusToString(v.status),
				wordToGuess[i],
				statusToString(statuses[i]),
			)
		}
	}

}

func TestAppendGuess(t *testing.T) {
	ws := newWordleState("HELLO")
	g := newGuess("HELLO")

	g.updateLettersWithWord(ws.word)
	ws.appendGuess(g)

	fmt.Printf("Guesses %d, guess %v\n", ws.currGuess, ws.guesses[ws.currGuess].string())

	if !ws.isWordGuessed() {
		t.Errorf("isWordGuessed() should return true")
	}
}
func TestShouldEngGameCorrect(t *testing.T) {
	ws := newWordleState("HELLO")
	g := newGuess("HELLO")

	g.updateLettersWithWord(ws.word)
	ws.appendGuess(g)

	if !ws.shouldEndGame() {
		t.Errorf("shouldEndGame should return true")
	} else {
		if !ws.isWordGuessed() {
			t.Errorf("isWordGuessed should return true")
		}
	}
}

func TestShouldEndGameMaxGuesses(t *testing.T) {
	ws := newWordleState("HELLO")

	for i := 0; i < maxGuesses; i++ {
		g := newGuess(words.GetWord())
		g.updateLettersWithWord(ws.word)
		ws.appendGuess(g)
	}

	if !ws.shouldEndGame() {
		fmt.Printf("test: %v\n", len(ws.guesses) < maxGuesses)
	} else {
		if len(ws.guesses) < maxGuesses {
			t.Errorf("ws.currGuess should be >= maxGuesses")
		}
	}
}
