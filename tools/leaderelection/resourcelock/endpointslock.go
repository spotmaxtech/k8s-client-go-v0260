/*
Copyright 2016 The Kubernetes Authors.

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

package resourcelock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spotmaxtech/k8s-api-v0260/core/v1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	corev1client "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes/typed/core/v1"
)

type endpointsLock struct {
	// EndpointsMeta should contain a Name and a Namespace of an
	// Endpoints object that the LeaderElector will attempt to lead.
	EndpointsMeta metav1.ObjectMeta
	Client        corev1client.EndpointsGetter
	LockConfig    ResourceLockConfig
	e             *v1.Endpoints
}

// Get returns the election record from a Endpoints Annotation
func (el *endpointsLock) Get(ctx context.Context) (*LeaderElectionRecord, []byte, error) {
	var record LeaderElectionRecord
	ep, err := el.Client.Endpoints(el.EndpointsMeta.Namespace).Get(ctx, el.EndpointsMeta.Name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, err
	}
	el.e = ep
	if el.e.Annotations == nil {
		el.e.Annotations = make(map[string]string)
	}
	recordStr, found := el.e.Annotations[LeaderElectionRecordAnnotationKey]
	recordBytes := []byte(recordStr)
	if found {
		if err := json.Unmarshal(recordBytes, &record); err != nil {
			return nil, nil, err
		}
	}
	return &record, recordBytes, nil
}

// Create attempts to create a LeaderElectionRecord annotation
func (el *endpointsLock) Create(ctx context.Context, ler LeaderElectionRecord) error {
	recordBytes, err := json.Marshal(ler)
	if err != nil {
		return err
	}
	el.e, err = el.Client.Endpoints(el.EndpointsMeta.Namespace).Create(ctx, &v1.Endpoints{
		ObjectMeta: metav1.ObjectMeta{
			Name:      el.EndpointsMeta.Name,
			Namespace: el.EndpointsMeta.Namespace,
			Annotations: map[string]string{
				LeaderElectionRecordAnnotationKey: string(recordBytes),
			},
		},
	}, metav1.CreateOptions{})
	return err
}

// Update will update and existing annotation on a given resource.
func (el *endpointsLock) Update(ctx context.Context, ler LeaderElectionRecord) error {
	if el.e == nil {
		return errors.New("endpoint not initialized, call get or create first")
	}
	recordBytes, err := json.Marshal(ler)
	if err != nil {
		return err
	}
	if el.e.Annotations == nil {
		el.e.Annotations = make(map[string]string)
	}
	el.e.Annotations[LeaderElectionRecordAnnotationKey] = string(recordBytes)
	e, err := el.Client.Endpoints(el.EndpointsMeta.Namespace).Update(ctx, el.e, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	el.e = e
	return nil
}

// RecordEvent in leader election while adding meta-data
func (el *endpointsLock) RecordEvent(s string) {
	if el.LockConfig.EventRecorder == nil {
		return
	}
	events := fmt.Sprintf("%v %v", el.LockConfig.Identity, s)
	subject := &v1.Endpoints{ObjectMeta: el.e.ObjectMeta}
	// Populate the type meta, so we don't have to get it from the schema
	subject.Kind = "Endpoints"
	subject.APIVersion = v1.SchemeGroupVersion.String()
	el.LockConfig.EventRecorder.Eventf(subject, v1.EventTypeNormal, "LeaderElection", events)
}

// Describe is used to convert details on current resource lock
// into a string
func (el *endpointsLock) Describe() string {
	return fmt.Sprintf("%v/%v", el.EndpointsMeta.Namespace, el.EndpointsMeta.Name)
}

// Identity returns the Identity of the lock
func (el *endpointsLock) Identity() string {
	return el.LockConfig.Identity
}
