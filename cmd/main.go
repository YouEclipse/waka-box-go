package main

import (
	"context"
	"os"
	"strings"

	wakabox "github.com/YouEclipse/waka-box-go/pkg"
	"github.com/google/go-github/github"
)

func main() {
	wakaAPIKey := os.Getenv("WAKATIME_API_KEY")

	ghToken := os.Getenv("GH_TOKEN")
	ghUsername := os.Getenv("GH_USER")
	box := wakabox.NewBox(wakaAPIKey, ghUsername, ghToken)

	lines, err := box.GetStats(context.Background())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	gistID := "d3798a7bc234087e75aed5716474f42a"
	filename := "📊 Weekly development breakdown"
	gist, err := box.GetGist(ctx, gistID)
	if err != nil {
		panic(err)
	}

	f := gist.Files[github.GistFilename(filename)]

	f.Content = github.String(strings.Join(lines, "\n"))
	gist.Files[github.GistFilename(filename)] = f
	err = box.UpdateGist(ctx, gistID, gist)
	if err != nil {
		panic(err)
	}
}
