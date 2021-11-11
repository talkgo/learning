---
title: 211104key
weight: 1
bookHidden: true
---

### 211104作业

* 示例代码

```
func tictactoe(board []string) (res string) {
	// 获取棋盘的大小
	isFull := false                // 0 default none space 1 exist space
	stringToNum := map[string]int{ // X:1 O:0 " ":-1000
		"X": 1,
		"O": 0,
		" ": -1000,
	}

	for i := 0; i < len(board); i++ { // 使用双层for循环来计算列和 列号
		var sum = 0                       // 列和初始化
		for j := 0; j < len(board); j++ { // 行号
			sum += stringToNum[string(board[j][i])]
		}
		if sum == len(board) { // winner : X
			return "X"
		} else if sum == 0 { // winner : O
			return "O"
		} else if sum < 0 {
			isFull = true
		}
	}

	for i := 0; i < len(board); i++ { // 使用双层for循环来计算列和 行号
		var sum = 0                       // 行和初始化
		for j := 0; j < len(board); j++ { // 列号
			sum += stringToNum[string(board[i][j])]
		}
		if sum == len(board) { // winner : X
			return "X"
		} else if sum == 0 { // winner : O
			return "O"
		} else if sum < 0 {
			isFull = true
		}
	}

	var sum1 = 0                      // 右上左下对角线之和
	var sum2 = 0                      // 左上右下对角线之和
	for i := 0; i < len(board); i++ { // 计算对角线之和
		sum1 = sum1 + stringToNum[string(board[len(board)-i-1][i])]
		sum2 = sum2 + stringToNum[string(board[i][i])]
	}
	if sum1 == len(board) || sum2 == len(board) { // winner : X
		return "X"
	} else if sum1 == 0 || sum2 == 0 { // winner : O
		return "O"
	} else if sum1 < 0 || sum2 < 0 {
		isFull = true
	}

	if isFull { // 如果不存在判断是否有空格
		return "Pending"
	}

	return "Draw"

}
```
