package interval

import (
	"fmt"
	"sort"
)

type Numeric interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | int
}

type Interval[T Numeric] struct {
	Min Limit[T]
	Max Limit[T]
}

type Limit[T Numeric] struct {
	Value T
	Open  bool
}

func Merge[T Numeric](intervals ...Interval[T]) ([]Interval[T], error) {
	// concideration: this will modify this initial slice: in order to avoid this a copy could be created which would cause additional allocation
	sort.SliceStable(intervals, func(a, b int) bool {
		intervalA := intervals[a]
		intervalB := intervals[b]
		if intervalA.Min.Value == intervalB.Min.Value {
			if intervalA.Min.Open && !intervalB.Min.Open {
				return true
			}
			return false
		}
		return intervalA.Min.Value < intervalB.Min.Value
	})

	return intervals, nil
}

func (interval Interval[T]) String() string {
	b0 := "["
	if interval.Min.Open {
		b0 = "("
	}
	b1 := "]"
	if interval.Max.Open {
		b1 = ")"
	}
	return fmt.Sprintf("%s %v, %v %s", b0, interval.Min.Value, interval.Max.Value, b1)
}

func (interval Interval[T]) merge(other Interval[T]) (Interval[T], bool) {

	return Interval[T]{}, false
}
