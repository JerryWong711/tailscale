// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build !linux

package main

import (
	"errors"
	"io"
	"net/netip"
	"time"
)

func getUDPConnKernelTimestamp() (io.ReadWriteCloser, error) {
	return nil, errors.New("unimplemented")
}

func measureSTUNRTTKernel(conn io.ReadWriteCloser, hostname string, dst netip.AddrPort) (rtt time.Duration, err error) {
	return 0, errors.New("unimplemented")
}

func getProtocolSupportInfo(p protocol) protocolSupportInfo {
	switch p {
	case protocolSTUN:
		return protocolSupportInfo{
			kernelTS:    false,
			userspaceTS: true,
			stableConn:  true,
		}
	case protocolHTTPS:
		return protocolSupportInfo{
			kernelTS:    false,
			userspaceTS: true,
			stableConn:  true,
		}
	case protocolTCP:
		return protocolSupportInfo{
			kernelTS:    true,
			userspaceTS: false,
			stableConn:  true,
		}
	}
	return protocolSupportInfo{}
}

func setSOReuseAddr(fd uintptr) error {
	return nil
}
