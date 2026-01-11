// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package jsonapi_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/nextmn/json-api/jsonapi"
)

func TestMessageWithError(t *testing.T) {
	const (
		msg_test     = "Example of message"
		err_test_msg = "example of error"
	)
	err_test := errors.New(err_test_msg)
	j2, err := json.Marshal(map[string]string{
		"message": msg_test,
		"error":   err_test_msg,
	})
	if err != nil {
		t.Fatalf("Could not marshal map to json")
	}

	u := &jsonapi.MessageWithError{
		Message: msg_test,
		Error:   err_test,
	}

	j1, err := json.Marshal(u)
	if err != nil {
		t.Error("Could not marshal MessageWithError to json")
	}

	if !bytes.Equal(j1, j2) {
		t.Errorf("Result of marshaling MessageWithError to json is incorrect")
	}

	unm := &jsonapi.MessageWithError{}
	if err := unm.UnmarshalJSON(j1); err != nil {
		t.Errorf("Unmarshal of MessageWithError failed: %s", err)
	}
}
