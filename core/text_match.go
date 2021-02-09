package core

import (
	"github.com/Knetic/govaluate"
	"github.com/textMatch/common"
)

func TextMatch(matchMap map[string]interface{}, matchTree MatchTree) (interface{}, error) {
	if matchTree.MatchFloors == nil || len(matchTree.MatchFloors) == 0 {
		return nil, common.MatchTreeEmpty
	}
	var nextFloorId string

	//deal first floor
	for _, floorMap := range matchTree.MatchFloors[0] {
		for _, line := range floorMap {
			result, err := matchOneLine(line, matchMap)
			if err != nil {
				return nil, err
			}
			if result {
				//todo test 如果第一层已经结束了，直接返回
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
		return nil,common.MatchTreeFirstNotMatch
	}

	//deal other floor
	floorLen := len(matchTree.MatchFloors)
	for i := 1; i < floorLen; i++ {
		lineMap := matchTree.MatchFloors[i]
		targetLines := lineMap[nextFloorId]
		for _, line := range targetLines {
			result, err := matchOneLine(line, matchMap)
			if err != nil {
				return nil, err
			}
			if result {
				if line.IsHaveNextNode() {
					return line.Data,nil
				}
				nextFloorId = line.NextFloorId
				break
			}
		}
	}

	return nil,common.MatchTreeNotMatch
}

func matchOneLine(line MatchTreeLine, matchMap map[string]interface{}) (bool, error) {
	expression, err := govaluate.NewEvaluableExpression(line.Expr)
	if err != nil {
		return false, err
	}
	result, err := expression.Evaluate(matchMap)
	if err != nil {
		return false, err
	}
	return result.(bool), nil
}
