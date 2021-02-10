package core

import (
	"github.com/textMatch/common"
)

type MatchTree struct {
	MatchFloors []MatchTreeFLoorMap `json:"match_floors"`
}

type MatchTreeFLoorMap map[string][]MatchTreeLine

type MatchTreeLine struct {
	CacheKey    []string    `json:"cache_key"` //not require，can use cache features
	Expr        string      `json:"expr"`
	NextFloorId string      `json:"next_floor_id"`
	Data        interface{} `json:"data"`
}

func (this *MatchTreeLine) IsHaveNextNode() bool {
	return this.NextFloorId == common.NextNodeStatusEOF
}

type MatchData map[string]interface{}

//return key1value1key2value2expr
func (this MatchData) GetStringAddExpr(line MatchTreeLine) string {
	cacheKey := []byte{}
	for _, v := range line.CacheKey {
		tempCacheKey := []byte{}
		tempCacheKey = append(tempCacheKey, []byte(v)...)
		tempCacheKey = append(tempCacheKey, []byte(common.TransferInterfaceToString(this[v]))...)
		cacheKey = append(cacheKey, tempCacheKey...)
	}
	cacheKey = append(cacheKey, []byte(line.Expr)...)
	return string(cacheKey)
}
