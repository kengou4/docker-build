// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -package mockhelmclient -destination=./mock/interface.go -self_package=. Client
//

// Package mockhelmclient is a generated GoMock package.
package mockhelmclient

import (
	context "context"
	reflect "reflect"

	helmclient "github.com/mittwald/go-helm-client"
	gomock "go.uber.org/mock/gomock"
	action "helm.sh/helm/v3/pkg/action"
	chart "helm.sh/helm/v3/pkg/chart"
	cli "helm.sh/helm/v3/pkg/cli"
	getter "helm.sh/helm/v3/pkg/getter"
	release "helm.sh/helm/v3/pkg/release"
	repo "helm.sh/helm/v3/pkg/repo"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddOrUpdateChartRepo mocks base method.
func (m *MockClient) AddOrUpdateChartRepo(entry repo.Entry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateChartRepo", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrUpdateChartRepo indicates an expected call of AddOrUpdateChartRepo.
func (mr *MockClientMockRecorder) AddOrUpdateChartRepo(entry any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateChartRepo", reflect.TypeOf((*MockClient)(nil).AddOrUpdateChartRepo), entry)
}

// GetChart mocks base method.
func (m *MockClient) GetChart(chartName string, chartPathOptions *action.ChartPathOptions) (*chart.Chart, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChart", chartName, chartPathOptions)
	ret0, _ := ret[0].(*chart.Chart)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetChart indicates an expected call of GetChart.
func (mr *MockClientMockRecorder) GetChart(chartName, chartPathOptions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChart", reflect.TypeOf((*MockClient)(nil).GetChart), chartName, chartPathOptions)
}

// GetProviders mocks base method.
func (m *MockClient) GetProviders() getter.Providers {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProviders")
	ret0, _ := ret[0].(getter.Providers)
	return ret0
}

// GetProviders indicates an expected call of GetProviders.
func (mr *MockClientMockRecorder) GetProviders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProviders", reflect.TypeOf((*MockClient)(nil).GetProviders))
}

// GetRelease mocks base method.
func (m *MockClient) GetRelease(name string) (*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelease", name)
	ret0, _ := ret[0].(*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelease indicates an expected call of GetRelease.
func (mr *MockClientMockRecorder) GetRelease(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelease", reflect.TypeOf((*MockClient)(nil).GetRelease), name)
}

// GetReleaseValues mocks base method.
func (m *MockClient) GetReleaseValues(name string, allValues bool) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReleaseValues", name, allValues)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReleaseValues indicates an expected call of GetReleaseValues.
func (mr *MockClientMockRecorder) GetReleaseValues(name, allValues any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReleaseValues", reflect.TypeOf((*MockClient)(nil).GetReleaseValues), name, allValues)
}

// GetSettings mocks base method.
func (m *MockClient) GetSettings() *cli.EnvSettings {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings")
	ret0, _ := ret[0].(*cli.EnvSettings)
	return ret0
}

// GetSettings indicates an expected call of GetSettings.
func (mr *MockClientMockRecorder) GetSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockClient)(nil).GetSettings))
}

