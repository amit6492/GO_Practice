package main

import (
    "fmt"
    "time"
)

func display(str string){
    for x:=0; x<6; x++ {
        time.Sleep(1 * time.Second) 
        fmt.Println(str)
    }
}

func  main(){

    go display("welcome")

    display("something")
}


