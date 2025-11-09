// This module provides functionality integrate the following dagger modules: docker, trivy, and cosign.
// It provides a high-level function to build, scan, push, and sign container images.

package main

import (
	"context"
	"dagger/build-image/internal/dagger"
)

type BuildImage struct{}

// BuildImage builds, scans, pushes, and signs a Docker image.
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
