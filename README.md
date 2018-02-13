# Internet Drum Pattern Validator

Validates base-64 encoded Internet Drum Pattern data and, if valid, return a formatted pattern. 

## Build status

[![Build Status](https://travis-ci.org/internetofdrums/internet-drum-pattern-validator.svg?branch=master)](https://travis-ci.org/internetofdrums/internet-drum-pattern-validator)

## Requirements

To build the binary for your system, you need a [Go distribution][1].

## Installing

To install the validator, do:

```bash
go install github.com/internetofdrums/internet-drum-pattern-validator
```

## Running

To validate a drum pattern, run the `internet-drum-pattern-validator` command with the drum pattern as its only 
argument, e.g.:

```bash
internet-drum-pattern-validator f38AAAAAAAB/fwAAAAAAAH9/AAAAAAAAf38AAAAAAAAAAAAAQEAAAAAAAABAQAAAAAAAAEBAAAAAAAAAQEAAAAAAAAAAAAAAf38AAAAAAAAAAAAAAAAAAH9/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=
```

The program should return successfully and print the decoded pattern, for example:

```text
The drum pattern is valid!

After decoding, the pattern looks like this, where (XX,YY) is one note with a length of XX and a velocity of YY:

(7F,7F)(00,00)(00,00)(00,00) (7F,7F)(00,00)(00,00)(00,00) (7F,7F)(00,00)(00,00)(00,00) (7F,7F)(00,00)(00,00)(00,00) 
(00,00)(00,00)(40,40)(00,00) (00,00)(00,00)(40,40)(00,00) (00,00)(00,00)(40,40)(00,00) (00,00)(00,00)(40,40)(00,00) 
(00,00)(00,00)(00,00)(00,00) (7F,7F)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (7F,7F)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) 
(00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00) (00,00)(00,00)(00,00)(00,00)
```

### Errors

If the drum pattern is missing or incorrect, the program exits with a non-zero exit code and some helpful instructions,
for example:

```text
Could not decode drum pattern: illegal base64 data at input byte 0.
```

or:

```text
Usage: internet-drum-pattern-validator <pattern>

where <pattern> is a standard Base64 (see RFC 4648) encoded byte array
following the Internet Drum Pattern Specification, see:

https://github.com/internetofdrums/internet-drum-pattern-spec#readme

If the pattern is valid, the pattern data is formatted and written to stdout.
```

[1]: https://golang.org/doc/install
