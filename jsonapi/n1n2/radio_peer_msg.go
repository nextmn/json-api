// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"net/netip"

	"github.com/nextmn/json-api/jsonapi"
)

// RadioPeerMsg is used to discover UEs/gNBs "radio" peers.
type RadioPeerMsg struct {
	Control jsonapi.ControlURI `json:"control"`
	Data    netip.AddrPort     `json:"data"`
}
