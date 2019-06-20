// Copyright 2019 The Meshery Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package octarine

import (
	"github.com/layer5io/meshery-octarine/meshes"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/informers"
)

type metaInformerFactory struct {
	k8s informers.SharedInformerFactory
}

func (m *metaInformerFactory) K8s() informers.SharedInformerFactory {
	return m.k8s
}

func (oClient *OctarineClient) runVet() error {
	kubeInformerFactory := informers.NewSharedInformerFactory(oClient.k8sClientset, 0)
//	informerFactory := &metaInformerFactory{
//		k8s: kubeInformerFactory,
//	}

	stopCh := make(chan struct{})

	kubeInformerFactory.Start(stopCh)
	oks := kubeInformerFactory.WaitForCacheSync(stopCh)
	for inf, ok := range oks {
		if !ok {
			err := errors.Errorf("Failed to sync: %s", inf)
			logrus.Error(err)
			return err
		}
	}
	return nil
}

func convertVetLevelToMesheryLevel(level string) meshes.EventType {
	switch level {
	// case "INFO":
	// 	return
	case "WARNING":
		return meshes.EventType_WARN
	case "ERROR":
		return meshes.EventType_ERROR
	default:
		return meshes.EventType_INFO
	}
}
