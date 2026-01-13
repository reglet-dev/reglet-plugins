// Package main provides a TCP plugin for Reglet.
// This is compiled to WASM and loaded by the Reglet runtime.
//go:build wasip1

package main

import (
	"context"
	"log/slog"

	regletsdk "github.com/reglet-dev/reglet-sdk/go"
	regletnet "github.com/reglet-dev/reglet-sdk/go/net"
)

// wrapDialTCP wraps the SDK's DialTCP to return the local TCPConnectResult type
func wrapDialTCP(ctx context.Context, host, port string, timeoutMs int, useTLS bool) (*TCPConnectResult, error) {
	result, err := regletnet.DialTCP(ctx, host, port, timeoutMs, useTLS)
	if err != nil {
		return nil, err
	}
	return &TCPConnectResult{
		Connected:       result.Connected,
		Address:         result.Address,
		ResponseTimeMs:  result.ResponseTimeMs,
		RemoteAddr:      result.RemoteAddr,
		LocalAddr:       result.LocalAddr,
		TLS:             result.TLS,
		TLSVersion:      result.TLSVersion,
		TLSCipherSuite:  result.TLSCipherSuite,
		TLSServerName:   result.TLSServerName,
		TLSCertSubject:  result.TLSCertSubject,
		TLSCertIssuer:   result.TLSCertIssuer,
		TLSCertNotAfter: result.TLSCertNotAfter,
	}, nil
}

func init() {
	slog.Info("TCP plugin init() started")
	regletsdk.Register(&tcpPlugin{DialTCP: wrapDialTCP})
	slog.Info("TCP plugin init() registered")
}

// main function for the WASM plugin.
func main() {}
