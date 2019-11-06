package utils

import (
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// OverWriteEnvFromEngine will over-ride the default variables from one provided in the chaosEngine
func OverWriteEnvFromEngine(appns string, chaosEngine string, config *rest.Config, m map[string]string, chaosExperiment string) {
	_, litmusClientSet, err := GenerateClientSets(config)
	if err != nil {
		log.Info(err)
	}
	engineSpec, err := litmusClientSet.LitmuschaosV1alpha1().ChaosEngines(appns).Get(chaosEngine, metav1.GetOptions{})
	envList := engineSpec.Spec.Experiments
	for i := range envList {
		if envList[i].Name == chaosExperiment {
			keyValue := envList[i].Spec.Components
			for j := range keyValue {
				m[keyValue[j].Name] = keyValue[j].Value
			}
		}

	}

}
