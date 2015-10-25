package main

import (
	"fmt"
	"github.com/SudoQ/robotjazz"
	"github.com/SudoQ/robotjazz/util"
)

var (
	QUIT = "exit"
)

var cmdMap = map[string]func(){
	"help":  help,
	"match": GetMatchingChords,
	"sim":   GetSimilarChords,
}

var helpMap = map[string]string{
	"help":  "Find out more about a given command",
	"match": "Match input notes to chords",
	"sim":   "Input chord and output similar chords",
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

var jr *robotjazz.Jazzrobot

func main() {
	jr = robotjazz.New()
	jr.Load("../resources/chords.json")

	menu()
	cmd := ""
	for cmd != QUIT {
		cmd = prompt("rj> ")
		if cmdFunc, ok := cmdMap[cmd]; ok {
			cmdFunc()
		}
	}
}

func GetMatchingChords() {
	reducedNotes := make([]int, 0)

	strNote := "not done"
	for strNote != "" {
		strNote = prompt("Enter a note or hit enter to continue: ")
		if strNote == "" {
			break
		}

		reducedNotes = append(reducedNotes, util.GetNoteValue(strNote))
	}

	extendedNotes := util.ExtendedNoteForm(reducedNotes)

	chords, err := jr.GetMatchingChords(extendedNotes)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, chord := range chords {
		fmt.Println(chord)
	}
}

func GetSimilarChords() {
	chordRoot := prompt("Enter the chord root note (e.g. C, F, Bb, ...): ")
	chordPattern := prompt("Enter a chord pattern (e.g. major, minor, ...): ")
	chordStr := fmt.Sprintf("%s %s", chordRoot, chordPattern)

	fmt.Println(chordStr)
	chords, err := jr.GetSimilarChords(chordStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, chord := range chords {
		fmt.Println(chord)
	}
}
