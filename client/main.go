package main

import(
	"log"
	"fmt"
	"net/rpc"
)
//Item ..
type Item struct{
	Title string
	Body string
}
func main(){
	var reply Item
	var db []Item
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("error registering API", err)
		
	}
	a := Item{"primera", "the first message"}
	b := Item{"segundo", "el segundo mensaje"}
	c := Item{"tress", "el tress mensaje"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("DB: ", db)

}