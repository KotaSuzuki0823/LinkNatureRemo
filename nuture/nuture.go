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

func Curltest() {
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
}
func GetAppliancesList() {
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
}
func GetRegistAppliancesSignal(id string) {
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
}

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
