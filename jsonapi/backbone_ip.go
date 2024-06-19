// Copyright 2024 Louis Royer and the NextMN-json-api contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT
package jsonapi

import (
	"fmt"
	"net/netip"
)

type BackboneIP struct {
	netip.Addr
}

func (b *BackboneIP) UnmarshalText(text []byte) error {
	if err := b.Addr.UnmarshalText(text); err != nil {
		return err
	}
	if !b.Addr.Is6() {
		return fmt.Errorf("Backbone IP must be an IPv6 address")
	}
	return nil
}
