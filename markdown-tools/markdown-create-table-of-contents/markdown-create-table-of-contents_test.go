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

func Test_createTOC(t *testing.T) {
	type args struct {
		inputFilename string
		addHeading3   bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test create TOC README.md",
			args: args{
				inputFilename: "README.md",
				addHeading3:   false,
			},
		},
		{
			name: "Test create TOC README.md -h3",
			args: args{
				inputFilename: "README.md",
				addHeading3:   true,
			},
		},
		{
			name: "Test create TOC",
			args: args{
				inputFilename: "README.md",
				addHeading3:   false,
			},
		},
		{
			name: "Test create TOC -h3",
			args: args{
				inputFilename: "README.md",
				addHeading3:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createTOC(tt.args.inputFilename, tt.args.addHeading3)
		})
	}
}

func Test_makeTOCEntry(t *testing.T) {
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
			makeTOCEntry(tt.args.heading, tt.args.headingNumber, tt.args.inputFilename)
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
