package main

import (
	"context"
	"fmt"

	"github.com/wenerme/astgo/ami"
)

func main() {
	boot := make(chan *ami.Message, 1)

	conn, err := ami.Connect(
		"127.0.0.1:5038",
		// ami.WithAuth("admin", "admin"), // AMI auth
		ami.WithAuth("callert", "Fo3shooyelizood"),
		// add predefined subscriber
		ami.WithSubscribe(ami.SubscribeFullyBootedChanOnce(boot)),
		ami.WithSubscribe(func(ctx context.Context, msg *ami.Message) bool {
			fmt.Println(msg.Format()) // log everything
			return true               // keep subscribe
		}, ami.SubscribeSend(), // subscribe send message - default recv only

		))
	if err != nil {
		panic(err)
	}
	<-boot
	// AMI now FullyBooted
	_ = conn
}
