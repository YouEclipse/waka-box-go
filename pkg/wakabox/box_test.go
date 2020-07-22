package wakabox

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
}

func TestGenerateBarChart(t *testing.T) {
	type args struct {
		ctx     context.Context
		percent float64
		size    int
		style   string
	}

	ctx := context.Background()
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "barchart-0%-Empty",
			args: args{
				ctx:     ctx,
				percent: 0,
				size:    21,
				style:   "EMPTY",
			},
			want: "                     ",
		},
		{
			name: "barchart-23.5%-Underscore",
			args: args{
				ctx:     ctx,
				percent: 23.5,
				size:    21,
				style:   "UNDERSCORE",
			},
			want: "████▉▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁",
		},
		{
			name: "barchart-0%",
			args: args{
				ctx:     ctx,
				percent: 0,
				size:    21,
				style:   "SOLIDLT",
			},
			want: "░░░░░░░░░░░░░░░░░░░░░",
		},
		{
			name: "barchart-23.5%",
			args: args{
				ctx:     ctx,
				percent: 23.5,
				size:    21,
				style:   "SOLIDLT",
			},
			want: "████▉░░░░░░░░░░░░░░░░",
		},
		{
			name: "barchart-72.5%",
			args: args{
				ctx:     ctx,
				percent: 72.5,
				size:    21,
				style:   "SOLIDLT",
			},
			want: "███████████████▏░░░░░",
		},
		{
			name: "barchart-100%",
			args: args{
				ctx:     ctx,
				percent: 100,
				size:    21,
				style:   "SOLIDLT",
			},
			want: "█████████████████████",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateBarChart(tt.args.ctx, tt.args.percent, tt.args.size, tt.args.style); got != tt.want {
				t.Errorf("GenerateBarChart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetStats(t *testing.T) {

	wakaAPIKey := os.Getenv("WAKATIME_API_KEY")

	ghToken := os.Getenv("GH_TOKEN")
	ghUsername := os.Getenv("GH_USER")
	style := BoxStyle{
		BarStyle:  os.Getenv("GIST_BARSTYLE"),
		BarLength: os.Getenv("GIST_BARLENGTH"),
		TimeStyle: os.Getenv("GIST_TIMESTYLE"),
	}
	fmt.Printf("%+v - %+v", style, BarStyle[style.BarStyle])
	box := NewBox(wakaAPIKey, ghUsername, ghToken, style)

	lines, err := box.GetStats(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log(strings.Join(lines, "\n"))

}

func TestBox_Gist(t *testing.T) {
	wakaAPIKey := os.Getenv("WAKATIME_API_KEY")

	ghToken := os.Getenv("GH_TOKEN")
	ghUsername := os.Getenv("GH_USER")
	gistID := os.Getenv("GIST_ID")

	style := BoxStyle{
		BarStyle:  os.Getenv("GIST_BARSTYLE"),
		BarLength: os.Getenv("GIST_BARLENGTH"),
		TimeStyle: os.Getenv("GIST_TIMESTYLE"),
	}

	box := NewBox(wakaAPIKey, ghUsername, ghToken, style)

	ctx := context.Background()
	filename := "📊 Weekly development breakdown"
	gist, err := box.GetGist(ctx, gistID)
	if err != nil {
		t.Error(err)
	}

	f := gist.Files[github.GistFilename(filename)]

	f.Content = github.String(time.Now().UTC().Format(time.RFC3339))
	gist.Files[github.GistFilename(filename)] = f
	err = box.UpdateGist(ctx, gistID, gist)
	if err != nil {
		t.Error(err)
	}
}
func Test_convertTime(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		want string
	}{
		{
			name: "10 hrs",
			want: "10h",
		},
		{
			name: "18 hrs 40 mins",
			want: "18h40m",
		},
		{
			name: "1 hr 13 mins",
			want: "1h13m",
		},
		{
			name: "2 mins",
			want: "2m",
		},
		{
			name: "0 secs",
			want: "0s",
		},
		{
			name: "99 hrs 99 mins 99 secs",
			want: "99h99m99s",
		},
		{
			name: "1 sec",
			want: "1s",
		},
		{
			name: "1 min",
			want: "1m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDuration(tt.name); got != tt.want {
				t.Errorf("convertTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_Readme(t *testing.T) {
	wakaAPIKey := os.Getenv("WAKATIME_API_KEY")

	ghToken := os.Getenv("GH_TOKEN")
	ghUsername := os.Getenv("GH_USER")

	style := BoxStyle{
		BarStyle:  os.Getenv("GIST_BARSTYLE"),
		BarLength: os.Getenv("GIST_BARLENGTH"),
		TimeStyle: os.Getenv("GIST_TIMESTYLE"),
	}

	box := NewBox(wakaAPIKey, ghUsername, ghToken, style)

	ctx := context.Background()

	filename := "test.md"
	title := `####  <a href="https://gist.github.com/YouEclipse/9bc7025496e478f439b9cd43eba989a4" target="_blank">📊 Weekly development breakdown</a>`
	content := []byte(`Go         🕓 18h3m ██████████████████████▉░░░░░ 82.1%
YAML       🕓 1h47m ██▎░░░░░░░░░░░░░░░░░░░░░░░░░  8.1%
JavaScript 🕓 40m   ▊░░░░░░░░░░░░░░░░░░░░░░░░░░░  3.1%
Markdown   🕓 34m   ▋░░░░░░░░░░░░░░░░░░░░░░░░░░░  2.6%
Other      🕓 32m   ▋░░░░░░░░░░░░░░░░░░░░░░░░░░░  2.5%`)

	err := box.UpdateMarkdown(ctx, title, filename, content)
	if err != nil {
		t.Error(err)
	}
	c, _ := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", c)
}
