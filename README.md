# slicemeta

Provides fast slice utilities that don't rely on "reflect" usage but generates package for specified element type.

Comparison with DeepEqual for int slice with 100k elements:

    pkg: github.com/azavorotnii/slicemeta/internal/testing/builtin_test
    BenchmarkEqual/intutil-8                   20000             61676 ns/op
    BenchmarkEqual/DeepEqual-8                   100          10452532 ns/op

# Install

    go get -u github.com/azavorotnii/slicemeta

# Examples

For usage examples, see tests in "internal/testing" folder. You will need to run first:

    $ cd ./internal/testing
    $ go generate ./...

Now tests are fully working and should be a good example of usage.

#Usage

For example we need utilities for time.Time slices. Lets generate it with:

    //go:generate slicemeta -type time.Time -import "time"

We will get package "timeutil" with following methods:

    func Contains(in []time.Time, value time.Time) bool
    func ContainsAny(in []time.Time, values ...time.Time) bool
    func ContainsFunc(in []time.Time, f func(time.Time) bool) bool
    func Count(in []time.Time, value time.Time) int
    func CountAny(in []time.Time, values ...time.Time) int
    func CountFunc(in []time.Time, f func(time.Time) bool) int
    func Equal(a, b []time.Time) bool
    func Filter(in []time.Time, f func(time.Time) bool) []time.Time
    func Index(in []time.Time, value time.Time) int
    func IndexAny(in []time.Time, values ...time.Time) int
    func IndexFunc(in []time.Time, f func(time.Time) bool) int
    func LastIndex(in []time.Time, value time.Time) int
    func LastIndexAny(in []time.Time, values ...time.Time) int
    func LastIndexFunc(in []time.Time, f func(time.Time) bool) int
    func Map(in []time.Time, f func(time.Time) time.Time) []time.Time
    func Reduce(in []time.Time, f func(time.Time, time.Time) time.Time) time.Time
    func Shuffle(in []time.Time)

By default, "==" operator used to compare if values are equal:

    func Contains(in []time.Time, value time.Time) bool {
		for _, v := range in {
			if v == value {
				return true
			}
		}
		return false
    }

This is fastest but not a correct way for time.Time objects as will not offset timestamps to same timezone for
comparison. Method "Equal" will be used if option "equal" provided as:

    //go:generate slicemeta -type time.Time -import "time" -equal method

As a result, generated code will look like:

    func Contains(in []time.Time, value time.Time) bool {
        for _, v := range in {
            if v.Equal(value) {
                return true
            }
        }
        return false
    }

When objects requires DeepEqual for comparison, option "equal" should be "deep":

    func Contains(in []time.Time, value time.Time) bool {
        for _, v := range in {
            if reflect.DeepEqual(v, value) {
                return true
            }
        }
        return false
    }

If custom method needed for equality check, format template can be provided:

    //go:generate slicemeta -type time.Time -import "time" -equal "%v.Unix() == %v.Unix()"

Result will be:

    func Contains(in []time.Time, value time.Time) bool {
        for _, v := range in {
            if v.Unix() == value.Unix() {
                return true
            }
        }
        return false
    }

Sorting wrapper created if "less" option provided. It can be "operator" to use "<" for comparison or format template string similar to one above:

    //go:generate slicemeta -type time.Time -import "time" -equal method -less "%v.Before(%v)"

As a result, wrapper that satisfies sort.Interface created:

    type TimeSlice []time.Time

    func (s TimeSlice) Len() int {
        return len(s)
    }
    func (s TimeSlice) Swap(i, j int) {
        s[i], s[j] = s[j], s[i]
    }
    func (s TimeSlice) Less(i, j int) bool {
        return s[i].Before(s[j])
    }

When option "methods" provided, only methods which names satisfy "methods" regexp will be generated:

    //go:generate slicemeta -type time.Time -import "time" -equal method -methods "Contains|Count|Equal"

Only following methods will be generated:

    func Contains(in []time.Time, value time.Time) bool
    func ContainsAny(in []time.Time, values ...time.Time) bool
    func ContainsFunc(in []time.Time, f func(time.Time) bool) bool
    func Count(in []time.Time, value time.Time) int
    func CountAny(in []time.Time, values ...time.Time) int
    func CountFunc(in []time.Time, f func(time.Time) bool) int
    func Equal(a, b []time.Time) bool

# Authors

 * Andrii Zavorotnii (andrii.zavorotnii@gmail.com)
