package main

import "testing"

func Test_makeTOC(t *testing.T) {
	type args struct {
		heading       string
		headingNumber string
		inputFilename string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeTOC(tt.args.heading, tt.args.headingNumber, tt.args.inputFilename)
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
