package nuture

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var (
	token = os.Getenv("NUTUREREMO_TOKEN")
)

// Dummy go-fmt対策
func Dummy() {
	println("DUMMY")
}

// Curltest natureRemo接続テスト（機器情報を取得）
func Curltest() []byte {
	url := "https://api.nature.global/1/devices"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("accept", " application/json")
	req.Header.Set("Authorization", " Bearer "+token)

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)
	return dumpResp
}

// GetAppliancesList 家電リストの取得
func GetAppliancesList() []byte {
	url := "https://api.nature.global/1/appliances"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("accept", " application/json")
	req.Header.Set("Authorization", " Bearer "+token)

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)

	return dumpResp
}

// GetRegistAppliancesSignal 家電IDで指定した家電のシグナル情報を取得
func GetRegistAppliancesSignal(id string) []byte {
	url := "https://api.nature.global/1/appliances/" + id + "signals"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("accept", " application/json")
	req.Header.Set("Authorization", " Bearer "+token)
	req.Header.Set("appliance", " "+id)

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)
	return dumpResp
}

// SendSignal 赤外線シグナルの送信
func SendSignal(id string) {
	url := "https://api.nature.global/1/signals/" + id + "/send"
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Set("accept", " application/json")
	req.Header.Set("Authorization", " Bearer "+token)
	//req.Header.Set("sinal", " "+id)

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)
}
