/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

type Joke struct{
	Id string `json:"id"`
	Joke string `json:"joke"`
	Status string `json:"status"`
}

// jokeCmd represents the joke command
var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "A bad joke",
	Long: `Not much to say!`,
	Run: func(cmd *cobra.Command, args []string) {
		getJoke()
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)
}

func getJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}

	json.Unmarshal(responseBytes, &joke)

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseApi string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseApi,
		nil,
	)

	if(err != nil) {
		fmt.Printf("could not create request %v", err)
	}

	request.Header.Add("Accept", "Application/json")
	request.Header.Add("User-Agent", "go-cli-joke")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}