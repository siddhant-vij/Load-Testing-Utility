package tester

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/siddhant-vij/Load-Testing-Utility/pkg/common"
)

func LoadWords(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			return nil, error
		}
		words = append(words, line[0])
	}
	return words, nil
}

func RunTests(words []string, numTestingThreads int) []common.Result {
	wordsPerThread := len(words) / numTestingThreads
	results := make([]common.Result, 0)

	for numServerThreads := 1; numServerThreads <= 25; numServerThreads++ {
		var totalLatency time.Duration
		var totalRequests int
		var mutex sync.Mutex
		var wg sync.WaitGroup

		startTime := time.Now()

		for i := 0; i < numTestingThreads; i++ {
			wg.Add(1)
			go func(threadID, startIdx int) {
				defer wg.Done()
				for j := startIdx; j < startIdx+wordsPerThread; j++ {
					start := time.Now()
					url := fmt.Sprintf("http://localhost:8080/frequency?word=%s&numServerThreads=%d", words[j], numServerThreads)

					resp, err := http.Get(url)
					if err != nil {
						fmt.Printf("Thread %d: Error making request: %v\n", threadID, err)
						continue
					}
					_, err = io.ReadAll(resp.Body)
					if err != nil {
						fmt.Printf("Thread %d: Error reading response body: %v\n", threadID, err)
						continue
					}
					resp.Body.Close()

					mutex.Lock()
					totalLatency += time.Since(start)
					totalRequests++
					mutex.Unlock()
				}
			}(i, i*wordsPerThread)
		}

		wg.Wait()

		endTime := time.Now()
		testDuration := endTime.Sub(startTime)

		avgLatency := float64(totalLatency.Milliseconds()) / float64(totalRequests)
		throughput := float64(totalRequests) / testDuration.Seconds()

		result:= common.Result{
			NumTestingThreads: numTestingThreads,
			NumServerThreads:  numServerThreads,
			Latency:           avgLatency,
			Throughput:        throughput,
		}

		fmt.Printf("Test Result: %+v\n", result)
		results = append(results, result)
	}

	return results
}
