package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AnnonaOrg/osenv"
	"github.com/gin-gonic/gin"
	"github.com/umfaka/umfaka_core/internal/constvar"
	"github.com/umfaka/umfaka_core/internal/db"
	_ "github.com/umfaka/umfaka_core/internal/dotenv"
	"github.com/umfaka/umfaka_core/internal/log"
	"github.com/umfaka/umfaka_core/router"
	"github.com/umfaka/umfaka_core/router/middleware"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("run time panic:%v\n", err)
		}
	}()

	fmt.Printf("%s %s\n%s\n",
		constvar.APPName(), constvar.APPVersion(), constvar.APPDesc(),
	)
	runAt := "运行在"
	if osenv.IsInDocker() {
		runAt = runAt + "(Docker)"
	}
	runAt = runAt + ": " + osenv.Getwd()
	fmt.Println(runAt)
	time.Sleep(time.Second * 3)

	// 数据库连接初始化
	if err := db.Init(); err != nil {
		log.Fatalf(
			"数据库( %s )连接初始化出错: %v",
			osenv.GetServerDbType(), err,
		)
	}
	defer db.Close()

	// Set gin mode.
	ginMode := osenv.GetServerGinRunmode()
	gin.SetMode(ginMode)
	//Create the Gin engine.
	g := gin.New()
	//Routes.
	router.Load(
		g,
		middleware.Logging(),
		middleware.RequestId(),
	)

	go mainTask()

	addr := ":" + osenv.GetServerPort()
	if err := http.ListenAndServe(addr, g); err != nil {
		log.Errorf(
			"(%s)出错了，需要重启: %v",
			constvar.APPName(), err,
		)
	}
}
