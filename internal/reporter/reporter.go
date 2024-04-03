package reporter

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/siddhant-vij/Load-Testing-Utility/pkg/common"
)

func ReportResults(results []common.Result) {
	for _, result := range results {
		saveResultToCSV(result)
	}
}

func saveResultToCSV(result common.Result) {
	file, err := os.OpenFile("resources/results.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening results.csv: %v", err)
		return
	}
	
	defer file.Close()
	
	writer := csv.NewWriter(file)
	defer writer.Flush()

	if stat, err := file.Stat(); err == nil && stat.Size() == 0 {
		writer.Write([]string{"NumTestingThreads", "NumServerThreads", "Latency", "Throughput"})
	}

	record := []string{
		fmt.Sprintf("%d", result.NumTestingThreads),
		fmt.Sprintf("%d", result.NumServerThreads),
		fmt.Sprintf("%.2f", result.Latency),
		fmt.Sprintf("%.2f", result.Throughput),
	}
	if err := writer.Write(record); err != nil {
		fmt.Printf("Error writing to results.csv: %v", err)
	}
}
