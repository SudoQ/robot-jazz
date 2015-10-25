package model

import (
	"github.com/SudoQ/robotjazz/data"
)

type Model struct {
	Centroids []*data.Data
}

func New() *Model {
	return &Model{
		Centroids: make([]*data.Data, 0),
	}
}

func (model *Model) AddCentroid(dataItem *data.Data) error {
	model.Centroids = append(model.Centroids, dataItem)
	return nil
}

func (model *Model) Classify(attributes []float64) (*data.Data, error) {
	dataItem := data.New(attributes, len(model.Centroids), "")
	dataItem.UpdateClassification(model.Centroids)
	return dataItem, nil
}

func (model *Model) UpdateCentroids() error {
	for _, centroid := range model.Centroids {
		centroid.UpdateClassification(model.Centroids)
	}
	return nil
}
