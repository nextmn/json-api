// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package n4tosrv6

import "github.com/nextmn/json-api/jsonapi"

type GtpLayer struct {
	Teid jsonapi.TEID `json:"udp.teid,omitzero"` // TEID
}

func (l GtpLayer) Match(Teid jsonapi.TEID) bool {
	return l.Teid == 0 || Teid == 0 || (Teid == l.Teid)
}
