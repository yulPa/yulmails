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
