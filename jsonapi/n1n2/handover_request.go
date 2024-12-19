// Copyright 2024 Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"net/netip"

	"github.com/nextmn/json-api/jsonapi"
)

// HandoverRequest is send by the CP to the target gNB during the handover preparation phase
type HandoverRequest struct {
	UeCtrl    jsonapi.ControlURI `json:"ue-ctrl"`
	Cp        jsonapi.ControlURI `json:"ue-ctrl"`
	TargetgNB jsonapi.ControlURI `json:"ue-ctrl"`
	Sessions  []struct {
		Addr        netip.Addr    `json:"ue-addr"`
		UplinkFteid jsonapi.Fteid `json:"uplink-fteid"`
	} `json:"sessions"`
}
