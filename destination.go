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

type DestinationAPI interface {
	List(ctx context.Context) (*v1.ListCommonServiceItemsResponse, error)
	Create(ctx context.Context, destName, description string, tags []string,
		setting v1.CommonServiceItemDestinationSettings) (*v1.CommonServiceItem, error)
	Read(ctx context.Context, id string) (*v1.CommonServiceItem, error)
	Update(ctx context.Context, id, destName, description string, tags []string,
		optSetting *v1.CommonServiceItemDestinationSettings) (*v1.CommonServiceItem, error)
	Delete(ctx context.Context, id string) error
	GetStatus(ctx context.Context, id string) (*v1.GetCommonServiceItemStatusResponseNotificationStatus, error)
}

var _ DestinationAPI = (*DestinationOp)(nil)

type DestinationOp struct {
	client *v1.Client
}

func NewDestinationOp(client *v1.Client) DestinationAPI {
	return &DestinationOp{client: client}
}

func (o *DestinationOp) List(ctx context.Context) (*v1.ListCommonServiceItemsResponse, error) {
	const methodName = "Destination.List"
	ctx = setContextProviderClass(ctx, v1.CommonServiceItemProviderClassSaknoticedestination)
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

func (o *DestinationOp) Create(ctx context.Context, destName, description string, tags []string,
	setting v1.CommonServiceItemDestinationSettings) (*v1.CommonServiceItem, error) {
	const methodName = "Destination.Create"
	req := v1.PostCommonServiceItemRequest{
		CommonServiceItem: v1.PostCommonServiceItemRequestCommonServiceItem{
			Name:        destName,
			Description: description,
			Tags:        tags,
			Provider: v1.PostCommonServiceItemRequestCommonServiceItemProvider{
				Class: v1.PostCommonServiceItemRequestCommonServiceItemProviderClassSaknoticedestination,
			},
			Settings: v1.PostCommonServiceItemRequestCommonServiceItemSettings{
				Type:                                 v1.CommonServiceItemDestinationSettingsPostCommonServiceItemRequestCommonServiceItemSettings,
				CommonServiceItemDestinationSettings: setting,
			},
		},
	}
	res, err := o.client.CreateCommonServiceItem(ctx, v1.OptPostCommonServiceItemRequest{Value: req, Set: true})
	if err != nil {
		var e *v1.ErrorStatusCode
		if errors.As(err, &e) {
			switch e.StatusCode {
			case http.StatusCreated:
				return &v1.CommonServiceItem{}, nil
			default:
				return nil, NewAPIError(methodName, e.StatusCode, err)
			}
		}
		return nil, NewError(methodName, err)
	}
	return &res.CommonServiceItem, nil
}

func (o *DestinationOp) Read(ctx context.Context, id string) (*v1.CommonServiceItem, error) {
	const methodName = "Destination.Read"
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

func (o *DestinationOp) Update(ctx context.Context, id, destName, description string, tags []string,
	optSetting *v1.CommonServiceItemDestinationSettings) (*v1.CommonServiceItem, error) {
	const methodName = "Destination.Update"
	setting := v1.CommonServiceItemDestinationSettings{}
	set := false

	if optSetting != nil {
		setting = *optSetting
		set = true
	}
	req := v1.PutCommonServiceItemRequest{
		CommonServiceItem: v1.PutCommonServiceItemRequestCommonServiceItem{
			Name:        destName,
			Description: description,
			Tags:        tags,
			Settings: v1.OptPutCommonServiceItemRequestCommonServiceItemSettings{
				Set: set,
				Value: v1.PutCommonServiceItemRequestCommonServiceItemSettings{
					Type:                                 v1.CommonServiceItemDestinationSettingsPutCommonServiceItemRequestCommonServiceItemSettings,
					CommonServiceItemDestinationSettings: setting,
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

func (o *DestinationOp) Delete(ctx context.Context, id string) error {
	const methodName = "Destination.Delete"
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

func (o *DestinationOp) GetStatus(ctx context.Context, id string) (*v1.GetCommonServiceItemStatusResponseNotificationStatus, error) {
	const methodName = "Destination.GetStatus"
	res, err := o.client.GetCommonServiceItemStatus(ctx, v1.GetCommonServiceItemStatusParams{ID: id})
	if err != nil {
		var e *v1.ErrorStatusCode
		if errors.As(err, &e) {
			return nil, NewAPIError(methodName, e.StatusCode, err)
		}
		return nil, NewError(methodName, err)
	}
	return &res.NotificationStatus, nil
}
