package config

import "github.com/bitrise-io/go-utils/freezable"

const (
	// DefaultPort ...
	DefaultPort = "3645"
	// DefaultFrontendPort ...
	DefaultFrontendPort = "4567"

	// MinimalValidSecrets ...
	MinimalValidSecrets = "{}"
	// MinimalValidBitriseYML ...
	MinimalValidBitriseYML = "format_version: 1.3.0"
)

var (
	// BitriseYMLPath ...
	BitriseYMLPath freezable.String
	// SecretsYMLPath ...
	SecretsYMLPath freezable.String
)
