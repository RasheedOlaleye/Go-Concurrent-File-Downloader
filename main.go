package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func readURLs(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url != "" {
			urls = append(urls, url)
		}
	}

	return urls, scanner.Err()
}

func downloadFile(url, fileName string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("❌ %s failed: %v", fileName, err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(filepath.Join("downloads", fileName))
	if err != nil {
		results <- fmt.Sprintf("❌ %s failed: %v", fileName, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		results <- fmt.Sprintf("❌ %s failed: %v", fileName, err)
		return
	}

	results <- fmt.Sprintf("✅ %s downloaded successfully", fileName)
}

func main() {
	// Create downloads folder if it doesn't exist
	if err := os.MkdirAll("downloads", 0755); err != nil {
		log.Fatal(err)
	}

	urls, err := readURLs("url.txt")
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan string)
	var wg sync.WaitGroup

	wg.Add(len(urls))

	for i, url := range urls {
		fileName := fmt.Sprintf("file%d.pdf", i+1)
		go downloadFile(url, fileName, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("\n🎉 All downloads completed.")
}