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

package simplenotification

import (
	"context"
	"errors"

	v1 "github.com/sacloud/simple-notification-api-go/apis/v1"
)

type SendNotificationMessageAPI interface {
	Send(ctx context.Context, id string, request v1.SendNotificationMessageRequest) (*v1.SendNotificationMessageResponse, error)
}

var _ SendNotificationMessageAPI = (*sendNotificationMessageOp)(nil)

type sendNotificationMessageOp struct {
	client *v1.Client
}

func NewSendNotificationMessageOp(client *v1.Client) SendNotificationMessageAPI {
	return &sendNotificationMessageOp{client: client}
}

func (o *sendNotificationMessageOp) Send(ctx context.Context, id string, request v1.SendNotificationMessageRequest) (*v1.SendNotificationMessageResponse, error) {
	const methodName = "SendNotificationMessageAPI.Send"
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
