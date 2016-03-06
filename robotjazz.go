package robotjazz

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SudoQ/robotjazz/chord"
	"github.com/SudoQ/robotjazz/data"
	"github.com/SudoQ/robotjazz/model"
	"github.com/SudoQ/robotjazz/noteset"
	"io/ioutil"
	"log"
	//"os"
	//"strconv"
)

type Jazzrobot struct {
	mainModel *model.Model
	chordMap  map[string]*noteset.Noteset
}

func New() *Jazzrobot {
	return &Jazzrobot{
		mainModel: model.New(),
		chordMap:  make(map[string]*noteset.Noteset),
	}
}

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

func extractSlice(jsonMap map[string]interface{}) ([]pattern, error) {
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

func loadPatternSlice(filename string) ([]pattern, error) {
	var err error

	jsonMap, err := loadJSON(filename)

	if err != nil {
		return nil, err
	}

	chordPatterns, err := extractSlice(jsonMap)

	if err != nil {
		return nil, err
	}
	return chordPatterns, nil
}

func (jr *Jazzrobot) Load(filename string) error {
	patterns, _ := loadPatternSlice(filename)
	//chrom := []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	//chords := make(map[string][]float64)
	k := len(patterns) * 12
	for _, pattern := range patterns {
		for noteValue := 0; noteValue < 12; noteValue++ {
			weights := make([]float64, 12)
			for noteIndex, patternEntry := range pattern.Notes {
				w := pattern.Weights[noteIndex]
				newNoteValue := (noteValue + patternEntry) % 12
				weights[newNoteValue] = pattern.Weights[noteIndex] * w
			}
			id := fmt.Sprintf("%d", len(jr.chordMap))
			root := noteValue
			noteWeights := weights
			patternId := pattern.Name
			patternNotes := pattern.Notes

			jr.chordMap[id] = noteset.New(id, root, noteWeights, patternId, patternNotes)
			jr.mainModel.AddCentroid(data.New(weights, k, id))
			//chordName := fmt.Sprintf("%s %s", noteName, pattern.Name)
			//chords[chordName] = chord
		}
	}
	/*
		csvfile, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer csvfile.Close()
		reader := csv.NewReader(csvfile)
		rawCSVdata, err := reader.ReadAll()
		if err != nil {
			return err
		}
		k := len(rawCSVdata)
		for _, v := range rawCSVdata {
			attributes := make([]float64, 0)
			var tag string
			for i := range v {
				if i == 0 {
					tag = v[i]
					continue
				}
				attr, _ := strconv.ParseFloat(v[i], 64)
				attributes = append(attributes, attr)
			}
			jr.mainModel.AddCentroid(data.New(attributes, k, tag))
		}*/
	jr.mainModel.UpdateCentroids()
	log.Printf("Loaded [%d, %d] number of chords\n", len(jr.chordMap), len(jr.mainModel.Centroids))
	return nil
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (jr *Jazzrobot) GetMatchingChords(notes []float64) ([]*chord.Chord, error) {
	if len(notes) != 12 {
		return nil, errors.New("First input argument must be a float64 slice of length 12")
	}
	dataItem, err := jr.mainModel.Classify(notes)
	if err != nil {
		return nil, err
	}
	return getTopMatches(dataItem)
	/*
		topTenChords := make([]*chord.Chord, 0)
		for i := 0; i < MinInt(len(dataItem.ClosestCentroids), 10); i++ {
			centroid := dataItem.ClosestCentroids[i]

			name := centroid.Tag
			noteWeights := centroid.Attributes
			chrd := chord.New(name, noteWeights)
			topTenChords = append(topTenChords, chrd)
		}
		return topTenChords, nil
	*/
}

func getTopMatches(dataItem *data.Data) ([]*chord.Chord, error) {
	topTenChords := make([]*chord.Chord, 0)
	for i := 0; i < MinInt(len(dataItem.ClosestCentroids), 10); i++ {
		centroid := dataItem.ClosestCentroids[i]

		name := centroid.Tag
		noteWeights := centroid.Attributes
		chrd := chord.New(name, noteWeights)
		topTenChords = append(topTenChords, chrd)
	}
	return topTenChords, nil
}

func (jr *Jazzrobot) GetSimilarChords(chordName string) ([]*chord.Chord, error) {
	centroids := jr.mainModel.Centroids
	var chrd *data.Data
	log.Println(chordName)
	for _, centroid := range centroids {
		if centroid.Tag == chordName {
			log.Println("Chord found")
			chrd = centroid
		}
	}
	if chrd == nil {
		return nil, errors.New("Chord not found")
	}
	topChords, err := getTopMatches(chrd)
	if err != nil {
		return nil, err
	}
	return topChords, nil
}
