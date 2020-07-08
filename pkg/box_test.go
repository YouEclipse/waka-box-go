package wakabox

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-github/github"
)

func TestGenerateBarChart(t *testing.T) {
	type args struct {
		ctx     context.Context
		percent float64
		size    int
	}

	ctx := context.Background()
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "barchart-0%",
			args: args{
				ctx:     ctx,
				percent: 0,
				size:    21,
			},
			want: "░░░░░░░░░░░░░░░░░░░░░",
		},
		{
			name: "barchart-23.5%",
			args: args{
				ctx:     ctx,
				percent: 23.5,
				size:    21,
			},
			want: "████▉░░░░░░░░░░░░░░░░",
		},
		{
			name: "barchart-72.5%",
			args: args{
				ctx:     ctx,
				percent: 72.5,
				size:    21,
			},
			want: "███████████████▏░░░░░",
		},
		{
			name: "barchart-100%",
			args: args{
				ctx:     ctx,
				percent: 100,
				size:    21,
			},
			want: "█████████████████████",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateBarChart(tt.args.ctx, tt.args.percent, tt.args.size); got != tt.want {
				t.Errorf("GenerateBarChart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetStats(t *testing.T) {
	wakaAPIKey := os.Getenv("WAKATIME_API_KEY")

	ghToken := os.Getenv("GH_TOKEN")
	ghUsername := os.Getenv("GH_USER")
	box := NewBox(wakaAPIKey, ghUsername, ghToken)

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
	box := NewBox(wakaAPIKey, ghUsername, ghToken)

	ctx := context.Background()
	gistID := os.Getenv("GIST_ID")
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
