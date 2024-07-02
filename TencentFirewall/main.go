package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func main() {
	// 创建一个Ticker，每隔30分钟触发一次
	ticker := time.NewTicker(30 * time.Minute)

	go func() {
		for {
			select {
			case <-ticker.C:
				ip, err := getIp()
				if err != nil {
					log.Fatal(err)
				}
				result, err := addIp(ip)
				if err != nil {
					fmt.Println("error:", err.Error())
				}
				fmt.Println(result)
				accessToken := "94a1a00f04343c3b74032c413c857a1d9b26b7aa55cfbb0aa649857ca27bfc28"
				secret := ""
				client := dingtalk.NewClient(accessToken, secret)

				msg := dingtalk.NewTextMessage().SetContent(fmt.Sprintf("Tencent:%s", result))

				if strings.Contains(result, "防火墙规则已经存在") {
					continue
				} else {
					client.Send(msg)
				}
			}
		}

	}()
	// 保持主程序运行，防止提前退出
	fmt.Println("定时任务已启动，按Ctrl+C退出...")
	select {}
}

type Params struct {
	InstanceId    string         `json:"InstanceId"`
	FirewallRules []FirewallRule `json:"FirewallRules"`
}
type FirewallRule struct {
	Protocol                string `json:"Protocol"`
	Port                    string `json:"Port"`
	CidrBlock               string `json:"CidrBlock"`
	Action                  string `json:"Action"`
	FirewallRuleDescription string `json:"FirewallRuleDescription"`
}

func addIp(ipAddress string) (string, error) {
	credential := common.NewCredential("AKID48Qdn0BqV5g8cOIiFlEnDiM7jFsbN4na", "2wlJzngXQEwLnTECDA119JXuTxJZq2BD")
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	cpf.HttpProfile.ReqMethod = "POST"
	client := common.NewCommonClient(credential, "ap-shanghai", cpf).WithLogger(log.Default())

	request := tchttp.NewCommonRequest("lighthouse", "2020-03-24", "CreateFirewallRules")

	firewallRule := FirewallRule{
		Protocol:                "TCP",
		Port:                    "ALL",
		CidrBlock:               ipAddress,
		Action:                  "ACCEPT",
		FirewallRuleDescription: "",
	}
	params := Params{
		InstanceId:    "lhins-rcaxix1i",
		FirewallRules: []FirewallRule{firewallRule},
	}
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return err.Error(), nil
	}
	fmt.Println("jsonParam", string(jsonParams))
	err = request.SetActionParameters(jsonParams)
	if err != nil {
		return err.Error(), nil
	}

	response := tchttp.NewCommonResponse()
	err = client.Send(request, response)
	if err != nil {
		if strings.Contains(err.Error(), "防火墙规则已经存在") {
			return fmt.Sprintf("防火墙规则已经存在%s", ipAddress), nil
		} else {
			fmt.Println("fail to invoke api:", err.Error())
		}
	}
	return string(response.GetBody()), nil

}

func getIp() (string, error) {
	cmd := exec.Command("curl", "4.ipw.cn")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	output := out.String()
	return output, nil
}
