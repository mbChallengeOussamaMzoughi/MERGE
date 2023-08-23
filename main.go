package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// ParseStringToNumbers converts a specific string pattern of number pairs to an array of integers.
// Complexity:
// Time: O(n + m) where n is the length of the input string and m is the number of matched pairs.
// Space: O(n) in terms of memory consumption.
func ParseStringOfIntervalsToNumbersArrayOfIntervals(input string) ([]int, error) {
	numbers := []int{} // This slice will hold the resulting numbers

	// Validate the input against a regex pattern to ensure it adheres to the expected format
	pattern := `(\[\d+,\d+\] ?)+`
	if matched, _ := regexp.MatchString("^"+pattern+"$", input); !matched {
		return nil, errors.New("input format is invalid")
	}

	// Use another regex to extract the individual pairs of numbers from the string
	regex := regexp.MustCompile(`\[(\d+),(\d+)\]`)
	matches := regex.FindAllStringSubmatch(input, -1)

	// Ensure that there aren't too many pairs
	if len(matches) > 20 {
		return nil, errors.New("no more than 20 pairs are allowed")
	}

	for _, match := range matches {
		// Convert the first number in the pair
		firstNum, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, err
		}

		// Convert the second number in the pair
		secondNum, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, err
		}

		// Ensure neither number in the pair exceeds 10,000
		if firstNum > 10000 || secondNum > 10000 {
			return nil, errors.New("numbers cannot be bigger than 10,000")
		}

		// Ensure the first number is less than the second
		if firstNum >= secondNum {
			return nil, errors.New("first number in the pair must be smaller than the second")
		}

		// Append both numbers to the result slice
		numbers = append(numbers, firstNum, secondNum)
	}

	return numbers, nil
}

// FindSmallestIntervalAndRemoveFromSlice identifies the smallest pair in the input slice (nums), and then removes this pair and
// returns along with this pair a slice without the pair.
// Complexity:
// Time: O(n) - where n is the length of the input slice. The function iterates over the slice once to find the smallest number.
// Space: O(n) - where n is the length of the input slice. The function creates a new slice (`newIntervalsSlice`) that can have up to n-2 elements (in the worst case).
func FindSmallestIntervalAndRemoveFromSlice(intervals []int) (first int, second int, newIntervalsSlice []int) {

	// Return zeros and the empty slice if intervals is empty.
	if len(intervals) == 0 {
		return 0, 0, intervals
	}

	// Find the index of the smallest number
	minIndex := 0
	for i, num := range intervals {
		if num < intervals[minIndex] {
			minIndex = i
		}
	}

	// Assign the smallest number to 'first'
	// which is the 'lower bound' of the interval
	first = intervals[minIndex]

	// If there's a number next to the smallest number, assign it to 'second'
	// second is the 'upper bound' of the interva
	if minIndex+1 < len(intervals) {
		second = intervals[minIndex+1]
	}

	// Create a new slice without the two numbers (without the smallest interval)
	newIntervalsSlice = append(intervals[:minIndex], intervals[minIndex+2:]...)

	return
}

// IntervalMerger merges every other pairs in the slice with the given pair (firstInput, secondInput)
// and deletes the merged pair from the slice.
// Complexity:
// Time: O(n), where n is the length of newNumsInput. The function goes through the newNumsInput slice once.
// Spacw: O(n), where n is the length of newNumsInput. In the worst case, all pairs are retained, making tempNums the same size as newNumsInput.
func IntervalMerger(firstInput int, secondInput int, newNumsInput []int) (firstOutput int, secondOutput int, newNumsOutput []int) {
	// Temporary slice to store pairs that aren't removed
	var tempNums []int

	// our slice for e.g. [2, 19, 25, 30, 17, 23] which is [2,19] [25,30] [17,23]
	// so 'i += 2' is used here to jump from interval to interval inside the slice
	for i := 0; i < len(newNumsInput); i += 2 {
		// If there's no next element in the slice, break
		if i+1 >= len(newNumsInput) {
			break
		}

		// If the current pair's first integer is less than secondInput
		// for e.g. [2,19] [17,23] and 17 < 19
		if newNumsInput[i] < secondInput {
			// If the pair's second integer is greater than secondInput, update secondInput
			// for e.g 23 > 19 than the second input is 23 now and the original interval is [2,23]
			if newNumsInput[i+1] > secondInput {
				secondInput = newNumsInput[i+1]
			}
		} else {
			// If the current pair doesn't meet the removal condition, add it to the tempNums slice
			tempNums = append(tempNums, newNumsInput[i], newNumsInput[i+1])
		}
	}

	return firstInput, secondInput, tempNums
}

// MERGE takes an input string in a specific format of number pairs, processes them
// to merge every other pairs according to specific rules (defined in ParseStringToNumbers,
// FindAndRemove, and FinChecker) and then returns the result as a string.
// Complexity:
// Time: O(n + m + k)
// Space: O(p + q + r)
func MERGE(inputStr string) (outputStr string) {
	outputStr = ""
	temp := ""
	// Convert the string representation of number pairs into a slice of integers
	nums, err := ParseStringOfIntervalsToNumbersArrayOfIntervals(inputStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	i := 0
	firstIteration := true
	// While there are numbers left in the nums slice
	for i < len(nums) {
		// Retrieve the smallest interval plus the "array" with the other intervals
		first, second, newNums := FindSmallestIntervalAndRemoveFromSlice(nums)

		// Apply the merging rule to the found pair and get the updated new slice after the merge operation
		firstWIP, secondWIP, newNumsWIP := IntervalMerger(first, second, newNums)

		// Convert the merged pair to its string representation and append it to the output string
		if firstIteration {
			temp = "[" + strconv.Itoa(firstWIP) + "," + strconv.Itoa(secondWIP) + "]"
		} else {
			temp = " [" + strconv.Itoa(firstWIP) + "," + strconv.Itoa(secondWIP) + "]"
		}
		outputStr = outputStr + temp

		// Update the nums slice to the new slice after the merge operation for the next iteration
		nums = newNumsWIP
		firstIteration = false
	}
	return outputStr
}

func main() {
	inputStr := "[25,30] [2,19] [14,23] [4,8]"
	outputStr := MERGE(inputStr)
	fmt.Println(outputStr)
}
