package controller

import (
	"html/template"
	"math/rand"
	"net/http"
	// "encoding/json"
)

func tracker() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			templates := template.Must(template.ParseGlob("view/*.html"))

			artists := DecodeArrayMap("https://groupietrackers.herokuapp.com/api/artists")

			var d map[string]interface{}
			j := 0
			i := 0
			nums := []int{}

			for range artists {
				i = rand.Intn(52)
				if j < 8 {
					if Valid(i, nums) == true {
						nums = append(nums, i)
						d = artists[i]
						templates.ExecuteTemplate(w, "templates.html", d)
						j++
					}
				}

			}

		}

	}
}

func Valid(i int, nums []int) bool {
	for j := range nums {
		if i == nums[j] {
			return false
		}
	}
	return true
}
