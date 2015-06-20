package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgutz/dat.v1"
)

func TestRealJSON(t *testing.T) {
	j, _ := dat.NewJSON([]int{1, 2, 3})
	var num int
	testDB.SQL("select $1->1", j).QueryScalar(&num)
	assert.Equal(t, 2, num)
}

func TestRealJSONInterpolated(t *testing.T) {
	j, _ := dat.NewJSON([]int{1, 2, 3})
	var num int
	testDB.SQL("select $1->1", j).SetIsInterpolated(true).QueryScalar(&num)
	assert.Equal(t, 2, num)
}
