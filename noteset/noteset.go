package noteset

type Noteset struct {
	id string
	root int
	noteWeights []float64
	patternId string
	patternNotes []int
}

func New(	id string,
					root int,
					noteWeights []float64,
					patternId string,
					patternNotes []int) (*Noteset) {
	return &Noteset{
		id: id,
		root: root,
		noteWeights: noteWeights,
		patternId: patternId,
		patternNotes: patternNotes,
	}
}
