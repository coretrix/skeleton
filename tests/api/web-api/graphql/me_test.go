package graphql

import (
	"coretrix/skeleton/tests"
	"testing"

	"github.com/tj/assert"
)

type queryMe struct {
	Me struct {
		ID       string
		Username string
		Age      int
	}
}

func TestQueryMe(t *testing.T) {
	got := &queryMe{}
	projectName, resolver := tests.GetWebAPIResolver()
	err := tests.CreateContextWebAPI(t, projectName, resolver, nil).HandleQuery(got, nil)

	assert.Nil(t, err)

	want := &queryMe{Me: struct {
		ID       string
		Username string
		Age      int
	}{ID: "1234", Username: "Me", Age: 69}}

	assert.Equal(t, want, got)
}
