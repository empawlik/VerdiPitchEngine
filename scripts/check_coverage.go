package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Configure coverage requirements
const defaultTarget = 80.0

var excludeList = []string{
	"github.com/empawlik/verdi-pitch-engine/cmd/",
}

var packageTargets = map[string]float64{
	"github.com/empawlik/verdi-pitch-engine/internal/converter": 80.0,
	"github.com/empawlik/verdi-pitch-engine/internal/fs":        80.0,
}

type pkgStats struct {
	statements int
	covered    int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run check_coverage.go <coverage.out>")
		os.Exit(1)
	}

	stats, err := parseCoverageFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error processing coverage: %v\n", err)
		os.Exit(1)
	}

	if err := evaluateCoverage(stats); err != nil {
		fmt.Println("\n❌ Test Coverage Thresholds Not Met:")
		fmt.Println(err.Error())
		fmt.Println("\nPlease write more tests to push coverage above the required thresholds.")
		os.Exit(1)
	}

	fmt.Println("✅ All test coverage thresholds met.")
	os.Exit(0)
}

func parseCoverageFile(path string) (map[string]*pkgStats, error) {
	cleanPath := filepath.Clean(path)
	file, err := os.Open(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	stats := make(map[string]*pkgStats)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mode:") || len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) < 3 {
			continue
		}

		filePart := strings.Split(parts[0], ":")[0]
		pkgName := filepath.Dir(filePart)
		statements, _ := strconv.Atoi(parts[1])
		count, _ := strconv.Atoi(parts[2])

		if _, exists := stats[pkgName]; !exists {
			stats[pkgName] = &pkgStats{}
		}
		stats[pkgName].statements += statements
		if count > 0 {
			stats[pkgName].covered += statements
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	return stats, nil
}

func evaluateCoverage(stats map[string]*pkgStats) error {
	var failed []string
	var defaultWarnings []string
	var lowTargetWarnings []string
	var passed []string

	for pkg, stat := range stats {
		if isExcluded(pkg) || stat.statements == 0 {
			continue
		}

		coverage := (float64(stat.covered) / float64(stat.statements)) * 100.0
		target := defaultTarget
		hasExplicitTarget := false
		if specificTarget, ok := packageTargets[pkg]; ok {
			target = specificTarget
			hasExplicitTarget = true
			if target < defaultTarget {
				lowTargetWarnings = append(lowTargetWarnings, fmt.Sprintf("   - %s: %.1f%%", pkg, target))
			}
		} else {
			defaultWarnings = append(defaultWarnings, fmt.Sprintf("   - %s", pkg))
		}

		if coverage < target {
			failed = append(failed, fmt.Sprintf("   - %s: %.1f%% (target: %.1f%%)", pkg, coverage, target))
		} else if hasExplicitTarget {
			passed = append(passed, fmt.Sprintf("   - %s: %.1f%% (target: %.1f%%)", pkg, coverage, target))
		}
	}

	if len(defaultWarnings) > 0 {
		fmt.Printf("\n⚠️  Warning: The following packages have no explicit target and are using the default %.1f%% threshold:\n", defaultTarget)
		fmt.Println(strings.Join(defaultWarnings, "\n"))
	}

	if len(lowTargetWarnings) > 0 {
		fmt.Printf("\n⚠️  Warning: The following explicit targets are below the default %.1f%% threshold:\n", defaultTarget)
		fmt.Println(strings.Join(lowTargetWarnings, "\n"))
	}

	if len(passed) > 0 {
		fmt.Println("\n✅ Packages meeting coverage thresholds:")
		fmt.Println(strings.Join(passed, "\n"))
	}

	if len(failed) > 0 {
		return fmt.Errorf("%s", strings.Join(failed, "\n"))
	}
	return nil
}

func isExcluded(pkg string) bool {
	for _, ex := range excludeList {
		if strings.HasPrefix(pkg, ex) {
			return true
		}
	}
	return false
}
