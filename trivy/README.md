# Trivy Dagger Module

This module provides Dagger functions to scan container images for vulnerabilities using [Trivy](https://github.com/aquasecurity/trivy) and analyze the results.

## Features

- Scan container images or containers for vulnerabilities (HIGH, CRITICAL)
- Output results in CycloneDX SBOM format
- Analyze scan results and summarize vulnerabilities
- Fail pipeline if CRITICAL vulnerabilities are found

## Usage

### Import

Import the module in your Go Dagger project:

```go
import "github.com/wouter2397/dagger-modules/trivy"
```

### Example

```go
trivyMod := &trivy.Trivy{}

// Scan an image and analyze results
output, err := trivyMod.ScanAndAnalyze(ctx, "alpine:latest")
if err != nil {
    // handle error (e.g., critical vulnerabilities found)
}
fmt.Println(output)
```

### Functions

- `ScanImage(ctx, imageRef)`: Scans a container image by reference and returns the SBOM file.
- `ScanContainer(ctx, ctr, imageRef)`: Scans a Dagger container and returns the SBOM file.
- `AnalyzeResults(ctx, sbom)`: Analyzes a CycloneDX SBOM file and summarizes vulnerabilities.
- `ScanAndAnalyze(ctx, imageRef)`: Scans an image and analyzes results in one step.