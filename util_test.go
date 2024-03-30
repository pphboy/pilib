package pglib_test

import (
	"pglib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFileExist(t *testing.T) {
	a := assert.New(t)

	a.Equal(true, pglib.IsFileExist("pack.go"), "./pack.go 文件不存在")
	a.Equal(true, pglib.IsDirExist("/usr"), "dir 不存在")
}
