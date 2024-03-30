
package main

import "fmt"
import "time"


func producer(c chan int){

    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("[producer]: pushing %d\n", i)
        // TODO: push real value to buffer
        select {
        case c <- i:
        }

        
    }

}



func consumer(c chan int){

    time.Sleep(1 * time.Second)
    for {
        i := 0 //TODO: get real value from buffer
        select {
        case i = <-c:
            fmt.Printf("[consumer]: %d\n", i)
            time.Sleep(50 * time.Millisecond)}
        }
    
}



func main(){

    buffer := make(chan int, 5)
    fmt.Println("Buffer created")

    go producer(buffer)
    go consumer(buffer)

    time.Sleep(10000 * time.Millisecond)
}