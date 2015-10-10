package main

import (
	"fmt"
	"github.com/SudoQ/robot-jazz/model"
	//"github.com/SudoQ/robot-jazz/data"
)

func MinInt(x,y int) int{
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println("Robot Jazz v0.1")
	m := model.New()
	m.Load("resources/chords-v1.csv")
	dataItem, _ := m.Classify([]float64{1.0,0.0,0.0,0.0,1.0,0.0,0.0,1.0,0.0,0.0,0.0,0.0})


	for i := 0; i < MinInt(len(dataItem.ClosestCentroids), 10); i++ {
		fmt.Println(dataItem.ClosestCentroids[i].Tag)
	}
}
