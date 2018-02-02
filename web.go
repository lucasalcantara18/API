/*
package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"encoding/json"
	"os"
	"io/ioutil"
	"strconv"
	"math/rand"
	"time"
)

type Response struct {
	Page    int    `json:"totalResults"`
	Beer []Beer `json:"data"`
}

// A Pokemon Struct to map every pokemon to.
type Beer struct {

	EntryNo string `json:"id"`
	Name string `json:"name"`
	ABV string `json:"abv"`
	IBU string `json:"ibu"`
	IsOrganic string `json:"isOrganic"`
	Status string `json:"status"`
	Data string `json:"createDate"`


}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // parse arguments, you have to call this by yourself
	fmt.Println(r.Form)  // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//fmt.Fprintf(w, "hahahahaha!") // send data to client side

	///---------------------------------------------------///
	response, err := http.Get("http://api.brewerydb.com/v2/beers/?key=8c8de43d4fb516428fc3baba71e8eca6&name=****")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Page)
	fmt.Println(len(responseObject.Beer))


	var prob int
	var contI int = 45
	var contS int = 5

	for i := 0; i < len(responseObject.Beer); i++ {

		if contS > 0 {
			prob = rand.Intn(1000)
		}else{
			prob = 0
		}


		fmt.Println(prob)
		if  prob%2 == 0 {
			if(contI > 0){
				fmt.Fprintf(w, strconv.Itoa(i))
				fmt.Fprintf(w, " - ")
				if responseObject.Beer[i].ABV != "" {
					fmt.Fprintln(w, responseObject.Beer[i].ABV)
				}else if responseObject.Beer[i].IBU != ""{
					fmt.Fprintln(w, responseObject.Beer[i].IBU)
				}
				contI--
			}
		}else if  prob%2 != 0 {
			if(contS > 0){
				fmt.Fprintf(w, strconv.Itoa(i))
				fmt.Fprintf(w, " - ")
				fmt.Fprintln(w, responseObject.Beer[i].Name)
				contS--
			}
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", sayhelloName) // set router
	err := http.ListenAndServe(":9126", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
*/