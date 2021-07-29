package service

import (
	"net/http"
	"testing"
)

func TestCreateRole(t *testing.T) {
	router := createTestRouter()

	// Check role doesn't exist
	NewResponseTester("GET", "http://example.com/v2/roles/testCreateRole").WithStatus(http.StatusNotFound).Test(t, router)

	// Create role
	NewResponseTester("PUT", "http://example.com/v2/roles/testCreateRole").WithStatus(http.StatusOK).Test(t, router)

	// Expect role exists
	NewResponseTester("GET", "http://example.com/v2/roles/testCreateRole").WithStatus(http.StatusOK).Test(t, router)
}

func TestDeleteRole(t *testing.T) {
	router := createTestRouter()

	// Create role
	NewResponseTester("PUT", "http://example.com/v2/roles/testDeleteRole").WithStatus(http.StatusOK).Test(t, router)

	// Delete role
	NewResponseTester("DELETE", "http://example.com/v2/roles/testDeleteRole").WithStatus(http.StatusOK).Test(t, router)

	// Check role doesn't exist
	NewResponseTester("GET", "http://example.com/v2/roles/testDeleteRole").WithStatus(http.StatusNotFound).Test(t, router)
}

func TestGrantRolePermission(t *testing.T) {
	router := createTestRouter()

	// Create role
	NewResponseTester("PUT", "http://example.com/v2/roles/testGrantRolePermission").WithStatus(http.StatusOK).Test(t, router)

	// Grant permission
	NewResponseTester("PUT", "http://example.com/v2/roles/testGrantRolePermission/bundles/testbundle/permissions/testpermission").WithStatus(http.StatusOK).Test(t, router)
}

func TestGrantRolePermissionInvalidRole(t *testing.T) {
	router := createTestRouter()

	// Create role
	NewResponseTester("PUT", "http://example.com/v2/roles/testGrantRolePermissionInvalidRole").WithStatus(http.StatusOK).Test(t, router)

	// Grant permission to different role
	NewResponseTester("PUT", "http://example.com/v2/roles/testGrantRolePermissionInvalidRole2/bundles/testbundle/permissions/testpermission").WithStatus(http.StatusNotFound).Test(t, router)
}

func TestRevokeRolePermission(t *testing.T) {
	router := createTestRouter()

	// Create role
	NewResponseTester("PUT", "http://example.com/v2/roles/testRevokeRolePermission").WithStatus(http.StatusOK).Test(t, router)

	// Grant permission
	NewResponseTester("PUT", "http://example.com/v2/roles/testRevokeRolePermission/bundles/testbundle/permissions/testpermission").WithStatus(http.StatusOK).Test(t, router)

	// Revoke permission
	NewResponseTester("DELETE", "http://example.com/v2/roles/testRevokeRolePermission/bundles/testbundle/permissions/testpermission").WithStatus(http.StatusOK).Test(t, router)
}

func TestRevokeRolePermissionInvalidRole(t *testing.T) {
	router := createTestRouter()

	// Create role
	NewResponseTester("PUT", "http://example.com/v2/roles/testRevokeRolePermissionInvalidRole").WithStatus(http.StatusOK).Test(t, router)

	// Grant permission
	NewResponseTester("PUT", "http://example.com/v2/roles/testRevokeRolePermissionInvalidRole/bundles/testbundle/permissions/testpermission").WithStatus(http.StatusOK).Test(t, router)

	// Revoke permission from different role
	NewResponseTester("DELETE", "http://example.com/v2/roles/testRevokeRolePermissionInvalidRole2/bundles/testbundle/permissions/testpermission").WithStatus(http.StatusNotFound).Test(t, router)
}

func TestRevokeRolePermissionInvalidPermission(t *testing.T) {
	router := createTestRouter()

	// Create role
	NewResponseTester("PUT", "http://example.com/v2/roles/testRevokeRolePermissionInvalidPermission").WithStatus(http.StatusOK).Test(t, router)

	// Grant permission
	NewResponseTester("PUT", "http://example.com/v2/roles/testRevokeRolePermissionInvalidPermission/bundles/testbundle/permissions/testpermission").WithStatus(http.StatusOK).Test(t, router)

	// Revoke permission that doesn't exist, will be ignored and return 200
	// NOTE: Should we check for this? There may be implications with versioning.
	NewResponseTester("DELETE", "http://example.com/v2/roles/testRevokeRolePermissionInvalidPermission/bundles/testbundle/permissions/testpermission2").WithStatus(http.StatusOK).Test(t, router)
}
