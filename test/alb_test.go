package test

import (
	"crypto/tls"
	"fmt"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"testing"
	"time"
)

func TestAlbExample(t *testing.T) {
	// 모듈 위치 지정
	opts := &terraform.Options{
		TerraformDir: "../module/alb",
	}

	// 테스트 종료시 모든 리소스 삭제
	defer terraform.Destroy(t, opts)

	// 테라폼 init 및 apply
	terraform.InitAndApply(t, opts)

	// ALB의 URL 정보 가져오기
	albDnsName := terraform.OutputRequired(t, opts, "alb_dns_name")
	url := fmt.Sprintf("http://%s", albDnsName)

	// ALB의 기본 동작이 작동하고 404를 반환하는지 테스트

	expectedStatus := 404
	expectedBody := "404: page not found"

	maxRetries := 10
	timeBetweenRetries := 10 * time.Second

	http_helper.HttpGetWithRetry(
		t,
		url,
		&tls.Config{},
		expectedStatus,
		expectedBody,
		maxRetries,
		timeBetweenRetries,
	)
}
