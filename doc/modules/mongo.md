# mongo

this is the main part of the api and yulmails configuration. In order to be able to mock the database, it means, call the database without actual calls, we need to split our struct into differents layer in order to inject the good dependence following the database context.

Example for the collection interface: 

we have the following interface: 

```go
type Collection interface {
	Count() (int, error)
	Find(interface{}) Query
	Insert(...interface{}) error
	Remove(interface{}) error
	Update(interface{}, interface{}) error
}
```

it means: each time we are using a `Collection` we need to pass to the caller, a struct who implements each of these functions. So it could be a classical `Collection` struct from `mongo` library, or our custom implementation (the mock).

we can try to implement the method `Count`, we need to create a method who counts the number of elements following the mongo query. Why not something like that: 

```go
func (mc MockCollection) Count() (n int, err error) {
	return 10, nil
}
```

so in our tests, we are using the MockMongo, so each time we'll call the `Count` method, we'll get a result of 10 and no error. After, you can customize this method to create errors, or custom behavior to verify how your code answer.

```go
func (mc MockCollection) Count() (n int, err error) {
    if toto {
    	return 10, nil
    }
    return 0, errors.New("not defined")
}
```

once we've setted up this way of coding, we can implement the mongo API: why we need to do this ? We don't want to manipulate mongo himself, we want to be able to request elements (ex: `ReadEntity`) directly, without handle errors etc from the database. (DAO).

example: 

we have have the following [method](https://github.com/yulPa/yulmails/blob/develop/pkg/mongo/mongo.go#L251)

```go
func (md MongoDatabase) DeleteEntity(entName string) error {
	/*
		Delete one entity from DB
		parameter: <string> Entity name
		return: <error> Nil if no error
	*/

	// We need first to remove all associated environments
	envs, err := md.ReadEnvironments(entName)
	if err != nil {
		return err
	}

	for _, env := range envs {
		err = md.DeleteEnvironment(entName, env.Name)
	}

	colEntity := md.C("entity")

	err = colEntity.Remove(bson.M{"name": entName})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
```

as you can see, the business logic is inside the method, so we can now just create a `MongoDatabase` struct and call our mothed to delete an entity !