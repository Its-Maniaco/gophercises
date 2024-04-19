package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	jp "github.com/Its-Maniaco/AdventureGame/JsonParse"
	h "github.com/Its-Maniaco/AdventureGame/story"
)

func main() {
	port := flag.Int("port", 3000, "port to run CYOA")
	fileName := flag.String("file", "gopher.json", "Json file with Choose Your Own Story.")
	flag.Parse()
	data := jp.JsonParse(*fileName)
	//tc.TemplateCreate(data)
	//story.StartStory(data)

	nh := h.NewStoryHandler(data)
	fmt.Println("Starting server at ", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nh))
}
