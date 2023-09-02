package interval

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
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

	//validate input
	for _, interval := range intervals {
		err := interval.Validate()
		if err != nil {
			return []Interval[T]{}, err
		}
	}

	// consideration: this will modify this initial slice: in order to avoid this a copy could be created which would cause additional allocation
	sort.Slice(intervals, func(a, b int) bool {
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

	var result []Interval[T]
	var current *Interval[T]
	for i, interval := range intervals {
		if current == nil {
			// important to use intervals[i] instead of interval since we don't want a pointer to the interval variable that will change in every iteration
			current = &intervals[i]
			continue
		}
		if merged, didMerge := current.merge(interval); didMerge {
			current = &merged
		} else {
			result = append(result, *current)
			current = &intervals[i]
		}
	}
	if current != nil {
		result = append(result, *current)
	}

	return result, nil
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
	return fmt.Sprintf("%s%v,%v%s", b0, interval.Min.Value, interval.Max.Value, b1)
}

func Parse[T Numeric](input string) (Interval[T], error) {
	input = strings.TrimSpace(input)
	iLen := len(input)
	if iLen < 5 { // minitmal input : "[1,2]"
		return Interval[T]{}, errors.New(fmt.Sprintf("invalid input %s", input))
	}

	minIsOpen := false
	first := input[0:1]
	if first == "(" {
		minIsOpen = true
	} else if first != "[" {
		return Interval[T]{}, errors.New(fmt.Sprintf("invalid input min limit %s : %s", first, input))
	}
	maxIsOpen := false
	last := input[iLen-1 : iLen]
	if last == ")" {
		maxIsOpen = true
	} else if last != "]" {
		return Interval[T]{}, errors.New(fmt.Sprintf("invalid input max limit %s : %s", last, input))
	}
	numbers := strings.Split(input[1:iLen-1], ",")
	if len(numbers) != 2 {
		return Interval[T]{}, errors.New(fmt.Sprintf("invalid input values %s", input))
	}
	minValue, err := strconv.ParseFloat(numbers[0], 64)
	if err != nil {
		return Interval[T]{}, errors.New(fmt.Sprintf("invalid input min value %s : %v", input, err))
	}
	maxValue, err := strconv.ParseFloat(numbers[1], 64)
	if err != nil {
		return Interval[T]{}, errors.New(fmt.Sprintf("invalid input max value %s : %v", input, err))
	}
	return Interval[T]{
		Min: Limit[T]{Value: T(minValue), Open: minIsOpen},
		Max: Limit[T]{Value: T(maxValue), Open: maxIsOpen},
	}, nil
}

func (interval Interval[T]) Validate() error {
	if interval.Max.Value < interval.Min.Value {
		return errors.New(fmt.Sprintf("invalid interval %s", interval))
	}
	if interval.Max.Value == interval.Min.Value &&
		(interval.Max.Open || interval.Min.Open) {
		return errors.New(fmt.Sprintf("invalid interval %s", interval))
	}
	return nil
}

func (interval Interval[T]) merge(other Interval[T]) (Interval[T], bool) {

	if interval.Max.Value < other.Min.Value {
		return Interval[T]{}, false
	} else if interval.Max.Value == other.Min.Value &&
		(interval.Max.Open || other.Min.Open) {
		return Interval[T]{}, false
	}

	//yes we do have an intersection
	if interval.Max.Value > other.Max.Value { // interval has the greater upper limit
		return Interval[T]{
			Min: interval.Min,
			Max: interval.Max,
		}, true
	} else if interval.Max.Value == other.Max.Value { // limits are semi equal
		return Interval[T]{
				Min: interval.Min,
				Max: Limit[T]{
					Value: other.Max.Value,
					Open:  interval.Max.Open && other.Max.Open,
				},
			},
			true
	}
	// other has the upper limit
	return Interval[T]{
		Min: interval.Min,
		Max: other.Max,
	}, true
}
