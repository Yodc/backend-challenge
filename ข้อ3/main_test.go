package main

import (
	"reflect"
	"testing"
)

func TestReplaceData(t *testing.T) {
	sampleText := "t-bone fatback pastrami t-bone pork meatloaf jowl .. enim t-bone"
	expected := "t-bone fatback pastrami t-bone pork meatloaf jowl  enim t-bone"
	if result := replaceData(sampleText); result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestCountBeef(t *testing.T) {
	sampleText := "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone"
	expected := map[interface{}]int{
		"t-bone":   4,
		"fatback":  1,
		"pastrami": 1,
		"pork":     1,
		"meatloaf": 1,
		"jowl":     1,
		"enim":     1,
		"bresaola": 1,
	}

	result := countBeef(replaceData(sampleText))
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected map %v, but got %v", expected, result)
	}
}
