// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

import (
	"net/netip"
)

type NextHop struct {
	netip.Addr
}

func ParseNextHop(nh string) (NextHop, error) {
	h, err := netip.ParseAddr(nh)
	if err != nil {
		return NextHop{}, err
	}
	if !h.Is6() {
		return NextHop{}, ErrNotIPv6Address
	}
	return NextHop{h}, nil
}

func (nh *NextHop) UnmarshalText(text []byte) error {
	if err := nh.Addr.UnmarshalText(text); err != nil {
		return err
	}
	if !nh.Addr.Is6() {
		return ErrNotIPv6Address
	}
	return nil
}
