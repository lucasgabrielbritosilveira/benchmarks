package fjkmeans

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type KMeans struct {
	Dimension      int
	ForkThreshold  int
	ClusterCount   int
	IterationCount int
}

func NewKMeans(dimension int, clusterCount int, iterationCount int) *KMeans {
	// Define a memory-based limit
	const sizeLimit = (256 / 2) * 1024 // Half of L2 cache, similar to Java
	elementSize := dimension*8 + 16    // Approximate size of an element
	forkThreshold := sizeLimit / elementSize

	return &KMeans{
		Dimension:      dimension,
		ForkThreshold:  forkThreshold,
		ClusterCount:   clusterCount,
		IterationCount: iterationCount,
	}
}

func (km *KMeans) Run(data [][]float64) [][]float64 {
	centroids := km.randomSample(data, km.ClusterCount)
	for i := 0; i < km.IterationCount; i++ {
		clusters := km.assignClusters(data, centroids)
		centroids = km.updateCentroids(data, clusters)
	}
	return centroids
}

func (km *KMeans) randomSample(data [][]float64, clusterCount int) [][]float64 {
	if len(data) == 0 {
		panic("randomSample: data is empty")
	}

	rand.Seed(time.Now().UnixNano())
	sample := make([][]float64, 0, clusterCount)
	seen := make(map[int]bool)

	for len(sample) < clusterCount {
		index := rand.Intn(len(data))
		if !seen[index] {
			seen[index] = true
			sample = append(sample, data[index])
		}
	}

	return sample
}

func (km *KMeans) assignClusters(data [][]float64, centroids [][]float64) map[int][][]float64 {
	clusters := make(map[int][][]float64)
	var mu sync.Mutex
	var wg sync.WaitGroup
	dataSize := len(data)
	taskSize := km.ForkThreshold

	for start := 0; start < dataSize; start += taskSize {
		end := start + taskSize
		if end > dataSize {
			end = dataSize
		}
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localClusters := make(map[int][][]float64)
			for i := start; i < end; i++ {
				closest := km.findNearestCentroid(data[i], centroids)
				localClusters[closest] = append(localClusters[closest], data[i])
			}
			mu.Lock()
			for k, v := range localClusters {
				clusters[k] = append(clusters[k], v...)
			}
			mu.Unlock()
		}(start, end)
	}
	wg.Wait()
	return clusters
}

func (km *KMeans) findNearestCentroid(point []float64, centroids [][]float64) int {
	minDist := math.MaxFloat64
	closest := -1
	for i, centroid := range centroids {
		dist := euclideanDistance(point, centroid)
		if dist < minDist {
			minDist = dist
			closest = i
		}
	}
	return closest
}

func (km *KMeans) updateCentroids(data [][]float64, clusters map[int][][]float64) [][]float64 {
	centroids := make([][]float64, km.ClusterCount)
	var wg sync.WaitGroup

	for clusterID := 0; clusterID < km.ClusterCount; clusterID++ {
		wg.Add(1)
		go func(clusterID int) {
			defer wg.Done()
			points := clusters[clusterID]
			if len(points) == 0 {
				centroids[clusterID] = data[rand.Intn(len(data))] // Reassign randomly
			} else {
				centroids[clusterID] = computeAverage(points, km.Dimension)
			}
		}(clusterID)
	}
	wg.Wait()
	return centroids
}

func computeAverage(points [][]float64, dimension int) []float64 {
	sum := make([]float64, dimension)
	for _, point := range points {
		for i := 0; i < dimension; i++ {
			sum[i] += point[i]
		}
	}
	for i := 0; i < dimension; i++ {
		sum[i] /= float64(len(points))
	}
	return sum
}

func euclideanDistance(x, y []float64) float64 {
	sum := 0.0
	for i := range x {
		diff := x[i] - y[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func generateData(count, dimension, clusterCount int) [][]float64 {
	if count <= 0 || dimension <= 0 || clusterCount <= 0 {
		panic("generateData: invalid parameters")
	}

	data := make([][]float64, count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		point := make([]float64, dimension)
		for j := 0; j < dimension; j++ {
			point[j] = rand.Float64() + float64((i+j)%clusterCount)
		}
		data[i] = point
	}
	return data
}

func Run() {
	dimension := 5
	clusterCount := 5
	iterationCount := 5
	loopCount := 5
	data := generateData(100000, dimension, clusterCount)

	kmeans := NewKMeans(dimension, clusterCount, iterationCount)

	var tmp [][]float64
	for i := 0; i < loopCount; i++ {
		tmp = kmeans.Run(data)
	}
	centroids := tmp

	fmt.Println("Final Centroids:")
	for _, centroid := range centroids {
		fmt.Println(centroid)
	}
}
