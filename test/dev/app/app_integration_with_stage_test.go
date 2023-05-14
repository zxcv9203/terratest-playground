package test

import (
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"testing"
)

/*
* skip 환경변수를 통해 특정환경에 대해 생성 및 삭제를 스킵시킬 수 있습니다.
* SKIP_deploy_db = true -> DB 배포 스킵
* SKIP_teardown_db -> DB 삭제 스킵
 */
func TestAppWithStage(t *testing.T) {
	t.Parallel()

	// 변수로 선언하여 짧은 구문으로 사용할 수 있도록 처리
	stage := test_structure.RunTestStage

	// MySQL 배포
	defer stage(t, "teardown_db", func() { teardownDb(t, devMysqlPath) })
	stage(t, "deploy_db", func() { deployDb(t, devMysqlPath) })

	// App 배포
	defer stage(t, "teardown_app", func() { teardownApp(t, devAppPath) })
	stage(t, "deploy_app", func() { deployApp(t, devMysqlPath, devAppPath) })

	// App 유효성 검증
	stage(t, "validate_app", func() { validateAppWithPath(t, devAppPath) })

	// App 재배포
	stage(t, "redeploy_app", func() { redeployApp(t, devAppPath) })
}
