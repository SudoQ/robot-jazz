package model

import (
	"encoding/csv"
	"fmt"
	"os"
	"github.com/SudoQ/robot-jazz/data"
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
		for i := range v {
			attr, _ := strconv.ParseFloat(v[i], 64)
			attributes = append(attributes, attr)
		}
		centroids = append(centroids, data.New(attributes, k))
		//a0, _ := strconv.ParseFloat(v[0], 64)
		//a1, _ := strconv.ParseFloat(v[1], 64)
		//a2, _ := strconv.ParseFloat(v[2], 64)
		//centroids = append(centroids, data.New([]float64{a0, a1, a2}, k))
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

func (model *Model) Classify(attributes []float64) error {
	dataItem := data.New(attributes, len(model.Centroids))
	dataItem.UpdateClassification(model.Centroids)
	fmt.Println(dataItem.Classification)
	return nil
}
