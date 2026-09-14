package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type ThirdPartyRole struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoleClient interface {
	FetchRoles(ctx context.Context) ([]ThirdPartyRole, error)
}

type roleClient struct {
	httpClient *http.Client
	baseURL    string
	rdb        *redis.Client
}

func NewRoleClient(rdb *redis.Client) RoleClient {
	return &roleClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    "https://6aa7b6a19b08676cd32b76c6.mockapi.io/roles",
		rdb:        rdb,
	}
}

func (c *roleClient) FetchRoles(ctx context.Context) ([]ThirdPartyRole, error) {
	cacheKey := "roles:thirdparty"

	if c.rdb != nil {
		cachedData, err := c.rdb.Get(ctx, cacheKey).Result()
		if err == nil && cachedData != "" {
			var roles []ThirdPartyRole
			if err := json.Unmarshal([]byte(cachedData), &roles); err == nil {
				return roles, nil
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed status: %d", resp.StatusCode)
	}

	var roles []ThirdPartyRole
	if err := json.NewDecoder(resp.Body).Decode(&roles); err != nil {
		return nil, err
	}

	if c.rdb != nil {
		bytes, err := json.Marshal(roles)
		if err == nil {
			c.rdb.Set(ctx, cacheKey, string(bytes), 5*time.Minute)
		}
	}

	return roles, nil
}
