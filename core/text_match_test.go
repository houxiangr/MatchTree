package core

import (
	"encoding/json"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/textMatch/common"
	"reflect"
	"testing"
	"time"
)

func TestTextMatchNormal(t *testing.T) {
	testMatchTreeJson := `{
	"match_floors": [{
			"one_floor_node1": [{
					"expr": "a==1",
					"next_floor_id": "two_floor_node1",
					"data": {}
				},
				{
					"expr": "a==2",
					"next_floor_id": "two_floor_node1",
					"data": {}
				},
				{
					"expr": "a==3",
					"next_floor_id": "EOF",
					"data": {
						"test": 3
					}
				}
			]
		},
		{
			"two_floor_node1": [{
				"expr": "b==1",
				"next_floor_id": "EOF",
				"data": {
					"test": 1
				}
			}, {
				"expr": "b==2",
				"next_floor_id": "EOF",
				"data": {
					"test": 2
				}
			}]
		}
	]
}`
	testMatchTree := MatchTree{}
	err := json.Unmarshal([]byte(testMatchTreeJson), &testMatchTree)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name    string
		args    map[string]interface{}
		want    string
		wanterr error
	}{
		{
			name: " a==1 && b==1，common logic",
			args: map[string]interface{}{
				"a": 1,
				"b": 1,
			},
			want: `{"test":1}`,
		},
		{
			name: " a==3 ，first floor get target",
			args: map[string]interface{}{
				"a": 3,
				"b": 1,
			},
			want: `{"test":3}`,
		},
		{
			name: " a==4 ，first not get target",
			args: map[string]interface{}{
				"a": 4,
				"b": 1,
			},
			wanterr: common.MatchTreeFirstNotMatch,
			want:    "",
		},
		{
			name: " a==1 && b==4 ，middle floor not get target",
			args: map[string]interface{}{
				"a": 1,
				"b": 4,
			},
			wanterr: common.MatchTreeNotMatch,
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := TextMatch(tt.args, testMatchTree)
			var wantObj interface{}
			if tt.want != "" {
				json.Unmarshal([]byte(tt.want), &wantObj)
			}

			if !reflect.DeepEqual(res, wantObj) {
				t.Errorf("TestTextMatchNormal() res = %v, want %v", res, tt.want)
			}
			if err != tt.wanterr {
				t.Errorf("TestTextMatchNormal() err = %v, wanterr %v", err, tt.wanterr)
			}
		})
	}
}

func TestTextMatchByBytes(t *testing.T) {
	testMatchTreeJson := `{
	"match_floors": [{
			"one_floor_node1": [{
					"expr": "a==1",
					"next_floor_id": "two_floor_node1",
					"data": {}
				},
				{
					"expr": "a==2",
					"next_floor_id": "two_floor_node1",
					"data": {}
				},
				{
					"expr": "a==3",
					"next_floor_id": "EOF",
					"data": {
						"test": 3
					}
				}
			]
		},
		{
			"two_floor_node1": [{
				"expr": "b==1",
				"next_floor_id": "EOF",
				"data": {
					"test": 1
				}
			}, {
				"expr": "b==2",
				"next_floor_id": "EOF",
				"data": {
					"test": 2
				}
			}]
		}
	]
}`

	tests := []struct {
		name    string
		args    map[string]interface{}
		want    string
		wanterr error
	}{
		{
			name: " a==1 && b==1，common logic",
			args: map[string]interface{}{
				"a": 1,
				"b": 1,
			},
			want: `{"test":1}`,
		},
		{
			name: " a==3 ，first floor get target",
			args: map[string]interface{}{
				"a": 3,
				"b": 1,
			},
			want: `{"test":3}`,
		},
		{
			name: " a==4 ，first not get target",
			args: map[string]interface{}{
				"a": 4,
				"b": 1,
			},
			wanterr: common.MatchTreeFirstNotMatch,
			want:    "",
		},
		{
			name: " a==1 && b==4 ，middle floor not get target",
			args: map[string]interface{}{
				"a": 1,
				"b": 4,
			},
			wanterr: common.MatchTreeNotMatch,
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := TextMatchByBytes(tt.args, []byte(testMatchTreeJson))
			var wantObj interface{}
			if tt.want != "" {
				json.Unmarshal([]byte(tt.want), &wantObj)
			}

			if !reflect.DeepEqual(res, wantObj) {
				t.Errorf("TestTextMatchNormal() res = %v, want %v", res, tt.want)
			}
			if err != tt.wanterr {
				t.Errorf("TestTextMatchNormal() err = %v, wanterr %v", err, tt.wanterr)
			}
		})
	}
}

func TestTextMatchError(t *testing.T) {
	testMatchTreeJson := `{}`
	testMatchTree := MatchTree{}
	err := json.Unmarshal([]byte(testMatchTreeJson), &testMatchTree)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name    string
		args    map[string]interface{}
		want    interface{}
		wanterr error
	}{
		{
			name: "match tree is empty",
			args: map[string]interface{}{
				"a": 1,
				"b": 1,
			},
			wanterr: common.MatchTreeEmpty,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := TextMatch(tt.args, testMatchTree)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("TestTextMatchNormal() res = %v, want %v", res, tt.want)
			}
			if err != tt.wanterr {
				t.Errorf("TestTextMatchNormal() err = %v, wanterr %v", err, tt.wanterr)
			}
		})
	}
}

// compare time consuming without match tree
func TestTextMatchCompare(t *testing.T) {
	dataMap := map[string]interface{}{
		"a": 1,
		"b": 1,
		"c": 1,
	}
	t1 := time.Now() // get current time
	testMatchTreeJson := `{
    "match_floors": [
        {
            "one_floor_node1": [
                {
                    "expr": "a==1",
                    "next_floor_id": "two_floor_node1",
					"cache_key":["a"],
                    "data": {
                        
                    }
                }
            ]
        },
        {
            "two_floor_node1": [
                {
                    "expr": "b==1",
                    "next_floor_id": "three_floor_node1",
					"cache_key":["b"],
                    "data": {
                        
                    }
                }
            ]
        },
        {
            "three_floor_node1": [
                {
                    "expr": "c==1",
                    "next_floor_id": "EOF",
					"cache_key":["c"],
                    "data": {
                        "test": 1
                    }
                }
            ]
        }
    ]
}`
	testMatchTree := MatchTree{}
	err := json.Unmarshal([]byte(testMatchTreeJson), &testMatchTree)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 100000; i++ {
		TextMatch(dataMap, testMatchTree)
	}
	fmt.Println("use match tree: ", time.Since(t1))
	fmt.Println("cache map value: ")
	cacheMap.Range(func(k, v interface{}) bool {
		fmt.Println(k, " ", v)
		return true
	})

	t2 := time.Now() // get current time
	for i := 0; i < 100000; i++ {
		expression1, _ := govaluate.NewEvaluableExpression("a==1 && b==3 && c==1")
		expression1.Evaluate(dataMap)
		expression2, _ := govaluate.NewEvaluableExpression("a==1 && b==1 && c==1")
		expression2.Evaluate(dataMap)
	}
	fmt.Println("not use match tree: ", time.Since(t2))
}
