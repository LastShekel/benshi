package worker

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"
)

// reduceFiles main function for reduce task. Estimating wordcounts and saving it to file.
func reduceFiles(ctx context.Context) {
	taskId := ctx.Value("taskId").(string)
	workerUrl := ctx.Value("workerUrl").(string)
	driverUrl := ctx.Value("driverUrl").(string) + "/done"
	entries, err := os.ReadDir(intermediatePath)
	if err != nil {
		log.Fatal(err)
	}
	var counts = make(map[string]int)
	log.Printf("Reducing task %s\n", taskId)
	for _, entry := range entries {
		if getBucketId(entry.Name()) == taskId {
			filename := path.Join(
				ctx.Value("folder").(string),
				entry.Name(),
			)
			file, err := os.ReadFile(filename)
			if err != nil {
				return
			}
			log.Printf("Processing file %s\n", filename)
			countWords(counts, file)
		}
	}
	filename := path.Join(
		ctx.Value("out").(string),
		"out-"+taskId,
	)
	log.Printf("Writing to %s\n", filename)
	for k, v := range counts {
		writeToFile(filename, fmt.Sprintf("%s %d", k, v), new(sync.Mutex))
	}
	log.Println("Reduce task processed")
	values := map[string]string{"url": workerUrl}
	_, err = SendJson(driverUrl, values)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Failed to send done message")
	}
}

// getBucketId returning bucketId for file
func getBucketId(entry string) string {
	flag := false
	for i, s := range entry {
		if s == '-' {
			if flag {
				return entry[i+1:]
			}
			flag = true
		}
	}
	return ""
}

// countWords adding word counts to existing counter
func countWords(wordsCount map[string]int, data []byte) {
	words := strings.Split(string(data), "\n")
	for _, word := range words {
		_, present := wordsCount[word]
		if present {
			wordsCount[word] += 1
		} else {
			wordsCount[word] = 1
		}
	}
}
