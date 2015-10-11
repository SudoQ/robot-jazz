package util

var chrom = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}

func GetNoteName(note int) string {
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

func GetNoteValue(note string) int {
	return noteValue[note]
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
