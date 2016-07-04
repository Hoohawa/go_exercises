// Run exaple
// $ go run ex_1_10.go http://google.com http://microsoft.com http://facebook.com
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const nruns = 10

var result = make(map[string][]float64)
var sucess = make(map[string][]bool)

func main() {
	start := time.Now()
	urls := os.Args[1:]

	fetchAll(urls)
	saveResultsToFile(urls)

	fmt.Printf("Completed 10 runs after %v seconds\n", time.Since(start).Seconds())
}

func fetchAll(urls []string) {
	for i := 0; i < nruns; i++ {
		ch := make(chan bool)
		for _, url := range urls {
			go fetch(url, ch)
		}
		for range urls {
			<-ch
		}
	}
}

func fetch(url string, done chan<- bool) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		sucess[url] = append(sucess[url], false)
		done <- true
		return
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		sucess[url] = append(sucess[url], false)
		done <- true
		return
	}
	secs := time.Since(start).Seconds()
	result[url] = append(result[url], secs)
	fmt.Printf("Downloaded %s after %.2f\n", url, secs)
	done <- true
}

func saveResultsToFile(urls []string) {
	logFile, err := os.Create("ex_1_10.log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open ex_1_10.log: %v\n", err)
	}
	logFile.WriteString(strings.Join(urls, "\t") + "\n") // Print header line
	for i := 0; i < nruns; i++ {                         // Print individual runs
		runTimes := []string{}
		for _, url := range urls {
			runTimes = append(runTimes, fmt.Sprintf("%.4f", result[url][i]))
		}
		logFile.WriteString(strings.Join(runTimes, "\t") + "\n")
	}
	logFile.Sync()
	logFile.Close()
}
