// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package fhttp

import (
	"github.com/synnaxlabs/freighter"
	"github.com/synnaxlabs/x/httputil"
)

type BindableTransport interface {
	BindTo(app *fiber.App)
}

var streamReporter = freighter.Reporter{
	Protocol:  "websocket",
	Encodings: httputil.SupportedContentTypes(),
}

var unaryReporter = freighter.Reporter{
	Protocol:  "http",
	Encodings: httputil.SupportedContentTypes(),
}
