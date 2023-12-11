// DO NOT EDIT LOCAL SDK - USE v3 okta-sdk-golang FOR API CALLS THAT DO NOT EXIST IN LOCAL SDK
package sdk

import (
	"context"
	"fmt"
	"net/http"
)

// FIXME calling undocumented public API
func (m *APISupplement) UpdateUserFactor(ctx context.Context, userId, factorId string, factorInstance Factor) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v", userId, factorId)
	req, err := m.RequestExecutor.
		WithAccept("application/json").
		WithContentType("application/json").
		NewRequest(http.MethodPut, url, factorInstance)
	if err != nil {
		return nil, err
	}
	return m.RequestExecutor.Do(ctx, req, factorInstance)
}
