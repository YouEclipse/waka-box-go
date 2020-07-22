package wakabox

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/YouEclipse/wakatime-go/pkg/wakatime"
	"github.com/google/go-github/github"
)

// maxLineLength is the visible number of characters in a pinned gist box
// (accounting for the clock emoji)
const maxLineLength = 53

// BarStyle defines valid styles for the progress bar
var BarStyle = map[string][]rune{
	"SOLIDLT":    []rune(`░▏▎▍▌▋▊▉█`),
	"SOLIDMD":    []rune(`▒▏▎▍▌▋▊▉█`),
	"SOLIDDK":    []rune(`▓▏▎▍▌▋▊▉█`),
	"EMPTY":      []rune(` ▏▎▍▌▋▊▉█`),
	"UNDERSCORE": []rune(`▁▏▎▍▌▋▊▉█`),
}

// BoxStyle contains information for initalizing a gist box style
type BoxStyle struct {
	BarStyle     string // Style of the progress bar as defined by BarStyle
	BarLength    string // Length of the bar as a string (gets converted to an Int)
	TimeStyle    string // Style of the time text. "SHORT" will be abbreviated.
	barLengthInt int    // Set automatically from the Length defined above
	maxLangLen   int    // Set automatically from the list of languages from wakatime
	maxTimeLen   int    // Set automatically from the list of times from wakatime

}

// Box contains a github and wakatime client and styling information for the gist box
type Box struct {
	github   *github.Client
	wakatime *wakatime.Client
	style    BoxStyle
}

// NewBox creates a box struct with appropriate wakatime and github information and gist styling information
func NewBox(wakaAPIKey, ghUsername, ghToken string, style BoxStyle) *Box {
	box := &Box{}

	box.wakatime = wakatime.NewClient(wakaAPIKey, nil)

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(ghUsername),
		Password: strings.TrimSpace(ghToken),
	}
	box.github = github.NewClient(tp.Client())

	length, err := strconv.Atoi(style.BarLength)
	if err != nil {
		length = 21 //Default to 21
	}
	style.barLengthInt = length
	if style.BarStyle == "" {
		style.BarStyle = "SOLIDLT" // Default to SOLIDLT
	}
	box.style = style

	return box
}

// GetStats gets the language stats form wakatime.com.
func (b *Box) GetStats(ctx context.Context) ([]string, error) {
	stats, err := b.wakatime.Stats.Current(ctx, wakatime.RangeLast7Days, &wakatime.StatsQuery{})
	if err != nil {
		return nil, fmt.Errorf("wakabox.GetStats: Error getting Current Stats: %w", err)
	}

	if languages := stats.Data.Languages; len(languages) > 0 {
		lines, err := b.GenerateGistLines(ctx, languages)
		if err != nil {
			return nil, fmt.Errorf("wakabox.GetStats: Error generating gist lines: %w", err)
		}

		return lines, nil
	}
	return []string{"Still Gathering Statistics..."}, nil
}

// GetGist gets the gist from github.com.
func (b *Box) GetGist(ctx context.Context, id string) (*github.Gist, error) {
	gist, _, err := b.github.Gists.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("wakabox.GetGist: Error getting gist from github: %w", err)
	}
	return gist, nil
}

// UpdateGist updates the gist.
func (b *Box) UpdateGist(ctx context.Context, id string, gist *github.Gist) error {
	_, _, err := b.github.Gists.Edit(ctx, id, gist)
	if err != nil {
		return fmt.Errorf("wakabox.UpdateGist: Error updating gist: %w", err)
	}
	return nil
}

func (b *Box) UpdateMarkdown(ctx context.Context, title, filename string, content []byte) error {
	md, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("wakabox.UpdateMarkdown: Error reade a file: %w", err)
	}

	start := []byte("<!-- waka-box start -->")
	before := md[:bytes.Index(md, start)+len(start)]
	end := []byte("<!-- waka-box end -->")
	after := md[bytes.Index(md, end):]

	newMd := bytes.NewBuffer(nil)
	newMd.Write(before)
	newMd.WriteString("\n" + title + "\n")
	newMd.WriteString("```text\n")
	newMd.Write(content)
	newMd.WriteString("\n")
	newMd.WriteString("```\n")
	newMd.WriteString("<!-- Powered by https://github.com/YouEclipse/waka-box-go . -->\n")
	newMd.Write(after)

	err = ioutil.WriteFile(filename, newMd.Bytes(), os.ModeAppend)
	if err != nil {
		return fmt.Errorf("wakabox.UpdateMarkdown: Error write a file: %w", err)
	}

	return nil
}

// GenerateGistLines takes an slice of wakatime.StatItems, and generates a line for the gist.
func (b *Box) GenerateGistLines(ctx context.Context, languages []wakatime.StatItem) ([]string, error) {
	max := 0
	lines := make([]string, 0)
	for _, stat := range languages {
		if b.style.TimeStyle == "SHORT" {
			*stat.Text = convertDuration(*stat.Text)
		}
		if b.style.maxTimeLen < len(*stat.Text) {
			b.style.maxTimeLen = len(*stat.Text)
		}
		if b.style.maxLangLen < len(*stat.Name) {
			b.style.maxLangLen = len(*stat.Name)
		}
	}
	if b.style.barLengthInt < 0 {
		b.style.barLengthInt = maxLineLength - (b.style.maxLangLen + b.style.maxTimeLen + 10)
	}
	for _, stat := range languages {
		if max >= 5 {
			break
		}
		lines = append(lines, b.ConstructLine(ctx, stat))
		max++
	}
	return lines, nil
}

// ConstructLine formats a gist line from stat infomation
func (b *Box) ConstructLine(ctx context.Context, stat wakatime.StatItem) string {
	return fmt.Sprintf("%-*s🕓 %-*s%s%5.1f%%",
		b.style.maxLangLen+1, *stat.Name,
		b.style.maxTimeLen+1, *stat.Text,
		GenerateBarChart(ctx, *stat.Percent, b.style.barLengthInt, b.style.BarStyle),
		*stat.Percent)
}

// GenerateBarChart generates a bar chart with the given percent and size.
// Percent is a float64 from 0-100 representing the progress bar percentage
// Size is an int representing the length of the progress bar in characters
// BarType is a BarType representing the type of barchart: It can be one of the following:
//    SOLIDLT SOLIDMD SOLIDDK: Block characters with a dotted background
//    UNDERSCORE: Block characters with an line accross the boottom
//    EMPTY: Block characters with an empty background
func GenerateBarChart(ctx context.Context, percent float64, size int, barType string) string {
	// using rune as for utf-8 encoding
	syms := BarStyle[barType]
	if len(syms) > 9 {
		panic("No Syms")
	}

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

func convertDuration(t string) string {
	r := strings.NewReplacer(
		"hr", "h",
		"min", "m",
		"sec", "s",
		" ", "",
		"s", "",
	)
	return r.Replace(t)
}
