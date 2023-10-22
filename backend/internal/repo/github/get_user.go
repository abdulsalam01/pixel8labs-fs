package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/abdulsalam/pixel8labs/internal/constant"
)

func (u *UserRepo) GetVisitorByID(ctx context.Context, id int64) (uint64, error) {
	var (
		total uint64
		err   error
	)

	key := fmt.Sprintf(constant.CacheUserByID, id)
	data, err := u.bigCache.Get(key)
	if err != nil {
		return total, err
	}

	// Conversion factors.
	dataStr := string(data)
	dataInt, err := strconv.Atoi(dataStr)
	if err != nil {
		return total, err
	}

	total = uint64(dataInt) // Wrapped.
	return total, nil
}
