// my-go-examples md5-hash-from-file.go

package main

import (
	"reflect"
	"testing"
)

func Test_setLogLevel(t *testing.T) {
	type args struct {
		logLevel string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Trace",
			args: args{
				logLevel: "trace",
			},
			wantErr: false,
		},
		{
			name: "Test Info",
			args: args{
				logLevel: "info",
			},
			wantErr: false,
		},
		{
			name: "Test Error",
			args: args{
				logLevel: "error",
			},
			wantErr: false,
		},
		{
			name: "Test Bad Log Level",
			args: args{
				logLevel: "whatever",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setLogLevel(tt.args.logLevel); (err != nil) != tt.wantErr {
				t.Errorf("setLogLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getFilename(t *testing.T) {
	type args struct {
		ssh bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFilename(tt.args.ssh)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFilename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test Read file",
			args: args{
				filename: "test/readfile_test.txt",
			},
			want:    []byte("Hi Jeff, how are you.\n\nYou can keep secrets in the file."),
			wantErr: false,
		},
		{
			name: "Test Read file error",
			args: args{
				filename: "fake.txt",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSSHFile(t *testing.T) {
	type args struct {
		ssh           bool
		plainTextByte []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSSHFile(tt.args.ssh, tt.args.plainTextByte)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSSHFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSSHFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateMD5Hash(t *testing.T) {
	type args struct {
		plainTextByte []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Calculate a md5 hash",
			args: args{
				plainTextByte: []byte("Thisisatest"),
			},
			want: "0480aa34aa3db358b37cde2ab6b65326",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMD5Hash(tt.args.plainTextByte); got != tt.want {
				t.Errorf("calculateMD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printReadableMD5(t *testing.T) {
	type args struct {
		md5Hash string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Print more readable form",
			args: args{
				md5Hash: "0480aa34aa3db358b37cde2ab6b65326",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printReadableMD5(tt.args.md5Hash)
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
