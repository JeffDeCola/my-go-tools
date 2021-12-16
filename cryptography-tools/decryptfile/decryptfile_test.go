package main

import (
	"io"
	"reflect"
	"strings"
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

func Test_getCipherText(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Get Cipher Text from file",
			args: args{
				filename: "test/encrypted_test.txt",
			},
			want:    "eceedcd4e9021647dfcc4628178b54aff7c83d0bd819d94dc753ccef9db0fd6aa1801cfa0d8a1c0d444657afee8990fb9120837da97d5f65fe96e74e8bd8cb00355950b771409a2d786bc6abc2d6d5aecb887000",
			wantErr: false,
		},
		{
			name: "Test Get Cipher Text from file - file does not exit",
			args: args{
				filename: "fake.txt",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCipherText(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCipherText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCipherText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getParaphrase(t *testing.T) {
	type args struct {
		r              io.Reader
		paraphraseFile string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test for read file",
			args: args{
				r:              strings.NewReader(""),
				paraphraseFile: "paraphrase.txt",
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "Test for read file error",
			args: args{
				r:              strings.NewReader(""),
				paraphraseFile: "fake.txt",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test for user input",
			args: args{
				r:              strings.NewReader("jeff"),
				paraphraseFile: "",
			},
			want:    "jeff",
			wantErr: false,
		},
		{
			name: "Test for user input error",
			args: args{
				r:              strings.NewReader("\n"),
				paraphraseFile: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getParaphrase(tt.args.r, tt.args.paraphraseFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("getParaphrase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
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

func Test_getUserInput(t *testing.T) {
	type args struct {
		r       io.Reader
		askUser string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Get User Input Error",
			args: args{
				r:       strings.NewReader("\n"),
				askUser: "What is your name?",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Get User Input",
			args: args{
				r:       strings.NewReader("jeff"),
				askUser: "What is your name?",
			},
			want:    "jeff",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUserInput(tt.args.r, tt.args.askUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUserInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getUserInput() = %v, want %v", got, tt.want)
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

func Test_decryptCipherText(t *testing.T) {
	type args struct {
		keyByte    []byte
		cipherText string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test get plainText from CipherText",
			args: args{
				keyByte:    []byte("098f6bcd4621d373cade4e832627b4f6"),
				cipherText: "6120622063206420652066205c7f8975643637854a78d414fd3334673b714ebbb55a34a00f230c2202b12a796df3365603884365e8a3e75b5d90b21df3919b372c5b0cf42d21f1cc4bb03aa45d63e19641823977",
			},
			want:    []byte("Hi Jeff, how are you.\n\nYou can keep secrets in the file."),
			wantErr: false,
		},
		{
			name: "Test get plainText from CipherText - Bad keybyte",
			args: args{
				keyByte:    []byte("09"),
				cipherText: "6120622063206420652066205c7f8975643637854a78d414fd3334673b714ebbb55a34a00f230c2202b12a796df3365603884365e8a3e75b5d90b21df3919b372c5b0cf42d21f1cc4bb03aa45d63e19641823977",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decryptCipherText(tt.args.keyByte, tt.args.cipherText)
			if (err != nil) != tt.wantErr {
				t.Errorf("decryptCipherText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decryptCipherText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writePlainText(t *testing.T) {
	type args struct {
		plainTextByte []byte
		filename      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test write plainText to file",
			args: args{
				plainTextByte: []byte("Test writing file data"),
				filename:      "test/output_test.txt",
			},
			wantErr: false,
		},
		{
			name: "Test write plainText to file - can't create file in fake directory",
			args: args{
				plainTextByte: []byte("Test writing file data"),
				filename:      "fake/output_test.txt",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writePlainText(tt.args.plainTextByte, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("writePlainText() error = %v, wantErr %v", err, tt.wantErr)
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
