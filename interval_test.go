package interval

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {

	intervalA := Interval[float64]{
		Min: Limit[float64]{
			Value: 0,
			Open:  false,
		},
		Max: Limit[float64]{
			Value: 2,
			Open:  false,
		},
	}

	intervalB := Interval[float64]{
		Min: Limit[float64]{
			Value: -1,
			Open:  false,
		},
		Max: Limit[float64]{
			Value: 2,
			Open:  true,
		},
	}

	res, err := Merge(intervalA, intervalB)

	if err != nil {
		t.Error("failed to sort")
	}

	fmt.Println(res)
}
