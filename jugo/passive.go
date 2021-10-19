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

var PassiveData PassiveJsonStruct

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
		// Outer Layer

		client := resty.New()

		switch key {
		case "beian":
			// None key Get
			url := fmt.Sprintf("%s%s", uri, target)
			if resp, err := client.R().Get(url); err == nil {

				if resp.StatusCode() != 200 {
					continue
				}
				PassiveData.Beian = json.Marshaler(resp)
				fmt.Println(PassiveData.Beian)
				fmt.Printf("[+] %-30s %c[1;0;34m%s%c[0m\n", "Beian information obtained successfully", 0x1b, "√", 0x1b)
			} else {
				log.Fatal(err)
			}

		}

	}
	fmt.Println("Passive")

}
