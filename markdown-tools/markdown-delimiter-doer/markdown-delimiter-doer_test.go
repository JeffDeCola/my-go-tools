package main

import (
	"os"
	"testing"
)

func Test_printLine(t *testing.T) {
	type args struct {
		line       string
		outputFile *os.File
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printLine(tt.args.line, tt.args.outputFile)
		})
	}
}

func Test_checkExpiredDates(t *testing.T) {
	type args struct {
		t *table
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkExpiredDates(tt.args.t)
		})
	}
}

func Test_buildTableStruct(t *testing.T) {
	type args struct {
		t     *table
		stuff []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildTableStruct(tt.args.t, tt.args.stuff)
		})
	}
}

func Test_makeHTMLTABLE(t *testing.T) {
	type args struct {
		stuff      []string
		outputFile *os.File
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeHTMLTABLE(tt.args.stuff, tt.args.outputFile)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
