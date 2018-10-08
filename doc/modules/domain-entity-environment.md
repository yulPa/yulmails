# domain, entity and environment

each structure has his associated JSON format. 

example: 

```go
type Entity struct {
	Name    string          `json:"name"`
	Abuse   string          `json:"abuse"`
	Options options.Options `json:"options"`
}
```

the associated JSON will be: 

```json
{
    "name": "toto",
    "abuse": "toto@tld",
    "options": {
        "quota": ...,
        "conservation": ...
    }
}
```

You can easily create an entity or a list of entities from a JSON format: 

```go
entites := entity.NewEntities(`JSON SOURCE`)
```

It's the same thing for all the structs in this project. 