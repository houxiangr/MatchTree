package core

import (
	"github.com/textMatch/common"
	"reflect"
	"testing"
)

func TestIsHaveNextNode(t *testing.T) {
	tests := []struct {
		name string
		args MatchTreeLine
		want bool
	}{
		{
			name: "is eof",
			args: MatchTreeLine{
				NextFloorId: common.NextNodeStatusEOF,
			},
			want: true,
		},{
			name: "not eof",
			args: MatchTreeLine{
				NextFloorId: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.args.IsHaveNextNode()
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("TestTextMatchNormal() res = %v, want %v", res, tt.want)
			}
		})

	}
}
