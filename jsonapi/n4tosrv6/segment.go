// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

import (
	"net/netip"
)

type Segment struct {
	netip.Addr
}

func ParseSegment(s string) (Segment, error) {
	seg, err := netip.ParseAddr(s)
	if err != nil {
		return Segment{}, err
	}
	if !seg.Is6() {
		return Segment{}, ErrNotIPv6Address
	}
	return Segment{seg}, nil
}

func (s *Segment) UnmarshalText(text []byte) error {
	if err := s.Addr.UnmarshalText(text); err != nil {
		return err
	}
	if !s.Addr.Is6() {
		return ErrNotIPv6Address
	}
	return nil
}
