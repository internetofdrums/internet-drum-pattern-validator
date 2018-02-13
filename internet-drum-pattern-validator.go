package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os"
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
`

const maximumNoteDataValue = 127
const numberOfInstruments = 16
const numberOfNotesPerBeat = 4
const numberOfBeatsPerBar = 4
const numberOfNotesPerInstrument = numberOfNotesPerBeat * numberOfBeatsPerBar
const numberOfDataPartsPerNote = 2
const numberOfDataPartsPerInstrument = numberOfNotesPerInstrument * numberOfDataPartsPerNote
const numberOfDataPartsPerDrumPattern = numberOfInstruments * numberOfDataPartsPerInstrument

var errorColor = color.New(color.FgRed).SprintFunc()
var successColor = color.New(color.FgGreen).SprintFunc()

func Decode(pattern string) ([]byte, error) {
	result, err := base64.StdEncoding.DecodeString(pattern)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func ValidatePattern(pattern []byte) error {
	numberOfDataParts := len(pattern)

	if numberOfDataParts != numberOfDataPartsPerDrumPattern {
		return errors.New(fmt.Sprintf(
			"The drum pattern does contains %d data parts (bytes), but should contain exactly %d bytes.",
			numberOfDataParts,
			numberOfDataPartsPerDrumPattern,
		))
	}

	for _, dataPart := range pattern {
		if dataPart < 0 || dataPart > maximumNoteDataValue {
			return errors.New(fmt.Sprintf(
				"Encountered data value of 0x%x, which exceeds allowed value of 0x%x.",
				dataPart,
				maximumNoteDataValue,
			))
		}
	}

	return nil
}

func Convert(pattern []byte) DrumPattern {
	instruments := make([]Instrument, numberOfInstruments)

	for i := 0; i < numberOfInstruments; i++ {
		patternDataChunkStartIndex := i * numberOfDataPartsPerInstrument
		patternDataChunkEndIndex := patternDataChunkStartIndex + numberOfDataPartsPerInstrument
		noteData := pattern[patternDataChunkStartIndex:patternDataChunkEndIndex]
		notes := make([]Note, numberOfNotesPerInstrument)

		for j := 0; j < numberOfDataPartsPerInstrument; j += 2 {
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
	var buffer bytes.Buffer

	for _, instrument := range pattern.instruments {
		AppendFormattedInstrument(instrument, &buffer)
	}

	return buffer.String()
}

func AppendFormattedInstrument(instrument Instrument, buffer *bytes.Buffer) {
	for index, note := range instrument.notes {
		AppendFormattedNote(note, buffer)

		if (index+1)%numberOfNotesPerBeat == 0 {
			buffer.WriteString(" ")
		}
	}

	buffer.WriteString(fmt.Sprintln())
}

func AppendFormattedNote(note Note, buffer *bytes.Buffer) {
	buffer.WriteString("(")
	AppendFormattedNoteDataPart(note.length, buffer)
	buffer.WriteString(",")
	AppendFormattedNoteDataPart(note.velocity, buffer)
	buffer.WriteString(")")
}

func AppendFormattedNoteDataPart(value byte, buffer *bytes.Buffer) {
	if value == 0x00 {
		buffer.WriteString("00")
	} else {
		buffer.WriteString(fmt.Sprintf("%X", value))
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, instructions)
		os.Exit(1)
	}

	pattern, decodingError := Decode(os.Args[1])

	if decodingError != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf(errorColor("Could not decode drum pattern: %s."), decodingError))
		os.Exit(1)
	}

	validationError := ValidatePattern(pattern)

	if validationError != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf(errorColor("The drum pattern is invalid: %s"), validationError))
		os.Exit(1)
	}

	fmt.Println(successColor("The drum pattern is valid!"))
	fmt.Println()

	drumPattern := Convert(pattern)
	formattedPattern := GetFormattedPattern(drumPattern)

	fmt.Println("After decoding, the pattern looks like this, " +
		"where (XX,YY) is one note with a length of XX and a velocity of YY:")
	fmt.Println()
	fmt.Print(formattedPattern)
}
