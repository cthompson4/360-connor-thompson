package main

import (
	"fmt"
	"testing"
	"example.com/validator/first"
	"example.com/validator/fourth"
	"example.com/validator/second"
	"example.com/validator/third"
)

func TestFirst(t *testing.T) {
	fmt.Println("Testing first:")
	result := first.FirstStandardTest()
	if (result != 2) {
		t.Error(
			"For", "test.txt",
			"expected", 2,
			"got", result,
		)
	}
}
func TestSecond(t *testing.T) {
	fmt.Println("Testing second:")
	result := second.SecondStandardTest()
	if (result != 1) {
		t.Error(
			"For", "testREADME.md",
			"expected", 1,
			"got", result,
		)
	}
}
func TestThird(t *testing.T) {
	fmt.Println("Testing third:")
	result := third.ThirdStandardTest()
	if (result != 0) {
		t.Error(
			"For", "LICENSE.lic",
			"expected", 0,
			"got", result,
		)
	}
}
func TestFourth(t *testing.T) {
	fmt.Println("Testing fourth:")
	result := fourth.FourthStandardTest()
	if (result != 3) {
		t.Error(
			"For", "test.txt",
			"expected", 3,
			"got", result,
		)
	}
}