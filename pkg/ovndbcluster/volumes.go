package ovndbcluster

import corev1 "k8s.io/api/core/v1"

// GetDBClusterVolumes -
// TODO: merge to GetVolumes when other controllers also switched to current config
// mechanism.
func GetDBClusterVolumes(name string) []corev1.Volume {
	var config0640AccessMode int32 = 0640
	var scriptsVolumeDefaultMode int32 = 0755

	return []corev1.Volume{
		{
			Name: "etc-machine-id",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/etc/machine-id",
				},
			},
		},
		{
			Name: "etc-localtime",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/etc/localtime",
				},
			},
		},
		{
			Name: "scripts",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					DefaultMode: &scriptsVolumeDefaultMode,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: name + "-scripts",
					},
				},
			},
		},
		{
			Name: "config-data",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					DefaultMode: &config0640AccessMode,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: name + "-config-data",
					},
				},
			},
		},
	}

}

// GetDBClusterVolumeMounts - OVN DBCluster VolumeMounts
func GetDBClusterVolumeMounts(name string) []corev1.VolumeMount {
	return []corev1.VolumeMount{
		{
			Name:      "etc-machine-id",
			MountPath: "/etc/machine-id",
			ReadOnly:  true,
		},
		{
			Name:      "etc-localtime",
			MountPath: "/etc/localtime",
			ReadOnly:  true,
		},
		{
			Name:      "scripts",
			MountPath: "/usr/local/bin/container-scripts",
			ReadOnly:  true,
		},
		{
			Name:      "config-data",
			MountPath: "/var/lib/config-data",
			ReadOnly:  false,
		},
		{
			Name:      name,
			MountPath: "/etc/ovn",
			ReadOnly:  false,
		},
	}

}
