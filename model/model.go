package model

import (
	"encoding/csv"
	"github.com/SudoQ/robotjazz/data"
	"os"
	"strconv"
)

type Model struct {
	Centroids []*data.Data
}

func New() *Model {
	return &Model{
		Centroids: make([]*data.Data, 0),
	}
}

func (model *Model) Load(filename string) error {
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
	centroids := make([]*data.Data, 0)
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
		centroids = append(centroids, data.New(attributes, k, tag))
	}
	model.Centroids = centroids
	model.UpdateCentroids()
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
