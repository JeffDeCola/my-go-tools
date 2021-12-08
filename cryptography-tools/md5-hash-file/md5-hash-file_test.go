// my-go-examples md5-hash-from-file.go

package main

import (
	"reflect"
	"testing"
)

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

func Test_getFilename(t *testing.T) {
	type args struct {
		ssh bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFilename(tt.args.ssh); got != tt.want {
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
		name string
		args args
		want []byte
	}{
		{
			name: "Test ReadFile",
			args: args{
				filename: "testfile.txt",
			},
			want: []byte("This is a test to get the md5 fingerprint from this file.\n" +
				"This is a test to get the md5 fingerprint from this file.\n" +
				"This is a test to get the md5 fingerprint from this file.\n" +
				"This is a test to get the md5 fingerprint from this file.\n" +
				"This is a test to get the md5 fingerprint from this file."),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
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
		name string
		args args
		want string
	}{
		{
			name: "Test parseSSHFile false",
			args: args{
				ssh:           false,
				plainTextByte: []byte("This is a test"),
			},
			want: "This is a test",
		},
		{
			name: "Test parseSSHFile false",
			args: args{
				ssh:           true,
				plainTextByte: []byte("ssh-rsa Thisisatest blahblah"),
			},
			want: "Thisisatest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSSHFile(tt.args.ssh, tt.args.plainTextByte); got != tt.want {
				t.Errorf("parseSSHFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateMD5Hash(t *testing.T) {
	type args struct {
		plainText string
		isSSH     bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Calculate MD5 Hash",
			args: args{
				plainText: "Thisisatest\n",
				isSSH:     false,
			},
			want: "85ceadf8a5789290094ac22b03170c6b",
		},
		{
			name: "Test Calculate MD5 Hash on ssh file",
			args: args{
				plainText: "Thisisatest\n",
				isSSH:     true,
			},
			want: "259da16013dc6713a6b0b7fa25bf22fe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMD5Hash(tt.args.plainText, tt.args.isSSH); got != tt.want {
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
			name: "Test Readable MD5",
			args: args{
				md5Hash: "2fe4a3",
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
