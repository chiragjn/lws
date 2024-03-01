/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	leaderworkersetv1 "sigs.k8s.io/lws/api/leaderworkerset/v1"
)

// LeaderWorkerTemplateApplyConfiguration represents an declarative configuration of the LeaderWorkerTemplate type for use
// with apply.
type LeaderWorkerTemplateApplyConfiguration struct {
	LeaderTemplate *v1.PodTemplateSpec                  `json:"leaderTemplate,omitempty"`
	WorkerTemplate *v1.PodTemplateSpec                  `json:"workerTemplate,omitempty"`
	Size           *int32                               `json:"size,omitempty"`
	RestartPolicy  *leaderworkersetv1.RestartPolicyType `json:"restartPolicy,omitempty"`
}

// LeaderWorkerTemplateApplyConfiguration constructs an declarative configuration of the LeaderWorkerTemplate type for use with
// apply.
func LeaderWorkerTemplate() *LeaderWorkerTemplateApplyConfiguration {
	return &LeaderWorkerTemplateApplyConfiguration{}
}

// WithLeaderTemplate sets the LeaderTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LeaderTemplate field is set to the value of the last call.
func (b *LeaderWorkerTemplateApplyConfiguration) WithLeaderTemplate(value v1.PodTemplateSpec) *LeaderWorkerTemplateApplyConfiguration {
	b.LeaderTemplate = &value
	return b
}

// WithWorkerTemplate sets the WorkerTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WorkerTemplate field is set to the value of the last call.
func (b *LeaderWorkerTemplateApplyConfiguration) WithWorkerTemplate(value v1.PodTemplateSpec) *LeaderWorkerTemplateApplyConfiguration {
	b.WorkerTemplate = &value
	return b
}

// WithSize sets the Size field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Size field is set to the value of the last call.
func (b *LeaderWorkerTemplateApplyConfiguration) WithSize(value int32) *LeaderWorkerTemplateApplyConfiguration {
	b.Size = &value
	return b
}

// WithRestartPolicy sets the RestartPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RestartPolicy field is set to the value of the last call.
func (b *LeaderWorkerTemplateApplyConfiguration) WithRestartPolicy(value leaderworkersetv1.RestartPolicyType) *LeaderWorkerTemplateApplyConfiguration {
	b.RestartPolicy = &value
	return b
}
