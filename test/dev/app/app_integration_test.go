package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"testing"
)

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
