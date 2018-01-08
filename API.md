## Create an Entity

Create a new Entity in order to define your infrastructure.

  * __URL__

  /api/v1/entity

  * __Method__:

  `POST`

  * __URL Params__:

  __Required__:

    name=[string]
    abuse=[string]
    sent=[int]
    unsent=[int]

  __Optional__:

    keep=[boolean]

  * __Data Params__:

    ```json
    {
      "name": "An entity",
      "abuse": "abuse@domain.tld",
      "conservation":{
        "sent": 5,
        "unsent": 2,
        "keep": true
      }
    }
    ```

  * __Success Response__:

    * __Code__: 200


  * __Error Response__:

    Not implemented

## Get entities list

Fetch all entity created

  * __URL__

  /api/v1/entity

  * __Method__:

  `GET`

  * __Success Response__:

    * __Code__: 200
      __Content__:
      ```json
      {
        "name": "An entity",
        "abuse": "abuse@domain.tld",
        "conservation":{
          "sent": 5,
          "unsent": 2,
          "keep": true
        }
      }
      ```

  * __Error Response__:

    Not implemented

## Create an environment

Add an environment to an existing entity

* __URL__

/api/v1/entity/{entity_name}/environment

* __Method__:

`POST`

* __URL Params__:

__Required__:

  ips=[[]string]
  abuse=[string]
  open=[boolean]

__Optional__:

  quote=[object]

* __Data Params__:

  ```json
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
  ```

* __Success Response__:

  * __Code__: 200


* __Error Response__:

  Not implemented
