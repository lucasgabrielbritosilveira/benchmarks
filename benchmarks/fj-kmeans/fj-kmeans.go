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
	// Defina um limite baseado no tamanho da memória
	const sizeLimit = (256 / 2) * 1024 // metade do cache L2, similar ao Java
	elementSize := dimension*8 + 16    // Tamanho aproximado de um elemento
	forkThreshold := sizeLimit / elementSize

	return &KMeans{
		Dimension:      dimension,
		ForkThreshold:  forkThreshold,
		ClusterCount:   clusterCount,
		IterationCount: iterationCount,
	}
}

// Função principal para executar o K-Means
func (km *KMeans) Run(data [][]float64) [][]float64 {
	centroids := km.randomSample(data, km.ClusterCount)
	for i := 0; i < km.IterationCount; i++ {
		clusters := km.assignClusters(data, centroids)
		centroids = km.updateCentroids(clusters)
	}
	return centroids
}

// Seleciona aleatoriamente `clusterCount` centroides
func (km *KMeans) randomSample(data [][]float64, clusterCount int) [][]float64 {
	if len(data) == 0 {
		panic("randomSample: data is empty")
	}

	rand.Seed(100)
	sample := make([][]float64, clusterCount)

	for i := 0; i < clusterCount; i++ {
		index := rand.Intn(len(data)) // Safe random index
		sample[i] = data[index]
	}

	return sample
}

// Atribui cada ponto ao cluster mais próximo
func (km *KMeans) assignClusters(data [][]float64, centroids [][]float64) map[int][][]float64 {
	clusters := make(map[int][][]float64)
	var wg sync.WaitGroup
	dataSize := len(data)
	taskSize := km.ForkThreshold

	// Cria goroutines para paralelizar a atribuição
	for start := 0; start < dataSize; start += taskSize {
		end := start + taskSize
		if end > dataSize {
			end = dataSize
		}
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				closest := km.findNearestCentroid(data[i], centroids)
				clusters[closest] = append(clusters[closest], data[i])
			}
		}(start, end)
	}
	wg.Wait()
	return clusters
}

// Calcula o centroide mais próximo
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

// Atualiza os centroides calculando a média dos clusters
func (km *KMeans) updateCentroids(clusters map[int][][]float64) [][]float64 {
	centroids := make([][]float64, km.ClusterCount)
	var wg sync.WaitGroup

	for clusterID, points := range clusters {
		wg.Add(1)
		go func(clusterID int, points [][]float64) {
			defer wg.Done()
			centroids[clusterID] = computeAverage(points, km.Dimension)
		}(clusterID, points)
	}
	wg.Wait()
	return centroids
}

// Calcula a média de um conjunto de pontos
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

// Distância Euclidiana
func euclideanDistance(x, y []float64) float64 {
	sum := 0.0
	for i := range x {
		diff := x[i] - y[i]
		sum += diff * diff
	}
	return sum // Sem sqrt para manter compatibilidade com o código Java
}

// Gera dados aleatórios para teste
func generateData(count, dimension, clusterCount int) [][]float64 {

	if count <= 0 || dimension <= 0 || clusterCount <= 0 {
		panic("generateData: invalid parameters")
	}

	data := make([][]float64, count)
	rand.New(rand.NewSource(time.Now().UnixNano()))
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
	data := generateData(1000, dimension, clusterCount)

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
