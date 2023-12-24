// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package customauthextension // import "github.com/r0mdau/customauthextension"

import (
	"context"
	"errors"
	"strings"

	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/auth"
)

var (
	errNoAuth             = errors.New("no basic auth provided")
	errInvalidCredentials = errors.New("invalid credentials")
)

type customAuth struct {
	SharedKey string
}

func newServerAuthExtension(cfg *Config) (auth.Server, error) {
	ca := customAuth{
		SharedKey: cfg.SharedKey,
	}
	return auth.NewServer(
		auth.WithServerStart(ca.serverStart),
		auth.WithServerAuthenticate(ca.authenticate),
	), nil
}

func (ca *customAuth) serverStart(_ context.Context, _ component.Host) error {
	return nil
}

func (ca *customAuth) authenticate(ctx context.Context, headers map[string][]string) (context.Context, error) {
	auth := getAuthHeader(headers)
	if auth == "" {
		return ctx, errNoAuth
	}

	if auth != ca.SharedKey {
		return ctx, errInvalidCredentials
	}

	cl := client.FromContext(ctx)
	authData := &authData{
		raw: auth,
	}
	cl.Auth = authData
	return client.NewContext(ctx, cl), nil
}

func getAuthHeader(h map[string][]string) string {
	const (
		canonicalHeaderKey = "Authorization"
		metadataKey        = "authorization"
	)

	authHeaders, ok := h[canonicalHeaderKey]

	if !ok {
		authHeaders, ok = h[metadataKey]
	}

	if !ok {
		for k, v := range h {
			if strings.EqualFold(k, metadataKey) {
				authHeaders = v
				break
			}
		}
	}

	if len(authHeaders) == 0 {
		return ""
	}

	return authHeaders[0]
}

var _ client.AuthData = (*authData)(nil)

type authData struct {
	raw string
}

func (a *authData) GetAttribute(name string) any {
	switch name {
	case "raw":
		return a.raw
	default:
		return nil
	}
}

func (*authData) GetAttributeNames() []string {
	return []string{"username", "raw"}
}
