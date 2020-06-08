package wakabox

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"

	wakatime "github.com/YouEclipse/wakatime-go/pkg"
	"github.com/google/go-github/github"
)

type Box struct {
	github   *github.Client
	wakatime *wakatime.Client
}

func NewBox(wakaAPIKey, ghUsername, ghToken string) *Box {
	box := &Box{}

	box.wakatime = wakatime.NewClient(wakaAPIKey, nil)

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(ghUsername),
		Password: strings.TrimSpace(ghToken),
	}
	box.github = github.NewClient(tp.Client())

	return box
}

// GetStats gets the language stats form wakatime.com.
func (b *Box) GetStats(ctx context.Context) ([]string, error) {
	stats, err := b.wakatime.Stats.Current(ctx, wakatime.RangeLast7Days, &wakatime.StatsQuery{})
	if err != nil {
		return nil, err
	}
	max := 0

	if languages := stats.Data.Languages; len(languages) > 0 {
		lines := make([]string, 0)
		for _, stat := range stats.Data.Languages {
			if max >= 5 {
				break
			}

			line := pad(*stat.Name, " ", 9) + " " +
				pad("🕓 "+*stat.Text, " ", 15) + " " +
				GenerateBarChart(ctx, *stat.Percent, 21) + " " +
				pad(fmt.Sprintf("%.1f%%", *stat.Percent), " ", 5)
			lines = append(lines, line)
			max++
		}
		return lines, nil
	} else {
		return nil, errors.New("Insufficient statistics")
	}
}

// GetGist gets the gist from github.com.
func (b *Box) GetGist(ctx context.Context, id string) (*github.Gist, error) {
	gist, _, err := b.github.Gists.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return gist, nil
}

// UpdateGist updates the gist.
func (b *Box) UpdateGist(ctx context.Context, id string, gist *github.Gist) error {
	_, _, err := b.github.Gists.Edit(ctx, id, gist)
	return err
}

// GenerateBarChart generates a bar chart with the given percent and size.
func GenerateBarChart(ctx context.Context, percent float64, size int) string {
	// using rune as for utf-8 encoding
	var syms = []rune(`░▏▎▍▌▋▊▉█`)

	frac := int(math.Floor((float64(size) * 8 * percent) / 100))
	barsFull := int(math.Floor(float64(frac) / 8))

	if barsFull >= size {
		return strings.Repeat(string(syms[8:9]), size)
	}

	var semi = frac % 8

	bar := strings.Repeat(string(syms[8:9]), barsFull) + string(syms[semi:semi+1])

	return pad(bar, string(syms[0:1]), size)
}

func pad(s, pad string, targetLength int) string {
	padding := targetLength - utf8.RuneCountInString(s)
	if padding <= 0 {
		return s
	}

	return s + strings.Repeat(pad, padding)
}
