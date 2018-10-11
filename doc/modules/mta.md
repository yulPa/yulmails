# mta

in order to make the first entrypoint easier, we decided to create a Mail Transfert Agent (MTA).

So this is the configuration of our MTA: 

```go
cfg = &guerrilla.AppConfig{
    LogFile:      "/tmp/guerrilla.log",
    AllowedHosts: []string{"."},
    BackendConfig: backends.BackendConfig{
        "save_process": "Hasher|RedisQueue",
        "redis_addr":   "redis:6379",
    },
    Servers: []guerrilla.ServerConfig{guerrilla.ServerConfig{
        ListenInterface: "0.0.0.0:25",
        IsEnabled:       true,
    }},
}
...
d.AddProcessor("RedisQueue", processor.RedisQueueProcessor)
```

As we can see, we are adding a custom backend processor `RedisQueue`. This custom processor, will load our fresh emails into a `redis` queue.

```go
return func(p backends.Processor) backends.Processor {
    return backends.ProcessWith(
        func(e *mail.Envelope, task backends.SelectTask) (backends.Result, error) {
            if task == backends.TaskSaveMail {
                connection := rmq.OpenConnection("emailsService", "tcp", config.Addr, 1)
                queue := connection.OpenQueue("emails")
                if ok := queue.Publish(e.Data.String()); !ok {
                    log.Println("unable to push email on queue")
                }
                return p.Process(e, task)
            }
            return p.Process(e, task)
        },
    )
}
```

So with this system, we are able to keep a trace of our emails in yulmails system. So we can read emails from the stack whenever we want and process them with anti-spam, etc. 