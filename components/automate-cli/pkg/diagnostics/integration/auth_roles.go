package integration

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pkg/errors"

	"github.com/chef/automate/components/automate-cli/pkg/diagnostics"
	"github.com/chef/automate/components/automate-cli/pkg/diagnostics/lbrequest"
)

const createRoleTemplate = `
{
	"id":"{{ .ID }}",
	"name":"{{ .ID }} test role",
	"actions": ["{{ .Action }}"]
}
`

// RoleInfo represents the role parameters, including v1 and v2 fields.
type RoleInfo struct {
	Role struct {
		ID      string   `json:"id"`
		Name    string   `json:"name"`
		Actions []string `json:"actions"`
	}
}

type authRoleSave struct {
	ID      string   `json:"id"`
	Actions []string `json:"actions"`
}

// CreateRandomRole creates a role
func CreateRandomRole(tstCtx diagnostics.TestContext) (*RoleInfo, error) {
	roleInfo := RoleInfo{}
	err := MustJSONDecodeSuccess(
		tstCtx.DoLBRequest(
			"/apis/iam/v2/roles",
			lbrequest.WithMethod("POST"),
			lbrequest.WithJSONStringTemplateBody(createRoleTemplate, struct {
				ID     string
				Action string
			}{
				ID:     TimestampName(),
				Action: "system:serviceVersion:get",
			}),
		)).WithValue(&roleInfo)

	if err != nil {
		return nil, errors.Wrap(err, "Could not create role")
	}

	return &roleInfo, nil
}

// GetRole fetches the given role
func GetRole(tstCtx diagnostics.TestContext, id string) (*RoleInfo, error) {
	roleInfo := RoleInfo{}
	err := MustJSONDecodeSuccess(tstCtx.DoLBRequest(
		fmt.Sprintf("/apis/iam/v2/roles/%s", id),
	)).WithValue(&roleInfo)

	if err != nil {
		return nil, errors.Wrap(err, "Could not fetch role")
	}
	return &roleInfo, nil
}

// DeleteRole deletes the role with the given id
func DeleteRole(tstCtx diagnostics.TestContext, id string) error {
	err := MustJSONDecodeSuccess(
		tstCtx.DoLBRequest(
			fmt.Sprintf("/apis/iam/v2/roles/%s", id),
			lbrequest.WithMethod("DELETE"),
		)).Error()

	if err != nil {
		return errors.Wrap(err, "Could not delete role")
	}
	return nil
}

// CreateAuthRolesDiagnostic create the diagnostic struct for auth roles
func CreateAuthRolesDiagnostic() diagnostics.Diagnostic {
	return diagnostics.Diagnostic{
		Name: "auth-roles",
		Tags: diagnostics.Tags{"auth"},
		Skip: func(tstCtx diagnostics.TestContext) (bool, string, error) {
			isV2, err := tstCtx.IsIAMV2()
			if err != nil {
				return false, "", err
			}
			return !isV2, "requires IAM v2", nil
		},
		Generate: func(tstCtx diagnostics.TestContext) error {
			roleInfo, err := CreateRandomRole(tstCtx)
			if err != nil {
				return err
			}

			tstCtx.SetValue("auth-roles", &authRoleSave{
				ID:      roleInfo.Role.ID,
				Actions: roleInfo.Role.Actions,
			})
			return nil
		},
		Verify: func(tstCtx diagnostics.VerificationTestContext) {
			loaded := authRoleSave{}
			err := tstCtx.GetValue("auth-roles", &loaded)
			require.NoError(tstCtx, err, "Could not find generated context")

			roleInfo, err := GetRole(tstCtx, loaded.ID)
			require.NoError(tstCtx, err)

			assert.Equal(tstCtx, loaded.ID, roleInfo.Role.ID)
			assert.Equal(tstCtx, loaded.Actions, roleInfo.Role.Actions)
		},
		Cleanup: func(tstCtx diagnostics.TestContext) error {
			loaded := authRoleSave{}
			err := tstCtx.GetValue("auth-roles", &loaded)
			if err != nil {
				return errors.Wrap(err, "Could not find generated context")
			}

			return DeleteRole(tstCtx, loaded.ID)
		},
	}
}

func init() {
	diagnostics.RegisterDiagnostic(
		CreateAuthRolesDiagnostic(),
	)
}
