# MatchTree
输入map[string]interface和决策树节点。
输出决策出的最后的结果

通过将表达式变成一个决策树，减少判断条件的次数

example:
`a==1 && b==1 && c==1`
to
`{
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
