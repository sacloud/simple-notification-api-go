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
	"context"
	"fmt"
	"net/http"
	"net/url"

	ht "github.com/ogen-go/ogen/http"
	v1 "github.com/sacloud/simple-notification-api-go/apis/v1"
)

var _ ht.Client = (*filterInjector)(nil)

type filterInjector struct {
	listAPIPath string
	client      ht.Client
}

func newFilterInjector(apiURL string, client ht.Client) (ht.Client, error) {
	u, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	return &filterInjector{
		listAPIPath: u.JoinPath("/commonserviceitem").Path,
		client:      client,
	}, nil
}

func (t *filterInjector) Do(req *http.Request) (*http.Response, error) {
	// NOTE: OpenAPIで表現できないクエリの書き込みを行う
	// 同じエンドポイントに3種類のProvider.Classでフィルタしたいため、生成コードの書き換えでなくclient middlewareで対応
	// `GET /commonserviceitem?{"Filter":{"Provider.Class":"saknoticedestination"}}`
	// `GET /commonserviceitem?{"Filter":{"Provider.Class":"saknoticegroup"}}`
	// `GET /commonserviceitem?{"Filter":{"Provider.Class":"saknoticerouting"}}`.
	if req.Method == http.MethodGet && req.URL.Path == t.listAPIPath {
		pc := getFilterProviderClass(req.Context())
		req.URL.RawQuery = fmt.Sprintf(`{"Filter":{"Provider.Class":"%s"}}`, pc)
	}

	return t.client.Do(req)
}

type ctxKeyFilterProviderClass struct{}

func setFilterProviderClass(ctx context.Context, v v1.ProviderClass) context.Context {
	return context.WithValue(ctx, ctxKeyFilterProviderClass{}, v)
}

func getFilterProviderClass(ctx context.Context) v1.ProviderClass {
	return ctx.Value(ctxKeyFilterProviderClass{}).(v1.ProviderClass)
}
