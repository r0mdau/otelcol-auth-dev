// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package customauthextension // import "github.com/r0mdau/customauthextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"

	"github.com/r0mdau/customauthextension/internal/metadata"
)

// NewFactory creates a factory for the static bearer token Authenticator extension.
func NewFactory() extension.Factory {
	return extension.NewFactory(
		metadata.Type,
		createDefaultConfig,
		createExtension,
		metadata.ExtensionStability,
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createExtension(_ context.Context, _ extension.CreateSettings, cfg component.Config) (extension.Extension, error) {
	return newServerAuthExtension(cfg.(*Config))
}
