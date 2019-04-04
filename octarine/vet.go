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
