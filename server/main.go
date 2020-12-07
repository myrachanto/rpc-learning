package main
import(
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/http"
)
//Item ..
type Item struct{
	Title string
	Body string
}
var database []Item
//API ...
type API int
func main(){
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
		
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("listening error ", err)
	}
	log.Printf("seving rpc on port %d", 4040)
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal("error serving  ", err)
	}
// fmt.Println("initial db: ", database)
// a := Item{"primera", "el primera mensaje"}
// b := Item{"segundo", "el segundo mensaje"}
// c := Item{"tress", "el tress mensaje"}

// AddItem(a)
// AddItem(b)
// AddItem(c)
// fmt.Println("second db: ", database)

// DeleteItem(b)
// fmt.Println("second db: ", database)
// EditItem("tress", Item{"tress", "third item"})
// fmt.Println("second db: ", database)
}
//GetDB ..
func (a *API) GetDB(title string, reply *[]Item) error{
	fmt.Println("database called!")
	*reply = database
	return nil
}
//GetByName ..
func (a *API)GetByName(title string, reply *Item) error {
	var getItem Item
	 for _, i := range database {
		 if(i.Title == title){
			getItem = i
		 }
	 }
	 *reply = getItem
	 return nil	 
}
//AddItem ...
func (a *API)AddItem(item Item, reply *Item) error {
 database = append(database, item)
 *reply = item
 return nil
}
//EditItem ..
func (a *API)EditItem(edit Item, reply *Item) error {
var changed Item
for item, v := range database{
	if v.Title == edit.Title {
		database[item] = edit
		changed = edit
	}
}
*reply = changed
return nil
}
//DeleteItem ...
func (a *API)DeleteItem(item Item, reply *Item) error {
var del Item
for i, v := range database {
	 if v.Title == item.Title && v.Body == item.Body{
		 database = append(database[:i], database[i+1:] ...)
		 del = item
		 break
	 }
}
*reply = del
return nil
}