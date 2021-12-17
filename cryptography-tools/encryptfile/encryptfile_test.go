package main

import (
	"io"
	"os"
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
			want:    "test\n",
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
func Test_encryptPlainText(t *testing.T) {
	type args struct {
		r             io.Reader
		keyByte       []byte
		plainTextByte []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get cipherText from plainText",
			args: args{
				// r needs to be long enough
				r:             strings.NewReader("a b c d e f g"),
				keyByte:       []byte("098f6bcd4621d373cade4e832627b4f6"),
				plainTextByte: []byte("Hi Jeff, how are you.\n\nYou can keep secrets in the file."),
			},
			want:    "6120622063206420652066205c7f8975643637854a78d414fd3334673b714ebbb55a34a00f230c2202b12a796df3365603884365e8a3e75b5d90b21df3919b372c5b0cf42d21f1cc4bb03aa45d63e19641823977",
			wantErr: false,
		},
		{
			// GET CIPHER BLOCK USING KEY - TEST ERROR
			name: "Get cipherText from plainText - error on get cipher block",
			args: args{
				// r needs to be long enough
				r:             strings.NewReader("a b c d e f g"),
				keyByte:       []byte("toshort"),
				plainTextByte: []byte("Hi Jeff, how are you.\n\nYou can keep secrets in the file."),
			},
			want:    "",
			wantErr: true,
		},
		/* {
			// GET GCM INSTANCE THAT USES THE AES CIPHER - TEST ERROR
			name: "Get cipherText from plainText - get gcm instance",
			args: args{
				r:             strings.NewReader("a b c d e f g"),
				keyByte:       []byte("098f6bcd4621d373cade4e832627b4f6"),
				plainTextByte: []byte("Hi Jeff, how are you.\n\nYou can keep secrets in the file."),
			},
			want:    "",
			wantErr: true,
		}, */
		{
			// CREATE A NONCE AND POPULATE - TEST ERROR
			name: "Get cipherText from plainText - error on populate nonce",
			args: args{
				// r needs to be long enough - here it is too short
				r:             strings.NewReader("a b c d e f"),
				keyByte:       []byte("098f6bcd4621d373cade4e832627b4f6"),
				plainTextByte: []byte("Hi Jeff, how are you.\n\nYou can keep secrets in the file."),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encryptPlainText(tt.args.r, tt.args.keyByte, tt.args.plainTextByte)
			if (err != nil) != tt.wantErr {
				t.Errorf("encryptPlainText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("encryptPlainText() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func Test_writeCipherTextFile(t *testing.T) {
	type args struct {
		cipherText string
		filename   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test write cipherText to file",
			args: args{
				cipherText: "8976ecbdeec83d82540e01466487acd6a5e841c5bfdb172502894270c84fcde7d12d644abce90b609def3b8eb10537160b6499c4696604cdc8217c9752569eb3535bbab3eed90b9ef01f3e7e213d3a683570dd82",
				filename:   "test/encrypted_test.txt",
			},
			wantErr: false,
		},
		{
			name: "Test write cipherText to file - Can;t write file to fake directory",
			args: args{
				cipherText: "8976ecbdeec83d82540e01466487acd6a5e841c5bfdb172502894270c84fcde7d12d644abce90b609def3b8eb10537160b6499c4696604cdc8217c9752569eb3535bbab3eed90b9ef01f3e7e213d3a683570dd82",
				filename:   "fake/encrypted_test.txt",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeCipherTextFile(tt.args.cipherText, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("writeCipherText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/

func Test_writeHeader(t *testing.T) {
	type args struct {
		outputFile os.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Writing Header",
			args: args{
				outputFile: os.File{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeHeader(tt.args.outputFile); (err != nil) != tt.wantErr {
				t.Errorf("writeHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_writeCipherText(t *testing.T) {
	type args struct {
		cipherText string
		outputFile os.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeCipherText(tt.args.cipherText, tt.args.outputFile); (err != nil) != tt.wantErr {
				t.Errorf("writeCipherText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_writeFooter(t *testing.T) {
	type args struct {
		outputFile os.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeFooter(tt.args.outputFile); (err != nil) != tt.wantErr {
				t.Errorf("writeFooter() error = %v, wantErr %v", err, tt.wantErr)
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
