# carbonize

[![Go Reference](https://pkg.go.dev/badge/interrato.dev/carbonize.svg)](https://pkg.go.dev/interrato.dev/carbonize)

A simple tool for opening any textual content with [Carbon].

It allows for custom JSON configurations, and it features UNIX-style composability.

```
$ carbonize example.go                    # A new tab will open in your default browser.
$ cat example.go | carbonize              # This is equivalent to the previous command.
$ carbonize -c my-config.json example.go  # Use a custom configuration.
```

## Usage

```
Usage:
    carbonize [-c PATH] [INPUT]

Options:
    -c, --configuration PATH  Use PATH as a configuration file.
    -n, --trailing-newline    Keep the trailing newline if it exists.

INPUT defaults to standard input.

A configuration file must be a JSON file. Ideally, it should be a configuration
exported from the Carbon website.
```

Keep in mind that due to browser limitations and other factors, a maximum input length do exists.
The current value used by carbonize is 45580 characters after URL-encoding.

### Example

Run carbonize on an example Go source code file.

```
$ carbonize example.go
```

The browser opens, then I modify the comment text color to yellow, and the following image is the result. ✨

<p align="center">
<img
    alt="The source code of a Go program which runs an HTTP server on port 8080 that randomly uses or not TLS."
    src="https://user-images.githubusercontent.com/26801023/189491857-1b7c864c-28cb-464b-b130-0d903ac72d26.png"
/>
</p>

#### Carbonize GitHub source files

You can easily run carbonize on GitHub raw source files by combining it with `curl`.

```
$ curl -s https://raw.githubusercontent.com/BuriedInTheGround/fine/main/fine.go | carbonize
```

## Installation

<!-- On Windows, Linux, and macOS, you can use [the pre-built binaries]. -->

If your system has [Go 1.18+], you can build from source:

```
git clone https://interrato.dev/carbonize && cd carbonize
go build -o . interrato.dev/carbonize/cmd/...
```

<!-- <table>
    <tr>
        <td>NixOS / Nix</td>
        <td>
            <code>TODO</code>
        </td>
    </tr>
    <tr>
        <td>TODO</td>
        <td>
            <code>TODO</code>
        </td>
    </tr>
</table> -->


<!-- References -->

[Carbon]: https://carbon.now.sh "Carbon official website"
[the pre-built binaries]: https://github.com/BuriedInTheGround/carbonize/releases "GitHub releases page for carbonize"
[Go 1.18+]: https://go.dev/dl "The Go programming language downloads page"
