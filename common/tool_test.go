package common

import (
	"reflect"
	"testing"
)

func TestTransferInterfaceToString(t *testing.T) {
	tests := []struct {
		name    string
		args    interface{}
		want    string
		wanterr error
	}{
		{
			name: "int to string",
			args: 1,
			wanterr: nil,
			want:    "1",
		},
		{
			name: "string to string",
			args: "1",
			wanterr: nil,
			want:    "1",
		},
		{
			name: "not match type to string",
			args: map[string]interface{}{
				"1":1,
			},
			wanterr: NotMatchInterfaceType,
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := TransferInterfaceToString(tt.args)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("TestTransferInterfaceToString() res = %v, want %v", res, tt.want)
			}
			if err != tt.wanterr {
				t.Errorf("TestTransferInterfaceToString() err = %v, wanterr %v", err, tt.wanterr)
			}
		})
	}
}
