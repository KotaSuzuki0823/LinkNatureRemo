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

//go fmt対策
func Dummy() {
	println("DUMMY")
}

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
