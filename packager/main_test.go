package packager

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestNewPackagerList(t *testing.T) {
	type args struct {
		config *viper.Viper
	}
	tests := []struct {
		name string
		args args
		want []PackageHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPackagerList(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPackagerList() = %v, want %v", got, tt.want)
			}
		})
	}
}
