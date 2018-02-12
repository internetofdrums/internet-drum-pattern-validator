package main

import (
	"fmt"
	"os"
	"encoding/base64"
	"errors"
	"bytes"
)

type DrumPattern struct {
	instruments []Instrument
}

type Instrument struct {
	notes []Note
}

type Note struct {
	length   byte
	velocity byte
}

const instructions = `
Usage: internet-drum-pattern-validator <pattern>

where <pattern> is a standard Base64 (see RFC 4648) encoded byte array
following the Internet Drum Pattern Specification, see:

https://github.com/internetofdrums/internet-drum-pattern-spec#readme

If the pattern is valid, the pattern data is formatted and written to stdout.
`;

const maximumNoteDataValue = 127
const numberOfInstruments = 16
const numberOfNotesPerBeat = 4
const numberOfBeatsPerBar = 4
const numberOfNotesPerInstrument = numberOfNotesPerBeat * numberOfBeatsPerBar
const numberOfDataPartsPerNote = 2
const numberOfBytesPerInstrument = numberOfNotesPerInstrument * numberOfDataPartsPerNote
const numberOfBytesPerDrumPattern = numberOfInstruments * numberOfBytesPerInstrument

func Decode(pattern string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(pattern)
}

func IsValidPattern(pattern []byte) (bool, error) {
	if len(pattern) != numberOfBytesPerDrumPattern {
		return false, errors.New("Drum pattern does not contain exactly 512 data parts.")
	}

	for _, dataPart := range pattern {
		if dataPart < 0 || dataPart > maximumNoteDataValue {
			return false, errors.New(fmt.Sprintf(
				"Encountered data value of 0x%x, which exceeds allowed value of 0x%x.",
				dataPart,
				maximumNoteDataValue,
			))
		}
	}

	return true, nil
}

func Convert(pattern []byte) DrumPattern {
	instruments := make([]Instrument, numberOfInstruments)

	for i := 0; i < numberOfInstruments; i++ {
		patternDataChunkStartIndex := i * numberOfBytesPerInstrument;
		patternDataChunkEndIndex := patternDataChunkStartIndex + numberOfBytesPerInstrument;
		noteData := pattern[patternDataChunkStartIndex:patternDataChunkEndIndex]
		notes := make([]Note, numberOfNotesPerInstrument)

		for j := 0; j < numberOfBytesPerInstrument; j += 2 {
			note := Note{
				length:   noteData[j],
				velocity: noteData[j+1],
			}

			notes[j/2] = note
		}

		instruments[i] = Instrument{notes}
	}

	return DrumPattern{instruments}
}

func GetFormattedPattern(pattern DrumPattern) string {
	buffer := bytes.NewBufferString("")

	for _, instrument := range pattern.instruments {
		for index, note := range instrument.notes {
			buffer.WriteString("[")
			buffer.WriteString(Format(note.length))
			buffer.WriteString(",")
			buffer.WriteString(Format(note.velocity))
			buffer.WriteString("]")

			if (index+1)%numberOfNotesPerBeat == 0 {
				buffer.WriteString(" ")
			}
		}

		buffer.WriteString(fmt.Sprintln())
	}

	return buffer.String()
}

func Format(value byte) string {
	if value == 0x00 {
		return "00"
	} else {
		return fmt.Sprintf("%X", value)
	}
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Fprint(os.Stderr, instructions)
		os.Exit(1)
	}

	encodedPattern := os.Args[1]
	pattern, decodingError := Decode(encodedPattern)

	if decodingError != nil {
		fmt.Fprintln(os.Stderr, "An error occured while decoding the pattern.")
	}

	result, validationError := IsValidPattern(pattern)

	if validationError != nil {
		fmt.Fprintln(os.Stderr, validationError)
		os.Exit(1)
	}

	if result != true {
		fmt.Fprintln(os.Stderr, "The pattern is not valid.")
		os.Exit(1)
	}

	drumPattern := Convert(pattern)
	formattedPattern := GetFormattedPattern(drumPattern)

	fmt.Print(formattedPattern)
}
