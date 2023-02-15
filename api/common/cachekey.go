package common

type CacheKey struct {
	TTL int64 `json:"ttl"`
}

var (
	CacheKeyList = map[string]CacheKey{
		"user_reg_key": {
			TTL: int64(30 * 60),
		},
	}
)
