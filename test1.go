package main

import "fmt"
import "time"

func main() {
	ch := make(chan int, 8)
	ch <- 1
	ch <- 2
	count := 0
	for{
	select{
	case ch <- 3:
	case i:=<-ch:
        fmt.Println(i)
	default:
	time.Sleep(100 * time.Millisecond)
	}
	count++
	if count > 20 {
		break
	}
	}
	
	time.Sleep(10000 * time.Millisecond)

}