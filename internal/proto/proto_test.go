package proto

import (
	"strings"
	"testing"
)

func TestEncodeParamsOrder(t *testing.T) {
	params := map[string]string{"B": "2", "A": "1"}
	encoded1, err := EncodeParams(params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// encode again with different map order
	encoded2, err := EncodeParams(map[string]string{"A": "1", "B": "2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if encoded1 != encoded2 {
		t.Fatalf("expected deterministic encode, got %s vs %s", encoded1, encoded2)
	}
}

func TestDecodeResponse(t *testing.T) {
	// "привет" in Cyrillic
	original := "привет"
	params := map[string]string{"MSG": original}
	enc, err := EncodeParams(params)
	if err != nil {
		t.Fatalf("encode params: %v", err)
	}
	dec, err := DecodeResponse(enc)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !strings.Contains(dec, original) {
		t.Fatalf("expected to contain %q, got %q", original, dec)
	}
}
