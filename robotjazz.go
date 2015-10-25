package robotjazz

import (
	"errors"
	"github.com/SudoQ/robotjazz/chord"
	"github.com/SudoQ/robotjazz/data"
	"github.com/SudoQ/robotjazz/model"
	"log"
	"encoding/csv"
	"os"
	"strconv"
)

type Jazzrobot struct {
	mainModel *model.Model
}

func New() *Jazzrobot {
	return &Jazzrobot{
		mainModel: model.New(),
	}
}

func (jr *Jazzrobot) Load(filename string) error {
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
	}
	jr.mainModel.UpdateCentroids()
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
