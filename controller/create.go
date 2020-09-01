package controller

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {

			artists := DecodeArrayMap("https://groupietrackers.herokuapp.com/api/artists")
			relation := DecodeMap("https://groupietrackers.herokuapp.com/api/relation")

			json.NewEncoder(w).Encode(artists)
			json.NewEncoder(w).Encode(relation)

			// index := relation["index"].([]interface{})

			// for i := range index {
			// 	if i == 0 {
			// 		json.NewEncoder(w).Encode(index[i])
			// 		fmt.Println(index[i])
			// 	}
			// }

			// var d map[string]interface{}

			// for i := range artists {
			// 	if i == 0 {
			// 		d = artists[i]
			// 		json.NewEncoder(w).Encode(d)
			// 		fmt.Println(d)
			// }
			// }
		}
	}
}
func DecodeArrayMap(url string) []map[string]interface{} {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var f []map[string]interface{}

	json.Unmarshal(data, &f)
	return f
}
func DecodeMap(url string) map[string]interface{} {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var f map[string]interface{}

	json.Unmarshal(data, &f)
	return f
}
// func NotR(i int) int {
// 	var nums []int
// 	nums = append(nums, i)

// 	for j := range nums {
// 		if i != nums[j] {
// 			return i
// 		}
// 	}
// 	fmt.Println(nums)
	
// }
