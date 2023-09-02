# interval-merge-go
## Demo Project for Mercedes-Benz Tech Innovation

This project has been created to demonstrate my development skills. Its functionality is to merge all overlapping intervals from a given set and output the result set of merged and unmerged intervals.

## Development Requirements

In oder to build this project the go1.19 is required. This project has been build and tested under Ubuntu 22.04.3 LTS. 
Building and testing this project on Windows and Mac may require changes to the Makefile.



## Build
    
The projects Makefile contains the following target to build an executable file.

```shell
make build
```
This will generate an executable file `interval-merge`

## Commandline Usage

Once the `interval-merge` is build it can be invoked with a command line argument that provides the intervals to merge.
Note: the interval definition has to be quoted in oder to shield it from bash interpretation.
Open or closed interval limits or endpoint can be defined using parentheses or square brackets
See mathematical notation https://en.wikipedia.org/wiki/Interval_(mathematics)#Notations_for_intervals


```shell
./interval-merge "(1,2.34) [2,34] (-5,10)"
```

## Code Usage

The following example illustrates how to invoke the merge function when used as a library.

```Go
intervals := []interval.Interval[int]{ // define a set of intervals to merge using generic numeric type
    {
        Min: interval.Limit[int]{Value: 3, Open: false},
        Max: interval.Limit[int]{Value: 5, Open: true},
    },
    {
        Min: interval.Limit[int]{Value: 5, Open: false},
        Max: interval.Limit[int]{Value: 10, Open: false},
    },
}

res, err := interval.Merge[int](intervals...)
```

### Errors

The merge function may return an error if the input contains invalid intervals. Intervals are considered invalid
once the minimum value is greater than the maximum value or the minimum and maximum values are equal but at least one of them is an open interval limit

## Unit Tests

Automated test can be executed form the projects root directory using 
```shell
go test ./interval -v
```
