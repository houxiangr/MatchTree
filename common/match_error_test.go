package common

import (
	"reflect"
	"testing"
)

func TestError_SetExtraMsg(t *testing.T) {
	type fields struct {
		ErrNo  int
		ErrMsg string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Error
	}{
		{
			name: "case1: 信息返回",
			fields: fields{
				ErrMsg: "test",
			},
			args: args{
				s: "suffix",
			},
			want: Error{
				ErrMsg: "test:suffix",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				ErrNo:  tt.fields.ErrNo,
				ErrMsg: tt.fields.ErrMsg,
			}
			if got := e.SetExtraMsg(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExtraMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewError(t *testing.T) {
	type args struct {
		err Error
		msg string
	}
	tests := []struct {
		name string
		args args
		want Error
	}{
		{
			name: "生成指定error",
			args: args{
				err: Error{
					ErrNo: 1,
				},
				msg: "test",
			},
			want: Error{1, "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewError(tt.args.err, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_GetErrNo(t *testing.T) {
	type fields struct {
		ErrNo  int
		ErrMsg string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "获取err msg",
			fields: fields{
				ErrNo:  1,
				ErrMsg: "test",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				ErrNo:  tt.fields.ErrNo,
				ErrMsg: tt.fields.ErrMsg,
			}
			if got := e.GetErrNo(); got != tt.want {
				t.Errorf("GetErrNo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		ErrNo  int
		ErrMsg string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "获取err msg",
			fields: fields{
				ErrNo:  1,
				ErrMsg: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				ErrNo:  tt.fields.ErrNo,
				ErrMsg: tt.fields.ErrMsg,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
