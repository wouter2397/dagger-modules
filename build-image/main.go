// A generated module for BuildImage functions
//
// This dagger module is created to integrate with the following dagger modules:
// - docker
// - trivy
// - cosign
// It provides a high-level function to build a Docker image, scan it for vulnerabilities,
// push it to a registry, and sign it.

package main

import (
	"context"
	"dagger/build-image/internal/dagger"
)

type BuildImage struct{}

// BuildImage builds a Docker image from the specified directory and Dockerfile,
// scans it for vulnerabilities using Trivy, pushes it to a container registry,
// and signs it using Cosign.
// It returns the output of the vulnerability analysis and the image digest.
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
