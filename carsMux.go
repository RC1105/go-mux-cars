package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Car struct {
	Name  string
	Price int
	Brand string
}

var ls []Car = []Car{}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", funcGet)
	mux.HandleFunc("GET /all", getCars)

	mux.HandleFunc("POST /add/{com}/{car}/{price}", postInsert)
	mux.HandleFunc("POST /del/{car}", postDel)
	mux.HandleFunc("POST /upd/{car}/{price}", postUpd)
	log.Fatal(http.ListenAndServe(":4000", mux))
}
func funcGet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello Guys")

}
func getCars(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ls)
}
func postDel(w http.ResponseWriter, r *http.Request) {
	car := r.PathValue("car")
	for i := 0; i < len(ls); i++ {
		if ls[i].Name == car {
			ls = append(ls[:i], ls[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(ls)


}
func postUpd(w http.ResponseWriter, r *http.Request) {
	car := r.PathValue("car")
	pr, _ := strconv.Atoi(r.PathValue("price"))
	for i := 0; i < len(ls); i++ {
		if ls[i].Name == car {
			ls[i].Price = pr
			break
		}
	}
	json.NewEncoder(w).Encode(ls)

}

func postInsert(w http.ResponseWriter, r *http.Request) {

	pr, _ := strconv.Atoi(r.PathValue("price"))
	car := r.PathValue("car")
	com := r.PathValue("com")
	c := Car{
		Name:  car,
		Price: pr,
		Brand: com,
	}
	fmt.Print(c)
	ls = append(ls, c)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(ls)
}
