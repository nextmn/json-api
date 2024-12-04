// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

import (
	"fmt"
	"net/netip"
)

type Segment struct {
	netip.Addr
}

func NewSegment(s string) (*Segment, error) {
	seg, err := netip.ParseAddr(s)
	if err != nil {
		return nil, err
	}
	if !seg.Is6() {
		return nil, fmt.Errorf("Segment must be an IPv6 address")
	}
	return &Segment{seg}, nil
}

func (s *Segment) UnmarshalText(text []byte) error {
	if err := s.Addr.UnmarshalText(text); err != nil {
		return err
	}
	if !s.Addr.Is6() {
		return fmt.Errorf("Segment must be an IPv6 address")
	}
	return nil
}
