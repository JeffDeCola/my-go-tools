package main

import "testing"

func Test_checkErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test checkErr nil",
			args: args{
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkErr(tt.args.err)
		})
	}
}

func Test_checkVersion(t *testing.T) {
	type args struct {
		version bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test check version false",
			args: args{
				version: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkVersion(tt.args.version)
		})
	}
}

func Test_setLogLevel(t *testing.T) {
	type args struct {
		debugTrace bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Debug False",
			args: args{
				debugTrace: false,
			},
		},
		{
			name: "Test Debug True",
			args: args{
				debugTrace: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setLogLevel(tt.args.debugTrace)
		})
	}
}

func Test_doer(t *testing.T) {
	type args struct {
		delimiter      string
		inputFilename  string
		outputFilename string
		htmlTableBool  bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test doer",
			args: args{
				delimiter:      "$$",
				inputFilename:  "input.md",
				outputFilename: "output_test.md",
				htmlTableBool:  false,
			},
		},
		{
			name: "Test doer -html switch",
			args: args{
				delimiter:      "$$",
				inputFilename:  "input.md",
				outputFilename: "output_test.md",
				htmlTableBool:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doer(tt.args.delimiter, tt.args.inputFilename, tt.args.outputFilename, tt.args.htmlTableBool)
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
