package testinfra

import (
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPerfTestJobPresubmit test presubmit jobs
func TestPerfTestJobPresubmit(t *testing.T) {
	// WHEN
	jobConfig, err := tester.ReadJobConfig("./../../../../prow/jobs/test-infra/performance-test.yaml")
	// THEN
	require.NoError(t, err)

	actualPresubmit := tester.FindPresubmitJobByNameAndBranch(jobConfig.Presubmits["kyma-project/test-infra"], "pre-master-perf-test", "master")
	require.NotNil(t, actualPresubmit)
	assert.Equal(t, 10, actualPresubmit.MaxConcurrency)
	assert.False(t, actualPresubmit.SkipReport)
	assert.True(t, actualPresubmit.Decorate)
	assert.Equal(t, "github.com/kyma-project/test-infra", actualPresubmit.PathAlias)
	tester.AssertThatHasPresets(t, actualPresubmit.JobBase, tester.PresetDindEnabled, tester.PresetDockerPushRepoKyma, tester.PresetGcrPush, tester.PresetBuildPr)
	tester.AssertThatJobRunIfChanged(t, *actualPresubmit, "performance-tools/performance-cluster/fix")
	assert.Equal(t, "^performance-tools/performance-cluster/", actualPresubmit.RunIfChanged)
	tester.AssertThatExecGolangBuildpack(t, actualPresubmit.JobBase, tester.ImageGolangBuildpackLatest, "/home/prow/go/src/github.com/kyma-project/test-infra/performance-tools/performance-cluster")
}

// TestPerfTestJobPostsubmit test postsubmit jobs
func TestPerfTestJobPostsubmit(t *testing.T) {
	// WHEN
	jobConfig, err := tester.ReadJobConfig("./../../../../prow/jobs/test-infra/perf-test.yaml")
	// THEN
	require.NoError(t, err)

	expName := "post-master-perf-test"
	actualPost := tester.FindPostsubmitJobByNameAndBranch(jobConfig.Postsubmits["kyma-project/test-infra"], expName, "master")
	require.NotNil(t, actualPost)

	assert.Equal(t, expName, actualPost.Name)
	assert.Equal(t, []string{"master"}, actualPost.Branches)

	assert.Equal(t, 10, actualPost.MaxConcurrency)
	assert.True(t, actualPost.Decorate)
	assert.Equal(t, "github.com/kyma-project/test-infra", actualPost.PathAlias)
	tester.AssertThatHasPresets(t, actualPost.JobBase, tester.PresetDindEnabled, tester.PresetDockerPushRepoKyma, tester.PresetGcrPush, tester.PresetBuildMaster)
	assert.Equal(t, "^performance-tools/performance-cluster/", actualPost.RunIfChanged)
	tester.AssertThatExecGolangBuildpack(t, actualPost.JobBase, tester.ImageGolangBuildpackLatest, "/home/prow/go/src/github.com/kyma-project/test-infra/performance-tools/performance-cluster")
}
