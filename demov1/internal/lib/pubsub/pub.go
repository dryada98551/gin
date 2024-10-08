package libpubsub

import (
	"context"
	"cloud.google.com/go/pubsub"

	configs "main/internal/lib/init"
	libmodel "main/internal/lib/model"
	libreturncode "main/internal/lib/returnCode"
)

func TopicPub(ctx context.Context, data []byte, attributes map[string]string) (libmodel.SystemReturn) {
	// Create a sysReturn
	var sysReturn libmodel.SystemReturn

	result := configs.Global.PubSub.Topic.Publish(ctx, &pubsub.Message{
		Data: data,
		Attributes: attributes,
	})

	// 等待消息發布結果
	_, err := result.Get(ctx)
	if err != nil {
		sysReturn.Code = libreturncode.Cobelib.ServiceFail.Publish.Code
		sysReturn.Message = libreturncode.Cobelib.ServiceFail.Publish.Message
		return sysReturn
	}

	sysReturn.Code = libreturncode.Cobelib.Success.Publish.Code
	sysReturn.Message = libreturncode.Cobelib.Success.Publish.Message
	return sysReturn
}
