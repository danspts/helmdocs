package main

import (
	"os"
	"testing"
)

func TestHelmdocs(t *testing.T) {
	t.Run("MissingSubcommand", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic, got none")
			}
		}()
		helmdocs([]string{"helmdocs"})
	})

	t.Run("InvalidSubcommand", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic, got none")
			}
		}()
		helmdocs([]string{"helmdocs", "invalid"})
	})

	t.Run("GenerateMissingSubcommand", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic, got none")
			}
		}()
		helmdocs([]string{"helmdocs", "generate"})
	})

	t.Run("ValidFlagParse", func(t *testing.T) {
		filename := "test_schema.json"
		content := `{"type": "object"}`
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer os.Remove(filename)
		command := "helmdocs generate readme"
		handleGenerateReadmeCommand(command, []string{"--schema-path", "test_schema.json", "--output", "./README.md"})
		handleGenerateValuesCommand(command, []string{"--schema-path", "test_schema.json", "--output", "./values.yaml", "--omit-default", "false"})
	})
}

func TestReadSchema(t *testing.T) {
	t.Run("ValidFile", func(t *testing.T) {
		filename := "test_schema.json"
		content := `{"type": "object"}`
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer os.Remove(filename)

		_, err = readSchema(filename)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("EmptyFile", func(t *testing.T) {
		filename := "empty_schema.json"
		err := os.WriteFile(filename, []byte(""), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer os.Remove(filename)

		_, err = readSchema(filename)
		if err == nil {
			t.Errorf("Expected error for empty file, got none")
		}
	})

	t.Run("InvalidFile", func(t *testing.T) {
		filename := "non_existent_file.json"
		_, err := readSchema(filename)
		if err == nil {
			t.Errorf("Expected error, got none")
		}
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		filename := "test_invalid.json"
		content := `{"type": "object"`
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer os.Remove(filename)

		_, err = readSchema(filename)
		if err == nil {
			t.Errorf("Expected error, got none")
		}
	})
}

func TestWriteOutput(t *testing.T) {
	t.Run("ValidWrite", func(t *testing.T) {
		filename := "test_output.md"
		content := "Generated Content"

		err := writeOutput(filename, content)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		os.Remove(filename)
	})

	t.Run("InvalidWrite", func(t *testing.T) {
		filename := "/invalid_path/test_output.md"
		content := "Generated Content"

		err := writeOutput(filename, content)
		if err == nil {
			t.Errorf("Expected error, got none")
		}
	})

	t.Run("EmptyContent", func(t *testing.T) {
		filename := "test_output_empty.md"
		content := ""

		err := writeOutput(filename, content)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		os.Remove(filename)
	})
}
