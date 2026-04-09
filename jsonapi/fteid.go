// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package jsonapi

import "net/netip"

type TEID uint32

type Fteid struct {
	Addr netip.Addr `json:"addr"`
	Teid TEID       `json:"teid"`
}

func NewFteid(addr netip.Addr, teid TEID) *Fteid {
	return &Fteid{
		Addr: addr,
		Teid: teid,
	}
}
