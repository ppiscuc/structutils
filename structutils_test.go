package structutils
import (
	"reflect"
	"testing"
)

func TestStructJSONTags(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"onetag",
			args{struct {
				B string `json:"b"`
			}{}},
			[]string{"b"},
		},
		{
			"omitempty",
			args{struct {
				C string `json:"b,omitemty"`
			}{}},
			[]string{"b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructJSONTags(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructJSONTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
