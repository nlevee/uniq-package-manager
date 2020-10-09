package gonpm

import (
	"testing"
)

func TestHasPackage(t *testing.T) {
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
			if got := HasPackage(tt.args.dir); got != tt.want {
				t.Errorf("HasPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPackageLock(t *testing.T) {
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
			if got := HasPackageLock(tt.args.dir); got != tt.want {
				t.Errorf("HasPackageLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
