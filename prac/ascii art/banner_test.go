package main

import (
	"os"
	"testing"
)

func createTestBannerFile(t *testing.T, filename string) {
	t.Helper()

	content := ""

	for i := 0; i < 95; i++ {
		ch := rune(32 + i)

		for j := 0; j < 8; j++ {
			content += string(ch) + "\n"
		}

		content += "\n"
	}

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
}

func TestLoadBannerSuccess(t *testing.T) {
	filename := "test_banner.txt"

	createTestBannerFile(t, filename)

	defer os.Remove(filename)

	result, err := LoadBanner(filename)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 95 {
		t.Fatalf("expected 95 characters, got %d", len(result))
	}

	if _, ok := result['A']; !ok {
		t.Fatalf("expected character A to exist")
	}

	if len(result['A']) != 8 {
		t.Fatalf("expected A to have 8 lines")
	}
}

func TestLoadBannerFileNotFound(t *testing.T) {
	result, err := LoadBanner("notfound.txt")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if result != nil {
		t.Fatalf("expected nil map on error")
	}
}

func TestLoadBannerSpaceCharacter(t *testing.T) {
	filename := "test_banner.txt"

	createTestBannerFile(t, filename)

	defer os.Remove(filename)

	result, err := LoadBanner(filename)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	spaceArt, ok := result[' ']

	if !ok {
		t.Fatalf("space character not found")
	}

	if len(spaceArt) != 8 {
		t.Fatalf("expected 8 lines for space")
	}
}

func TestLoadBannerTildeCharacter(t *testing.T) {
	filename := "test_banner.txt"

	createTestBannerFile(t, filename)

	defer os.Remove(filename)

	result, err := LoadBanner(filename)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, ok := result['~']; !ok {
		t.Fatalf("expected ~ character to exist")
	}
}
