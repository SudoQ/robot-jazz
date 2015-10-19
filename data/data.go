package data

import (
	"math"
)

type Data struct {
	Tag              string
	Attributes       []float64
	Distances        []float64
	Classification   int
	ClosestCentroids []*Data
}

func New(attr []float64, numClass int, tag string) *Data {
	return &Data{
		Tag:              tag,
		Attributes:       attr,
		Distances:        make([]float64, numClass),
		Classification:   0,
		ClosestCentroids: make([]*Data, 0),
	}
}

func (data *Data) updateDistances(centroids []*Data) {
	for i, centroid := range centroids {
		sumOfSquares := 0.0
		for j := range centroid.Attributes {
			sumOfSquares += math.Pow(data.Attributes[j]-centroid.Attributes[j], 2)
		}
		distance := math.Sqrt(sumOfSquares)
		data.Distances[i] = distance
	}
}

func (data *Data) UpdateClassification(centroids []*Data) {
	data.updateDistances(centroids)

	closestCentroids := make([]int, 0)
	for centroidId, distance := range data.Distances {
		if len(closestCentroids) == 0 {
			closestCentroids = append(closestCentroids, centroidId)
		} else {
			for i, compareId := range closestCentroids {
				currentDistance := data.Distances[compareId]
				if distance < currentDistance {

					// The following three lines is a slice insert
					closestCentroids = append(closestCentroids, 0)
					copy(closestCentroids[i+1:], closestCentroids[i:])
					closestCentroids[i] = centroidId

					break
				}
			}
		}
	}
	for _, closeCentroidId := range closestCentroids {
		data.ClosestCentroids = append(data.ClosestCentroids, centroids[closeCentroidId])
	}
}

func (data *Data) Waverage(item *Data, weigth float64) {
	for i := range data.Attributes {
		data.Attributes[i] = data.Attributes[i]*(1.0-weigth) + item.Attributes[i]*weigth
	}
}
