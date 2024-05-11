package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateAttahcment(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error ConnectToPq", err)
	}
	as := NewAttachmentStore(db)
	param := types.Attachment{
		CompanyId: 2,
		UserId:    4,
		FileName:  "testname.png",
		MimeType:  "image/x-png",
	}
	attch, err := as.Create(param)
	if err != nil {
		t.Error("error creating Attachment", err)
	}
	defer as.DeleteById(attch.ID)

	assert.Greater(t, attch.ID, 0)
	assert.Greater(t, attch.CreatedAt, time.Time{})
}

func TestUpdateAttahcmentByid(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error ConnectToPq", err)
	}
	as := NewAttachmentStore(db)

	attch, err := createAttachment()
	if err != nil {
		t.Error("error createAttachment in TestUpdateAttahcmentByid", err)
	}
	defer as.DeleteById(attch.ID)

	param := types.Attachment{
		FileName: "testnamee.png",
		MimeType: "imagee/x-png",
	}
	if err := as.UpdateByid(param, attch.ID); err != nil {
		t.Error("error Updating Attachment", err)
	}
}

func TestGetById(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error ConnectToPq", err)
	}
	as := NewAttachmentStore(db)
	attch, err := createAttachment()
	if err != nil {
		t.Error("error createAttachment in TestUpdateAttahcmentByid", err)
	}
	defer as.DeleteById(attch.ID)

	expected, err := as.GetById(attch.ID)
	if err != nil {
		t.Error("error Getting Attachment by id", err)
	}
	assert.Equal(t, expected.ID, attch.ID)
}

// helping func
func createAttachment() (types.Attachment, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return types.Attachment{}, nil
	}
	as := NewAttachmentStore(db)
	param := types.Attachment{
		CompanyId: 2,
		UserId:    4,
		FileName:  "testname.png",
		MimeType:  "image/x-png",
	}
	attch, err := as.Create(param)
	if err != nil {
		return types.Attachment{}, nil
	}

	return attch, nil
}
