package main

import (
	"reflect"
	"testing"
)

func check[T comparable](f, s []T) bool {
	return reflect.DeepEqual(f, s)
}

func errorf(r, need any, t *testing.T) {
	t.Errorf("not equal - expected %v - got %v", need, r)
}

func TestUnique(t *testing.T) {
	{
		r := Unique([]int{1, 2, 3, 4, 5})
		need := []int{1, 2, 3, 4, 5}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]string{"a", "b", "c"})
		need := []string{"a", "b", "c"}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]string{"a", "a", "b", "c"})
		need := []string{"a", "b", "c"}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]string{"a", "a", "b", "b", "c"})
		need := []string{"a", "b", "c"}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]string{"a", "a", "b", "c", "c"})
		need := []string{"a", "b", "c"}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]string{"a", "a", "b", "c", "c", "a", "b"})
		need := []string{"a", "b", "c"}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]string{"a", "a", "b", "c", "c", "a", "b"})
		need := []string{"a", "b", "c"}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]string{"a", "a", "a", "b", "c", "c", "a", "b"})
		need := []string{"a", "b", "c"}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]string{})
		need := []string{}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]int{1, 1, 1})
		need := []int{}
		if check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]int{1, 1, 1})
		need := []int{1}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]string{"a", "c"})
		need := []string{"c", "a"}
		if check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]float64{})
		need := []float64{}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]float64{1., 2., 3.})
		need := []float64{}
		if check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]float64{1., 2., 3., 1., 1, 1, 1, 1, 1, 1, 2, 0})
		need := []float64{1, 2, 3}
		if check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]float64{1., 2., 3., 1., 1, 1, 1, 1, 1, 1, 2, 0})
		need := []float64{1, 2, 3, 0}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]int{100, 200, 300})
		need := []int{}
		if check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]int{})
		need := []int{100, 200, 300}
		if check(r, need) {
			errorf(r, need, t)
		}
	}

	{
		r := Unique([]string{})
		need := []string{}
		if !check(r, need) {
			errorf(r, need, t)
		}
	}
}

func TestMD5(t *testing.T) {

}
