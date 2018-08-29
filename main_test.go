package main

import (
	"reflect"
	"testing"
)

// Test the calculation results from a test.csv file
func TestCalculate(t *testing.T) {
	var filePath = "./test.csv"
	cal := new(Calculator)
	cal.readFile(filePath)
	cal.calculate()
	cal.unchargeHighestCaller()

	for idx, vCaller := range cal.callers {
		switch idx {
		case 0:
			if vCaller.numberId != "+351914374373" {
				t.Error("Expected +351914374373, got ", vCaller.numberId)
			}
			if vCaller.totalSeconds != 558 {
				t.Error("Expected 558, got ", vCaller.totalSeconds)
			}
			if vCaller.totalPrice != 0.452 {
				t.Error("Expected 0.452, got ", vCaller.totalPrice)
			}
		case 1:
			if vCaller.numberId != "+351914374375" {
				t.Error("Expected +351914374375, got ", vCaller.numberId)
			}
			if vCaller.totalSeconds != 292 {
				t.Error("Expected 292, got ", vCaller.totalSeconds)
			}
			if vCaller.totalPrice != 0.24333333333333335 {
				t.Error("Expected 0.24333333333333335, got ", vCaller.totalPrice)
			}
		case 2:
			if vCaller.numberId != "+351217538223" {
				t.Error("Expected +351217538223, got ", vCaller.numberId)
			}
			if vCaller.totalSeconds != 3825 {
				t.Error("Expected 3825, got ", vCaller.totalSeconds)
			}
			if vCaller.totalPrice != 0 {
				t.Error("Expected 0, got ", vCaller.totalPrice)
			}
		}
	}
}

// Test the calculation for call duration
func TestCalculateCallDuration(t *testing.T) {
	var startTime = "21:03:04"
	var finishTime = "21:06:01"
	callDuration := calculateCallDuration(startTime, finishTime)
	if callDuration != 177.0 {
		t.Error("Expected 177.0, got ", callDuration)
	}
}

// Test the calculation for call price
func TestCalculateCallPrice(t *testing.T) {
	var callDuration = 200.0
	callPrice := calculateCallPrice(callDuration)
	if callPrice != 0.16666666666666666 {
		t.Error("Expected 0.16666666666666666, got ", callPrice)
	}

	callDuration = 500.0
	callPrice = calculateCallPrice(callDuration)
	if callPrice != 0.31666666666666665 {
		t.Error("Expected 0.31666666666666665, got ", callPrice)
	}
}

// Test the type results to convert string time to an array of ints
func TestConvertTimeStringToInt(t *testing.T) {
	var stringTime = "12:12:12"
	var typeCheck = 12
	timeInt := convertTimeStringToInt(stringTime)
	hour := timeInt[0]
	minute := timeInt[1]
	second := timeInt[2]
	if reflect.TypeOf(hour) != reflect.TypeOf(typeCheck) {
		t.Error("Expected int, got ", reflect.TypeOf(hour))
	}
	if reflect.TypeOf(minute) != reflect.TypeOf(typeCheck) {
		t.Error("Expected int, got ", reflect.TypeOf(minute))
	}
	if reflect.TypeOf(hour) != reflect.TypeOf(typeCheck) {
		t.Error("Expected int, got ", reflect.TypeOf(second))
	}
}
