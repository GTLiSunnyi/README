package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	consulapi "github.com/hashicorp/consul/api"

	"go-kit-demo/go-kit-client/service"
	"go-kit-demo/go-kit-server/types"
)

// 只需要写 EncodeUserRequest 和 DecodeUserResponse，其他基本不用修改
func main() {
	serviceName := flag.String("name", "", "服务名")
	flag.Parse()

	if *serviceName == "" {
		log.Panicln("请指定服务名!")
	}

	// 1. 创建 consul client
	config := consulapi.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", types.IP, types.CONSUL_PROT)
	api_client, _ := consulapi.NewClient(config)
	client := consul.NewClient(api_client)

	logger := kitlog.NewLogfmtLogger(os.Stdout)

	// 2. 创建 consul 的实例
	// 最后的 true 表示只有通过健康检查的服务才能被得到
	instancer := consul.NewInstancer(client, logger, *serviceName, []string{*serviceName}, true)

	// factory 定义了如何获得服务端的 endpoint, 这里的 service_url 是从 consul 中读取到的 service 的 address
	factory := func(service_url string) (endpoint.Endpoint, io.Closer, error) {
		target, _ := url.Parse("http://" + service_url)
		return httptransport.NewClient("GET", target, service.EncodeUserRequest, service.DecodeUserResponse).Endpoint(), nil, nil
	}

	endpointer := sd.NewEndpointer(instancer, factory, logger)

	// 负载均衡
	// 1. 轮询方式
	// mylb := lb.NewRoundRobin(endpointer)
	// 2. 随机方式
	mylb := lb.NewRandom(endpointer, time.Now().Unix())
	for {
		// 3. 执行
		getName, err := mylb.Endpoint()
		if err != nil {
			log.Panicln(err.Error())
		}
		res, err := getName(context.Background(), service.UserRequest{Uid: 101})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 4. 断言，得到响应值
		userinfo := res.(service.UserResponse)
		fmt.Println(userinfo.Result)

		time.Sleep(3 * time.Second)
	}
}
