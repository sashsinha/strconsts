<p align="center">
  <a href="https://github.com/sashsinha/strconsts"><img alt="strconsts logo" src="https://raw.githubusercontent.com/sashsinha/strconsts/main/logo.png"></a>
</p>

<h1 align="center">String Constants</h1>

<h3 align="center">Go module that provides the same string constants available in the <a href="https://docs.python.org/3/library/string.html#string-constants">Python string</a> module</h3>

<p align="center">
  <a href="https://raw.githubusercontent.com/sashsinha/strconsts/main/LICENCE"><img alt="License: MIT" src="https://raw.githubusercontent.com/sashsinha/strconsts/main/licence.svg"></a>
  <a href="https://godoc.org/github.com/sashsinha/strconsts"><img alt="Godoc" src="https://godoc.org/github.com/sashsinha/strconsts?status.svg"></a>
</p>

## Avaliable Constants

### `strconsts.ASCIILetters`
The concatenation of the `ASCIILowercase` and `ASCIIUppercase` constants described below. This value is not locale-dependent.

### `strconsts.ASCIILowercase`
The lowercase letters `'abcdefghijklmnopqrstuvwxyz'`. This value is not locale-dependent and will not change.

### `strconsts.ASCIIUppercase`
The uppercase letters `'ABCDEFGHIJKLMNOPQRSTUVWXYZ'`. This value is not locale-dependent and will not change.

### `strconsts.Digits`
The string `'0123456789'`.

### `strconsts.HexDigits`
The string `'0123456789abcdefABCDEF'`.

### `strconsts.OctDigits`
The string `'01234567'`.

### `strconsts.Punctuation`
String of ASCII characters which are considered punctuation characters in the C locale: 
``!"#$%&'()*+,-./:;<=>?@[\]^`{|}~``.

### `strconsts.Printable`
String of ASCII characters which are considered printable. This is a combination of [`Digits`](#strconstsdigits), [`ASCIILetters`](#strconstsasciiletters), [`Punctuation`](#strconstspunctuation), and [`Whitespace`](#strconstswhitespace).

### `strconsts.Whitespace`
A string containing all ASCII characters that are considered whitespace. This includes the characters space, tab, linefeed, return, formfeed, and vertical tab.

## Installation

Standard go get:

```shell
$ go get github.com/sashsinha/strconsts
```

## Example Usage

```go
package main

import (
    "fmt"
    "github.com/sashsinha/strconsts"
)

func main() {
    fmt.Println("Digits:", strconsts.Digits) // Digits: 0123456789
}
```
