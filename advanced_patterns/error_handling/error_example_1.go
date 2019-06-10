package main

import (
	"fmt"
	"net/http"
)

func errorExample() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)

		go func() {
			defer fmt.Println("Closed responses")
			defer close(responses)

			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}

				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()

		return responses
	}

	urls := []string{"http://www.google.com", "http://badhost"}
	done := make(chan interface{})

	defer fmt.Printf("Closing done chan")
	defer close(done)

	responses := checkStatus(done, urls...)

	for resp := range responses {
		fmt.Printf("Response: %v \n", resp.Status)
	}
}

func errorExampleFixed() {
	type Result struct {
		Error    error
		Response *http.Response
	}

	checkStatus := func(done chan interface{}, urls ...string) <-chan Result {
		resultChan := make(chan Result)

		go func() {
			defer close(resultChan)

			for _, url := range urls {
				resp, err := http.Get(url)
				result := Result{Response: resp, Error: err}

				select {
				case <-done:
					return
				case resultChan <- result:
				}
			}
		}()

		return resultChan
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"http://www.google.com", "http://badhost"}

	responses := checkStatus(done, urls...)

	for result := range responses {
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v \n", result.Response.Status)
	}
}

func main() {
	// errorExample()
	errorExampleFixed()
}
