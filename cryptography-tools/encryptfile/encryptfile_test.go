package main

import (
	"bytes"
	"io"
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeCipherTextFile(tt.args.cipherText, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("writeCipherTextFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// WE ARE WRITING TO A BUFFER - AND THEN CHECKING THE BEGINING OF IT
func Test_writeHeader(t *testing.T) {
	tests := []struct {
		name       string
		wantHandle string
		wantErr    bool
	}{
		{
			name:       "Test Write Footer",
			wantHandle: "\nThis secret file was created by Jeff DeCola\n",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handle := &bytes.Buffer{}
			if err := writeHeader(handle); (err != nil) != tt.wantErr {
				t.Errorf("writeHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// I EDITED THIS TO JUST CHECK THE PREFIX OF THE BUFFER
			if gotHandle := handle.Bytes(); !bytes.HasPrefix(gotHandle, []byte(tt.wantHandle)) {
				t.Errorf("writeHeader() = %v, want %v", gotHandle, tt.wantHandle)
			}
		})
	}
}

func Test_writeCipherText(t *testing.T) {
	type args struct {
		cipherText string
	}
	tests := []struct {
		name       string
		args       args
		wantHandle string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handle := &bytes.Buffer{}
			if err := writeCipherText(tt.args.cipherText, handle); (err != nil) != tt.wantErr {
				t.Errorf("writeCipherText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHandle := handle.String(); gotHandle != tt.wantHandle {
				t.Errorf("writeCipherText() = %v, want %v", gotHandle, tt.wantHandle)
			}
		})
	}
}

/*
func Test_writeHeader(t *testing.T) {
	type args struct {
		outputFH io.ReadWriter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Writing Header",
			args: args{
				// This is a thing that can written too - writing to a slice
				outputFH: bytes.NewBuffer(nil),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeHeader(tt.args.outputFH); (err != nil) != tt.wantErr {
				t.Errorf("writeHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
			b, err := io.ReadAll(tt.args.outputFH)
			if err != nil {
				t.Errorf("writeHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !bytes.HasPrefix(b, []byte("\nThis secret file was created by Jeff DeCola\n")) {
				t.Errorf("writeHeader() wrote %q but wanted %q, ", string(b), "This secret file....")
			}
		})
	}
} */

// WE ARE WRITING TO A BUFFER - AND THEN CHECKING IT
func Test_writeFooter(t *testing.T) {
	tests := []struct {
		name string
		// AUTO ADDED
		wantHandle string
		wantErr    bool
	}{
		{
			name:       "Test Write Footer",
			wantHandle: myFileDelimiter + "\n\n",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// AUTO ADDED - THIS IS COOL HOW IT PUTS A BUFFER HERE SO WE CAN WRITE TO IT
			handle := &bytes.Buffer{}
			if err := writeFooter(handle); (err != nil) != tt.wantErr {
				t.Errorf("writeFooter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHandle := handle.String(); gotHandle != tt.wantHandle {
				t.Errorf("writeFooter() = %v, want %v", gotHandle, tt.wantHandle)
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
