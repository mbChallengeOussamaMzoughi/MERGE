# Go Basic Project

This project is designed to parse and merge pairs of integers provided as string intervals. The primary functionality is parsing the given string into intervals, finding and removing the smallest interval, and finally merging the intervals based on a set of rules. Along with the core logic, there are utility functions and a Makefile to assist in building, running, and testing the project.

## Requirements

- Go (1.x) - [Installation guide](https://golang.org/doc/install)
- Make (Optional) - Most Unix-based systems (like MacOS and Linux) have this pre-installed.

## Directory Structure

.
├── main.go # Contains the core logic for parsing and merging intervals

├── main_test.go # Contains the test cases of the main.go functionalities

└── Makefile # Make commands for building and running the project


## Getting Started

### Clone the Repository

git clone https://github.com/mbChallengeOussamaMzoughi/MERGE.git

cd ./MERGE



### Building the Project

To build the project, run:

make build


This will produce a binary named `mergeapp`.

### Running the Project

To compile and run the project, use:

make run


### Testing

Although no tests are currently available, when added, you can run them with:

make test

### Cleaning Up

To remove the generated binary and clean up, run:

make clean


## Core Logic

The main program file (main.go) contains functions that together provide the ability to merge pairs of integers represented as intervals in a string.

- **ParseStringOfIntervalsToNumbersArrayOfIntervals**: Converts a string pattern of number pairs into an array of integers.

- **FindSmallestIntervalAndRemoveFromSlice**: Identifies the smallest interval from a slice and returns the slice without this interval.

- **IntervalMerger**: Merges pairs within a slice based on a given pair's parameters.

- **MERGE**: The core function that takes a string of intervals, parses it, identifies the smallest interval, and then merges them according to the set rules.

For a more in-depth understanding of each function, refer to the comments in main.go, which provide an overview of their responsibilities, complexities, and inner workings.