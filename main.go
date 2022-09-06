package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//language: go

var URL string = "https://api.dictionaryapi.dev/api/v2/entries/en_US/"

func main() {

	fmt.Println("Welcome to the Dictionary")
	fmt.Println("--------------------------")

	menu()
}

// menu used to select what a user wants to do
func menu() {
	fmt.Println("1. Find a definition")
	fmt.Println("2. Get an example")
	fmt.Println("3. Get a synonym")
	fmt.Println("4. Exit")

	var choice int

	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		findDefinition()
	case 2:
		getExample()
	case 3:
		getSynonym()
	case 4:
		exit()
	default:
		fmt.Println("Invalid choice")
	}
	//recursive call until user exits
	//2 second delay
	time.Sleep(2 * time.Second)
	menu()
}

func findDefinition() {
	//requesting a word from the user
	fmt.Println("Finding a definition")

	var word string

	fmt.Print("Enter a word: ")
	fmt.Scanln(&word)

	fmt.Println("Finding definition for: ", word)
	time.Sleep(1 * time.Second)

	//make a request to the dictionary api to get the definition of the word and print it
	//https://api.dictionaryapi.dev/api/v2/entries/en_US/{word}

	//make a request to the dictionary api
	res, err := http.Get(URL + word)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println("Response status: ", res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//create nested struct to store the response
	type response struct {
		Word      string `json:"word"`
		Phonetics []struct {
			Text  string
			Audio string
		}
		Meanings []struct {
			PartOfSpeech string
			Definitions  []struct {
				Definition string
				Example    string
				Synonyms   []string
				Antonyms   []string
			}
		}
	}

	//!-main Figure out how to fucking parse this shit right
	//! currently its just printing empty brackets
	/*//*in theory this should be parsing the response
	//* into a struct which can then be printed */
	var responseObject response
	json.Unmarshal(body, &responseObject)
	fmt.Println(string(responseObject.Word))

}

func getExample() {
	fmt.Println("Getting an example")
}

func getSynonym() {
	fmt.Println("Getting a synonym")
}

func exit() {
	fmt.Println("Exiting...")
	//END PROGRAM
	os.Exit(0)
}
