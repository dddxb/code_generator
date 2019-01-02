package biz

import (
	"code_generator/service/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {

	//正常插入日志
	create := model.Create{
		Address: "D:/code/go/src",
		Name:    "autoCreate",
	}

	assert.Equal(t, nil, Create(create))

}
