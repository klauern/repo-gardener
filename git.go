package gardener

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/google/go-github/v33/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

type GardenClient struct {
	GitHub *github.Client
}

func NewGardenClient(token string) *GardenClient {
	return &GardenClient{}
}

type PullRequester struct {
	GitHubPR      *github.PullRequest
	GitHubRepo    *github.Repository
	JiraTicketIDs []string
	Reason        string
}

const PullRequestDescriptionTemplate = `# Description

{{.Reason}}
{{if .JiraTicketIDs}}# References

## JIRA Tickets
{{- end}}
{{range .JiraTicketIDs}}  - {{.}}
{{end}}`

func (p *PullRequester) DescriptionString() (string, error) {
	var b strings.Builder
	t := template.Must(template.New("Pull Request").Parse(PullRequestDescriptionTemplate))
	err := t.Execute(&b, p)
	if err != nil {
		return "", fmt.Errorf("building PR template: %w", err)
	}
	return b.String(), nil
}
