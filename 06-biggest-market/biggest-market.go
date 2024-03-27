package main

import (
	"cmp"
	"encoding/json"
	"log"
	"os"
	"slices"
)

// User represents a user record.
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

const path = "users.json"

// getBiggestMarket takes in the slice of users and
// returns the biggest market.
func getBiggestMarket(users []User) (string, int) {
	dataMap := make(map[string]int)

	for _, user := range users {
		dataMap[user.Country] += 1
	}

	countries := make([]string, 0, len(dataMap))
	for country := range dataMap {
		countries = append(countries, country)
	}

	slices.SortFunc(countries, func(a, b string) int {
		return cmp.Compare(dataMap[b], dataMap[a])
	})

	return countries[0], dataMap[countries[0]]
}

func main() {
	users := importData()
	country, count := getBiggestMarket(users)
	log.Printf("The biggest user market is %s with %d users.\n",
		country, count)
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []User {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []User
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
