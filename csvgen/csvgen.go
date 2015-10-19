package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type pattern struct {
	Name    string
	Notes   []int
	Weights []float64
}

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

func extractMap(jsonMap map[string]interface{}) ([]pattern, error) {
	resultSlice := make([]pattern, 0)
	for _, jmValue := range jsonMap {
		valueSlice := jmValue.([]interface{})
		for _, patternMap := range valueSlice {

			mapv := patternMap.(map[string]interface{})

			patternName := mapv["name"].(string)
			notes := make([]int, 0)
			for _, mv := range mapv["notes"].([]interface{}) {
				notes = append(notes, int(mv.(float64)))
			}
			weights := make([]float64, 0)
			for _, weight := range mapv["weights"].([]interface{}) {
				weights = append(weights, weight.(float64))
			}

			resultSlice = append(resultSlice, pattern{Name: patternName, Notes: notes, Weights: weights})
		}
	}

	return resultSlice, nil
}

func loadPatternMap(filename string) ([]pattern, error) {
	var err error

	jsonMap, err := loadJSON(filename)

	if err != nil {
		return nil, err
	}

	chordPatterns, err := extractMap(jsonMap)

	if err != nil {
		return nil, err
	}
	return chordPatterns, nil
}

func main() {
	patterns, _ := loadPatternMap("../resources/chords.json")
	chrom := []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	chords := make(map[string][]float64)
	for _, pattern := range patterns {
		for noteValue, noteName := range chrom {
			chord := make([]float64, 12)
			for noteIndex, patternEntry := range pattern.Notes {
				w := pattern.Weights[noteIndex]
				newNoteValue := (noteValue + patternEntry) % 12
				chord[newNoteValue] = pattern.Weights[noteIndex] * w
			}
			chordName := fmt.Sprintf("%s %s", noteName, pattern.Name)
			chords[chordName] = chord
		}
	}

	csvLines := ""
	for name, chord := range chords {
		csvLine := name
		for _, note := range chord {
			csvLine = csvLine + fmt.Sprintf(",%.1f", note)
		}
		csvLine += "\n"
		csvLines += csvLine
	}
	ioutil.WriteFile("out.csv", []byte(csvLines), 0644)
}
