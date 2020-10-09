package tools

import (
	"testing"

	"github.com/docker/docker/client"
)

func TestGetFilePath(t *testing.T) {
	type args struct {
		dir      string
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFilePath(tt.args.dir, tt.args.filename); got != tt.want {
				t.Errorf("GetFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCleanContainer(t *testing.T) {
	type args struct {
		cli         *client.Client
		containerID string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CleanContainer(tt.args.cli, tt.args.containerID)
		})
	}
}

func TestImagePull(t *testing.T) {
	type args struct {
		cli       *client.Client
		nodeImage string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ImagePull(tt.args.cli, tt.args.nodeImage)
		})
	}
}
