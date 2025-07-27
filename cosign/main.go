// A generated module for Cosign functions
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
	"dagger/cosign/internal/dagger"
)

type Cosign struct{}

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
