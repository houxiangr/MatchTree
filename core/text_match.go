package core

import (
	"encoding/json"
	"github.com/Knetic/govaluate"
	"github.com/textMatch/common"
)

//use by bytes config
func TextMatchByBytes(matchData MatchData, matchTreeBytes []byte) (interface{}, error) {
	var matchTree MatchTree
	err := json.Unmarshal(matchTreeBytes, &matchTree)
	if err != nil {
		return nil, err
	}
	return TextMatch(matchData, matchTree)
}

//use by MatchTree struct
func TextMatch(matchData MatchData, matchTree MatchTree) (interface{}, error) {
	if matchTree.MatchFloors == nil || len(matchTree.MatchFloors) == 0 {
		return nil, common.MatchTreeEmpty
	}
	var nextFloorId string

	//deal first floor
	for _, floorMap := range matchTree.MatchFloors[0] {
		for _, line := range floorMap {
			result, err := matchOneLine(line, matchData)
			if err != nil {
				return nil, err
			}
			if result {
				// is first floor get target,return
				if line.IsHaveNextNode() {
					return line.Data, nil
				}
				nextFloorId = line.NextFloorId
				break
			}
		}
	}
	//first floor not match
	if nextFloorId == "" {
		return nil, common.MatchTreeFirstNotMatch
	}

	//deal other floor
	floorLen := len(matchTree.MatchFloors)
	for i := 1; i < floorLen; i++ {
		lineMap := matchTree.MatchFloors[i]
		targetLines := lineMap[nextFloorId]
		for _, line := range targetLines {
			result, err := matchOneLine(line, matchData)
			if err != nil {
				return nil, err
			}
			if result {
				if line.IsHaveNextNode() {
					return line.Data, nil
				}
				nextFloorId = line.NextFloorId
				break
			}
		}
	}

	return nil, common.MatchTreeNotMatch
}

func matchOneLine(line MatchTreeLine, matchData MatchData) (bool, error) {
	if IsHaveCache(matchData,line) {
		return GetCache(matchData,line)
	}
	expression, err := govaluate.NewEvaluableExpression(line.Expr)
	if err != nil {
		return false, err
	}
	result, err := expression.Evaluate(matchData)
	if err != nil {
		return false, err
	}
	SetCache(matchData,line,result)
	return result.(bool), nil
}
