package test_service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	v1 "test_server/gen/go/test_server_interface/v1/test"

	"github.com/bufbuild/connect-go"
)

type TestService struct {
	NotificationChannel chan string
	IsSubscribed        bool
}

func NewTestService() *TestService {
	return &TestService{
		NotificationChannel: make(chan string, 20),
		IsSubscribed:        false,
	}
}

func (s *TestService) Subscribe(
	ctx context.Context,
	req *connect.Request[v1.SubscribeRequest],
	stream *connect.ServerStream[v1.Notification],
) error {
	log.Println("Test Service - start Subscribe")
	defer log.Println("Test Service - end Subscribe")

	if s.IsSubscribed {
		return errors.New("already subscribed")
	}

	go s.PushRoutine()
	s.IsSubscribed = true

	for {
		notif := <-s.NotificationChannel
		stream.Send(&v1.Notification{
			Content: notif,
		})
	}
}

func (s *TestService) GetTestData(
	ctx context.Context,
	req *connect.Request[v1.GetTestDataRequest],
) (*connect.Response[v1.TestData], error) {
	log.Println("Test Service - start GetTestData")
	defer log.Println("Test Service - end GetTestData")

	log.Printf("%s", req.Msg.Content)

	return connect.NewResponse(&v1.TestData{
		Content: fmt.Sprintf("hoge-%s", req.Msg.Content),
	}), nil
}

// 延々とプッシュ通知しまくる。
func (s *TestService) PushRoutine() {
	for {
		log.Println("sending notifications...")
		s.NotificationChannel <- "りんご"
		time.Sleep(300 * time.Millisecond)
		s.NotificationChannel <- "ゴリラ"
		// s.NotificationChannel <- "ラッパ"
		// s.NotificationChannel <- "パンツ"
		// s.NotificationChannel <- "ツバル"
		// s.NotificationChannel <- "ルルーシュ"
		// s.NotificationChannel <- "ゆべし"
		// s.NotificationChannel <- "しおがま"
		// s.NotificationChannel <- "まろうど"
		// s.NotificationChannel <- "ドグラマグラ"
		log.Println("end sending")

		time.Sleep(1000 * time.Millisecond)
	}
}
