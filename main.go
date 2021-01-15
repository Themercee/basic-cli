package main

import (
	"fmt"
	prompt "github.com/c-bata/go-prompt"
)

// For a good example, see https://github.com/c-bata/go-prompt/blob/master/_example/http-prompt/main.go

var suggestions = []prompt.Suggest {
	{"exit", "Exit the prompt"},
	{"another", "Some description"},
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Println("Basic prompt")
	t := prompt.Input("> ", completer, )
	fmt.Println("You selected " + t)
}