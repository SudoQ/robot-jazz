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
	return nil
}

/*
func (model *Model) Save(filename string) error {
	csvfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	lines := make([][]string, 0)
	for _, c := range model.Centroids {
		line := make([]string, 12)
		for j, a := range c.Attributes {
			line[j] = fmt.Sprintf("%d.0", uint8(a))
		}
		lines = append(lines, line)
	}
	err = writer.WriteAll(lines)
	if err != nil {
		return err
	}
	return nil
}
*/

func (model *Model) Classify(attributes []float64) (*data.Data, error) {
	dataItem := data.New(attributes, len(model.Centroids), "")
	dataItem.UpdateClassification(model.Centroids)
	return dataItem, nil
}
