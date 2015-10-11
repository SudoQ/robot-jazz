package chord

import (
	"github.com/SudoQ/robotjazz/util"
	"fmt"
)

type Chord struct {
	name string
	noteWeights []float64
}

func New(name string, noteWeights []float64) *Chord {
	return &Chord {
		name: name,
		noteWeights: noteWeights,
	}
}

func (chrd * Chord) Name() string {
	return chrd.name
}

func (chrd * Chord) NoteWeights() []float64 {
	return chrd.noteWeights
}

func (chrd * Chord) String() string {
	line := "%s\t%s"
	chordStr := ""
	for _, note := range util.ReducedNoteForm(chrd.noteWeights) {
		chordStr += fmt.Sprintf("%s ", util.GetNoteName(note))
	}
	return fmt.Sprintf(line, chrd.name, chordStr)
}
