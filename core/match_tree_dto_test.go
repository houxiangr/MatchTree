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
		}, {
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

func TestGetStringAddExpr(t *testing.T){
	matchData := MatchData{
		"a":1,
		"b":2,
		"c":3,
	}

	tests := []struct {
		name string
		args MatchTreeLine
		want string
		wanterr error
	}{
		{
			name: "have cache keys",
			args: MatchTreeLine{
				CacheKey:[]string{"a","b"},
				Expr:"a>1",
			},
			want: "a1b2a>1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res,err := matchData.GetStringAddExpr(tt.args)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("TestTextMatchNormal() res = %v, want %v", res, tt.want)
			}
			if err != tt.wanterr {
				t.Errorf("TestTextMatchNormal() err = %v, wanterr %v", err, tt.wanterr)
			}
		})

	}
}
