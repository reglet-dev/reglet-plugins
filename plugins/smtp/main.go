// Package main provides an SMTP plugin for Reglet.
// This is compiled to WASM and loaded by the Reglet runtime.
//go:build wasip1

package main

import (
	"context"
	"log/slog"

	regletsdk "github.com/reglet-dev/reglet-sdk/go"
	regletnet "github.com/reglet-dev/reglet-sdk/go/net"
)

// wrapDialSMTP wraps the SDK's DialSMTP to return the local SMTPConnectResult type
func wrapDialSMTP(ctx context.Context, host, port string, timeoutMs int, useTLS bool, useStartTLS bool) (*SMTPConnectResult, error) {
	result, err := regletnet.DialSMTP(ctx, host, port, timeoutMs, useTLS, useStartTLS)
	if err != nil {
		return nil, err
	}
	return &SMTPConnectResult{
		Connected:      result.Connected,
		Address:        result.Address,
		ResponseTimeMs: result.ResponseTimeMs,
		Banner:         result.Banner,
		TLS:            result.TLS,
		TLSVersion:     result.TLSVersion,
		TLSCipherSuite: result.TLSCipherSuite,
		TLSServerName:  result.TLSServerName,
	}, nil
}

func init() {
	slog.Info("SMTP plugin init() started")
	regletsdk.Register(&smtpPlugin{
		DialSMTP: wrapDialSMTP,
	})
	slog.Info("SMTP plugin init() registered")
}

// main function for the WASM plugin.
func main() {}
