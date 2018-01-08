package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yulPa/yulmails/entity"
)

var router = GetRouterV1()

func TestGetEntitys(t *testing.T) {

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity"))
	body, _ := ioutil.ReadAll(res.Body)

	entity := entity.NewEntity(body)

	assert.Equal(t, "abuse@domain.tld", entity.Abuse)
}
