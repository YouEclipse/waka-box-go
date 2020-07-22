package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/YouEclipse/waka-box-go/pkg/wakabox"
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
	updateOption := os.Getenv("UPDATE_OPTION") // options for update: gist,markdown,gist_and_markdown
	markdownFile := os.Getenv("MARKDOWN_FILE") // the filename of markdown

	var updateGist, updateMarkdown bool
	if updateOption == "markdown" {
		updateMarkdown = true
	} else if updateOption == "gist_and_markdown" {
		updateGist = true
		updateMarkdown = true
	} else {
		updateGist = true
	}

	style := wakabox.BoxStyle{
		BarStyle:  os.Getenv("GIST_BARSTYLE"),
		BarLength: os.Getenv("GIST_BARLENGTH"),
		TimeStyle: os.Getenv("GIST_TIMESTYLE"),
	}

	box := wakabox.NewBox(wakaAPIKey, ghUsername, ghToken, style)

	ctx := context.Background()
	lines, err := box.GetStats(ctx)
	if err != nil {
		panic(err)
	}
	filename := "📊 Weekly development breakdown"
	if updateGist {

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

	if updateMarkdown && markdownFile != "" {
		title := filename
		if updateGist {
			title = fmt.Sprintf(`####  <a href="https://gist.github.com/%s" target="_blank">%s</a>`, gistID, title)
		}

		content := bytes.NewBuffer([]byte(filename))
		content.WriteString("\n")
		content.WriteString(strings.Join(lines, "\n"))

		err = box.UpdateMarkdown(ctx, title, markdownFile, content.Bytes())
		if err != nil {
			fmt.Println(err)
		}
	}

}
