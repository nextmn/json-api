// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"net/netip"

	"github.com/nextmn/json-api/jsonapi"
)

// Handovercommand is sent by the CP to the source gNB to start the execution of handover, and forwarded to the UE
type HandoverCommand struct {
	UeCtrl    jsonapi.ControlURI `json:"ue-ctrl"`
	Cp        jsonapi.ControlURI `json:"cp"`
	Sessions  []netip.Addr       `json:"sessions"`
	GNBSource jsonapi.ControlURI `json:"gnb-source"`
	GNBTarget jsonapi.ControlURI `json:"gnb-target"`
}
