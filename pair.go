package main

import (
	"sort"
)

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value string
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { tem := p[i]; p[i]=p[j];p[j]=tem}
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return len(p[i].Value) > len(p[j].Value) }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]string) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	return p
}