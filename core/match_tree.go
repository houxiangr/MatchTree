package core

import "github.com/textMatch/common"

type MatchTree struct {
	MatchFloors []MatchTreeFLoorMap `json:"match_floors"`
}

type MatchTreeFLoorMap map[string][]MatchTreeLine

type MatchTreeLine struct {
	Expr        string      `json:"expr"`
	NextFloorId string      `json:"next_floor_id"`
	Data        interface{} `json:"data"`
}

func (this *MatchTreeLine) IsHaveNextNode() bool {
	return this.NextFloorId == common.NextNodeStatusEOF
}
