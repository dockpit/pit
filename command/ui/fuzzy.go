package ui

import (
	"github.com/dockpit/pit/model"
)

var DefaultCost = &Cost{
	Del:  10,
	Ins:  1,
	Subs: 10}

type Entry struct {
	Distance  int
	Isolation *model.Isolation
}

type ByDistance []*Entry

func (s ByDistance) Len() int {
	return len(s)
}
func (s ByDistance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByDistance) Less(i, j int) bool {
	return s[i].Distance < s[j].Distance
}

type Cost struct {
	Del, Ins, Subs int
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minmin(a, b, c int) int {
	return min(min(a, b), c)
}

func Levenshtein(needle string, hays string, cost *Cost) int {
	if len(needle) == 0 {
		return len(hays) * cost.Ins
	}
	if len(hays) == 0 {
		return len(needle) * cost.Del
	}

	d := make([]int, len(hays)+1)
	e := make([]int, len(hays)+1)
	for i := 0; i < len(needle); i++ {
		e[0] = d[0] + 1
		for j := 0; j < len(hays); j++ {
			c := 0
			if needle[i] != hays[j] {
				c = cost.Subs
			}
			e[j+1] = minmin(d[j+1]+cost.Del, e[j]+cost.Ins, d[j]+c)
		}
		d, e = e, d
	}

	return d[len(d)-1]
}
