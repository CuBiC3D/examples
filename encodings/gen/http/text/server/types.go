// Code generated by goa v3.0.2, DO NOT EDIT.
//
// text HTTP server types
//
// Command:
// $ goa gen goa.design/examples/encodings/design -o
// $(GOPATH)/src/goa.design/examples/encodings

package server

import (
	text "goa.design/examples/encodings/gen/text"
)

// NewConcatstringsPayload builds a text service concatstrings endpoint payload.
func NewConcatstringsPayload(a string, b string) *text.ConcatstringsPayload {
	return &text.ConcatstringsPayload{
		A: a,
		B: b,
	}
}

// NewConcatbytesPayload builds a text service concatbytes endpoint payload.
func NewConcatbytesPayload(a string, b string) *text.ConcatbytesPayload {
	return &text.ConcatbytesPayload{
		A: a,
		B: b,
	}
}

// NewConcatstringfieldPayload builds a text service concatstringfield endpoint
// payload.
func NewConcatstringfieldPayload(a string, b string) *text.ConcatstringfieldPayload {
	return &text.ConcatstringfieldPayload{
		A: a,
		B: b,
	}
}

// NewConcatbytesfieldPayload builds a text service concatbytesfield endpoint
// payload.
func NewConcatbytesfieldPayload(a string, b string) *text.ConcatbytesfieldPayload {
	return &text.ConcatbytesfieldPayload{
		A: a,
		B: b,
	}
}
