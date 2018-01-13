package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yulPa/yulmails/entity"
	"github.com/yulPa/yulmails/environment"
	"github.com/yulPa/yulmails/mocks"
)

func TestReadEntities(t *testing.T) {

	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/%s", ts.URL, "api/v1/entities"))
	fmt.Println(res)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(body)

	entities := entity.NewEntities(body)

	assert.Equal(t, "abuse@domain.tld", entities[0].Abuse)
}

func TestCreateANewEnvironment(t *testing.T) {

	var sess = mocks.NewMockSession()
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

func TestCreateANewEnvironmentWithMissingArg(t *testing.T) {
	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	resp, _ := http.Post(
		fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity/an_entity/environment"),
		"application/json",
		bytes.NewBuffer([]byte(`
			{
			 aefs
			}
			`)),
	)
	assert.Equal(t, resp.StatusCode, 500)
}

func TestCreateANewEntity(t *testing.T) {

	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	resp, _ := http.Post(
		fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity"),
		"application/json",
		bytes.NewBuffer([]byte(`
			{
			  "name": "an_entity",
			  "abuse": "abuse@domain.tld",
			  "options": {
			    "conservation":{
			      "sent": 5,
			      "unsent": 2,
			      "keep": true
			    }
			  }
			}
			`)),
	)
	assert.Equal(t, resp.StatusCode, 201)
}

func TestReadOneEntity(t *testing.T) {

	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity/an_entity"))
	body, _ := ioutil.ReadAll(res.Body)

	ent, _ := entity.NewEntity(body)
	assert.Equal(t, ent.Name, "an_entity")

}

func TestReadANonExistingEntity(t *testing.T) {

	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity/an_bad_entity"))

	assert.Equal(t, res.StatusCode, 500)
}

func TestCreateANewEntityWithMissingArg(t *testing.T) {

	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	resp, _ := http.Post(
		fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity"),
		"application/json",
		bytes.NewBuffer([]byte(`
			{
				sdfaec
			}
			`)),
	)
	assert.Equal(t, resp.StatusCode, 500)
}

func TestReadOneEnvironment(t *testing.T) {
	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity/an_entity/environment/an_environment"))
	body, _ := ioutil.ReadAll(res.Body)

	env, _ := environment.NewEnvironment(body)
	assert.Equal(t, env.Name, "an_environment")
	assert.Equal(t, env.EntityId, "an_entity")
}

func TestReadOneNonExistingEnvironment(t *testing.T) {
	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity/an_entity/environment/toto"))
	assert.Equal(t, res.StatusCode, 500)
}

func TestDeleteAnEntity(t *testing.T) {
	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity/an_entity"), nil)
	res, _ := http.DefaultClient.Do(req)
	assert.Equal(t, res.StatusCode, 200)

}

func TestDeleteANonExistingEntity(t *testing.T) {
	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity/yolo"), nil)
	res, _ := http.DefaultClient.Do(req)
	assert.Equal(t, res.StatusCode, 500)

}

func TestUpdateAnEntity(t *testing.T) {
	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity/an_entity"), nil)
	res, _ := http.DefaultClient.Do(req)
	assert.Equal(t, res.StatusCode, 200)
}

func TestUpdateANonExistingEntity(t *testing.T) {
	var sess = mocks.NewMockSession()
	var router = GetRouterV1(sess)
	var ts = httptest.NewServer(router)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", ts.URL, "api/v1/entity/hello"), nil)
	res, _ := http.DefaultClient.Do(req)
	assert.Equal(t, res.StatusCode, 500)
}
