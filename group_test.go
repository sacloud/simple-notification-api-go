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
	"testing"

	"github.com/sacloud/packages-go/testutil"
	"github.com/sacloud/saclient-go"
	simplenotification "github.com/sacloud/simple-notification-api-go"

	v1 "github.com/sacloud/simple-notification-api-go/apis/v1"
)

func groupAPISetup(t *testing.T) (ctx context.Context, api simplenotification.GroupAPI) {
	testutil.PreCheckEnvsFunc("SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")(t)

	ctx = t.Context()

	var saClient saclient.Client

	client, err := simplenotification.NewClient(&saClient)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	api = simplenotification.NewGroupOp(client)

	return ctx, api
}

func TestGroupOp_Create(t *testing.T) {
	ctx, groupAPI := groupAPISetup(t)

	groupname := "test-group"
	description := "test-group-description"
	tags := []string{"test"}
	setting := v1.CommonServiceItemGroupSettings{
		Destinations: []string{"destination-id"},
	}

	resp, err := groupAPI.Create(ctx, groupname, description, tags, setting)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response but got nil")
	}
	t.Logf("GroupOp.Create response: %+v", resp)
}

func TestGroupOp_List(t *testing.T) {
	ctx, groupAPI := groupAPISetup(t)

	resp, err := groupAPI.List(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response but got nil")
	}
	t.Logf("GroupOp.List response: %+v", resp)
}

func TestGroupOp_Read(t *testing.T) {
	ctx, groupAPI := groupAPISetup(t)

	id := "your-group-id" // set your pre-created group ID here

	resp, err := groupAPI.Read(ctx, id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response but got nil")
	}
	t.Logf("GroupOp.Read response: %+v", resp)
}

func TestGroupOp_Update(t *testing.T) {
	ctx, groupAPI := groupAPISetup(t)

	id := "your-group-id" // set your pre-created group ID here
	groupname := "updated-group"
	description := "updated-description"
	tags := []string{"updated"}
	setting := v1.CommonServiceItemGroupSettings{
		Destinations: []string{"destination-id"},
	}

	resp, err := groupAPI.Update(ctx, id, groupname, description, tags, &setting)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response but got nil")
	}
	t.Logf("GroupOp.Update response: %+v", resp)
}

func TestGroupOp_UpdateWithoutSetting(t *testing.T) {
	ctx, groupAPI := groupAPISetup(t)

	id := "your-group-id" // set your pre-created group ID here
	groupname := "updated-group"
	description := "updated-description"
	tags := []string{"updated"}

	resp, err := groupAPI.Update(ctx, id, groupname, description, tags, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response but got nil")
	}
	t.Logf("GroupOp.UpdateWithoutSetting response: %+v", resp)
}

func TestGroupOp_Delete(t *testing.T) {
	ctx, groupAPI := groupAPISetup(t)

	id := "your-group-id" // set your pre-created group ID here

	err := groupAPI.Delete(ctx, id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Log("GroupOp.Delete succeeded")
}

func TestGroupOp_SendMessage(t *testing.T) {
	ctx, groupAPI := groupAPISetup(t)

	id := "your-group-id" // set your pre-created group ID here
	request := v1.SendNotificationMessageRequest{
		Message: "test message from GroupOp.SendMessage",
	}

	resp, err := groupAPI.SendMessage(ctx, id, request)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response but got nil")
	}
	t.Logf("GroupOp.SendMessage response: %+v", resp)
}
