package main

import (
	"net"

	"HuaTug.com/config/cache"
	"HuaTug.com/cmd/publish/dal"
	publish "HuaTug.com/kitex_gen/publishs/uploadvideoservice"
	"HuaTug.com/pkg/bound"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	//trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	//tracer2.InitJaeger(constants.UserServiceName)
	dal.Init()
}

func main() {
	//r, err := etcd.NewEtcdRegistry([]string{config.ConfigInfo.Etcd.Addr})
	r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})
	if err != nil {
		panic(err)
	}
	ip, err := constants.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ip+":8890")
	if err != nil {
		panic(err)
	}
	Init()
	cache.Init()
	//当出现了UserServiceImpl报错时 说明当前该接口的方法没有被完全实现

	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Publish"}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                             // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		//server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()), // BoundHandler
		server.WithRegistry(r),                              // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
