package utility

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/entgigi/bundle-operator/common"
)

// GetWatchNamespace returns the Namespace the operator should be watching for changes
func GetWatchNamespace() (string, error) {
	// WatchNamespaceEnvVar is the constant for env variable WATCH_NAMESPACE
	// which specifies the Namespace to watch.
	// An empty value means the operator is running with cluster scope.

	ns, found := os.LookupEnv(common.WatchNamespaceEnvVar)
	if !found {
		return "", fmt.Errorf("%s must be set", common.WatchNamespaceEnvVar)
	}
	return ns, nil
}

func GetOperatorDeploymentType() string {
	operatorType, found := os.LookupEnv(common.OperatorTypeEnvVar)
	if found {
		return operatorType
	} else {
		// default
		return common.OperatorTypeStandard
	}
}

func TruncateString(s string, max int) string {
	if max > len(s) {
		return s
	}
	return s[:max]
}

func GenerateSha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))

}
