package podman

import (
	"fmt"

	"github.com/skupperproject/skupper/pkg/utils"

	"github.com/skupperproject/skupper/api/types"
)

const (
	SharedTlsCertificates = "skupper-router-certs"
)

var (
	Username                = utils.ReadUsername()
	SkupperContainerVolumes = []string{"skupper-services", "skupper-local-server", "skupper-internal", "skupper-site-server", SharedTlsCertificates,
		types.ConsoleServerSecret, types.ConsoleUsersSecret, "prometheus-server-config", "prometheus-storage-volume"}
)

func OwnedBySkupper(resource string, labels map[string]string) error {
	notOwnedErr := fmt.Errorf("%s is not owned by Skupper", resource)
	if labels == nil {
		return notOwnedErr
	}
	if app, ok := labels["application"]; !ok || app != types.AppName {
		return notOwnedErr
	}
	return nil
}
