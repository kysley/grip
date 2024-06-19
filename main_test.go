// package main

// import (
// 	"io/ioutil"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"strings"
// 	"testing"
// )

// func TestMain(t *testing.T) {
// 	// Create test directory and files
// 	testDir := "test_dir"
// 	os.Mkdir(testDir, 0755)
// 	defer os.RemoveAll(testDir)

// 	libDir := filepath.Join(testDir, "lib")
// 	os.Mkdir(libDir, 0755)

// 	file1 := filepath.Join(libDir, "front.go")
// 	file2 := filepath.Join(libDir, "back.go")
// 	ioutil.WriteFile(file1, []byte("package lib\n\nfunc Front() {\n    // front function\n}"), 0644)
// 	ioutil.WriteFile(file2, []byte("package lib\n\nfunc Back() {\n    // back function\n}"), 0644)

// 	headFile := "head.txt"
// 	ioutil.WriteFile(headFile, []byte("Header content"), 0644)
// 	defer os.Remove(headFile)

// 	// Test cases
// 	tests := []struct {
// 		args        []string
// 		expected    string
// 		shouldError bool
// 	}{
// 		{
// 			args:        []string{testDir, "--out=out.txt"},
// 			expected:    "// lib/front.go\npackage lib\n\nfunc Front() {\n    // front function\n}\n\n// lib/back.go\npackage lib\n\nfunc Back() {\n    // back function\n}\n",
// 			shouldError: false,
// 		},
// 		{
// 			args:        []string{testDir, "--out=out.txt", "--match=lib/*.go"},
// 			expected:    "// lib/front.go\npackage lib\n\nfunc Front() {\n    // front function\n}\n\n// lib/back.go\npackage lib\n\nfunc Back() {\n    // back function\n}\n",
// 			shouldError: false,
// 		},
// 		{
// 			args:        []string{testDir, "--out=out.txt", "--head=Inline header"},
// 			expected:    "Inline header\n\n// lib/front.go\npackage lib\n\nfunc Front() {\n    // front function\n}\n\n// lib/back.go\npackage lib\n\nfunc Back() {\n    // back function\n}\n",
// 			shouldError: false,
// 		},
// 		{
// 			args:        []string{testDir, "--out=out.txt", "--head-file=" + headFile},
// 			expected:    "Header content\n\n// lib/front.go\npackage lib\n\nfunc Front() {\n    // front function\n}\n\n// lib/back.go\npackage lib\n\nfunc Back() {\n    // back function\n}\n",
// 			shouldError: false,
// 		},
// 		{
// 			args:        []string{testDir, "--out=out.txt", "--head=Inline header", "--head-file=" + headFile},
// 			expected:    "",
// 			shouldError: true,
// 		},
// 	}

// 	for _, test := range tests {
// 		cmd := exec.Command("go", "run", ".", test.args...)
// 		output, err := cmd.CombinedOutput()
// 		if test.shouldError {
// 			if err == nil {
// 				t.Fatalf("expected an error but got none, output: %s", output)
// 			}
// 		} else {
// 			if err != nil {
// 				t.Fatalf("unexpected error: %v, output: %s", err, output)
// 			}
// 			content, err := ioutil.ReadFile("out.txt")
// 			if err != nil {
// 				t.Fatalf("error reading output file: %v", err)
// 			}
// 			if strings.TrimSpace(string(content)) != strings.TrimSpace(test.expected) {
// 				t.Fatalf("expected %q but got %q", test.expected, string(content))
// 			}
// 		}
// 	}
// }
