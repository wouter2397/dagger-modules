// A generated module for DaggerScan functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/scan/internal/dagger"
	"encoding/json"
	"fmt"
	"strings"
)

type Rating struct {
	Severity string `json:"severity"`
}

type Affected struct {
	Ref string `json:"ref"`
}

type Vulnerability struct {
	ID     string `json:"id"`
	Source struct {
		Name string `json:"name"`
	} `json:"source"`
	Ratings []Rating   `json:"ratings"`
	Affects []Affected `json:"affects"`
}

type CycloneDXReport struct {
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
	BOMFormat       string          `json:"bomFormat"`
}

type Trivy struct{}

func (t *Trivy) Base() *dagger.Container {
	return dag.Container().
		From(fmt.Sprintf("aquasec/trivy")).
		WithMountedCache("/root/.cache/trivy", dag.CacheVolume("trivy-db-cache"))
}

func (t *Trivy) ScanImage(ctx context.Context, imageRef string) *dagger.File {

	return t.Base().
		WithExec([]string{
			"trivy",
			"image",
			"--scanners", "vuln",
			"--quiet",
			"--severity", "HIGH,CRITICAL",
			"--format", "cyclonedx",
			"--output", "sbom.json",
			imageRef}).File("sbom.json")
}

func (t *Trivy) ScanContainer(ctx context.Context, ctr *dagger.Container, imageRef string) *dagger.File {

	return t.Base().
		WithMountedFile("/scan/"+imageRef, ctr.AsTarball()).
		WithExec([]string{
			"trivy",
			"image",
			"--scanners", "vuln",
			"--quiet",
			"--severity", "HIGH,CRITICAL",
			"--format", "cyclonedx",
			"--output", "sbom.json",
			"--input", "/scan/" + imageRef}).File("sbom.json")
}

func (t *Trivy) AnalyzeResults(ctx context.Context, sbom *dagger.File) (string, error) {
	var report CycloneDXReport
	content, err := sbom.Contents(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	if err := json.Unmarshal([]byte(content), &report); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	hasCritical := false
	total := len(report.Vulnerabilities)

	output := fmt.Sprintf("🔍 Found %d total vulnerabilities\n", total)

	for _, v := range report.Vulnerabilities {
		severity := "UNKNOWN"
		if len(v.Ratings) > 0 {
			severity = strings.ToUpper(v.Ratings[0].Severity)
		}

		if severity == "CRITICAL" {
			hasCritical = true
		}

		affected := "(unknown component)"
		if len(v.Affects) > 0 {
			affected = v.Affects[0].Ref
		}

		output += fmt.Sprintf("- [%s] %s in %s (via %s)\n", severity, v.ID, affected, v.Source.Name)
	}

	if hasCritical {
		output += "❌ CRITICAL vulnerabilities found. Failing pipeline.\n"
		return output, fmt.Errorf("critical vulnerabilities found")
	}

	output += "✅ No CRITICAL vulnerabilities found. Passing pipeline.\n"
	return output, nil
}

func (t *Trivy) ScanAndAnalyze(ctx context.Context, imageRef string) (string, error) {
	sbom2 := t.ScanImage(ctx, imageRef)
	return t.AnalyzeResults(ctx, sbom2)
}
