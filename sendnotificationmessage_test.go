package simplenotification_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/sacloud/packages-go/testutil"
	"github.com/sacloud/saclient-go"
	simplenotification "github.com/sacloud/simple-notification-api-go"

	v1 "github.com/sacloud/simple-notification-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func apiSetup(t *testing.T) (ctx context.Context, api simplenotification.SendNotificationMessageAPI) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")(t)

	ctx = t.Context()

	var saClient saclient.Client

	client, err := simplenotification.NewClient(&saClient)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	require.NoError(t, err)
	simpleNotificationAPI := simplenotification.NewSendNotificationMessageOp(client)

	return ctx, simpleNotificationAPI
}

func TestSendNotificationMessageOp_Create(t *testing.T) {

	// テスト用の通知グループを作成
	ctx, simpleNotificationAPI := apiSetup(t)

	id := "your-notification-group-id" // 事前に作成した通知グループのIDを指定してください
	request := v1.SendNotificationMessageRequest{Message: "sacloud simple-notification-api sdk test message"}
	wantResp := &v1.SendNotificationMessageResponse{IsOk: true}

	resp, err := simpleNotificationAPI.Create(ctx, id, request)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(resp, wantResp) {
		t.Errorf("response mismatch: got %+v, want %+v", resp, wantResp)
	}
	t.Logf("SendNotificationMessageOp.Create response: %+v", resp)
}
