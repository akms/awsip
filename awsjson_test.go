package awsip

import (
	"testing"
)

func TestNewAwsIpadder(t *testing.T) {
	var t_aws *AwsJson
	aws := NewAwsIpadder()
	if aws == t_aws {
		t.Errorf("want type: %T\n get type: %T", aws, t_aws)
	}
	if aws.SyncToken == "" {
		t.Errorf("want numeric string. example:1417984028")
	}
	if aws.CreateDate == "" {
		t.Errorf("want date string. example:2014-12-07-20-31-01")
	}
	if aws.Prefixes == nil {
		t.Errorf("want map array. example:[{ ip_prefix:xx.xx.xx.xx/yy, region:us-east-1, service:AMAZON }]")
	}
	if aws.Headder != "" {
		t.Errorf("want nil.got %s", aws.Headder)
	}
}

func TestGetIpadder(t *testing.T) {
	aws := NewAwsIpadder()
	var (
		region  string = aws.Prefixes[0]["region"]
		service string = aws.Prefixes[0]["service"]
	)
	b := aws.GetIpadder(region, service)
	if b == nil {
		t.Errorf("want *bytes.Buffer value. got nil value.")
	}
}

func TestGetIpadderRegion(t *testing.T) {
	aws := NewAwsIpadder()
	var region string = aws.Prefixes[0]["region"]
	b := aws.GetIpadderRegion(region)
	if b == nil {
		t.Errorf("want *bytes.Buffer value. got nil value.")
	}
}

func TestGetIpadderService(t *testing.T) {
	aws := NewAwsIpadder()
	var service string = aws.Prefixes[0]["service"]
	b := aws.GetIpadderService(service)
	if b == nil {
		t.Errorf("want *bytes.Buffer value. got nil value.")
	}
}

func TestSetHeadder(t *testing.T) {
	aws := NewAwsIpadder()
	aws.SetHeadder("hoge")
	if aws.Headder != "hoge" {
		t.Errorf("want hoge. got %s", aws.Headder)
	}
}
