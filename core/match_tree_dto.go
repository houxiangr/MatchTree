package core

import (
	"github.com/houxiangr/MatchTree/common"
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
func (this MatchData) GetStringAddExpr(line MatchTreeLine) (string, error) {
	cacheKey := []byte{}
	for _, v := range line.CacheKey {
		tempCacheKey := []byte{}
		tempCacheKey = append(tempCacheKey, []byte(v)...)
		value, err := common.TransferInterfaceToString(this[v])
		if err != nil {
			return "", err
		}
		tempCacheKey = append(tempCacheKey, []byte(value)...)
		cacheKey = append(cacheKey, tempCacheKey...)
	}
	cacheKey = append(cacheKey, []byte(line.Expr)...)
	return string(cacheKey), nil
}
