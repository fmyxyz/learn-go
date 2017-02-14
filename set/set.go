package set

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	} else {
		return false
	}
}

func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

func (set *HashSet) Len() int {
	return len(set.m)
}

func (set *HashSet) Same(o *HashSet) bool {
	if o == nil {
		return false
	}
	if set.Len() != o.Len() {
		return false
	}
	for key := range set.m {
		if !o.Contains(key) {
			return false
		}

	}
	return true
}

func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

func (set *HashSet) IsSuperset(other *HashSet) bool {
	if other == nil {
		return false
	}
	setLen := set.Len()
	oLen := other.Len()
	if setLen == 0 || setLen == oLen {
		return false
	}
	if setLen > 0 && oLen == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

func (set *HashSet) Union(other *HashSet) *HashSet {
	union := NewHashSet()
	if other == nil {
		union.m = set.m
	} else {
		union.m = set.m
		for _, e := range other.Elements() {
			union.Add(e)
		}
	}
	return union
}

func (set *HashSet) Intersect(other *HashSet) *HashSet {
	intersect := NewHashSet()
	if other != nil {
		for _, e := range other.Elements() {
			if set.Contains(e) {
				intersect.Add(e)
			}
		}
	}
	return intersect
}

func (set *HashSet) Difference(other *HashSet) *HashSet {
	difference := NewHashSet()
	difference.m = set.m
	if other != nil {
		for _, v := range other.Elements() {
			if set.Contains(v) {
				difference.Remove(v)
			}
		}

	}
	return difference
}

func (set *HashSet) SymmetricDifference(other *HashSet) *HashSet {
	differenceSet := set.Difference(other)
	differenceOther := other.Difference(set)
	return differenceSet.Union(differenceOther)
}
