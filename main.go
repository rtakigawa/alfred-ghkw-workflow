package main

import (
	"strings"

	aw "github.com/deanishe/awgo"
)

var (
	wf        *aw.Workflow
	languages = []string{
		"javascript",
		"python",
		"java",
		"go",
		"typescript",
		"c++",
		"ruby",
		"php",
		"c#",
		"c",
		"shell",
		"scala",
		"dart",
		"rust",
		"kotlin",
		"swift",
		"groovy",
		"dm",
		"objective-c",
		"elixir",
		"perl",
		"coffeescript",
		"powershell",
		"clojure",
		"tsql",
		"lua",
		"vim script",
		"haskell",
		"emacs lisp",
		"ocaml",
		"jsonnet",
		"erlang",
		"r",
		"puppet",
		"julia",
		"fortran",
		"coq",
		"abap",
		"matlab",
		"systemverilog",
		"visual basic .net",
		"webassembly",
		"f#",
		"objective-c++",
		"elm",
		"vala",
		"smalltalk",
		"haxe",
		"common lisp",
		"roff",
	}
)

func init() {
	wf = aw.New()
}

func run() {
	args := wf.Args()

	if len(args) == 0 {
		createLanguages()
	}

	if len(args) == 1 {
		lang := args[0]

		createLanguages()
		wf.Filter(lang)

		if validLanguage(lang) {
			wf.NewItem("search keywords").
				Title("Please enter the keyword to search")
		}
	}

	if len(args) >= 2 {
		lang := args[0]
		keywords := strings.Join(args[1:], " ")
		cmd := "ghkw --language=" + lang + " " + keywords
		wf.NewItem("Search by keywords").
			Title("Search").
			Subtitle("Run '" + cmd + "' in terminal").
			Arg(cmd).
			Valid(true)
	}

	wf.WarnEmpty("No matching languages found", "Try a different query?")

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}

func createLanguages() {
	for _, l := range languages {
		wf.NewItem(l).
			Title(l).
			Autocomplete(l + " ").
			Subtitle("language")
	}
}

func validLanguage(lang string) bool {
	for _, l := range languages {
		if l == lang {
			return true
		}
	}
	return false
}
