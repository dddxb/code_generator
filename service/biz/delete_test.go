package biz

import (
	"code_generator/service/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Delete(t *testing.T) {

	//正常登陆
	delete := model.Delete{
		Address: "D:/code/go/src",
		Name: "autoCreate",
	}


	assert.Equal(t, nil, Delete(delete))

}

