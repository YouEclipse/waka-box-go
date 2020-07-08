package main

import (
	"context"
	"os"
	"strings"

	wakabox "github.com/YouEclipse/waka-box-go/pkg"
	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	wakaAPIKey := os.Getenv("WAKATIME_API_KEY")
	ghToken := os.Getenv("GH_TOKEN")
	ghUsername := os.Getenv("GH_USER")
	gistID := os.Getenv("GIST_ID")

	box := wakabox.NewBox(wakaAPIKey, ghUsername, ghToken)

	lines, err := box.GetStats(context.Background())
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

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
