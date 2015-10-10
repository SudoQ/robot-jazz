package robotjazz

import (
	"errors"
	"github.com/SudoQ/robotjazz/model"
)

/*
func main() {
	fmt.Println("Robot Jazz v0.1")
	m := model.New()
	m.Load("resources/chords-v1.csv")
	dataItem, _ := m.Classify([]float64{1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0})

	// Prints top ten most relevant chords
	for i := 0; i < MinInt(len(dataItem.ClosestCentroids), 10); i++ {
		fmt.Println(dataItem.ClosestCentroids[i].Tag)
	}
}
*/
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

var mainModel *model.Model

func init() {
	mainModel = model.New()
	mainModel.Load("resources/chords-v1.csv")
}

func GetMatchingChords(notes []float64) (map[string][]int, error) {
	if len(notes) != 12 {
		return nil, errors.New("First input argument must be a float64 slice of length 12")
	}
	dataItem, err := mainModel.Classify(notes)
	if err != nil {
		return nil, err
	}
	topTen := make(map[string][]int)
	for i := 0; i < MinInt(len(dataItem.ClosestCentroids), 10); i++ {
		centroid := dataItem.ClosestCentroids[i]
		topTen[centroid.Tag] = ReducedNoteForm(centroid.Attributes)
	}
	return topTen, nil
}

func ReducedNoteForm(notes []float64) []int {
	res := make([]int, 0)
	for note, weight := range notes {
		if weight != 0.0 {
			res = append(res, note)
		}
	}
	return res
}

func ExtendedNoteForm(notes []int) []float64 {
	res := make([]float64, 12)
	for _, note := range notes {
		res[note] = 1.0
	}
	return res
}
