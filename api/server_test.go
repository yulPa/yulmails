package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"

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

func TestCreateANewEnvironment(t *testing.T) {

	var sess = mongo.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	resp, _ := http.Post(
		fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity/an_entity/environment"),
		"application/json",
		bytes.NewBuffer([]byte(`
			{
			  "ips": [
			    "192.168.0.1",
			    "192.168.0.2",
			    "192.168.0.3"
			  ],
			  "abuse": "abuse@domain.tld",
			  "open": false,
			  "quota": {
			    "tenlastminutes": 150,
			    "sixtylastminutes": 200,
			    "lastday": 1000,
			    "lastweek": 3000,
			    "lastmonth": 10000
			  }
			}
			`)),
	)
	assert.Equal(t, resp.StatusCode, 201)
}
