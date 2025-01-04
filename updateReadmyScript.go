package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	baseDir       = "problems"
	readmeFile    = "README.md"
	startSection  = "## 📜 Problem List"
	headingPrefix = "## "

	// Folders corresponding to difficulty
	folderEasy   = "easy"
	folderMedium = "medium"
	folderHard   = "hard"
)

type Problem struct {
	Title, Difficulty, FolderName string
	Number                        int
}

func updateReadmy() {
	// We’ll parse subfolders in the order: hard, medium, easy
	// so that they appear grouped in the final table.
	difficulties := []string{folderHard, folderMedium, folderEasy}

	var allProblems []Problem
	count := 1

	for _, difficultyFolder := range difficulties {
		// For each difficulty (hard/medium/easy), read the subfolders
		dirPath := filepath.Join(baseDir, difficultyFolder)

		entries, err := os.ReadDir(dirPath)
		if err != nil {
			// If the folder doesn't exist or is empty, we skip it
			continue
		}

		coloredDifficulty := colorDifficulty(difficultyFolder)

		for _, entry := range entries {
			if entry.IsDir() {
				folderName := entry.Name()
				title := parseTitle(folderName)

				problem := Problem{
					Number:     count,
					Title:      title,
					Difficulty: coloredDifficulty,
					FolderName: folderName,
				}
				allProblems = append(allProblems, problem)
				count++
			}
		}
	}

	// Build the new table lines
	newTable := buildTable(allProblems)

	if err := updateReadme(newTable); err != nil {
		log.Fatalf("Failed to update the README: %v", err)
	}

	fmt.Println("README updated successfully!")
}

func colorDifficulty(difficultyFolder string) string {
	switch difficultyFolder {
	case folderEasy:
		return "🟢 Easy"
	case folderMedium:
		return "🟡 Medium"
	case folderHard:
		return "🔴 Hard"
	default:
		return difficultyFolder
	}
}

// parseTitle splits a CamelCase or PascalCase string into a space-separated title.
// Example: "PalindromeNumber" → "Palindrome Number"
func parseTitle(folderName string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	parts := re.FindAllString(folderName, -1)
	return strings.Join(parts, " ")
}

// buildTable constructs the Problem List section as Markdown lines.
func buildTable(problems []Problem) []string {
	tableLines := []string{
		startSection,
		"",
		"| Problem # | Title                                  | Difficulty                 | Folder Name                      |",
		"|-----------|----------------------------------------|----------------------------|----------------------------------|",
	}
	for _, p := range problems {
		line := fmt.Sprintf(
			"| %d         | %s | %s | `%s` |",
			p.Number,
			p.Title,
			p.Difficulty,
			p.FolderName,
		)
		tableLines = append(tableLines, line)
	}

	tableLines = append(tableLines, "\n---\n")
	return tableLines
}

// updateReadme reads my README file and replaces all lines
// from the "## 📜 Problem List" heading up to the next heading ("## ")
// with the new table lines.
func updateReadme(newTable []string) error {
	file, err := os.Open(readmeFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var updatedLines []string
	scanner := bufio.NewScanner(file)

	inProblemSection := false

	for scanner.Scan() {
		line := scanner.Text()

		// Trying to find the start of the Problem List section
		if strings.HasPrefix(line, startSection) {
			// Replacing old table lines with new ones
			updatedLines = append(updatedLines, newTable...)
			inProblemSection = true
			continue
		}

		// If we've started skipping lines in the Problem List section,
		// we keep skipping until we encounter the next heading
		if inProblemSection {
			// Once we see a new heading (## ...) that isn't the same as startSection,
			// we stop skipping and resume copying lines
			if strings.HasPrefix(line, headingPrefix) && line != startSection {
				inProblemSection = false
				// Add the heading to updatedLines
				updatedLines = append(updatedLines, line)
			}
			// Otherwise, ignore everything in between
			continue
		}

		// Outside the Problem List section, just keep lines as is
		updatedLines = append(updatedLines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Write updated lines back to README
	newContent := strings.Join(updatedLines, "\n") + "\n"
	return os.WriteFile(readmeFile, []byte(newContent), 0o644)
}
