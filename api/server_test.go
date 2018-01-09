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

func TestReadEntities(t *testing.T) {

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/%s", ts.URL, "api/v1/entities"))
	body, _ := ioutil.ReadAll(res.Body)

	entities := entity.NewEntities(body)

	assert.Equal(t, "abuse@domain.tld", entities.List[0].Abuse)
}
