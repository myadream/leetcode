package kidsWithCandies

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	//demo1
	var candies = []int{2,3,5,1,3}
	var extracandies = 	3
	var target = []bool{true,true,true,false,true}

	var result = kidsWithCandies(candies, extracandies)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	//demo2
	candies = []int{4,2,1,1,2}
	extracandies = 	1
	target = []bool{true,false,false,false,false}

	result = kidsWithCandies(candies, extracandies)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	//demo3
	candies = []int{12,1,12}
	extracandies = 	10
	target = []bool{true,false,true}

	result = kidsWithCandies(candies, extracandies)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}
}