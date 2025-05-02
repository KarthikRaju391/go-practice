package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Hello, world!")
	ctx := context.WithValue(context.Background(), "foo", "bar")
	val, err := fetchResponse(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
	fmt.Println(time.Since(start))
}

type ResponseStruct struct {
	resultVal string
	errVal    error
}

func fetchResponse(ctx context.Context) (string, error) {
	fmt.Println("Here is the value from the context:", ctx.Value("foo"))
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	resCh := make(chan ResponseStruct)

	go func() {
		res, err := fetchSomethingWhichTakesTime()
		resCh <- ResponseStruct{
			resultVal: res,
			errVal:    err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("Request took longer time to process than expected.")
		case resp := <-resCh:
			return resp.resultVal, resp.errVal
		}
	}
}

func fetchSomethingWhichTakesTime() (string, error) {
	// Increase the timeout to be more than 200 milliseconds to see trigger an error
	time.Sleep(100 * time.Millisecond)

	return "Here lies the response", nil
}
