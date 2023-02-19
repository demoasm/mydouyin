package testm

import (
	"context"
	"fmt"
	"mydouyin/kitex_gen/message"
	"mydouyin/kitex_gen/message/messageservice"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/mw"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var messageClient messageservice.Client

func initFavorite() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := messageservice.NewClient(
		consts.MessageServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

// func TestCreateMessage(t *testing.T) {
// 	initFavorite()
// 	resp, err := messageClient.CreateMessage(context.Background(), &message.CreateMessageRequest{
// 		FromUserId: 1,
// 		ToUserId:   2,
// 		Content:    "1111",
// 	})
// 	fmt.Println(resp, err)
// }

// func TestCreateMessage(t *testing.T) {
// 	initFavorite()
// 	resp, err := messageClient.GetFirstMessage(context.Background(), &message.GetFirstMessageRequest{
// 		Id:        1,
// 		FriendIds: []int64{1, 2, 3, 4},
// 	})
// 	fmt.Println(resp.FirstMessageList, err)
// }

func TestCreateMessage(t *testing.T) {
	initFavorite()
	resp, err := messageClient.GetMessageList(context.Background(), &message.GetMessageListRequest{
		FromUserId: 1,
		ToUserId:   2,
		PreMsgTime: 0,
	})
	fmt.Println(resp, err)
}
