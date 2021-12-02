package main

import (
	"os"
	"testing"
)

func Test_createKey(t *testing.T) {
	type args struct {
		paraphrase string
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
			got, err := createKey(tt.args.paraphrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("createKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encrypt(t *testing.T) {
	type args struct {
		keyByte       []byte
		plainTextByte []byte
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
			if got := encrypt(tt.args.keyByte, tt.args.plainTextByte); got != tt.want {
				t.Errorf("encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeCipherText(t *testing.T) {
	type args struct {
		cipherText string
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
			writeCipherText(tt.args.cipherText, tt.args.outputFile)
		})
	}
}

func Test_checkErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkErr(tt.args.err)
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
