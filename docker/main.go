// This module provides functionality to build and push Docker images

package main

import (
	"context"
	"dagger/docker/internal/dagger"
)

type Docker struct{}

// DockerBuild builds a Docker image from the specified directory and Dockerfile.
func (d *Docker) DockerBuild(ctx context.Context, dir *dagger.Directory, file string) *dagger.Container {
	return dir.DockerBuild(dagger.DirectoryDockerBuildOpts{Dockerfile: file})
}

// PushImage pushes a Docker image to the specified address and returns the image digest.
func (d *Docker) PushImage(ctx context.Context, container *dagger.Container, address string) (string, error) {
	string, error := container.Publish(ctx, address)
	if error != nil {
		return "", error
	}
	return string, nil
}

// BuildAndPush builds and pushes a Docker image using the DockerBuild and PushImage functions.
func (d *Docker) BuildAndPush(ctx context.Context, dir *dagger.Directory, file, address string) (string, error) {
	container := d.DockerBuild(ctx, dir, file)
	imageRef, err := d.PushImage(ctx, container, address)
	if err != nil {
		return "", err
	}
	return imageRef, nil
}
