package common

import (
	"os"
	"testing"

	"github.com/Xuanwo/navvy"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/yunify/qsctl/v2/pkg/types"
	"github.com/yunify/qsctl/v2/storage"
	"github.com/yunify/qsctl/v2/utils"
)

func TestMultipartInitTask_Run(t *testing.T) {
	x := &mockMultipartInitTask{}

	store := storage.NewMockObjectStorage()
	x.SetStorage(store)

	pool, err := navvy.NewPool(10)
	if err != nil {
		t.Fatal(err)
	}
	x.SetPool(pool)

	key := uuid.New().String()
	x.SetKey(key)

	offset := int64(0)
	x.SetCurrentOffset(&offset)
	x.SetTotalSize(1024)

	id := uuid.New().String()

	fn := func(task types.Todoist) navvy.Task {
		s := int64(1024)
		x.SetCurrentOffset(&s)

		t := &utils.EmptyTask{}
		t.SetID(id)
		t.SetPool(pool)
		return t
	}
	x.SetScheduler(types.NewScheduler(fn))

	task := NewMultipartInitTask(x)
	task.Run()

	// There must be only one task in wg, so the first should be ok, and the next should panic.
	assert.NotPanics(t, func() {
		x.GetScheduler().Done(id)
	})
	assert.Panics(t, func() {
		x.GetScheduler().Done(id)
	})
}

func TestMultipartFileUploadTask_Run(t *testing.T) {
	x := &mockMultipartFileUploadTask{}

	store := storage.NewMockObjectStorage()
	x.SetStorage(store)
	key := uuid.New().String()
	x.SetKey(key)
	uploadID, err := store.InitiateMultipartUpload(key)
	if err != nil {
		t.Fatal(err)
	}
	x.SetUploadID(uploadID)

	name, size, md5sum := utils.GenerateTestFile()
	defer os.Remove(name)

	x.SetPath(name)
	x.SetOffset(0)
	x.SetSize(size)
	x.SetPartNumber(0)
	x.SetMD5Sum(md5sum)
	x.SetID(uuid.New().String())

	sche := types.NewMockScheduler(nil)
	sche.New(nil)
	x.SetScheduler(sche)

	task := NewMultipartFileUploadTask(x)
	task.Run()

	err = store.CompleteMultipartUpload(key, uploadID, 1)
	assert.NoError(t, err)

	om, err := store.HeadObject(key)
	assert.NoError(t, err)
	assert.Equal(t, size, om.ContentLength)
}

func TestMultipartStreamUploadTask_Run(t *testing.T) {
	x := &mockMultipartStreamUploadTask{}

	store := storage.NewMockObjectStorage()
	x.SetStorage(store)
	key := uuid.New().String()
	x.SetKey(key)
	uploadID, err := store.InitiateMultipartUpload(key)
	if err != nil {
		t.Fatal(err)
	}
	x.SetUploadID(uploadID)

	buf, size, md5sum := utils.GenerateTestStream()

	x.SetSize(size)
	x.SetPartNumber(0)
	x.SetContent(buf)
	x.SetMD5Sum(md5sum)
	x.SetID(uuid.New().String())

	sche := types.NewMockScheduler(nil)
	sche.New(nil)
	x.SetScheduler(sche)

	task := NewMultipartStreamUploadTask(x)
	task.Run()

	err = store.CompleteMultipartUpload(key, uploadID, 1)
	assert.NoError(t, err)

	om, err := store.HeadObject(key)
	assert.NoError(t, err)
	assert.Equal(t, size, om.ContentLength)
}

func TestMultipartCompleteTask_Run(t *testing.T) {
	x := &mockMultipartCompleteTask{}

	store := storage.NewMockObjectStorage()
	x.SetStorage(store)
	key := uuid.New().String()
	x.SetKey(key)
	uploadID, err := store.InitiateMultipartUpload(key)
	if err != nil {
		t.Fatal(err)
	}
	x.SetUploadID(uploadID)

	buf, size, md5sum := utils.GenerateTestStream()
	err = store.UploadMultipart(key, uploadID, size, 0, md5sum, buf)
	if err != nil {
		t.Fatal(err)
	}

	partNumber := int32(1)
	x.SetCurrentPartNumber(&partNumber)

	task := NewMultipartCompleteTask(x)
	task.Run()

	om, err := store.HeadObject(key)
	assert.NoError(t, err)
	assert.Equal(t, size, om.ContentLength)
}