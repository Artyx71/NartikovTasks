package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println("Usage: go run main.go <file1> [<file2>] [<resultFile>]")
        return
    }

    var contents []string
    for _, file := range args[:len(args)] {
        content, err := ioutil.ReadFile(file)
        if err != nil {
            fmt.Printf("Error reading file %s: %v\n", file, err)
            return
        }
        contents = append(contents, string(content))
    }

    result := strings.Join(contents, " ")

    if len(args) == 1 {
        fmt.Println(result)
    } else if len(args) >= 2 {
        fmt.Println(result)
        if len(args) == 3 {
            resultFile := args[2]
            err := ioutil.WriteFile(resultFile, []byte(result), 0644)
            if err != nil {
                fmt.Printf("Error writing to file %s: %v\n", resultFile, err)
                return
            }
        }
    }
}