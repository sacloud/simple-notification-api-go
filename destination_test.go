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
	"github.com/stretchr/testify/require"

	v1 "github.com/sacloud/simple-notification-api-go/apis/v1"
)

func destinationAPISetup(t *testing.T) (ctx context.Context, api simplenotification.DestinationAPI) {

	ctx = t.Context()
	var saClient saclient.Client

	client, err := simplenotification.NewClient(&saClient)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	api = simplenotification.NewDestinationOp(client)

	return ctx, api
}

func TestDestinationOp(t *testing.T) {
	testutil.PreCheckEnvsFunc("SAKURA_ACCESS_TOKEN", "SAKURA_ACCESS_TOKEN_SECRET", "SAKURA_DESTINATION_TEST_MAILADDRESS")(t)
	ctx, destinationAPI := destinationAPISetup(t)
	id := ""
	destName := "test-destination-1"
	description := "test-destination-description"
	tags := []string{"test"}
	mailAddress := os.Getenv("SAKURA_DESTINATION_TEST_MAILADDRESS")
	setting := v1.CommonServiceItemDestinationSettings{
		Type:  v1.CommonServiceItemDestinationSettingsTypeEmail,
		Value: mailAddress,
	}
	result := t.Run("Create", func(t *testing.T) {
		resp, err := destinationAPI.Create(ctx, destName, description, tags, setting)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		t.Logf("DestinationOp.Create response: %+v", resp)
		id = resp.ID
	})
	defer func() {
		if id != "" {
			err := destinationAPI.Delete(ctx, id)
			if err != nil {
				t.Errorf("unexpected error on cleanup: %v", err)
			}
			t.Log("DestinationOp.Delete succeeded on cleanup")
		}
	}()

	if !result {
		t.Fatal("skipping rest of tests due to Create failure")
	}

	t.Run("List", func(t *testing.T) {
		resp, err := destinationAPI.List(ctx)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		t.Logf("DestinationOp.List response: %+v", resp)
	})
	t.Run("Read", func(t *testing.T) {
		resp, err := destinationAPI.Read(ctx, id)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		t.Logf("DestinationOp.Read response: %+v", resp)
	})
	t.Run("Update", func(t *testing.T) {
		destNameUpdate := "updated-destination"
		descriptionUpdate := "updated-description"
		tagsUpdate := []string{"updated"}
		resp, err := destinationAPI.Update(ctx, id, destNameUpdate, descriptionUpdate, tagsUpdate, &setting)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		require.Equal(t, destNameUpdate, resp.Name, "destination name should be updated")
		require.Equal(t, descriptionUpdate, resp.Description, "description should be updated")
		require.Equal(t, tagsUpdate, resp.Tags, "tags should be updated")
		t.Logf("DestinationOp.Update response: %+v", resp)
	})
	t.Run("UpdateWithoutSetting", func(t *testing.T) {
		destNameUpdateWithoutsetting := "updated-destination-withoutsetting"
		descriptionUpdateWithoutsetting := "updated-description-withoutsetting"
		tagsUpdateWithoutsetting := []string{"updated-withoutsetting"}
		resp, err := destinationAPI.Update(ctx, id, destNameUpdateWithoutsetting, descriptionUpdateWithoutsetting, tagsUpdateWithoutsetting, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		require.Equal(t, destNameUpdateWithoutsetting, resp.Name, "destination name should be updated")
		require.Equal(t, descriptionUpdateWithoutsetting, resp.Description, "description should be updated")
		require.Equal(t, tagsUpdateWithoutsetting, resp.Tags, "tags should be updated")
		t.Logf("DestinationOp.UpdateWithoutSetting response: %+v", resp)
	})
	t.Run("GetStatus", func(t *testing.T) {
		resp, err := destinationAPI.GetStatus(ctx, id)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp == nil {
			t.Fatal("expected response but got nil")
		}
		t.Logf("DestinationOp.GetStatus response: %+v", resp)
	})
	t.Run("Delete", func(t *testing.T) {
		err := destinationAPI.Delete(ctx, id)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		t.Log("DestinationOp.Delete succeeded")
		id = "" // prevent double delete in defer
	})
}
