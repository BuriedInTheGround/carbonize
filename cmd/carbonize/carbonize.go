package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net/url"
	"os"
	"runtime/debug"

	"github.com/pkg/browser"
	"interrato.dev/carbonize/internal/carbonize"
)

const usage = `Usage:
    carbonize [-c PATH] [INPUT]

Options:
    -c, --configuration PATH  Use PATH as a configuration file.
    -n, --trailing-newline    Keep the trailing newline if it exists.

INPUT defaults to standard input.

A configuration file must be a JSON file. Ideally, it should be a configuration
exported from the Carbon website.

Example:
    $ carbonize ./path/to/file.go`

const maxCodeLength = 45580

var Version string

//go:embed default-config.json
var defaultConfig embed.FS

func main() {
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	var (
		configurationFileFlag            string
		versionFlag, trailingNewlineFlag bool
	)

	flag.BoolVar(&versionFlag, "version", false, "print the version")
	flag.StringVar(&configurationFileFlag, "c", "", "configuration file")
	flag.StringVar(&configurationFileFlag, "configuration", "", "configuration file")
	flag.BoolVar(&trailingNewlineFlag, "n", false, "keep trailing newline")
	flag.BoolVar(&trailingNewlineFlag, "trailing-newline", false, "keep trailing newline")
	flag.Parse()

	if versionFlag {
		if Version != "" {
			fmt.Println(Version)
			return
		}
		if buildInfo, ok := debug.ReadBuildInfo(); ok {
			fmt.Println(buildInfo.Main.Version)
			return
		}
		fmt.Println("(unknown)")
		return
	}

	var in io.Reader = os.Stdin
	if name := flag.Arg(0); name != "" && name != "-" {
		f, err := os.Open(name)
		if err != nil {
			errorf("failed to open input file %q: %v", name, err)
		}
		defer f.Close()
		in = f
	}

	carbon(in, configurationFileFlag, trailingNewlineFlag)
}

func carbon(in io.Reader, configPath string, keepTrailingNewline bool) {
	var f fs.File

	if configPath == "-" {
		errorf("reading configuration file from stdin not supported")
	}

	if configPath != "" {
		var err error
		f, err = os.Open(configPath)
		if err != nil {
			errorf("failed to open configuration file %q: %v", configPath, err)
		}
		defer f.Close()
	} else {
		var err error
		f, err = defaultConfig.Open("default-config.json")
		if err != nil {
			errorf("failed to open default configuration file: %v", err)
		}
		defer f.Close()
	}

	config, err := carbonize.ParseConfig(f)
	if err != nil {
		errorf("failed to parse configuration: %v", err)
	}

	queryString, err := config.QueryString()
	if err != nil {
		errorf("failed to convert configuration to query string: %v", err)
	}

	inputBytes, err := io.ReadAll(in)
	if err != nil {
		errorf("failed to read input: %v", err)
	}

	// Remove the trailing newline if it exists.
	if !keepTrailingNewline {
		inputBytes = bytes.TrimRight(inputBytes, "\r\n")
	}
	code := url.QueryEscape(string(inputBytes))

	if len(code) > maxCodeLength {
		errorWithHint("input length too long", "note that Carbon breaks when the input is too long",
			"current maximum length for input after encoding is "+fmt.Sprintf("%d", maxCodeLength)+" characters")
	}

	err = browser.OpenURL("https://carbon.now.sh/?" + queryString + "&code=" + code)
	if err != nil {
		errorf("cannot open browser: %v", err)
	}
}
