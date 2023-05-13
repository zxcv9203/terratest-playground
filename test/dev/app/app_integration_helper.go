package test

import (
	"crypto/tls"
	"fmt"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"strings"
	"testing"
	"time"
)

const devPath = "../../../dev"
const devMysqlPath = devPath + "/mysql"
const devAppPath = devPath + "/app"

func validateApp(t *testing.T, opts *terraform.Options) {
	albDnsName := terraform.OutputRequired(t, opts, "alb_dns_name")
	url := fmt.Sprintf("http://%s", albDnsName)

	maxRetries := 10
	timeBetweenRetries := 10 * time.Second

	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		url,
		&tls.Config{},
		maxRetries,
		timeBetweenRetries,
		func(status int, body string) bool {
			return status == 200 && strings.Contains(body, "Hello, World")
		},
	)
}

func createAppOpts(opts *terraform.Options, path string) *terraform.Options {
	return &terraform.Options{
		TerraformDir: path,

		Vars: map[string]interface{}{
			"db_remote_state_bucket": opts.BackendConfig["bucket"],
			"db_remote_state_key":    opts.BackendConfig["key"],
			"environment":            opts.Vars["db_name"],
		},

		Reconfigure: true,
	}
}

func createDbOpts(t *testing.T, dir string) *terraform.Options {
	uniqueId := random.UniqueId()

	testBucketName := "yongc-s3-bucket-test"
	testBucketRegion := "us-west-1"
	testStateKey := fmt.Sprintf("%s/%s/terraform.tfstate", t.Name(), uniqueId)

	return &terraform.Options{
		TerraformDir: dir,

		Vars: map[string]interface{}{
			"db_name":     fmt.Sprintf("test%s", uniqueId),
			"db_password": "password",
		},

		BackendConfig: map[string]interface{}{
			"bucket":  testBucketName,
			"region":  testBucketRegion,
			"key":     testStateKey,
			"encrypt": true,
		},

		Reconfigure: true,
	}
}

func validateAppWithPath(t *testing.T, path string) {
	opts := test_structure.LoadTerraformOptions(t, path)
	validateApp(t, opts)
}

func teardownApp(t *testing.T, path string) {
	opts := test_structure.LoadTerraformOptions(t, path)
	defer terraform.Destroy(t, opts)
}

func deployApp(t *testing.T, dbPath string, appPath string) {
	dbOpts := test_structure.LoadTerraformOptions(t, dbPath)
	appOpts := createAppOpts(dbOpts, appPath)

	test_structure.SaveTerraformOptions(t, appPath, appOpts)

	terraform.InitAndApply(t, appOpts)
}

func teardownDb(t *testing.T, path string) {
	opts := test_structure.LoadTerraformOptions(t, path)
	defer terraform.Destroy(t, opts)
}

func deployDb(t *testing.T, path string) {
	opts := createDbOpts(t, path)

	// 나중에 실행되는 다른 테스트 단계에서 데이터를 다시 읽을 수 있도록 데이터를 디스크에 저장
	test_structure.SaveTerraformOptions(t, path, opts)

	terraform.InitAndApply(t, opts)
}
