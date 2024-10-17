// Copyright 2024 Louis Royer and the NextMN-json-api contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT
package jsonapi

import "net/netip"

type MatchGtp struct {
	Src  netip.Addr `json:"source-ip"`
	Teid uint32     `json:"teid"`
}
