// my-go-examples sha256-hash-from-file.go

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
			want: []byte("This is a test to get the sha256 fingerprint from this file.\n" +
				"This is a test to get the sha256 fingerprint from this file.\n" +
				"This is a test to get the sha256 fingerprint from this file.\n" +
				"This is a test to get the sha256 fingerprint from this file.\n" +
				"This is a test to get the sha256 fingerprint from this file."),
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

func Test_calculatesha256Hash(t *testing.T) {
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
			name: "Test Calculate sha256 Hash",
			args: args{
				plainText: "Thisisatest\n",
				isSSH:     false,
			},
			want: "fqlmuyxQj/RVeycSaFxnw0pRoaEP1/KgC8mTGJGAkAs",
		},
		{
			name: "Test Calculate sha256 Hash on ssh file",
			args: args{
				plainText: "Thisisatest\n",
				isSSH:     true,
			},
			want: "hdktgKv7nYMKf1llotaRYdE3oryMeCpXTSJ9X3UK3Nk",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatesha256Hash(tt.args.plainText, tt.args.isSSH); got != tt.want {
				t.Errorf("calculatesha256Hash() = %v, want %v", got, tt.want)
			}
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
