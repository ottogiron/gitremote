package server

import "testing"
import "io/ioutil"
import "io"

var onOutput = ioutil.Discard

func Test_gitService_Execute(t *testing.T) {

	type args struct {
		dir      string
		command  string
		onOutput io.Writer
	}
	tests := []struct {
		name string
		args args

		wantErr bool
	}{
		{"Execute git status", args{".", "git status", onOutput}, false},
		{"Empty command", args{".", "", onOutput}, true},
		{"Empty non git command", args{".", "got status", onOutput}, true},
		{"Command failing to execute", args{".", "git something", onOutput}, true},
	}
	for _, tt := range tests {
		g := &gitService{}
		t.Run(tt.name, func(t *testing.T) {
			err := g.Execute(tt.args.dir, tt.args.command, tt.args.onOutput)
			if (err != nil) != tt.wantErr {
				t.Errorf("%q. gitService.Execute() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}

		})

	}
}