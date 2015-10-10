package main

import (
	"fmt"
	"github.com/SudoQ/robotjazz"
)

var (
	QUIT = "exit"
)

var cmdMap = map[string]func(){
	"help":  help,
	"match": GetMatchingChords,
}

var helpMap = map[string]string{
	"help":  "Find out more about a given command",
	"match": "Match input notes to chords",
}

func menu() {
	fmt.Println("Robot Jazz CLI version 0.1")
	fmt.Println("Enter \"exit\" to exit")
	fmt.Println("Enter \"help\" to know more about the commands")
}

func help() {
	fmt.Println("Available commands:")
	for cmd := range helpMap {
		fmt.Printf("\t%s\n", cmd)
	}
	cmd := prompt("Enter command to know more: ")
	if helpString, ok := helpMap[cmd]; ok {
		fmt.Printf("%s: %s\n", cmd, helpString)
	} else {
		fmt.Printf("%s has no help section\n", cmd)
	}
}

func prompt(text string) string {
	fmt.Print(text)
	input := ""
	fmt.Scanln(&input)
	return input
}

func main() {
	menu()
	cmd := ""
	for cmd != QUIT {
		cmd = prompt("rj> ")
		if cmdFunc, ok := cmdMap[cmd]; ok {
			cmdFunc()
		}
	}
}

var chrom = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}

func getNoteName(note int) string {
	return chrom[note]
}

var noteValue = map[string]int{
	"C":  0,
	"C#": 1,
	"Db": 1,
	"D":  2,
	"D#": 3,
	"Eb": 3,
	"E":  4,
	"F":  5,
	"F#": 6,
	"Gb": 6,
	"G":  7,
	"G#": 8,
	"Ab": 8,
	"A":  9,
	"A#": 10,
	"Bb": 10,
	"B":  11,
}

func getNoteValue(note string) int {
	return noteValue[note]
}

func GetMatchingChords() {
	reducedNotes := make([]int, 0)

	strNote := "not done"
	for strNote != "" {
		strNote = prompt("Enter a note or hit enter to continue: ")
		if strNote == "" {
			break
		}

		reducedNotes = append(reducedNotes, getNoteValue(strNote))
	}

	extendedNotes := robotjazz.ExtendedNoteForm(reducedNotes)

	chords, err := robotjazz.GetMatchingChords(extendedNotes)
	if err != nil {
		fmt.Println(err)
		return
	}
	for name, notes := range chords {
		line := "%s\t%s\n"
		chordStr := ""
		for _, note := range notes {
			chordStr += fmt.Sprintf("%s ", getNoteName(note))
		}
		fmt.Printf(line, name, chordStr)
	}
}