// InstallChart mocks base method.
func (m *MockClient) InstallChart(ctx context.Context, spec *helmclient.ChartSpec, opts *helmclient.GenericHelmOptions) (*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallChart", ctx, spec, opts)
	ret0, _ := ret[0].(*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallChart indicates an expected call of InstallChart.
func (mr *MockClientMockRecorder) InstallChart(ctx, spec, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallChart", reflect.TypeOf((*MockClient)(nil).InstallChart), ctx, spec, opts)
}

// InstallOrUpgradeChart mocks base method.
func (m *MockClient) InstallOrUpgradeChart(ctx context.Context, spec *helmclient.ChartSpec, opts *helmclient.GenericHelmOptions) (*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallOrUpgradeChart", ctx, spec, opts)
	ret0, _ := ret[0].(*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallOrUpgradeChart indicates an expected call of InstallOrUpgradeChart.
func (mr *MockClientMockRecorder) InstallOrUpgradeChart(ctx, spec, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallOrUpgradeChart", reflect.TypeOf((*MockClient)(nil).InstallOrUpgradeChart), ctx, spec, opts)
}

// LintChart mocks base method.
func (m *MockClient) LintChart(spec *helmclient.ChartSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LintChart", spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// LintChart indicates an expected call of LintChart.
func (mr *MockClientMockRecorder) LintChart(spec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LintChart", reflect.TypeOf((*MockClient)(nil).LintChart), spec)
}

// ListDeployedReleases mocks base method.
func (m *MockClient) ListDeployedReleases() ([]*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeployedReleases")
	ret0, _ := ret[0].([]*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeployedReleases indicates an expected call of ListDeployedReleases.
func (mr *MockClientMockRecorder) ListDeployedReleases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployedReleases", reflect.TypeOf((*MockClient)(nil).ListDeployedReleases))
}

// ListReleaseHistory mocks base method.
func (m *MockClient) ListReleaseHistory(name string, max int) ([]*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReleaseHistory", name, max)
	ret0, _ := ret[0].([]*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReleaseHistory indicates an expected call of ListReleaseHistory.
func (mr *MockClientMockRecorder) ListReleaseHistory(name, max any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleaseHistory", reflect.TypeOf((*MockClient)(nil).ListReleaseHistory), name, max)
}

// ListReleasesByStateMask mocks base method.
func (m *MockClient) ListReleasesByStateMask(arg0 action.ListStates) ([]*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReleasesByStateMask", arg0)
	ret0, _ := ret[0].([]*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReleasesByStateMask indicates an expected call of ListReleasesByStateMask.
func (mr *MockClientMockRecorder) ListReleasesByStateMask(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleasesByStateMask", reflect.TypeOf((*MockClient)(nil).ListReleasesByStateMask), arg0)
}

// RollbackRelease mocks base method.
func (m *MockClient) RollbackRelease(spec *helmclient.ChartSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackRelease", spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackRelease indicates an expected call of RollbackRelease.
func (mr *MockClientMockRecorder) RollbackRelease(spec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackRelease", reflect.TypeOf((*MockClient)(nil).RollbackRelease), spec)
}

// RunChartTests mocks base method.
func (m *MockClient) RunChartTests(releaseName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunChartTests", releaseName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunChartTests indicates an expected call of RunChartTests.
func (mr *MockClientMockRecorder) RunChartTests(releaseName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunChartTests", reflect.TypeOf((*MockClient)(nil).RunChartTests), releaseName)
}

// SetDebugLog mocks base method.
func (m *MockClient) SetDebugLog(debugLog action.DebugLog) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDebugLog", debugLog)
}

// SetDebugLog indicates an expected call of SetDebugLog.
func (mr *MockClientMockRecorder) SetDebugLog(debugLog any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDebugLog", reflect.TypeOf((*MockClient)(nil).SetDebugLog), debugLog)
}

// TemplateChart mocks base method.
func (m *MockClient) TemplateChart(spec *helmclient.ChartSpec, options *helmclient.HelmTemplateOptions) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TemplateChart", spec, options)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TemplateChart indicates an expected call of TemplateChart.
func (mr *MockClientMockRecorder) TemplateChart(spec, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TemplateChart", reflect.TypeOf((*MockClient)(nil).TemplateChart), spec, options)
}

// UninstallRelease mocks base method.
func (m *MockClient) UninstallRelease(spec *helmclient.ChartSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallRelease", spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallRelease indicates an expected call of UninstallRelease.
func (mr *MockClientMockRecorder) UninstallRelease(spec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallRelease", reflect.TypeOf((*MockClient)(nil).UninstallRelease), spec)
}

// UninstallReleaseByName mocks base method.
func (m *MockClient) UninstallReleaseByName(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallReleaseByName", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallReleaseByName indicates an expected call of UninstallReleaseByName.
func (mr *MockClientMockRecorder) UninstallReleaseByName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallReleaseByName", reflect.TypeOf((*MockClient)(nil).UninstallReleaseByName), name)
}

// UpdateChartRepos mocks base method.
func (m *MockClient) UpdateChartRepos() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChartRepos")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateChartRepos indicates an expected call of UpdateChartRepos.
func (mr *MockClientMockRecorder) UpdateChartRepos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChartRepos", reflect.TypeOf((*MockClient)(nil).UpdateChartRepos))
}

// UpgradeChart mocks base method.
func (m *MockClient) UpgradeChart(ctx context.Context, spec *helmclient.ChartSpec, opts *helmclient.GenericHelmOptions) (*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeChart", ctx, spec, opts)
	ret0, _ := ret[0].(*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeChart indicates an expected call of UpgradeChart.
func (mr *MockClientMockRecorder) UpgradeChart(ctx, spec, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeChart", reflect.TypeOf((*MockClient)(nil).UpgradeChart), ctx, spec, opts)
}

// MockRollBack is a mock of RollBack interface.
type MockRollBack struct {
	ctrl     *gomock.Controller
	recorder *MockRollBackMockRecorder
}

// MockRollBackMockRecorder is the mock recorder for MockRollBack.
type MockRollBackMockRecorder struct {
	mock *MockRollBack
}

// NewMockRollBack creates a new mock instance.
func NewMockRollBack(ctrl *gomock.Controller) *MockRollBack {
	mock := &MockRollBack{ctrl: ctrl}
	mock.recorder = &MockRollBackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRollBack) EXPECT() *MockRollBackMockRecorder {
	return m.recorder
}

// RollbackRelease mocks base method.
func (m *MockRollBack) RollbackRelease(spec *helmclient.ChartSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackRelease", spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackRelease indicates an expected call of RollbackRelease.
func (mr *MockRollBackMockRecorder) RollbackRelease(spec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackRelease", reflect.TypeOf((*MockRollBack)(nil).RollbackRelease), spec)
}