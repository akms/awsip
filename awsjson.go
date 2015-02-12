package awsip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AwsJson struct {
	SyncToken  string              `json:syncToken`
	CreateDate string              `json:createDate`
	Prefixes   []map[string]string `json:prefixes`
	Headder    string
}

func NewAwsIpadder() *AwsJson {
	var (
		resp       *http.Response
		json_body  []byte
		err        error
		json_url   string = "https://ip-ranges.amazonaws.com/ip-ranges.json"
		awsip_json AwsJson
	)

	resp, err = http.Get(json_url)
	if err != nil {
		log.Fatal(err)
	}

	json_body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	err = json.Unmarshal(json_body, &awsip_json)
	if err != nil {
		log.Fatal(err)
	}
	return &awsip_json
}

func (ip *AwsJson) GetIpadder(region string, service string) (ip_buffer *bytes.Buffer) {
	if region == "" || service == "" {
		var geterr error = fmt.Errorf("nil value region or service.\n region: %s, service: %s\n", region, service)
		log.Fatal(geterr)
	}
	ip_buffer = bytes.NewBufferString("")
	for _, val := range ip.Prefixes {
		if val["region"] == region && val["service"] == service {
			ip_buffer.WriteString(ip.Headder)
			ip_buffer.WriteString(" ")
			ip_buffer.WriteString(val["ip_prefix"])
			ip_buffer.WriteString("\n")
		}
	}
	return
}

func (ip *AwsJson) GetIpadderRegion(region string) (ip_buffer *bytes.Buffer) {
	if region == "" {
		var geterr error = fmt.Errorf("nil value region .\n")
		log.Fatal(geterr)
	}
	ip_buffer = bytes.NewBufferString("")
	for _, val := range ip.Prefixes {
		if val["region"] == region {
			ip_buffer.WriteString(ip.Headder)
			ip_buffer.WriteString(" ")
			ip_buffer.WriteString(val["ip_prefix"])
			ip_buffer.WriteString("\n")
		}
	}
	return
}

func (ip *AwsJson) GetIpadderService(service string) (ip_buffer *bytes.Buffer) {
	if service == "" {
		var geterr error = fmt.Errorf("nil value service.\n")
		log.Fatal(geterr)
	}
	ip_buffer = bytes.NewBufferString("")
	for _, val := range ip.Prefixes {
		if val["service"] == service {
			ip_buffer.WriteString(ip.Headder)
			ip_buffer.WriteString(" ")
			ip_buffer.WriteString(val["ip_prefix"])
			ip_buffer.WriteString("\n")
		}
	}
	return
}

func (ip *AwsJson) SetHeadder(head_string string) {
	if head_string == "" {
		var geterr error = fmt.Errorf("No value headder string.\n")
		log.Fatal(geterr)
	}
	ip.Headder = head_string
}
