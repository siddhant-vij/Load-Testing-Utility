package common

type Result struct {
	NumTestingThreads int
	NumServerThreads  int
	Latency           float64
	Throughput        float64
}
