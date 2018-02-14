package v2v3action

import (
	"code.cloudfoundry.org/cli/actor/v3action"
)

//go:generate counterfeiter . V3Actor

type V3Actor interface {
	GetOrganizationByName(orgName string) (v3action.Organization, v3action.Warnings, error)
	ShareServiceInstanceToSpaces(serviceInstanceGUID string, spaceGUIDs []string) (v3action.RelationshipList, v3action.Warnings, error)

	CloudControllerAPIVersion() string
}