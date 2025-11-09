// This module provides functionality to sign container images using Cosign.
package main

import (
	"context"
	"dagger/cosign/internal/dagger"
)

type Cosign struct{}

// Sign signs a container image using Cosign.
func (c *Cosign) Sign(
	ctx context.Context,
	// Cosign private key
	privateKey dagger.Secret,
	// Cosign password
	password dagger.Secret,
	// registry username
	registryUsername *string,
	// name of the image
	registryPassword *dagger.Secret,
	// Container image digests to sign
	digest string) ([]string, error) {
	stdouts := []string{}
	cmd := []string{"cosign", "sign", digest, "--key", "env://COSIGN_PRIVATE_KEY", "--tlog-upload=false"}
	if registryUsername != nil && registryPassword != nil {
		pwd, err := registryPassword.Plaintext(ctx)
		if err != nil {
			return nil, err
		}

		cmd = append(
			cmd,
			"--registry-username",
			*registryUsername,
			"--registry-password",
			pwd,
		)
	}
	cosign := dag.
		Container().
		From("chainguard/cosign:latest").
		WithUser("nonroot").
		WithEnvVariable("COSIGN_YES", "true").
		WithSecretVariable("COSIGN_PASSWORD", &password).
		WithSecretVariable("COSIGN_PRIVATE_KEY", &privateKey).
		WithExec(cmd)

	stdout, err := cosign.Stdout(ctx)
	if err != nil {
		return nil, err
	}

	stdouts = append(stdouts, stdout)

	return stdouts, nil
}
