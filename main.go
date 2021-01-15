package main

import (
	"fmt"
	"os"
	"strings"
	prompt "github.com/c-bata/go-prompt"
)

// For a good example, see https://github.com/c-bata/go-prompt/blob/master/_example/http-prompt/main.go

var suggestions = []prompt.Suggest{
	{"exit", "Exit the prompt"},
	{"another", "Some description"},
}

var context string

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func executor(in string) {
	in = strings.TrimSpace(in)

	// var method, body string
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	case "cd":
		fmt.Println("Do something and prompt again")
		return
	}

}

func livePrefix() (string, bool) {
	if context == "/" {
		return "", false
	}
	return context + "> ", true
}

func main() {
	// Basic Use case
	fmt.Println("Basic prompt")
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)

	// More advanced
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(context + "> "),
		prompt.OptionLivePrefix(livePrefix),
		prompt.OptionTitle("Basic cli"),
	)
	p.Run()

}
