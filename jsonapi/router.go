// Copyright 2024 Louis Royer and the NextMN-json-api contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT
package jsonapi

type Router struct {
	Locator  Locator    `json:"locator"`  // locator (ipv6 prefix)
	Backbone BackboneIP `json:"backbone"` // data plane backbone ipv6 address
	Control  ControlURI `json:"control"`  // url used for control
}
