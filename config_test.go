package gardener

import (
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewGardenConfig(t *testing.T) {
	tests := []struct {
		name string
		want *GardenConfig
	}{
		{
			name: "empty",
			want: &GardenConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGardenConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGardenConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGardenConfig_Template(t *testing.T) {
	testFs := afero.NewMemMapFs()
	type fields struct {
		Name        string
		PullRequest *pullRequest
		Actions     []repoAction
	}
	type args struct {
		path string
		fs   afero.Fs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "new file",
			args: args{
				path: "garden.yaml",
				fs:   testFs,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GardenConfig{
				Name:        tt.fields.Name,
				PullRequest: tt.fields.PullRequest,
				Actions:     tt.fields.Actions,
			}
			if err := g.Template(tt.args.path, tt.args.fs); (err != nil) != tt.wantErr {
				t.Errorf("GardenConfig.Template() error = %v, wantErr %v", err, tt.wantErr)
			}
			file, err := afero.ReadFile(testFs, tt.args.path)
			if err != nil {
				t.Errorf("GardenConfig.Template() file open error = %v", err)
			}
			if !reflect.DeepEqual(string(file), GardenConfigTemplate) {
				t.Errorf("GardenConfig.Template() file comparison error = %v, want %v", string(file), GardenConfigTemplate)
			}
		})
	}
}
