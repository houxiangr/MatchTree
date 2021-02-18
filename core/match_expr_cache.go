package core

import (
	"github.com/MatchTree/common"
	"sync"
)

var cacheMap = sync.Map{}

func SetCache(matchData MatchData, line MatchTreeLine, result interface{}) error {
	var err error
	var cacheKey string

	if len(line.CacheKey) == 0 {
		return nil
	}
	cacheKey, err = matchData.GetStringAddExpr(line)
	if err != nil {
		return err
	}
	cacheMap.Store(cacheKey, result)
	return nil
}

func IsHaveCache(matchData MatchData, line MatchTreeLine) (bool, error) {
	var err error
	var cacheKey string

	if len(line.CacheKey) == 0 {
		return false, nil
	}

	cacheKey, err = matchData.GetStringAddExpr(line)
	if err != nil {
		return false, err
	}
	if _, ok := cacheMap.Load(cacheKey); ok {
		return true, nil
	}
	return false, nil
}

func GetCache(matchData MatchData, line MatchTreeLine) (bool, error) {
	var err error
	var cacheKey string
	cacheKey, err = matchData.GetStringAddExpr(line)
	if err != nil {
		return false, err
	}
	value, ok := cacheMap.Load(cacheKey)
	if !ok {
		return false, common.MatchCacheGetErr
	}
	return value.(bool), nil
}
