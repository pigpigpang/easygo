package internal

import (
	"context"
	"easygo/internal/common/config"
	"easygo/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type options struct {
	Mode    string
	Version string
}

// Option 定义配置项
type Option func(*options)

// SetMode 设定运行模式
func SetMode(s string) Option {
	return func(o *options) {
		if s == "" {
			o.Mode = "dev"
			return
		}
		o.Mode = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

// Init 初始化应用
func Init(ctx context.Context, opts ...Option) (func(), error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	// 初始化配置
	config.InitConfig(o.Mode)

	// 初始化日志
	loggerCleanFunc := InitLogger()

	// 初始化依赖注入器
	injector, injectorCleanFunc, err := BuildInjector()
	if err != nil {
		return nil, err
	}

	// 初始化HTTP服务
	httpServerCleanFunc := InitHTTPServer(ctx, injector.Engine)

	logger.WithContext(ctx).Infof("服务启动，运行模式：%s，版本号：%s，进程号：%d", config.C.Mode, o.Version, os.Getpid())

	return func() {
		httpServerCleanFunc()
		injectorCleanFunc()
		loggerCleanFunc()
	}, nil
}

// Run 运行服务
func Run(ctx context.Context, opts ...Option) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx, opts...)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("接收到信号[%v]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	logger.WithContext(ctx).Infof("服务退出")
	cleanFunc()
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
