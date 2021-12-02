package main

import (
	"os"
	"reflect"
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

func Test_decrypt(t *testing.T) {
	type args struct {
		keyByte    []byte
		cipherText string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decrypt(tt.args.keyByte, tt.args.cipherText); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCipherText(t *testing.T) {
	type args struct {
		inputFile *os.File
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
			if got := getCipherText(tt.args.inputFile); got != tt.want {
				t.Errorf("getCipherText() = %v, want %v", got, tt.want)
			}
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
