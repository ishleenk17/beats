// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package panw

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "panw", asset.ModuleFieldsPri, AssetPanw); err != nil {
		panic(err)
	}
}

// AssetPanw returns asset data.
// This is the base64 encoded zlib format compressed contents of module/panw.
func AssetPanw() string {
	return "eJzUz0FOBCEQheE9p3iZfV+AnRfQSTwBDm8mZKAgUNhye9NGO21rXNtvSSVffibcOSyKk9kAGjTS4nR2Mp8M4NkuNRQNWSzOD4/T0zNS9j3SAJWRrtHiheoMcA2MvlkDABPEJa7uMh2FFreae/l8+VNfthW3ahBlvboL23r6zV+2j/zant7yfHOprBHf/TvHnKvf3XxIlPbxD62dP4Jr7hrkdpDaNpoyHST2tcg/LX0PAAD//4Q63bA="
}
