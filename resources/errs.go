package resources

import (
	"errors"
)

var ErrInvalidResourceName = errors.New("invalid ResourceName name. Resource name must start with type of ResourceName. E.g postgres_users, redis_cart ")
