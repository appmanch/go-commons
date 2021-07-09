# go-commons - turbo

The go-commons `http` package provides enterprise grade http routing capabilities.

Usage
```
router := turbo.RegisterTurboEngine()
router.RegisterTurboRoute("GET", "/api/v1/healthCheck", healthCheck)
router.RegisterTurboRoute("GET,PUT,POST", "/api/v1/getData", getData)

srv := &http.Server{
    Handler:        router,
    Addr:           ":8080",
    ReadTimeout:    20 * time.Second,
    WriteTimeout:   20 * time.Second,
}

if err := srv.ListenAndServe(); err != nil {
    log.Fatalln(err)
}
```