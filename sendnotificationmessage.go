package simplenotification

import (
	"context"

	v1 "github.com/sacloud/simple-notification-api-go/apis/v1"
)

type SendNotificationMessageAPI interface {
	Create(ctx context.Context, id string, request v1.SendNotificationMessageRequest) (*v1.SendNotificationMessageResponse, error)
}

var _ SendNotificationMessageAPI = (*sendNotificationMessageOp)(nil)

type sendNotificationMessageOp struct {
	client *v1.Client
}

func NewSendNotificationMessageOp(client *v1.Client) SendNotificationMessageAPI {
	return &sendNotificationMessageOp{client: client}
}

func (o *sendNotificationMessageOp) Create(ctx context.Context, id string, request v1.SendNotificationMessageRequest) (*v1.SendNotificationMessageResponse, error) {
	response, err := o.client.SendNotificationMessage(ctx, v1.OptSendNotificationMessageRequest{Value: request, Set: true}, v1.SendNotificationMessageParams{ID: id})
	if err != nil {
		return nil, err
	}
	return response, nil
}
