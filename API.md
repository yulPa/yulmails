## Entity

  * [Create](https://github.com/yulPa/yulmails/master/API.md#create-an-entity)
  * [Update](https://github.com/yulPa/yulmails/master/API.md#update-entity)
  * [Delete](https://github.com/yulPa/yulmails/master/API.md#delete-entity)
  * [Read](https://github.com/yulPa/yulmails/master/API.md#read-an-entity)
  * [Get entities](https://github.com/yulPa/yulmails/master/API.md#get-entities)

## Environment

  * [Create](https://github.com/yulPa/yulmails/master/API.md#create-an-environment)
  * [Read](https://github.com/yulPa/yulmails/master/API.md#read-an-environment)
  * [Delete](https://github.com/yulPa/yulmails/master/API.md#delete-environment)


## Create an Entity

	Create a new Entity in order to define your infrastructure.

  * __URL__

		/api/v1/entity

  * __Method__:

  	`POST`

  * __URL Params__:

		__Required__:

			```
			name=[string]
			abuse=[string]
			sent=[int]
			unsent=[int]
			```

		__Optional__:

			`keep=[boolean]`

  * __Data Params__:

    ```json
    {
      "name": "An entity",
      "abuse": "abuse@domain.tld",
      "options": {
        "conservation":{
          "sent": 5,
          "unsent": 2,
          "keep": true
        }
      }
    }
    ```

  * __Success Response__:

		* __Code__: 200


  * __Error Response__:

		* __Code__: 500

## Get entities

	Read all entity created

  * __URL__

		/api/v1/entities

  * __Method__:

  	`GET`

  * __Success Response__:

		* __Code__: 200
		* __Content__:

			```json
			[
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
				},
				{
					"name": "another_entity",
					"abuse": "another_abuse@domain.tld",
					"options": {
						"conservation":{
							"sent": 5,
							"unsent": 3,
							"keep": true
						}
					}
				}
			]
				```

	* __Error Response__:

		* __Code__: 500

## Update entity

	Update a selected entity

  * __URL__;

		/api/v1/entity/{entity_name}

  * __Method__:

  	`POST`

  * __URL Params__:

			__Required__:

			```
			name=[string]
			abuse=[string]
			sent=[int]
			unsent=[int]
			```

			__Optional__:

			`keep=[boolean]`

 * __Data Params__:

  ```json
  {
    "name": "An entity",
    "abuse": "abuse@domain.tld",
    "options": {
      "conservation":{
        "sent": 5,
        "unsent": 2,
        "keep": true
      }
    }
  }
  ```


  * __Success Response__:

		* __Code__: 200

  * __Error Response__:

		* __Code__: 500
		* __Content__:

			```json
			{"error": "not found"}
			```

## Delete entity

  Delete a selected entity

  * __URL__:

		/api/v1/entity/{an_entity}

  * __Method__:

  	`DELETE`

  * __Success Response__:

		* __Code__: 200

  * __Error Response__:

		* __Code__: 500

## Create an environment

  Add an environment to an existing entity

  * __URL__:

    /api/v1/entity/{entity_name}/environment

  * __Method__:

    `POST`

  * __URL Params__:

    __Required__:

    ```
    ips=[[]string]
    abuse=[string]
    open=[boolean]
    ```

    __Optional__:

    `quote=[object]`

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
        "options":{
          "quota": {
            "tenlastminutes": 150,
            "sixtylastminutes": 200,
            "lastday": 1000,
            "lastweek": 3000,
            "lastmonth": 10000
          }
        }
      }
      ```

    * __Success Response__:

      * __Code__: 200


    * __Error Response__:

      * __Code__: 500

## Delete environment

  Delete a selected environment

  * __URL__:

    /api/v1/entity/{entity_name}/environment/{environment_name}

  * __Method__:

    `GET`

  * __Success Response__:

    * __Code__: 200

  * __Error Response__:

    * __Code__: 500
    * __Content__:<

    ```json
    {"error": "not found"}
    ```
