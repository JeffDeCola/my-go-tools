package main

import (
	"os"
	"testing"
)

func Test_convertInlineMarkdownToHTML(t *testing.T) {
	tests := []struct {
		name string
		line string
		want string
	}{
		{
			name: "markdown link",
			line: "[ADDIE-PC](https://example.com/addie)",
			want: `<a href="https://example.com/addie">ADDIE-PC</a>`,
		},
		{
			name: "markdown image stripped to alt text",
			line: `![Addie](https://example.com/addie.png)`,
			want: `Addie`,
		},
		{
			name: "inline emphasis and code",
			line: "Use **bold**, *italic*, and `code`.",
			want: "Use <strong>bold</strong>, <em>italic</em>, and <code>code</code>.",
		},
		{
			name: "raw html preserved",
			line: `<i>Some notes</i>`,
			want: `<i>Some notes</i>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertInlineMarkdownToHTML(tt.line)
			if got != tt.want {
				t.Fatalf("convertInlineMarkdownToHTML() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_makeHTMLTABLE(t *testing.T) {
	type args struct {
		stuff      []string
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
			makeHTMLTABLE(tt.args.stuff, tt.args.outputFile)
		})
	}
}

func Test_buildTableStruct(t *testing.T) {
	type args struct {
		t     *table
		stuff []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildTableStruct(tt.args.t, tt.args.stuff)
		})
	}
}

func Test_checkExpiredDates(t *testing.T) {
	type args struct {
		t *table
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkExpiredDates(tt.args.t)
		})
	}
}

func Test_printLine(t *testing.T) {
	type args struct {
		line       string
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
			printLine(tt.args.line, tt.args.outputFile)
		})
	}
}
