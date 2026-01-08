// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n1n2

import (
	"net/netip"

	"github.com/nextmn/json-api/jsonapi"
)

type Session struct {
	Addr          netip.Addr     `json:"ue-addr"`
	Dnn           string         `json:"dnn"`
	UplinkFteid   *jsonapi.Fteid `json:"uplink-fteid,omitempty"`
	DownlinkFteid *jsonapi.Fteid `json:"downlink-fteid,omitempty"`

	// when ForwardDownlinkFteid is not empty,
	// PDUs received on DownlinkFteid must be forwarded to it
	ForwardDownlinkFteid *jsonapi.Fteid `json:"forward-fteid,omitempty"`
}
