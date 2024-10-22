package main

import (
	"fmt"
	"time"

	"github.com/AnnonaOrg/osenv"
	"github.com/umfaka/umfaka_core/internal/constvar"
	"github.com/umfaka/umfaka_core/internal/initialize"
	"github.com/umfaka/umfaka_core/internal/log"
	"github.com/umfaka/umfaka_core/internal/service"
	"github.com/umfaka/umfaka_core/internal/tasks"
	"github.com/umfaka/umfaka_core/internal/utils"
)

func mainTask() {
	go func() {
		if err := pingServer(); err != nil {
			log.Fatalf(
				"(%s)没有响应，请检查配置及网络状态: %v",
				constvar.APPName(), err,
			)
		}
		log.Infof("(%s)成功部署，服务地址:%s", constvar.APPName(), osenv.GetServerUrl())
	}()

	if err := initialize.Init(); err != nil {
		log.Errorf("initialize.Init() : %v", err)
	}

	if osenv.GetBotTelegramWebhookURL() == "" {
		go mainBot()
	} else {
		go service.SetBotFatherWebhook()
	}

	go doTask()
}

// 自检openAPI服务是否正常运行
func pingServer() error {
	apiURL := osenv.GetServerUrl()
	for i := 0; i < 10; i++ {

		if utils.CheckPingBaseURL(apiURL) {
			return nil
		}

		log.Debugf(
			"(%s)等待自检, 1秒后重试(%d) %s",
			constvar.APPName(), i, apiURL,
		)
		time.Sleep(time.Second * 2)
	}
	return fmt.Errorf(
		"(%s)自检失败 %s.",
		constvar.APPName(), apiURL,
	)
}

func doTask() {
	tasks.Init()
}
