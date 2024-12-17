// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package jsonapi

import "net/netip"

type Fteid struct {
	Addr netip.Addr `json:"addr"`
	Teid uint32     `json:"teid"`
}

func NewFteid(addr netip.Addr, teid uint32) *Fteid {
	return &Fteid{
		Addr: addr,
		Teid: teid,
	}
}
