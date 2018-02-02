
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

)

// A Response struct to map the Entire Response
type Response struct {
	Page    int    `json:"totalResults"`
	Beer []Beer `json:"data"`
}

type ResponseB struct {
	BeerB []BeerB `json:"data"`
}

type BeerB struct {

	EntryNo string `json:"id"`
	Name string `json:"name"`
	Descri string `json:"description"`
	Year string `json:"established"`

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

// A struct to map our Pokemon's Species which includes it's name


func main() {
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

	for i := 0; i < len(responseObject.Beer); i++ {
		verificaString(responseObject.Beer[i].EntryNo)
		verificaString(responseObject.Beer[i].Name)
		verificaNum(responseObject.Beer[i].ABV)
		verificaNum(responseObject.Beer[i].IBU)
		verificaString(responseObject.Beer[i].IsOrganic)
		verificaString(responseObject.Beer[i].Status)
		verificaString(responseObject.Beer[i].Data)
	}
	//------------------------------------------------------//





	for i := 0; i < len(responseObject.Beer); i++ {
		var link string = "http://api.brewerydb.com/v2/beer/"
		var codigo string = responseObject.Beer[i].EntryNo
		var rlink string = "/breweries/?key=8c8de43d4fb516428fc3baba71e8eca6"
		out := fmt.Sprint(link, codigo, rlink)
		fmt.Println(out)

		responseb, err := http.Get(out)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseDatab, err := ioutil.ReadAll(responseb.Body)
		if err != nil {
			log.Fatal(err)
		}

		var responseObjectb ResponseB
		json.Unmarshal(responseDatab, &responseObjectb)

		for j := 0; j < len(responseObjectb.BeerB); j++ {
			verificaString(responseObjectb.BeerB[j].EntryNo)
			//fmt.Println(responseObjectb.BeerB[j].EntryNo)
			verificaString(responseObjectb.BeerB[j].Name)
			//fmt.Println(responseObjectb.BeerB[j].Name)
			verificaString(responseObjectb.BeerB[j].Descri)
			//fmt.Println(responseObjectb.BeerB[j].Descri)
			verificaNum(responseObjectb.BeerB[j].Year)
			//fmt.Println(responseObjectb.BeerB[j].Year)


		}



	}



}

	func verificaString(text string){

		_, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
		}else{
			fmt.Printf("data error: expected only string")
			fmt.Print(" - ")
			fmt.Println(text)
			os.Exit(0)
		}
	}

	func verificaNum(text string){

		if text != "" {

			_, err := strconv.ParseFloat(text, 64)
			if err != nil {
				fmt.Printf("data error: expected only numbers")
				fmt.Printf(" - ")
				fmt.Println(text)
				os.Exit(0)
			}else{
			}

		}

	}

