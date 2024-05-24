package story

import (
	"fmt"

	jp "github.com/Its-Maniaco/AdventureGame/JsonParse"
)

func StartStory(data jp.Story) {
	storyContinuation("intro", data)
	fmt.Println("Story Time is Over.")
}

func storyContinuation(storyTitle string, data jp.Story) {
	currStoryData := data[storyTitle]
	fmt.Println("Title of story: ", currStoryData.Title)
	fmt.Println(currStoryData.Story)
	if storyTitle == "home" {
		return
	}
	//Take user input and then match option after displaying options.
	for i, option := range currStoryData.Options {
		fmt.Println(i+1, ": ", option.Text)
	}

	var choice int
	fmt.Printf("Enter your choice: ")
	fmt.Scanf("%d", &choice)
	if choice > len(currStoryData.Options) && choice <= 0 {
		storyContinuation("home", data)
	} else {
		storyContinuation(currStoryData.Options[choice-1].Arc, data)
	}
}
