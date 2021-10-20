package jugo

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type PassiveJsonStruct struct {
	Beian interface{} `json:beian`
}

var (
	PassiveData PassiveJsonStruct
	json_data   map[string]interface{}
)

func init() {
	// 代理探测

}

func Passive(target string) {

	// Cli Output
	fmt.Printf("[*] %-30s\n", "Start performing passive information collection...")
	defer fmt.Printf("[*] %-30s\n", "Passive mode information collection is complete...")
	fmt.Printf("[*] %-30s: %c[1;4;40;32m%s%c[0m\n", "The purpose of this mission is", 0x1b, target, 0x1b)
	apikey_map := map[string]string{
		// {Apiurl : Toml File Node}
		"beian": "https://tapi.66sj.cn/api/url_icp?url=",
	}
	for key, uri := range apikey_map {

		switch key {
		case "beian":
			// Start
			url := fmt.Sprintf("%s%s", uri, target)
			client := resty.New()
			if resp, err := client.R().SetHeader("Accept", "application/json").Get(url); err == nil {
				// 正常请求
				if resp.StatusCode() == 200 {
					// 正常返回

					if err1 := json.Unmarshal([]byte(resp.Body()), &json_data); err1 != nil {
						log.Fatalf("unmarshal err = %v\n", err)

					}

					PassiveData.Beian = json_data["data"]
					fmt.Printf("[+] %-30s %c[1;0;34m%s%c[0m\n", "Beian information obtained successfully", 0x1b, "√", 0x1b)

					fmt.Println(PassiveData.Beian)
				} else {
					// 异常返回
					log.Fatalf("status code error: %d %s", resp.StatusCode(), resp.Status())
					fmt.Printf("[+] %-30s %c[1;0;31m%s%c[0m\n", "Beian information api interface returns an exception", 0x1b, "x", 0x1b)

				}
			} else {
				// 异常请求
				log.Fatal(err)
			}
			// End
		}

	}
	fmt.Println("Passive")
}
