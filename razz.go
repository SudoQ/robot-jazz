package main

import (
	"fmt"
	"github.com/SudoQ/robot-jazz/model"
)

func main() {
	fmt.Println("Robot Jazz v0.1")
	m := model.New()
	m.Load("resources/default.csv")
	m.Classify([]float64{1.0,0.0,0.0,0.0,1.0,0.0,0.0,1.0,0.0,0.0,0.0,0.0})
}
