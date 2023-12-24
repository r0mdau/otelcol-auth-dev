// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package customauthextension // import "github.com/r0mdau/customauthextension"

import (
	"errors"
)

var (
	errNoCredentialSource = errors.New("no shared key provided")
)

type Config struct {

	// SharedKey symetric.
	SharedKey string `mapstructure:"shared_key,omitempty"`
}

func (cfg *Config) Validate() error {
	if cfg.SharedKey == "" {
		return errNoCredentialSource
	}

	return nil
}
