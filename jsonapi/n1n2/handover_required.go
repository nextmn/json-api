// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"net/netip"

	"github.com/nextmn/json-api/jsonapi"
)

// HandoverRequired is send by the source gNB to the CP to start the handover preparation phase
type HandoverRequired struct {
	Ue          jsonapi.ControlURI `json:"ue"`
	PduSessions []netip.Addr       `json:"pdu-sessions"` // list of all pdu sessions of the UE to be moved
	TargetgNB   jsonapi.ControlURI `json:"target-gnb"`
}
