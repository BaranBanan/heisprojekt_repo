package main

import (
    "fmt"
    
)

func main() {
    i:=0

    // Create a channel to signal completion
    done := make(chan bool)

    // Spawn thread_1
    go func() {
        for j := 0; j < 1000000; j++ {
            i++
        }
        done <- true
    }()

    // Spawn thread_2
    go func() {
        for j := 0; j < 1000000; j++ {
            i--
        }
        done <- true
    }()

    // Wait for both goroutines to finish
    <-done
    <-done

    // Print i
    fmt.Println(i)
    

    
}
