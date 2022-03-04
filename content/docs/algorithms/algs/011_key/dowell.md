---
title: 211223key
weight: 1
bookHidden: true
---

### 211223作业

```

type BrowserHistory struct {
	history []string
	cur     int
}
// 初始化
func Constructor(homepage string) (b BrowserHistory) {
	b.cur = 0
	b.history = []string{homepage}
	return b
}

// 访问
func (b *BrowserHistory) Visit(url string) {
	sum := len(b.history)
	if b.cur == sum-1 {
		b.history = append(b.history, url)
	} else if b.cur < sum-1 {
		b.history = b.history[:b.cur+1]
		b.history = append(b.history, url)
	}
	b.cur++
}

// 后退
func (b *BrowserHistory) Back(steps int) string {
	if steps > b.cur {
		b.cur = 0
		return b.history[0]
	}
	b.cur -= steps
	return b.history[b.cur]
}

// 前进
func (b *BrowserHistory) Forward(steps int) string {
	sum := len(b.history)
	if b.cur == sum-1 {
		return b.history[b.cur]
	} else if b.cur+steps > sum-1 {
		b.cur = sum - 1
		return b.history[sum-1]
	} else {
		b.cur += steps
		return b.history[b.cur]
	}
}
```