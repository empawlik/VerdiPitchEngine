//go:build mage
// +build mage

package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Help

var Aliases = map[string]interface{}{
	"lint-markdown":  LintMarkdown,
	"fix-markdown":   FixMarkdown,
	"fmt":            Fmt,
	"vet":            Vet,
	"vulncheck":      Vulncheck,
	"lint":           Lint,
	"test":           Test,
	"ci-test":        CiTest,
	"check-coverage": CheckCoverage,
	"check":          Check,
	"deps":           Deps,
}

var (
	Go = "go"
)

// Help displays the available targets.
func Help() {
	fmt.Println("Run 'mage -l' to see available targets.")
}

// LintMarkdown lints Markdown files for 1000-KEYS compliance using central script.
func LintMarkdown() error {
	if _, err := os.Stat("../AIgorLabs-github/scripts/lint-markdown.py"); err == nil {
		fmt.Println("Running central markdown linter from local workspace...")
		return sh.RunV("python3", "../AIgorLabs-github/scripts/lint-markdown.py")
	}
	fmt.Println("WARNING: ../AIgorLabs-github not found. Skipping local markdown linting.")
	return nil
}

// FixMarkdown automatically adds or updates the 1000-KEYS YAML frontmatter on all markdown files.
func FixMarkdown() error {
	fmt.Println("Running Markdown fixer...")
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if path == ".git" || path == "vendor" || path == "node_modules" || path == "bin" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".md") {
			return nil
		}
		if info.Name() == "CHANGELOG.md" {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		contentStr := string(content)
		var body string
		var frontmatter string
		hasFrontmatter := false

		// Regex to match existing frontmatter
		re := regexp.MustCompile(`(?s)^---\n(.*?)\n---(?:\n|$)(.*)`)
		matches := re.FindStringSubmatch(contentStr)

		if len(matches) == 3 {
			hasFrontmatter = true
			frontmatter = matches[1]
			body = matches[2]
		} else {
			body = contentStr
		}

		// Calculate body hash
		hash := sha256.Sum256([]byte(body))
		bodyHash := fmt.Sprintf("%x", hash)[:16]

		dateStr := time.Now().Format("2006-01-02")

		if !hasFrontmatter {
			// Generate new frontmatter
			frontmatter = fmt.Sprintf(`project_name: VerdiPitchEngine
version: 0.1.0
status: active
priority: high
dev_stage: development
agent_role: lead_engineer
agent_weight: 1.0
asset_scope: backend
platform: qnap
tech_stack: [go, ffmpeg]
dependencies: []
created: %s
updated: %s
tags: [audio, dsp]
body_hash: %s`, dateStr, dateStr, bodyHash)
		} else {
			// Update body_hash in existing frontmatter
			reHash := regexp.MustCompile(`(?m)^body_hash\s*:.*$`)
			if reHash.MatchString(frontmatter) {
				frontmatter = reHash.ReplaceAllString(frontmatter, fmt.Sprintf(`body_hash: %s`, bodyHash))
			} else {
				frontmatter += fmt.Sprintf("\nbody_hash: %s", bodyHash)
			}
			// Update updated date
			reUpdated := regexp.MustCompile(`(?m)^updated\s*:.*$`)
			if reUpdated.MatchString(frontmatter) {
				frontmatter = reUpdated.ReplaceAllString(frontmatter, fmt.Sprintf(`updated: %s`, dateStr))
			}
		}

		newContent := fmt.Sprintf("---\n%s\n---\n%s", frontmatter, body)
		if newContent != contentStr {
			fmt.Printf("Fixed frontmatter for %s\n", path)
			return ioutil.WriteFile(path, []byte(newContent), info.Mode())
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("fix-markdown failed: %w", err)
	}
	fmt.Println("Markdown fixing complete!")
	return nil
}

// Deps installs project dependencies.
func Deps() error {
	fmt.Println("Installing dependencies...")
	return sh.RunV(Go, "mod", "download")
}

// Fmt formats code.
func Fmt() error {
	fmt.Println("Formatting code...")
	if err := sh.RunV(Go, "fmt", "./cmd/...", "./internal/..."); err != nil {
		return err
	}
	fmt.Println("Formatting complete!")
	return nil
}

// Vet runs go vet.
func Vet() error {
	fmt.Println("Running go vet...")
	if err := sh.RunV(Go, "vet", "./cmd/...", "./internal/..."); err != nil {
		return err
	}
	fmt.Println("Vet complete!")
	return nil
}

// Vulncheck runs govulncheck for dependency vulnerability scanning.
func Vulncheck() error {
	fmt.Println("Running govulncheck...")
	err := sh.RunV("govulncheck", "./...")
	if err != nil {
		fmt.Println("govulncheck failed or not installed. Installing...")
		if installErr := sh.RunV(Go, "install", "golang.org/x/vuln/cmd/govulncheck@latest"); installErr != nil {
			return installErr
		}
		if runErr := sh.RunV("govulncheck", "./..."); runErr != nil {
			return runErr
		}
	}
	fmt.Println("Vulncheck complete!")
	return nil
}

// Lint runs linters using strict .golangci.yml.
func Lint() error {
	fmt.Println("Running linters...")
	if err := sh.RunV("golangci-lint", "run", "./..."); err != nil {
		fmt.Println("golangci-lint failed or not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest")
		return err
	}
	fmt.Println("Lint complete!")
	return nil
}

// Test runs all tests.
func Test() error {
	fmt.Println("Running tests...")
	if err := sh.RunV(Go, "test", "-v", "-race", "-coverprofile=coverage.out", "./cmd/...", "./internal/..."); err != nil {
		return err
	}
	if err := sh.RunV(Go, "tool", "cover", "-html=coverage.out", "-o", "coverage.html"); err != nil {
		return err
	}
	fmt.Println("Tests complete! Coverage report: coverage.html")
	return nil
}

// CiTest runs tests without race detector for CI speed.
func CiTest() error {
	fmt.Println("Running CI tests...")
	return sh.RunV(Go, "test", "-v", "-coverprofile=coverage.out", "./cmd/...", "./internal/...")
}

// CheckCoverage enforces test coverage minimum targets across the workspace.
func CheckCoverage() error {
	mg.Deps(Test)
	fmt.Println("Enforcing test coverage targets...")
	return sh.RunV(Go, "run", "scripts/check_coverage.go", "coverage.out")
}

// Check runs all checks (format, vet, lint, vulncheck, check-coverage).
func Check() {
	mg.Deps(Fmt, Vet, Lint, LintMarkdown, Vulncheck, CheckCoverage)
}
