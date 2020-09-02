package main

import(
    "fmt"
    "log"
    proto "github.com/golang/protobuf/proto"
)

func main(){
    fmt.Println("hello");
    elliot := &Person{
        Name: "Elliot",
        Age: 24,
        SocialFollowers: &SocialFollowers{
            Twitter: 200,
            Youtube: 2500,
        },
    }

    data, err := proto.Marshal(elliot)
    if err != nil{
        log.Fatal("Marshalling error: ", err)
    }

    fmt.Println(data)

    newElliot := &Person{}
    err = proto.Unmarshal(data, newElliot)

    fmt.Println(newElliot.GetAge())
    fmt.Println(newElliot.SocialFollowers.GetTwitter())

}
