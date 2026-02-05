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

package simplenotification

import (
	"context"
	"errors"
	"net/http"

	v1 "github.com/sacloud/simple-notification-api-go/apis/v1"
)

type GroupAPI interface {
	List(ctx context.Context) (*v1.ListCommonServiceItemsResponse, error)
	Create(ctx context.Context, groupname, description string, tags []string,
		setting v1.CommonServiceItemGroupSettings) (*v1.CommonServiceItem, error)
	Read(ctx context.Context, id string) (*v1.CommonServiceItem, error)
	Update(ctx context.Context, id, groupname, description string, tags []string,
		optSetting *v1.CommonServiceItemGroupSettings) (*v1.CommonServiceItem, error)
	Delete(ctx context.Context, id string) error
	SendMessage(ctx context.Context, id string,
		request v1.SendNotificationMessageRequest) (*v1.SendNotificationMessageResponse, error)
}

var _ GroupAPI = (*GroupOp)(nil)

type GroupOp struct {
	client *v1.Client
}

func NewGroupOp(client *v1.Client) GroupAPI {
	return &GroupOp{client: client}
}

func (o *GroupOp) List(ctx context.Context) (*v1.ListCommonServiceItemsResponse, error) {
	const methodName = "Group.List"

	ctx = setContextProviderClass(ctx, v1.CommonServiceItemProviderClassSaknoticegroup)
	res, err := o.client.ListCommonServiceItems(ctx)
	if err != nil {
		var e *v1.ErrorStatusCode
		if errors.As(err, &e) {
			return nil, NewAPIError(methodName, e.StatusCode, err)
		}
		return nil, NewError(methodName, err)
	}
	return res, nil
}

func (o *GroupOp) Create(ctx context.Context, groupname, description string, tags []string,
	setting v1.CommonServiceItemGroupSettings) (*v1.CommonServiceItem, error) {
	const methodName = "Group.Create"
	req := v1.PostCommonServiceItemRequest{
		CommonServiceItem: v1.PostCommonServiceItemRequestCommonServiceItem{
			Name:        groupname,
			Description: description,
			Tags:        tags,
			ServiceClass: v1.OptString{
				Value: "cloud/saknoticegroup/2",
				Set:   true,
			},
			Provider: v1.PostCommonServiceItemRequestCommonServiceItemProvider{
				Class:        v1.PostCommonServiceItemRequestCommonServiceItemProviderClassSaknoticegroup,
				ServiceClass: v1.OptString{Value: "cloud/saknotice", Set: true},
			},
			Settings: v1.PostCommonServiceItemRequestCommonServiceItemSettings{
				Type:                           v1.CommonServiceItemGroupSettingsPostCommonServiceItemRequestCommonServiceItemSettings,
				CommonServiceItemGroupSettings: setting,
			},
		},
	}
	res, err := o.client.CreateCommonServiceItem(ctx, v1.OptPostCommonServiceItemRequest{Value: req, Set: true})
	if err != nil {
		var e *v1.ErrorStatusCode
		if errors.As(err, &e) {
			switch e.StatusCode {
			case http.StatusCreated:
				return &res.CommonServiceItem, nil
			default:
				return nil, NewAPIError(methodName, e.StatusCode, err)
			}
		}
		return nil, NewError(methodName, err)
	}
	return &res.CommonServiceItem, nil
}

func (o *GroupOp) Read(ctx context.Context, id string) (*v1.CommonServiceItem, error) {
	const methodName = "Group.Read"
	res, err := o.client.GetCommonServiceItem(ctx, v1.GetCommonServiceItemParams{ID: id})
	if err != nil {
		var e *v1.ErrorStatusCode
		if errors.As(err, &e) {
			return nil, NewAPIError(methodName, e.StatusCode, err)
		}
		return nil, NewError(methodName, err)
	}
	return &res.CommonServiceItem, nil
}

func (o *GroupOp) Update(ctx context.Context, id, groupname, description string, tags []string,
	optSetting *v1.CommonServiceItemGroupSettings) (*v1.CommonServiceItem, error) {
	const methodName = "Group.Update"
	setting := v1.CommonServiceItemGroupSettings{}
	set := false

	if optSetting != nil {
		setting = *optSetting
		set = true
	}
	req := v1.PutCommonServiceItemRequest{
		CommonServiceItem: v1.PutCommonServiceItemRequestCommonServiceItem{
			Name:        groupname,
			Description: description,
			Tags:        tags,
			Settings: v1.OptPutCommonServiceItemRequestCommonServiceItemSettings{
				Set: set,
				Value: v1.PutCommonServiceItemRequestCommonServiceItemSettings{
					Type:                           v1.CommonServiceItemGroupSettingsPutCommonServiceItemRequestCommonServiceItemSettings,
					CommonServiceItemGroupSettings: setting,
				},
			},
		},
	}
	res, err := o.client.UpdateCommonServiceItem(ctx, v1.OptPutCommonServiceItemRequest{Value: req, Set: true}, v1.UpdateCommonServiceItemParams{ID: id})
	if err != nil {
		var e *v1.ErrorStatusCode
		if errors.As(err, &e) {
			return nil, NewAPIError(methodName, e.StatusCode, err)
		}
		return nil, NewError(methodName, err)
	}
	return &res.CommonServiceItem, nil
}

func (o *GroupOp) Delete(ctx context.Context, id string) error {
	const methodName = "Group.Delete"
	_, err := o.client.DeleteCommonServiceItem(ctx, v1.DeleteCommonServiceItemParams{ID: id})
	if err != nil {
		var e *v1.ErrorStatusCode
		if errors.As(err, &e) {
			return NewAPIError(methodName, e.StatusCode, err)
		}
		return NewError(methodName, err)
	}
	return nil
}

func (o *GroupOp) SendMessage(ctx context.Context, id string, request v1.SendNotificationMessageRequest) (*v1.SendNotificationMessageResponse, error) {
	const methodName = "Group.SendMessage"
	res, err := o.client.SendNotificationMessage(ctx, v1.OptSendNotificationMessageRequest{Value: request, Set: true}, v1.SendNotificationMessageParams{ID: id})
	if err != nil {
		var e *v1.ErrorStatusCode
		if errors.As(err, &e) {
			return nil, NewAPIError(methodName, e.StatusCode, err)
		}
		return nil, NewError(methodName, err)
	}
	return res, nil
}
