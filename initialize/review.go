package initialize

import (
	"github.com/LINQQ1212/common2/global"
	"github.com/LINQQ1212/common2/review"
	"path"
)

func InitReview(f string) (*review.Review, error) {
	return review.NewReview(path.Join(global.CONFIG.System.MainDir, "review.db"), path.Join(global.CONFIG.System.MainDir, f))
}
