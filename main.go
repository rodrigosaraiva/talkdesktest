package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const firstMinutesPrice float64 = 0.05
const finalMinutesPrice float64 = 0.02
const minuteInSeconds float64 = 60
const firstMinutesInSeconds float64 = 300
const highestCallerDiscount float64 = 0

// Struct to store data from callers
type Caller struct {
	numberId     string
	totalSeconds float64
	totalPrice   float64
}

// Struct calculator that receives the file and group the callers
type Calculator struct {
	file    os.File
	callers []*Caller
}

// Main function
func main() {
	filePathFlag := flag.String("filepath", "", "Example of use: $ ./main -filepath=/dir/filename")
	flag.Parse()
	filePath := *filePathFlag

	cal := new(Calculator)
	cal.readFile(filePath)
	cal.calculate()
	cal.unchargeHighestCaller()
	cal.printResults()
}

// Method to read file
func (calc *Calculator) readFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	calc.file = *file
}

// Method that do the calculation duration and price for calls
func (calc *Calculator) calculate() {
	scanner := bufio.NewScanner(&calc.file)
	defer calc.file.Close()

	for scanner.Scan() {
		call := strings.Split(scanner.Text(), ";")
		callDuration := calculateCallDuration(call[0], call[1])
		callPrice := calculateCallPrice(callDuration)
		calc.updateCaller(call[2], callDuration, callPrice)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Method that update the caller at callers slice with calls prices and duration
func (calc *Calculator) updateCaller(callerId string, callDuration float64, callPrice float64) {
	callerFounded := false
	for _, vCaller := range calc.callers {
		if vCaller.numberId == callerId {
			vCaller.totalSeconds += callDuration
			vCaller.totalPrice += callPrice
			callerFounded = true
		}
	}
	if callerFounded == false {
		caller := &Caller{callerId, callDuration, callPrice}
		calc.callers = append(calc.callers, caller)
	}
}

// Method that find the highest caller and apply the discount
func (calc *Calculator) unchargeHighestCaller() {
	highest := 0.0
	var highestCaller *Caller
	for _, vCaller := range calc.callers {
		if vCaller.totalSeconds > highest {
			highest = vCaller.totalSeconds
			highestCaller = vCaller
		}
	}
	highestCaller.totalPrice = highestCallerDiscount
}

// Method to print the results
func (calc *Calculator) printResults() {
	for _, vCaller := range calc.callers {
		fmt.Println(fmt.Sprintf("Caller: %s - Time in Seconds: %.0f - Total Price: %.2f",
			vCaller.numberId, vCaller.totalSeconds, vCaller.totalPrice))
	}
}

// Method to calculate call durations
func calculateCallDuration(startTime string, finishTime string) float64 {
	now := time.Now()
	st := convertTimeStringToInt(startTime)
	ft := convertTimeStringToInt(finishTime)
	startT := time.Date(now.Year(), now.Month(), now.Day(), st[0], st[1], st[2], 0, time.UTC)
	finishT := time.Date(now.Year(), now.Month(), now.Day(), ft[0], ft[1], ft[2], 0, time.UTC)
	timeCall := finishT.Sub(startT).Seconds()
	return timeCall
}

// Method to calculate call prices
func calculateCallPrice(callDuration float64) float64 {
	callRemainder := 0.0
	if callDuration > firstMinutesInSeconds {
		callRemainder = callDuration - firstMinutesInSeconds
		callDuration = 300
	}
	callPrice := (callDuration * firstMinutesPrice / minuteInSeconds) +
		(callRemainder * finalMinutesPrice / minuteInSeconds)
	return callPrice
}

// Method to convert time received strings in an array of ints
func convertTimeStringToInt(convertTime string) [3]int {
	var arrayTimeInt [3]int
	arrayTimeString := strings.Split(convertTime, ":")
	for i := 0; i < 3; i++ {
		arrayTimeInt[i], _ = strconv.Atoi(arrayTimeString[i])
	}
	return arrayTimeInt
}
