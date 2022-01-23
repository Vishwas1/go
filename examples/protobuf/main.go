package main

import (
	"fmt"
	"log"

	"go/examples/protobuf/demo"

	"google.golang.org/protobuf/proto"
)

func main() {
	vishwas := &demo.Person{
		Name: "Vishwas",
		Age:  24,
		Followers: &demo.Followers{
			Youtube: 1400,
			Twitter: 2330,
		},
	}
	fmt.Println("Actual message: ", vishwas)

	// Marshalling into proto
	data, err := proto.Marshal(vishwas)
	if err != nil {
		log.Fatal("Marshalling err : ", err)
	}

	fmt.Println("Serialised message: ", data)

	// Unmarshalling
	newVishwas := &demo.Person{}
	err = proto.Unmarshal(data, newVishwas)
	if err != nil {
		log.Fatal("UnMarshalling err : ", err)
	}

	fmt.Println("My name is", newVishwas.GetName(), "and my age is", newVishwas.GetAge())
	fmt.Println("I have", newVishwas.Followers.GetYoutube(), "followers in Youtube and", newVishwas.Followers.GetTwitter(), "in Twitter")
}
