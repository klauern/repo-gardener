package gardener

import (
	"fmt"
	"text/template"

	"github.com/spf13/afero"
)

type pullRequest struct {
	Title        string `yaml:"title"`
	Template     string `yaml:"template"`
	TemplateFile string `yaml:"template_file"`
}

type repoAction struct {
	ID          string    `yaml:"id"`
	RelativeDir string    `yaml:"relative_dir"`
	GlobPattern string    `yaml:"glob_pattern"`
	Cmds        []repoCmd `yaml:"cmds"`
}

type repoCmd struct {
	Cmd    string `yaml:"cmd,omitempty"`
	Script string `yaml:"script,omitempty"`
}

type GardenConfig struct {
	Name        string       `yaml:"name"`
	PullRequest *pullRequest `yaml:"pull_request"`
	Actions     []repoAction `yaml:"actions"`
}

type GardenInstructionSet []GardenConfig

const GardenConfigTemplate = `
name: Thing
pull_request:
  title: something
  template: raw template here
  template_file: file path here
actions:
  - id: optional
    relative_dir: .github/
    glob_pattern: "thing*.yaml"
    cmds:
      - cmd: thing (runs with shell)
      - script: path/to/script.sh
`

func (g *GardenConfig) Template(path string, fs afero.Fs) error {
	t := template.New("config")
	templ, err := t.Parse(GardenConfigTemplate)
	if err != nil {
		return fmt.Errorf("parsing GardenConfigTemplate: %w", err)
	}
	exists, err := afero.Exists(fs, path)
	if err != nil {
		return fmt.Errorf("Error getting file status: %w", err)
	}
	if exists {
		return fmt.Errorf("File exists already")
	}
	file, err := fs.Create(path)
	if err != nil {
		return fmt.Errorf("creating template file: %w", err)
	}
	err = templ.Execute(file, g)
	if err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	return nil
}

func NewGardenConfig() *GardenConfig {
	return &GardenConfig{}
}
