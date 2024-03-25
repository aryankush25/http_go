package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")

	fmt.Println("err", err)
	fmt.Println("res", res)

	if err != nil {
		log.Fatal(err)
	}

	lw := logWriter{}

	io.Copy(lw, res.Body)

	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()

	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s", body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
