package core

import (
	"github.com/textMatch/common"
	"sync"
)

var cacheMap = sync.Map{}

func SetCache(matchData MatchData, line MatchTreeLine, result interface{}) {
	if len(line.CacheKey) == 0 {
		return
	}
	cacheMap.Store(matchData.GetStringAddExpr(line), result)
}

func IsHaveCache(matchData MatchData, line MatchTreeLine) bool {
	if len(line.CacheKey) == 0 {
		return false
	}
	if _, ok := cacheMap.Load(matchData.GetStringAddExpr(line)); ok {
		return true
	}
	return false
}

func GetCache(matchData MatchData, line MatchTreeLine) (bool, error) {
	value, ok := cacheMap.Load(matchData.GetStringAddExpr(line))
	if !ok {
		return false, common.MatchCacheGetErr
	}
	return value.(bool), nil
}
