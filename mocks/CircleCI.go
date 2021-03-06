// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import circleci "github.com/armakuni/circleci-workflow-dashboard/circleci"
import mock "github.com/stretchr/testify/mock"

// CircleCI is an autogenerated mock type for the CircleCI type
type CircleCI struct {
	mock.Mock
}

// CreateProjectEnvVar provides a mock function with given fields: _a0, _a1, _a2
func (_m *CircleCI) CreateProjectEnvVar(_a0 string, _a1 string, _a2 string) (circleci.ProjectEnvVar, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 circleci.ProjectEnvVar
	if rf, ok := ret.Get(0).(func(string, string, string) circleci.ProjectEnvVar); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(circleci.ProjectEnvVar)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProjectEnvVar provides a mock function with given fields: _a0, _a1
func (_m *CircleCI) DeleteProjectEnvVar(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPipelines provides a mock function with given fields: _a0
func (_m *CircleCI) GetAllPipelines(_a0 circleci.Project) (circleci.Pipelines, error) {
	ret := _m.Called(_a0)

	var r0 circleci.Pipelines
	if rf, ok := ret.Get(0).(func(circleci.Project) circleci.Pipelines); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(circleci.Pipelines)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(circleci.Project) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllProjects provides a mock function with given fields:
func (_m *CircleCI) GetAllProjects() (circleci.Projects, error) {
	ret := _m.Called()

	var r0 circleci.Projects
	if rf, ok := ret.Get(0).(func() circleci.Projects); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(circleci.Projects)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobsForWorkflow provides a mock function with given fields: _a0
func (_m *CircleCI) GetJobsForWorkflow(_a0 circleci.Workflow) (circleci.Jobs, error) {
	ret := _m.Called(_a0)

	var r0 circleci.Jobs
	if rf, ok := ret.Get(0).(func(circleci.Workflow) circleci.Jobs); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(circleci.Jobs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(circleci.Workflow) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectEnvVars provides a mock function with given fields: _a0
func (_m *CircleCI) GetProjectEnvVars(_a0 string) (circleci.ProjectEnvVars, error) {
	ret := _m.Called(_a0)

	var r0 circleci.ProjectEnvVars
	if rf, ok := ret.Get(0).(func(string) circleci.ProjectEnvVars); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(circleci.ProjectEnvVars)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowsForPipeline provides a mock function with given fields: _a0
func (_m *CircleCI) GetWorkflowsForPipeline(_a0 circleci.Pipeline) (circleci.Workflows, error) {
	ret := _m.Called(_a0)

	var r0 circleci.Workflows
	if rf, ok := ret.Get(0).(func(circleci.Pipeline) circleci.Workflows); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(circleci.Workflows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(circleci.Pipeline) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JobLink provides a mock function with given fields: _a0, _a1
func (_m *CircleCI) JobLink(_a0 circleci.Project, _a1 circleci.Job) string {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(circleci.Project, circleci.Job) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PreviousCompleteWorkflowState provides a mock function with given fields: _a0, _a1
func (_m *CircleCI) PreviousCompleteWorkflowState(_a0 circleci.Pipelines, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(circleci.Pipelines, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(circleci.Pipelines, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WorkflowLink provides a mock function with given fields: _a0, _a1, _a2
func (_m *CircleCI) WorkflowLink(_a0 circleci.Project, _a1 circleci.Pipeline, _a2 circleci.Workflow) string {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(circleci.Project, circleci.Pipeline, circleci.Workflow) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WorkflowStatus provides a mock function with given fields: _a0, _a1
func (_m *CircleCI) WorkflowStatus(_a0 circleci.Pipelines, _a1 circleci.Workflow) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(circleci.Pipelines, circleci.Workflow) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(circleci.Pipelines, circleci.Workflow) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
