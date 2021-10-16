package _05_CDS

import (
	"strconv"
	"strings"
)

type IPV4 struct {
	ans []string
}

func NewIPV4() *IPV4 {
	return &IPV4{}
}

func (I *IPV4) IPFilter(data string) (res []string) {
	// 1 不在范围内的
	if len(data) > 12 || len(data) < 4 {
		return []string{}
	}
	// 2 不是纯数字的
	for _, v := range data {
		if v < '0' || v > '9' {
			return []string{}
		}
	}
	I.ans = make([]string, 0)
	path := make([]string, 0, 4)
	I.dfs(data, 0, path, 0)
	return I.ans
}

func (I *IPV4) dfs(str string, start int, path []string, segmentSum int) {
	if len(str) == start && segmentSum == 4 {
		I.ans = append(I.ans, strings.Join(path, "."))
		return
	}

	for i := start; i < start+3; i++ {
		if i == len(str) {
			break
		}
		if !I.isAvailableIPSegment(segmentSum, len(str)-i) {
			continue
		}
		if I.isIPSegment(str[start : i+1]) {
			path = append(path, str[start:i+1])
			I.dfs(str, i+1, path, segmentSum+1)
			path = append(path[:len(path)-1], path[len(path):]...)
		}
	}
}

func (I *IPV4) isIPSegment(str string) bool {
	if len(str) == 0 || len(str) > 3 {
		return false
	}
	if len(str) > 1 && str[0] == '0' {
		return false
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return num >= 0 && num <= 255
}

func (I *IPV4) isAvailableIPSegment(segmentSum int, surplusStrLen int) bool {
	if segmentSum == 0 {
		return true
	}
	if segmentSum == 4 && surplusStrLen > 0 {
		return false
	}
	min := (4 - segmentSum) * 1
	max := (4 - segmentSum) * 3
	return surplusStrLen >= min && surplusStrLen <= max
}
