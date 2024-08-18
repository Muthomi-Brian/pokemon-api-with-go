package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	// "net/http"
	"os"
)

type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonAbility struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Pokemon struct {
	Name      string           `json:"name"`
	ID        int              `json:"id"`
	Height    int              `json:"height"`
	Weight    int              `json:"weight"`
	Abilities []PokemonAbility `json:"abilities"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("pokemon name please:")
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", input)
	fmt.Println(url)
	if err != nil {
		log.Fatal("you didnt input anything")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("an error occured", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("err", err)
	}
	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("ID: %d\n", pokemon.ID)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	// Print the abilities
	for _, ability := range pokemon.Abilities {
		fmt.Printf("Ability: %s\n (URL: %s\n), Hidden: %t\n, Slot: %d\n",
			ability.Ability.Name, ability.Ability.URL, ability.IsHidden, ability.Slot)
	}
}
