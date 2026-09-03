// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package jsonapi

import (
	"net/netip"
)

type TEID uint32

// FTEID represents an F-TEID. It is composed of an IP Adress and a TEID.
type Fteid struct {
	Addr netip.Addr `json:"addr"`
	Teid TEID       `json:"teid"`
}

// IsValid reports whether the [Fteid] is an initialized Fteid (non-zero TEID, and not the zero Fteid).
func (f Fteid) IsValid() bool {
	return f.Teid != 0 && f.Addr.IsValid()
}
