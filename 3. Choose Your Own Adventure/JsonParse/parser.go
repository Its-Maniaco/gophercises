package JsonParse

import (
	"encoding/json"
	"fmt"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JsonParse(filename string) Story {
	//os.Chdir("../")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error in opening file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	var data Story
	jsonDecoder := json.NewDecoder(file)
	if err := jsonDecoder.Decode(&data); err != nil {
		fmt.Println("Error in decoding file: ", err)
		os.Exit(1)
	}
	return data
}
