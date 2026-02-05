// Copyright 2026- The sacloud/simple-notification-api-go Authors
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
	"os"
	"testing"

	"github.com/sacloud/packages-go/testutil"
	"github.com/sacloud/saclient-go"
	simplenotification "github.com/sacloud/simple-notification-api-go"

	v1 "github.com/sacloud/simple-notification-api-go/apis/v1"
)

func groupAPISetup(t *testing.T) (ctx context.Context, api simplenotification.GroupAPI) {
	testutil.PreCheckEnvsFunc("SAKURA_ACCESS_TOKEN", "SAKURA_ACCESS_TOKEN_SECRET", "DESTINATION_TEST_ID")(t)

	ctx = t.Context()

	var saClient saclient.Client

	client, err := simplenotification.NewClient(&saClient)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	api = simplenotification.NewGroupOp(client)

	return ctx, api
}

func TestGroupOp(t *testing.T) {
	ctx, groupAPI := groupAPISetup(t)

	id := "" // set your pre-created group ID here
	groupname := "test-group-1"
	description := "test-group-description"
	tags := []string{"test"}
	destinationID := os.Getenv("DESTINATION_TEST_ID")
	setting := v1.CommonServiceItemGroupSettings{
		Destinations: []string{destinationID},
	}

	result := t.Run("Create", func(t *testing.T) {
		resp, err := groupAPI.Create(ctx, groupname, description, tags, setting)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		t.Logf("GroupOp.Create response: %+v", resp)
		id = resp.ID
	})
	defer func() {
		if id != "" {
			err := groupAPI.Delete(ctx, id)
			if err != nil {
				t.Fatalf("unexpected error on cleanup: %v", err)
			}
			t.Log("GroupOp.Delete succeeded on cleanup")
		}
	}()

	if !result {
		t.Fatal("skipping rest of tests due to Create failure")
	}

	t.Run("List", func(t *testing.T) {
		resp, err := groupAPI.List(ctx)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		t.Logf("GroupOp.List response: %+v", resp)
	})
	t.Run("Read", func(t *testing.T) {
		resp, err := groupAPI.Read(ctx, id)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		t.Logf("GroupOp.Read response: %+v", resp)
	})
	t.Run("Update", func(t *testing.T) {
		groupnameUpdate := "updated-group"
		descriptionUpdate := "updated-description"
		tagsUpdate := []string{"updated"}
		resp, err := groupAPI.Update(ctx, id, groupnameUpdate, descriptionUpdate, tagsUpdate, &setting)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		t.Logf("GroupOp.Update response: %+v", resp)
	})
	t.Run("UpdateWithoutSetting", func(t *testing.T) {
		groupnameUpdateWithoutsetting := "updated-group-withoutsetting"
		descriptionUpdateWithoutsetting := "updated-description-withoutsetting"
		tagsUpdateWithoutsetting := []string{"updated-withoutsetting"}
		resp, err := groupAPI.Update(ctx, id, groupnameUpdateWithoutsetting, descriptionUpdateWithoutsetting, tagsUpdateWithoutsetting, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		t.Logf("GroupOp.UpdateWithoutSetting response: %+v", resp)
	})
	t.Run("SendMessage", func(t *testing.T) {
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
	})
	t.Run("Delete", func(t *testing.T) {
		err := groupAPI.Delete(ctx, id)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		t.Log("GroupOp.Delete succeeded")
		id = "" // prevent double delete in defer
	})
}
