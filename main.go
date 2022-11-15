package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string // `json:"Title"`
	Desc    string // `json:"Desc"`
	Content string // `json:"Content"`
}

type Articles []Article

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Article 1", Desc: "Description 1", Content: "Content 1"},
		Article{Title: "Article 2", Desc: "Description 2", Content: "Content 2"},
		Article{Title: "Article 3", Desc: "Description 3", Content: "Content 3"},
		Article{Title: "Article 4", Desc: "Description 4", Content: "Content 4"},
	}

	articles = append(articles, Article{Title: "Article 5", Desc: "Description 5", Content: "Content 5"})
	fmt.Println("Article Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func currency(w http.ResponseWriter, r *http.Request) {
	nested := map[string]map[string]int{
		"Great Britain Pound": {"GBP": 1},
		"Euro":                {"EUR": 2},
		"USA Dollar":          {"USD": 3},
	}

	fmt.Println("Currency Endpoint")
	json.NewEncoder(w).Encode(nested)
	// return nested
}

func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/currency", currency)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
	// handleRequest()

	s := []int{10, 20, 30, 40}

	// We can loop through this slice in two ways:

	// 1. using "range"
	for key, value := range s {
		fmt.Println(key, value)
	}
	//If we dont want the key, we do, replace "key" with "_":
	for _, value := range s {
		fmt.Println(value)
	}

	// 2. Using traditional forloop:
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) //get the value at index "i"
	}

	currency := map[string]string{
		"AUD": "Australia Dollar",
		"GBP": "Great Britain Pound",
		"JPY": "Japan Yen",
		"CHF": "Switzerland Franc",
	}

	//a. Adding to the map:
	currency["USD"] = "USA Dollar"

	fmt.Println("Currency with USD added: ", currency)

	//b. Remove from the map:
	delete(currency, "GBP")
	fmt.Println("Currency with GBP deleted: ", currency)

	//c. Replacing one entry with another:
	currency["AUD"] = "New Zealand Dollar"
	fmt.Println("Currency with AUD value replaced with NZD: ", currency)

	//Ranging through the map:
	for key, value := range currency {
		fmt.Printf("%v might be equal to: %v\n", key, value)
	}

	nested := map[string]map[string]int{
		"Great Britain Pound": {"GBP": 1},
		"Euro":                {"EUR": 2},
		"USA Dollar":          {"USD": 3},
	}

	for key, value := range nested {
		fmt.Printf("Currency Name: %v\n", key)
		for k, v := range value {
			fmt.Printf("\t Currency Code: %v\n\t\t\t Ranking: %v\n\n", k, v)
		}
	}

	type animal struct {
		name            string
		characteristics []string
	}

	//A herbivore is an animal, so it can have the animal struct as a field
	type herbivore struct {
		animal
		eatHuman bool
	}

	herb := herbivore{
		animal: animal{
			name: "Lion",
			characteristics: []string{"Lacks sense",
				"Lazy",
				"Eat grass",
			},
		},
		// relative: relative{
		// 	name: "Goat",
		// 	characteristics: []string{"Lacks sense",
		// 		"Lazy",
		// 		"Eat grass",
		// 	},
		// },
		eatHuman: false, //maybe
	}

	//We use dot(.) to acces each field in the struct
	fmt.Println("Animal name:", herb.animal.name)
	fmt.Println("Eats human? ", herb.eatHuman)
	fmt.Println("Characteristics:")
	for _, v := range herb.animal.characteristics {
		fmt.Printf("\t %v\n", v)
	}

	arr := []string{"Assurance"}
	arr = append(arr, "James")
	for index, value := range arr {
		fmt.Println(index, value)
	}

	c := make([]int, len(s))
	copy(c, s)

	// p[1] = 100
	fmt.Println(c)
}
