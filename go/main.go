package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func isJSON(s string) bool {
    var js map[string]interface{}
    return json.Unmarshal([]byte(s), &js) == nil
}

func main() {
	fptr := flag.String("file", "test.txt", "file path to read from")
	uptr := flag.Int("update", 1000000, "How many lines should be checked before printing an update")
    flag.Parse()

    f, err := os.Open(*fptr)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = f.Close(); err != nil {
        	log.Fatal(err)
    	}
	}()
	
	s := bufio.NewScanner(f)

	lineNumber := 1
    for s.Scan() {
		if isJSON(s.Text()) == false {
			fmt.Printf("Error: line %v\n", lineNumber)
		}
		lineNumber++
		if lineNumber % *uptr == 0 {
			fmt.Printf("%v lines processed\n", lineNumber)
		}
    }
    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}