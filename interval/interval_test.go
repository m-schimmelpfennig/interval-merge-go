package interval

import (
	"fmt"
	"math"
	"testing"
)

func TestIntervalMerge(t *testing.T) {
	runMergeTest[int](t, math.MaxInt, math.MaxInt)
	runMergeTest[int8](t, math.MinInt8, math.MaxInt8)
	runMergeTest[int16](t, math.MinInt8, math.MaxInt8)
	runMergeTest[int32](t, math.MinInt8, math.MaxInt8)
	runMergeTest[int64](t, math.MinInt8, math.MaxInt8)
	runMergeTest[float32](t, math.MinInt8, math.MaxInt8)
	runMergeTest[float64](t, math.MinInt8, math.MaxInt8)
}

type mergeTest[T Numeric] struct {
	name       string
	intervals  []Interval[T]
	wantResult []Interval[T]
	wantErr    bool
}

func runMergeTest[T Numeric](t *testing.T, minValue T, maxValue T) {

	tests := []mergeTest[T]{
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
		{
			name: fmt.Sprintf("NegativeValues %T", *new(T)),
			intervals: []Interval[T]{
				{
					Min: Limit[T]{Value: T(-20), Open: false},
					Max: Limit[T]{Value: T(-10), Open: false},
				},
				{
					Min: Limit[T]{Value: T(-10), Open: false},
					Max: Limit[T]{Value: T(-5), Open: false},
				},
			},
			wantResult: []Interval[T]{
				{
					Min: Limit[T]{Value: T(-20), Open: false},
					Max: Limit[T]{Value: T(-5), Open: false},
				},
			},
			wantErr: false,
		},
		{
			name: fmt.Sprintf("MinimalInervals %T", *new(T)),
			intervals: []Interval[T]{
				{
					Min: Limit[T]{Value: T(6), Open: false},
					Max: Limit[T]{Value: T(6), Open: false},
				},
				{
					Min: Limit[T]{Value: T(6), Open: false},
					Max: Limit[T]{Value: T(7), Open: false},
				},
			},
			wantResult: []Interval[T]{
				{
					Min: Limit[T]{Value: T(6), Open: false},
					Max: Limit[T]{Value: T(7), Open: false},
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

type validationTest[T Numeric] struct {
	name         string
	interval     Interval[T]
	wantErr      bool
	errorMessage string
}

func TestValidation(t *testing.T) {
	runValidationTest[int](t)
}

func runValidationTest[T Numeric](t *testing.T) {

	tests := []validationTest[T]{
		{
			name: "NoErrorExpected",
			interval: Interval[T]{
				Min: Limit[T]{Value: 25, Open: false},
				Max: Limit[T]{Value: 30, Open: false},
			},
			wantErr:      false,
			errorMessage: "",
		},
		{
			name: "BasicErrorCase",
			interval: Interval[T]{
				Min: Limit[T]{Value: 2, Open: false},
				Max: Limit[T]{Value: -1, Open: false},
			},
			wantErr:      true,
			errorMessage: "invalid interval [2,-1]",
		},
		{
			name: "MaxClosedMinOpen",
			interval: Interval[T]{
				Min: Limit[T]{Value: -2, Open: false},
				Max: Limit[T]{Value: -2, Open: true},
			},
			wantErr:      true,
			errorMessage: "invalid interval [-2,-2)",
		},
		{
			name: "MinOpenClosedMax",
			interval: Interval[T]{
				Min: Limit[T]{Value: -5, Open: true},
				Max: Limit[T]{Value: -5, Open: false},
			},
			wantErr:      true,
			errorMessage: "invalid interval (-5,-5]",
		},
		{
			name: "BothOpen",
			interval: Interval[T]{
				Min: Limit[T]{Value: 7, Open: true},
				Max: Limit[T]{Value: 7, Open: true},
			},
			wantErr:      true,
			errorMessage: "invalid interval (7,7)",
		},
		{
			name: "BothClosed",
			interval: Interval[T]{
				Min: Limit[T]{Value: 0, Open: false},
				Max: Limit[T]{Value: 0, Open: false},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.interval.Validate()
			if tt.wantErr {
				if err == nil {
					t.Error("no error returned but an error was expected")
				} else {
					if err.Error() != tt.errorMessage {
						t.Errorf("expected error %s did not match returend error %v", tt.errorMessage, err)
					}
				}
			} else {
				if err != nil {
					t.Error("error returned but no error was expected")
				}
			}
		})
	}
}
