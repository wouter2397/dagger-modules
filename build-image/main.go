// A generated module for BuildImage functions
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
	"dagger/build-image/internal/dagger"
)

type BuildImage struct{}

// Returns a container that echoes whatever string argument is provided
func (m *BuildImage) BuildImage(ctx context.Context, dir *dagger.Directory, file string, imageRef string, privateKey *dagger.Secret, registryUsername string, registryPassword *dagger.Secret, password *dagger.Secret) (string, error) {
	container := dag.Docker().DockerBuild(dir, file)
	sbom := dag.Trivy().ScanContainer(container, imageRef)
	output, err := dag.Trivy().AnalyzeResults(ctx, sbom)
	if err != nil {
		return "", err
	}
	imageDigest, err2 := dag.Docker().PushImage(ctx, container, imageRef)
	output += imageDigest
	if err2 != nil {
		return "", err2
	}
	_, err3 := dag.Cosign().Sign(ctx, privateKey, password, registryPassword, imageDigest, dagger.CosignSignOpts{
		RegistryUsername: registryUsername,
	})
	if err3 != nil {
		return "", err3
	}
	return output, nil
}
