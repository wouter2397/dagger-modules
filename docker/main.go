// A generated module for Docker functions
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
	"dagger/docker/internal/dagger"
)

type Docker struct{}

func (d *Docker) DockerBuild(ctx context.Context, dir *dagger.Directory, file string) *dagger.Container {
	return dag.Container(dagger.ContainerOpts{}).Build(dir, dagger.ContainerBuildOpts{Dockerfile: file})
}

func (d *Docker) PushImage(ctx context.Context, container *dagger.Container, address string) (string, error) {
	string, error := container.Publish(ctx, address)
	if error != nil {
		return "", error
	}
	return string, nil
}

func (d *Docker) BuildAndPush(ctx context.Context, dir *dagger.Directory, file, address string) (string, error) {
	container := d.DockerBuild(ctx, dir, file)
	imageRef, err := d.PushImage(ctx, container, address)
	if err != nil {
		return "", err
	}
	return imageRef, nil
}
