package interval

import (
	"fmt"
	"math"
	"testing"
)

func TestIntervalMerge(t *testing.T) {
	run[int](t, math.MaxInt, math.MaxInt)
	run[int8](t, math.MinInt8, math.MaxInt8)
	run[int16](t, math.MinInt8, math.MaxInt8)
	run[int32](t, math.MinInt8, math.MaxInt8)
	run[int64](t, math.MinInt8, math.MaxInt8)
	run[float32](t, math.MinInt8, math.MaxInt8)
	run[float64](t, math.MinInt8, math.MaxInt8)
}

type test[T Numeric] struct {
	name       string
	intervals  []Interval[T]
	wantResult []Interval[T]
	wantErr    bool
}

func run[T Numeric](t *testing.T, minValue T, maxValue T) {

	tests := []test[T]{
		{
			name: fmt.Sprintf("OnlyOneElement %T", *new(T)),
			intervals: []Interval[T]{
				{
					Min: Limit[T]{Value: T(10), Open: false},
					Max: Limit[T]{Value: T(20), Open: false},
				},
			},
			wantErr: false,
			wantResult: []Interval[T]{
				{
					Min: Limit[T]{Value: T(10), Open: false},
					Max: Limit[T]{Value: T(20), Open: false},
				},
			},
		},
		{
			name: fmt.Sprintf("InvalidInput %T", *new(T)),
			intervals: []Interval[T]{
				{
					Min: Limit[T]{Value: T(20), Open: false},
					Max: Limit[T]{Value: T(10), Open: false},
				},
			},
			wantErr: true,
		},
		{
			name: fmt.Sprintf("ClosedOpenIntersection %T", *new(T)),
			intervals: []Interval[T]{
				{
					Min: Limit[T]{Value: T(3), Open: false},
					Max: Limit[T]{Value: T(5), Open: true},
				},
				{
					Min: Limit[T]{Value: T(5), Open: false},
					Max: Limit[T]{Value: T(10), Open: false},
				},
			},
			wantResult: []Interval[T]{
				{
					Min: Limit[T]{Value: T(3), Open: false},
					Max: Limit[T]{Value: T(5), Open: true},
				},
				{
					Min: Limit[T]{Value: T(5), Open: false},
					Max: Limit[T]{Value: T(10), Open: false},
				},
			},
			wantErr: false,
		},
		{
			name: fmt.Sprintf("OpenClosedIntersection %T", *new(T)),
			intervals: []Interval[T]{
				{
					Min: Limit[T]{Value: T(3), Open: false},
					Max: Limit[T]{Value: T(5), Open: false},
				},
				{
					Min: Limit[T]{Value: T(5), Open: true},
					Max: Limit[T]{Value: T(10), Open: false},
				},
			},
			wantResult: []Interval[T]{
				{
					Min: Limit[T]{Value: T(3), Open: false},
					Max: Limit[T]{Value: T(5), Open: false},
				},
				{
					Min: Limit[T]{Value: T(5), Open: true},
					Max: Limit[T]{Value: T(10), Open: false},
				},
			},
			wantErr: false,
		},
		{
			name: fmt.Sprintf("ClosedClosedIntersection %T", *new(T)),
			intervals: []Interval[T]{
				{
					Min: Limit[T]{Value: T(3), Open: false},
					Max: Limit[T]{Value: T(5), Open: false},
				},
				{
					Min: Limit[T]{Value: T(5), Open: false},
					Max: Limit[T]{Value: T(10), Open: false},
				},
			},
			wantResult: []Interval[T]{
				{
					Min: Limit[T]{Value: T(3), Open: false},
					Max: Limit[T]{Value: T(10), Open: false},
				},
			},
			wantErr: false,
		},
		{
			name: fmt.Sprintf("ClosedClosedIntersectionFlippedOrder %T", *new(T)),
			intervals: []Interval[T]{
				{
					Min: Limit[T]{Value: T(5), Open: false},
					Max: Limit[T]{Value: T(10), Open: false},
				},
				{
					Min: Limit[T]{Value: T(3), Open: false},
					Max: Limit[T]{Value: T(5), Open: false},
				},
			},
			wantResult: []Interval[T]{
				{
					Min: Limit[T]{Value: T(3), Open: false},
					Max: Limit[T]{Value: T(10), Open: false},
				},
			},
			wantErr: false,
		},
		{
			name: fmt.Sprintf("MinMaxValues %T", *new(T)),
			intervals: []Interval[T]{
				{
					Min: Limit[T]{Value: minValue, Open: false},
					Max: Limit[T]{Value: T(0), Open: false},
				},
				{
					Min: Limit[T]{Value: T(0), Open: false},
					Max: Limit[T]{Value: maxValue, Open: false},
				},
			},
			wantResult: []Interval[T]{
				{
					Min: Limit[T]{Value: minValue, Open: false},
					Max: Limit[T]{Value: maxValue, Open: false},
				},
			},
			wantErr: false,
		},
		{
			name: fmt.Sprintf("Example %T", *new(T)),
			intervals: []Interval[T]{
				{
					Min: Limit[T]{Value: 25, Open: false},
					Max: Limit[T]{Value: 30, Open: false},
				},
				{
					Min: Limit[T]{Value: 2, Open: false},
					Max: Limit[T]{Value: 19, Open: false},
				},
				{
					Min: Limit[T]{Value: 14, Open: false},
					Max: Limit[T]{Value: 23, Open: false},
				},
				{
					Min: Limit[T]{Value: 4, Open: false},
					Max: Limit[T]{Value: 8, Open: false},
				},
			},
			wantResult: []Interval[T]{
				{
					Min: Limit[T]{Value: 2, Open: false},
					Max: Limit[T]{Value: 23, Open: false},
				},
				{
					Min: Limit[T]{Value: 25, Open: false},
					Max: Limit[T]{Value: 30, Open: false},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			res, err := Merge[T](tt.intervals...)

			if tt.wantErr {
				if err == nil {
					t.Error("no error returned but an error was expected")
				}
			}

			for i, r := range res {
				exp := tt.wantResult[i]
				if r.Min.Open != exp.Min.Open {
					t.Errorf("expected min open %v did not match returend open %v", exp.Min.Open, r.Min.Open)
				}
				if r.Min.Value != exp.Min.Value {
					t.Errorf("expected min value %v did not match returend value %v", exp.Max.Value, r.Max.Value)
				}
				if r.Max.Open != exp.Max.Open {
					t.Errorf("expected max open %v did not match returend open %v", exp.Max.Open, r.Max.Open)
				}
				if r.Max.Value != exp.Max.Value {
					t.Errorf("expected max value %v did not match returend value %v", exp.Max.Value, r.Max.Value)
				}
			}
		})
	}
}

//TODO implement validation test cases
