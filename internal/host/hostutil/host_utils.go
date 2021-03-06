package hostutil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-openapi/swag"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/thoas/go-funk"

	"github.com/pkg/errors"

	"github.com/openshift/assisted-service/models"
)

const (
	MaxHostnameLength = 253
)

func GetCurrentHostName(host *models.Host) (string, error) {
	var inventory models.Inventory
	if host.RequestedHostname != "" {
		return host.RequestedHostname, nil
	}
	err := json.Unmarshal([]byte(host.Inventory), &inventory)
	if err != nil {
		return "", err
	}
	return inventory.Hostname, nil
}

func GetHostnameForMsg(host *models.Host) string {
	hostName, err := GetCurrentHostName(host)
	// An error here probably indicates that the agent didn't send inventory yet, fall back to UUID
	if err != nil || hostName == "" {
		return host.ID.String()
	}
	return hostName
}

func GetEventSeverityFromHostStatus(status string) string {
	switch status {
	case models.HostStatusDisconnected:
		return models.EventSeverityWarning
	case models.HostStatusInstallingPendingUserAction:
		return models.EventSeverityWarning
	case models.HostStatusInsufficient:
		return models.EventSeverityWarning
	case models.HostStatusError:
		return models.EventSeverityError
	default:
		return models.EventSeverityInfo
	}
}

func ValidateHostname(hostname string) error {
	if len(hostname) > MaxHostnameLength {
		return common.NewApiError(http.StatusBadRequest, errors.New("hostname is too long"))
	}
	pattern := "^[a-z0-9][a-z0-9-]{0,62}(?:[.][a-z0-9-]{1,63})*$"
	b, err := regexp.MatchString(pattern, hostname)
	if err != nil {
		return common.NewApiError(http.StatusInternalServerError, errors.Wrapf(err, "Matching hostname"))
	}
	if !b {
		return common.NewApiError(http.StatusBadRequest, errors.Errorf("Hostname does not pass required regex validation: %s. Hostname: %s", pattern, hostname))
	}
	return nil
}

func IgnitionFileName(host *models.Host) string {
	return fmt.Sprintf("%s-%s.ign", host.Role, host.ID)
}

var allowedFlags = []string{"--append-karg", "--delete-karg", "-n", "--copy-network", "--network-dir", "--save-partlabel", "--save-partindex", "--image-url"}

func ValidateInstallerArgs(args []string) error {
	re := regexp.MustCompile("^-+.*")
	for _, arg := range args {
		if !re.MatchString(arg) {
			continue
		}

		found := false
		for _, f := range allowedFlags {
			if arg == f {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("found unexpected flag %s for installer - allowed flags are %v", arg, allowedFlags)
		}
	}

	return nil
}

func FindDiskPathByID(ID string, host *models.Host) string {
	//Currently, the disk is identified by its path. In the future, a disk may be
	//identified by a UUID and then it should be looked up in the inventory
	return ID
}

func GetDeviceFullName(name string) string {
	return fmt.Sprintf("/dev/%s", name)
}

func IsDay2Host(h *models.Host) bool {
	day2HostKinds := []string{models.HostKindAddToExistingClusterHost,
		models.HostKindAddToExistingClusterOCPHost}
	return funk.ContainsString(day2HostKinds, swag.StringValue(h.Kind))
}
