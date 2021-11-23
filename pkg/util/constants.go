package util

const (
	prefix                         = "harvesterhci.io"
	RemovedPVCsAnnotationKey       = prefix + "/removedPersistentVolumeClaims"
	AdditionalCASecretName         = "harvester-additional-ca"
	AdditionalCAFileName           = "additional-ca.pem"
	AnnotationMigrationTarget      = prefix + "/migrationTargetNodeName"
	AnnotationMigrationUID         = prefix + "/migrationUID"
	AnnotationMigrationState       = prefix + "/migrationState"
	AnnotationTimestamp            = prefix + "/timestamp"
	AnnotationVolumeClaimTemplates = prefix + "/volumeClaimTemplates"
	AnnotationImageID              = prefix + "/imageId"
	AnnotationHash                 = prefix + "/hash"

	BackupTargetSecretName      = "harvester-backup-target-secret"
	LonghornSystemNamespaceName = "longhorn-system"
	CattleSystemNamespaceName   = "cattle-system"
	KubeSystemNamespace         = "kube-system"

	HTTPProxyEnv  = "HTTP_PROXY"
	HTTPSProxyEnv = "HTTPS_PROXY"
	NoProxyEnv    = "NO_PROXY"
)
