package main

import(
    "fmt"
    "time"
)

func savage() {
    arr1 := [3]string{"MF", "FU", "BAZINGA!!!"}

    for x1:=0 ;x1<3; x1++{
        time.Sleep(150 * time.Millisecond)
        fmt.Printf("%s\n", arr1[x1])
    }
}

func winto(){
    arr2 := [6]int{1,2,3,4,5,6}

    for x2:=0; x2< 5; x2++{
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("%d\n", arr2[x2])
    }
}

func main(){

    fmt.Println("Strat")
    go savage()

    go winto()
    time.Sleep(499 * time.Millisecond)
    fmt.Println("END")
}
