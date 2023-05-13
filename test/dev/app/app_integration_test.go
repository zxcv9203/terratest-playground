package test

import (
	"crypto/tls"
	"fmt"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"strings"
	"testing"
	"time"
)

const devPath = "../../../dev"
const devMysqlPath = devPath + "/mysql"
const devAppPath = devPath + "/app"

func TestHelloWorldAppDevelop(t *testing.T) {
	t.Parallel()

	// MySQL DB 배포
	dbOpts := createDbOpts(t, devMysqlPath)
	defer terraform.Destroy(t, dbOpts)
	terraform.InitAndApply(t, dbOpts)

	// App 배포
	appOpts := createAppOpts(dbOpts, devAppPath)
	defer terraform.Destroy(t, appOpts)
	terraform.InitAndApply(t, appOpts)

	// App 동작 확인
	validateApp(t, appOpts)
}

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
	}
}
