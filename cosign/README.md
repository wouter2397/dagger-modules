# Cosign Dagger Module

This module provides a Dagger function to sign container images using [Cosign](https://github.com/sigstore/cosign).

## Features

- Sign container images with a private key and password
- Supports registry authentication
- Runs Cosign in a secure, non-root container

## Usage

### Import

Import the module in your Go Dagger project:

```go
import "github.com/wouter2397/dagger-modules/cosign"
```

### Example

```go
cos := &cosign.Cosign{}
output, err := cos.Sign(
    ctx,
    privateKey,         // dagger.Secret: Cosign private key
    password,           // dagger.Secret: Cosign password
    &registryUsername,  // *string: Registry username (optional)
    &registryPassword,  // *dagger.Secret: Registry password (optional)
    digest,             // string: Image digest to sign
)
if err != nil {
    // handle error
}
fmt.Println(output)
```

### Parameters

- `privateKey` (`dagger.Secret`): Cosign private key
- `password` (`dagger.Secret`): Cosign password
- `registryUsername` (`*string`): Registry username (optional)
- `registryPassword` (`*dagger.Secret`): Registry password (optional)
- `digest` (`string`): Container image digest to sign