# BuildImage Dagger Module

This module provides a Dagger function to build, scan, push, and sign container images in a single workflow.

## Features

- Build a Docker image from a directory and Dockerfile
- Scan the built image for vulnerabilities using Trivy
- Analyze and summarize scan results
- Push the image to a container registry
- Sign the pushed image with Cosign

## Usage

### Import

Import the module in your Go Dagger project:

```go
import "github.com/wouter2397/dagger-modules/build-image"
```

### Example

```go
buildImageMod := &buildimage.BuildImage{}

output, err := buildImageMod.BuildImage(
    ctx,
    dir,              // *dagger.Directory: build context directory
    "Dockerfile",     // string: Dockerfile name
    "registry.example.com/myimage:latest", // string: image reference
    privateKey,       // *dagger.Secret: Cosign private key
    "username",       // string: registry username
    registryPassword, // *dagger.Secret: registry password
    password,         // *dagger.Secret: Cosign password
)
if err != nil {
    // handle error
}
fmt.Println(output)
```

### Parameters

- `dir` (`*dagger.Directory`): Build context directory
- `file` (`string`): Dockerfile name
- `imageRef` (`string`): Image reference (e.g., registry URL)
- `privateKey` (`*dagger.Secret`): Cosign private key
- `registryUsername` (`string`): Registry username
- `registryPassword` (`*dagger.Secret`): Registry password
- `password` (`*dagger.Secret`): Cosign password