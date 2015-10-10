package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func loadJSON(filename string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(data, &jsonMap)
	if err != nil {
		return nil, err
	}

	return jsonMap, nil
}

func extractMap(jsonMap map[string]interface{}) (map[string][]int, error) {
	resultMap := make(map[string][]int)
	for _, jmValue := range jsonMap {
		valueSlice := jmValue.([]interface{})
		for _, patternMap := range valueSlice {

			mapv := patternMap.(map[string]interface{})

			patternName := mapv["name"].(string)
			notes := make([]int, 0)
			for _, mv := range mapv["notes"].([]interface{}) {
				notes = append(notes, int(mv.(float64)))
			}
			resultMap[patternName] = notes
		}
	}

	return resultMap, nil
}

func loadPatternMap(filename string) (map[string][]int, error) {
	var err error

	jsonMap, err := loadJSON(filename)

	if err != nil {
		return nil, err
	}

	chordPatternMap, err := extractMap(jsonMap)

	if err != nil {
		return nil, err
	}
	return chordPatternMap, nil
}

//type Chord struct {
//	root int
//	patternName string
//	notes []int
//}
//
//type ChordPattern struct {
//	name string
//	notes []int
//}

func main(){
	patternMap, _ := loadPatternMap("../resources/chords.json")
	//chordMap = make(map[int][]int)
	chrom := []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	chords := make(map[string][]int)
	for patternName, pattern := range patternMap {
		for noteValue, noteName := range chrom {
			chord := make([]int, 0)
			for _, patternEntry := range pattern {
				chord = append(chord, ((noteValue + patternEntry) % 12))
			}
			chordName := fmt.Sprintf("%s %s", noteName, patternName)
			chords[chordName] = chord
		}
	}
	// Apply generic weights
	weights := []float64{
		1.0, // 0
		0.5, // 3
		0.5, // 5
		0.5, // 7
		0.5, // 2
		0.5, // 4
		0.5, // 6
		0.5, // ?
		0.5, // ?
		0.5, // ?
		0.5, // ?
		0.5} // ?

	wChords := make(map[string][]float64)
	for chordName, chord := range chords {
		wChord := make([]float64,12)
		for i, note := range chord {
			wChord[note] = weights[i]
		}
		wChords[chordName] = wChord
	}

	// Format as csv
	csvLines := ""
	for name, chord := range wChords {
		csvLine := name
		for _,note := range chord {
			csvLine = csvLine + fmt.Sprintf(",%.1f",note)
		}
		csvLine += "\n"
		csvLines += csvLine
	}
	ioutil.WriteFile("out.csv", []byte(csvLines), 0644)
	// Output csv lines, with weigthed notes and chord name
}
