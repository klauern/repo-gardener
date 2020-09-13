package gardener

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/cli/cli/api"
)

type PullRequester struct {
	GitHubPR      *api.PullRequest
	GitHubRepo    *api.Repository
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
