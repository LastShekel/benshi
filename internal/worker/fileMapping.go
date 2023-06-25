package worker

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

// mapFiles mapping all files from context and saving result to files
func mapFiles(ctx context.Context) {
	//M int, files []string
	M := ctx.Value("M").(int)
	taskId := ctx.Value("taskId").(string)
	files := ctx.Value("files").([]string)
	for i := 0; i < M; i++ {
		filename := path.Join(
			ctx.Value("folder").(string),
			fmt.Sprintf("mr-%s-%d", taskId, i),
		)
		file, err := os.Create(filename)
		if err != nil {
			return
		}
		err = file.Close()
		if err != nil {
			return
		}
		ctx = context.WithValue(ctx, filename, new(sync.Mutex))

	}
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			toBucket(ctx, file)
		}(file)
	}
	wg.Wait()
	log.Println("All files processed")
	values := map[string]string{"url": ctx.Value("workerUrl").(string)}
	_, err := SendJson(ctx.Value("driverUrl").(string)+"/done", values)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Failed to send done message")
	}
}

// toBucket selecting and saving data to file
func toBucket(ctx context.Context, data string) {
	log.Printf("Processing file\n")
	taskId := ctx.Value("taskId").(string)
	M := ctx.Value("M").(int)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			word = clearString(word)
			if len(word) == 0 {
				continue
			}
			bucketId := int(word[0]) % M
			filename := path.Join(
				ctx.Value("folder").(string),
				fmt.Sprintf("mr-%s-%d", taskId, bucketId),
			)
			writeToFile(filename, word, ctx.Value(filename).(*sync.Mutex))
		}
	}
	log.Printf("File processed\n")
}

// toBucket writes exact word to file
func writeToFile(filename string, word string, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	w := bufio.NewWriter(file)
	_, err = w.Write([]byte(word))
	if err != nil {
		return
	}
	_, err = w.WriteRune('\n')
	if err != nil {
		return
	}
	err = w.Flush()
	if err != nil {
		return
	}
}

// clearString clears word from extra symbols
func clearString(str string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9\- ]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
