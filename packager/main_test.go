package packager

import (
	"reflect"
	"testing"
)

func TestNewPackagerList(t *testing.T) {
	tests := []struct {
		name string
		want []PackageHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPackagerList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPackagerList() = %v, want %v", got, tt.want)
			}
		})
	}
}
