package gardener

import (
	"testing"

	"github.com/cli/cli/api"
)

func TestPullRequester_DescriptionString(t *testing.T) {
	type fields struct {
		GitHubPR      *api.PullRequest
		GitHubRepo    *api.Repository
		JiraTicketIDs []string
		Reason        string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "dummy",
			fields: fields{
				GitHubPR:      nil,
				GitHubRepo:    nil,
				JiraTicketIDs: []string{"SECURE-1234", "SECURE-2345"},
				Reason:        "because I said so",
			},
			want: `# Description

because I said so
# References

## JIRA Tickets
  - SECURE-1234
  - SECURE-2345
`,
			wantErr: false,
		},
		{
			name: "no tickets",
			fields: fields{
				GitHubPR:      nil,
				GitHubRepo:    nil,
				JiraTicketIDs: nil,
				Reason:        "some reason",
			},
			want: `# Description

some reason

`,
			wantErr: false,
		},
		{
			name: "no tickets 2",
			fields: fields{
				GitHubPR:      nil,
				GitHubRepo:    nil,
				JiraTicketIDs: []string{},
				Reason:        "some reason",
			},
			want: `# Description

some reason

`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PullRequester{
				GitHubPR:      tt.fields.GitHubPR,
				GitHubRepo:    tt.fields.GitHubRepo,
				JiraTicketIDs: tt.fields.JiraTicketIDs,
				Reason:        tt.fields.Reason,
			}
			got, err := p.DescriptionString()
			if (err != nil) != tt.wantErr {
				t.Errorf("PullRequester.DescriptionString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PullRequester.DescriptionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
