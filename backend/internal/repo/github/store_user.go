package github

import (
	"context"
	"fmt"

	"github.com/abdulsalam/pixel8labs/internal/constant"
)

func (u *UserRepo) SetVisitorByID(ctx context.Context, id int64) (uint64, error) {
	var err error

	key := fmt.Sprintf(constant.CacheUserByID, id)
	total, err := u.GetVisitorByID(ctx, id)
	if err != nil {
		total = 0
	}

	total++
	err = u.bigCache.Set(key, []byte(fmt.Sprintf("%d", total)))
	return total, err
}
