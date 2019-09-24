package task

import (
	"github.com/Xuanwo/navvy"

	"github.com/yunify/qsctl/v2/task/common"
)

// NewMakeBucketTask will create a make bucket task.
func NewMakeBucketTask(fn func(t *MakeBucketTask)) *MakeBucketTask {
	t := &MakeBucketTask{}

	pool, err := navvy.NewPool(10)
	if err != nil {
		panic(err)
	}
	t.SetPool(pool)

	fn(t)
	t.AddTODOs(common.NewBucketCreateTask)
	return t
}
