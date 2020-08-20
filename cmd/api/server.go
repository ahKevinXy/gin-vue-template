package api

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"learn-go/conf"

	"learn-go/router"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"learn-go/database"
)

var (
	configYml string
	port      string
	mode      string
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "go-admin server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

var echoTimes int

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func setup() {

	//1. 读取配置
	conf.Setup(configYml)
	//2. 设置日志
	//3. 初始化数据库链接
	database.Setup(conf.DatabaseConfig.Driver)

}

func run() error {
	if viper.GetString("application.mode") == "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.InitRouter()
	defer database.Eloquent.Close()

	srv := &http.Server{
		Addr:    conf.ApplicationConfig.Host + ":" + conf.ApplicationConfig.Port,
		Handler: r,
	}

	go func() {
		// 服务连接

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)

		}
	}()

	tip()

	fmt.Printf("-  Local:   http://localhost:%s/ \r\n", conf.ApplicationConfig.Port)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}
	fmt.Println("Server exiting")

	return nil
}

func tip() {
	usageStr := `欢迎使用 ` + ` 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}
