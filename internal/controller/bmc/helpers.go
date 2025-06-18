package bmc

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
)

// equalMap compares two string maps.
func equalMap(a, b map[string]string) bool {
	return equality.Semantic.DeepEqual(a, b)
}

// equalPorts compares two sets of container or service ports.
func equalPorts(a, b []corev1.ServicePort) bool {
	return equality.Semantic.DeepEqual(a, b)
}

// equalEnv compares two sets of EnvVars.
func equalEnv(a, b []corev1.EnvVar) bool {
	return equality.Semantic.DeepEqual(a, b)
}

func equalMounts(a, b []corev1.VolumeMount) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Name != b[i].Name || a[i].MountPath != b[i].MountPath || a[i].SubPath != b[i].SubPath {
			return false
		}
	}
	return true
}
