// Copyright 2022-2025 The sacloud/simple-notification-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package simplenotification

import (
	"fmt"
	"runtime"

	"github.com/sacloud/saclient-go"
	v1 "github.com/sacloud/simple-notification-api-go/apis/v1"
)

const DefaultAPIRootURL = "https://secure.sakura.ad.jp/cloud/zone/is1a/api/cloud/1.0/"

var UserAgent = fmt.Sprintf(
	"simple-notification-api-go/%s (%s/%s; +https://github.com/sacloud/simple-notification-api-go)",
	Version,
	runtime.GOOS,
	runtime.GOARCH,
)

func NewClient(client *saclient.Client) (*v1.Client, error) {
	return NewClientWithAPIRootURL(client, DefaultAPIRootURL)
}

func NewClientWithAPIRootURL(client *saclient.Client, apiRootURL string) (*v1.Client, error) {
	cli, err := client.DupWith(saclient.WithUserAgent(UserAgent))
	if err != nil {
		return nil, err
	}

	cc, err := newFilterInjector(apiRootURL, cli)
	if err != nil {
		return nil, err
	}
	return v1.NewClient(apiRootURL, v1.WithClient(cc))
}
