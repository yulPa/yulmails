package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yulPa/yulmails/entity"
	"github.com/yulPa/yulmails/mongo"
)

func TestReadEntities(t *testing.T) {

	var sess = mongo.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/%s", ts.URL, "api/v1/entities"))
	body, _ := ioutil.ReadAll(res.Body)

	entities := entity.NewEntities(body)

	assert.Equal(t, "abuse@domain.tld", entities.List[0].Abuse)
}
