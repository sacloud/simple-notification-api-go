// Copyright 2025- The sacloud/simple-notification-api-go Authors
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

package simplenotification_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/sacloud/packages-go/testutil"
	"github.com/sacloud/saclient-go"
	simplenotification "github.com/sacloud/simple-notification-api-go"

	v1 "github.com/sacloud/simple-notification-api-go/apis/v1"
)

func SendNotificationMessageAPISetup(t *testing.T) (ctx context.Context, api simplenotification.SendNotificationMessageAPI) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")(t)

	ctx = t.Context()

	var saClient saclient.Client

	client, err := simplenotification.NewClient(&saClient)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	API := simplenotification.NewSendNotificationMessageOp(client)

	return ctx, API
}

func TestSendNotificationMessageOp_Create(t *testing.T) {

	// テスト用の通知グループを作成
	ctx, sendNotificationMessageAPI := SendNotificationMessageAPISetup(t)

	id := "your-notification-group-id" // 事前に作成した通知グループのIDを指定してください
	request := v1.SendNotificationMessageRequest{Message: "sacloud simple-notification-api sdk test message"}
	wantResp := &v1.SendNotificationMessageResponse{IsOk: true}

	resp, err := sendNotificationMessageAPI.Send(ctx, id, request)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(resp, wantResp) {
		t.Errorf("response mismatch: got %+v, want %+v", resp, wantResp)
	}
	t.Logf("SendNotificationMessageOp.Send response: %+v", resp)
}
