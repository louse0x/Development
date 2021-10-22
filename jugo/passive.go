package jugo

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type PassiveJsonStruct struct {
	Beian     interface{} `json:"beian"`
	Cdn       interface{} `json:"cdn_1"`
	MultiNode interface{} `json:"cdn_2"`
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
	fmt.Printf("%c[1;0;32m%s%c[0m %-30s\n", 0x1b, "[*]", 0x1b, "Start performing passive information collection...")
	defer fmt.Printf("%c[1;0;32m%s%c[0m %-30s\n", 0x1b, "[*]", 0x1b, "Passive mode information collection is complete...")
	fmt.Printf("%c[1;0;32m%s%c[0m %s %c[1;4;40;32m%s%c[0m\n", 0x1b, "[*]", 0x1b, "The purpose of this mission is:", 0x1b, target, 0x1b)
	apikey_map := map[string]string{
		// {Apiurl : Toml File Node}
		"Beian":     "https://tapi.66sj.cn/api/url_icp?url=",
		"Cdn":       "http://tools.bugscaner.com/api/whichcdn/",
		"MultiNode": "https://myssl.com/api/v1/tools/cdn_check?domain=",
	}
	for key, uri := range apikey_map {
		// Start switch
		switch key {
		case "Beian":
			// Start Beian
			url := fmt.Sprintf("%s%s", uri, target)
			client := resty.New()
			if resp, err := client.R().SetHeaders(
				map[string]string{
					"User-Agent": RandomHeader(),
					"Accept":     "application/json",
				}).Get(url); err == nil {
				// 正常请求
				if resp.StatusCode() == 200 {
					// 正常返回
					if err1 := json.Unmarshal([]byte(resp.Body()), &json_data); err1 != nil {
						log.Fatalf("unmarshal err = %v\n", err)
					}
					PassiveData.Beian = json_data["data"]
					fmt.Printf("%c[1;0;34m%s%c[0m %-30s\n", 0x1b, "[√]", 0x1b, "Beian information OK!")
					// fmt.Println(PassiveData.Beian)
				} else {
					// 异常返回
					log.Fatalf("Beian status code error: %d %s", resp.StatusCode(), resp.Status())
					fmt.Printf("%c[1;0;31m%s%c[0m %-30s\n", 0x1b, "[x]", 0x1b, "Beian information Error!")
				}
			} else {
				// 异常请求
				log.Fatal(err)
			}
			// End Beian

		case "Cdn":
			// Start Cdn
			client := resty.New()
			if resp, err := client.R().SetHeaders(
				map[string]string{
					"User-Agent":   RandomHeader(),
					"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
					"Cookie":       "Hm_lvt_81c5c6d2c74d56c9ab654aec4c11e078=1634868138; __TOKEN_STR__EXIAOHUI=RyYf66W75CQKMBzn4kx3ZrK22FpDSJPm; __QRCODE_WEICHAT_URL__EXIAOHUI=https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=gQHe7zwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAyWnNuaGtsNlljeEYxR0dFcE54YzgAAgSqG3JhAwQAjScA; __QRCODE_SHOW_TOKEN__EXIAOHUI=RyYf66W75CQKMBzn4kx3ZrK22FpDSJPm; __LEAVE_TIME__EXIAOHUI=123; Hm_lpvt_81c5c6d2c74d56c9ab654aec4c11e078=1634870006",
				}).SetFormData(map[string]string{"url": target}).Post(uri); err == nil {
				// 正常请求
				if resp.StatusCode() == 200 {
					// 正常返回
					if err1 := json.Unmarshal([]byte(resp.Body()), &json_data); err1 != nil {
						log.Fatalf("unmarshal err = %v\n", err)
					}
					PassiveData.Cdn = json_data["info"]
					fmt.Printf("%c[1;0;34m%s%c[0m %-30s\n", 0x1b, "[√]", 0x1b, "Cdn information OK!")
					//fmt.Println(PassiveData.Cdn_1)
				} else {
					// 异常返回
					log.Fatalf("CDN status code error: %d %s", resp.StatusCode(), resp.Status())
					fmt.Printf("%c[1;0;31m%s%c[0m %-30s\n", 0x1b, "[x]", 0x1b, "Cdn information Error!")
				}
			} else {
				// 异常请求
				log.Fatal(err)
			}
			// End Cdn

		case "MultiNode":
			// Start MultiNode
			url := fmt.Sprintf("%s%s", uri, target)
			client := resty.New()
			if resp, err := client.R().SetHeaders(
				map[string]string{
					"User-Agent": RandomHeader(),
				}).Get(url); err == nil {
				// 正常请求
				if resp.StatusCode() == 200 {
					// 正常返回
					if err1 := json.Unmarshal([]byte(resp.Body()), &json_data); err1 != nil {
						log.Fatalf("unmarshal err = %v\n", err)
					}
					PassiveData.MultiNode = json_data["data"]
					// fmt.Println(resp)
					fmt.Printf("%c[1;0;34m%s%c[0m %-30s\n", 0x1b, "[√]", 0x1b, "MultiNode information OK!")
					// fmt.Println(PassiveData.MultiNode)
				} else {
					// 异常返回
					log.Fatalf("MultiNode status code error: %d %s", resp.StatusCode(), resp.Status())
					fmt.Printf("%c[1;0;31m%s%c[0m %-30s\n", 0x1b, "[x]", 0x1b, "MultiNode information Error!")
				}
			} else {
				// 异常请求
				log.Fatal(err)
			}
			// End MultiNode
		}
		// Start switch
	}
	// fmt.Println("Passive")
}
