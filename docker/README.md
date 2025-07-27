# Docker Dagger Module

This module provides Dagger functions to build and push Docker images using the Dagger SDK.

## Features

- Build Docker images from a directory and Dockerfile
- Push images to a registry
- Combined build and push workflow

## Usage

### Import

Import the module in your Go Dagger project:

```go
import "github.com/wouter2397/dagger-modules/docker"
```

### Example

```go
dockerMod := &docker.Docker{}

// Build a Docker image
container := dockerMod.DockerBuild(ctx, dir, "Dockerfile")

// Push the image to a registry
imageRef, err := dockerMod.PushImage(ctx, container, "registry.example.com/myimage:latest")
if err != nil {
    // handle error
}

// Build and push in one step
imageRef, err = dockerMod.BuildAndPush(ctx, dir, "Dockerfile", "registry.example.com/myimage:latest")
if err != nil {
    // handle error
}
```

### Functions

- `DockerBuild(ctx, dir, file)`: Builds a Docker image from the specified directory and Dockerfile.
- `PushImage(ctx, container, address)`: Pushes the built image to the specified registry address.
- `BuildAndPush(ctx, dir, file, address)`: Builds and pushes the image in one step.