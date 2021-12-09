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

func Test_getCipherText(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test get Cipher Text",
			args: args{
				filename: "encrypted.txt",
			},
			want: "eceedcd4e9021647dfcc4628178b54aff7c83d0bd819d94dc753ccef9db0fd6aa1801cfa0d8a1c0d444657afee8990fb9120837da97d5f65fe96e74e8bd8cb00355950b771409a2d786bc6abc2d6d5aecb887000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCipherText(tt.args.filename); got != tt.want {
				t.Errorf("getCipherText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getParaphrase(t *testing.T) {
	type args struct {
		paraphraseFile string
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
			if got := getParaphrase(tt.args.paraphraseFile); got != tt.want {
				t.Errorf("getParaphrase() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
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

func Test_decryptCipherText(t *testing.T) {
	type args struct {
		keyByte    []byte
		cipherText string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test decrypt cipherText",
			args: args{
				keyByte:    []byte("098f6bcd4621d373cade4e832627b4f6"),
				cipherText: "814ba70e7e9de663778f9c8743c728ab0cfc22a703344b9d72fd786b9930f0af21c00818e7fd48b70dadbb1accf8ce6d9a5c4e39afaa6f483d7131798f4b920217a5e7fa28ac73f0e61b7070fb68b05be8a3bfe2",
			},
			want: []byte("Hi Jeff, how are you.\n\nYou can keep secrets in the file."),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decryptCipherText(tt.args.keyByte, tt.args.cipherText); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decryptCipherText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writePlainTextByte(t *testing.T) {
	type args struct {
		plainTextByte []byte
		filename      string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				plainTextByte: []byte("Hi Jeff, how are you.\n\nYou can keep secrets in the file."),
				filename:      "output_test.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writePlainTextByte(tt.args.plainTextByte, tt.args.filename)
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
