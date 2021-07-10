# go-commons - turbo

The go-commons `http` package provides enterprise grade http routing capabilities.

###Common Usage
```go
func main() {
	router := turbo.RegisterTurboEngine()
    router.
    }RegisterTurboRoute("/api/v1/healthCheck", healthCheck)
    router.RegisterTurboRoute("/api/v1/getData", getData)
    
    srv := &http.Server{
        Handler:        router,
        Addr:           ":8080",
        ReadTimeout:    20 * time.Second,
        WriteTimeout:   20 * time.Second,
    }
    
    if err := srv.ListenAndServe(); err != nil {
        log.Fatalln(err)
    }
}
```

###Functionalities exposed
* RegisterTurboRoute - Registers the new route in the HTTP Server for the API
  * Methods - Function helps in defining the respective supported methods required by the API
    FYI: the HTTP methods are case in sensitive, so the methods can be added in any form like below
    ```go
    func main() {
        router := turbo.RegisterTurboEngine()
        router.RegisterTurboRoute("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("hello from turbo"))
        } ).StoreTurboMethod("get", "Post")
    }
    ```
* GetRoutes - Returns the List of all Registered Routes in the Server