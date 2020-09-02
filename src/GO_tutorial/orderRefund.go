package main

import (
    "fmt"
    "time"
)

func process(item string, _listner chan string, 
sleepTimer time.Duration){
    defer close(_listner)
    for i:= 1; i<= 5; i++ {
        time.Sleep(sleepTimer)
        _listner <- item
    }
}

func main(){
    orderListner := make(chan string)
    refundListner := make(chan string)

    go func(){
        process("order Processed", orderListner, time.Second/2)
    }()

    go func() {
        process("refund Processed", refundListner, time.Second)
    }()

    for i := 0; i<=10; i++{
        select{
            case msg := <-orderListner:
                fmt.Println(msg)
            case msg := <- refundListner:
                fmt.Println(msg)
        }
    }
}
