package main

import(
	"fmt"
	"time"
)

func main(){

	//var wg sync.WaitGroup
	//wg.Add(1)
	

	//go func(){
	//	count("lava")
	//	wg.Done()
	//}()

	//wg.Wait()

	c := make(chan string)

	go count("sheep", c)
	//for{
	//	msg, ope := <- c

	//	if !ope{
	//		break
	//	}
	//	fmt.Println(msg)
	//}

	for msg := range c{
		fmt.Println(msg)
	}
	
}

func count(thing string, c chan string){
	for i := 1;i<=5;i++ {
		//fmt.Println(i,
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)

	
}
