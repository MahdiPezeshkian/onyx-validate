package legalid_test

import (
	"errors"
	"testing"

	"github.com/MahdiPezeshkian/onyx-validate/iran/identity/legalid"
	"github.com/MahdiPezeshkian/onyx-validate/validate"
)

func TestLegalID(t *testing.T) {
	if err := legalid.Validate("10380284790"); err != nil {
		t.Fatalf("known valid: %v", err)
	}
	if err := legalid.Validate("۱۰۳۸۰۲۸۴۷۹۰"); err != nil {
		t.Fatalf("persian digits: %v", err)
	}
	cases := []struct {
		in   string
		code string
	}{
		{"", validate.CodeRequired},
		{"123", validate.CodeInvalidLength},
		{"00000000000", validate.CodeInvalidFormat},
		{"10300000000", validate.CodeInvalidFormat}, // middle zeros + likely checksum fail, format first
		{"10380284791", validate.CodeInvalidChecksum},
	}
	for _, tc := range cases {
		err := legalid.Validate(tc.in)
		var ve *validate.ValidationError
		if !errors.As(err, &ve) || ve.Code != tc.code {
			t.Fatalf("%q: got %#v want %s", tc.in, err, tc.code)
		}
	}
}
