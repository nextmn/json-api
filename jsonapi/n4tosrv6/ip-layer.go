// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

import (
	"net/netip"
	"slices"
)

type IpLayer struct {
	Src []netip.Prefix `json:"src,omitzero"`
	Dst []netip.Prefix `json:"dst,omitzero"`
}

func prefixContainsFunc(addr netip.Addr) func(p netip.Prefix) bool {
	return func(p netip.Prefix) bool {
		return p.Contains(addr)
	}
}

func (l IpLayer) Match(src, dst netip.Addr) bool {
	if src.IsValid() && l.Src != nil && !slices.ContainsFunc(l.Src, prefixContainsFunc(src)) {
		return false
	}
	if dst.IsValid() && l.Dst != nil && !slices.ContainsFunc(l.Src, prefixContainsFunc(dst)) {
		return false
	}
	return true
}
