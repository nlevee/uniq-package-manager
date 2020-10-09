package gocomposer

import (
	"testing"
)

func TestHasComposer(t *testing.T) {
	type args struct {
		dir string
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
			if got := HasComposer(tt.args.dir); got != tt.want {
				t.Errorf("HasComposer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasComposerLock(t *testing.T) {
	type args struct {
		dir string
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
			if got := HasComposerLock(tt.args.dir); got != tt.want {
				t.Errorf("HasComposerLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
