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
			name: "Test Read",
			args: args{
				filename: "input.txt",
			},
			want: []byte("Hi Jeff, how are you.\n\nYou can keep secrets in the file."),
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

func Test_getParaphrase(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getParaphrase(); got != tt.want {
				t.Errorf("getParaphraseHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getKeyByte(t *testing.T) {
	type args struct {
		paraphrase string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Get the keybyte from paraphrase 'test'",
			args: args{
				paraphrase: "test",
			},
			want: []byte("098f6bcd4621d373cade4e832627b4f6"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKeyByte(tt.args.paraphrase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKeyByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encryptFileData(t *testing.T) {
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
			if got := encryptFileData(tt.args.keyByte, tt.args.plainTextByte); got != tt.want {
				t.Errorf("encryptFileData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeCipherText(t *testing.T) {
	type args struct {
		cipherText string
		filename   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test write sipherText to file",
			args: args{
				cipherText: "8976ecbdeec83d82540e01466487acd6a5e841c5bfdb172502894270c84fcde7d12d644abce90b609def3b8eb10537160b6499c4696604cdc8217c9752569eb3535bbab3eed90b9ef01f3e7e213d3a683570dd82",
				filename:   "encrypted_test.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeCipherText(tt.args.cipherText, tt.args.filename)
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
