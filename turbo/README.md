# turbo (go-commons)

The go-commons `turbo` package provides enterprise grade http routing capabilities. The lightweight router works well with all the necessary Use Cases and at the same time scales well.

---

- [Installation](#installation)
- [Benchmarking Results](#benchmarking-results)
- [Quick Start Guide](#quick-start-guide)
- [Functionalities Exposed](#functionalities-exposed)

---

### Installation

```bash
go get go.appmanch.org/commons/turbo
```

### Benchmarking Results

```bash
To be released soon
```

### Quick Start Guide

Being a Lightweight HTTP Router, it comes with a simple usage as explained below, just import the package, and you are good to go.

```go
func main() {
	router := turbo.New()
	router.Get("/api/v1/healthCheck", healthCheck) // healthCheck is the handler Function
	router.Get("/api/v1/getData", getData) // getData is the handler Function
	
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

### Functionalities Exposed

1. Router lets you register routes based on the common HTTP Methods such as
    1. GET
    2. POST
    3. PUT
    4. DELETE
2. Routes can be registered in the following ways
    1. Registering Static Routes
        ```go
        router.Get("/api/v1/getCustomers", getCustomers) 
        ```

    2. Registering with Path Variables
       
       _The path variables need to be registered with their type upfront_
        ```go
        router.Get("/api/v1/getCustomer/:id:int32", getCustomer)
        ```

3. Path Params can be fetched with the built-in wrapper provided by the framework
    * The framework exposes a number of functions based on the type of variable that has been registered with the route
        * To fetch string parameters
            ```go
            getPathParms(id string, r *http.Request) string {}
            ```
        * To Fetch Int parameters
            ```go
            getIntPathParms(id string, r *http.Request) int {}
            ```
        * To fetch Float64 parameters
           ```go
           getFloatPathParms(id string, r *http.Request) float64 {}
           ```
        * To Fetch Boolean parameters
           ```go
           getBoolPathParms(id string, r *http.Request) bool {}
           ```
          
4. Query Parameters can also be fetched with a built-in wrapper functions provided by the framework
    * The Framework exposes a number of wrapper functions which lets you fetch the query params of specific data type required
        * To fetch string query params
            ```go
            GetQueryParams(id string, r *http.Request) string {}
            ```
        * To Fetch Int query params
            ```go
            GetIntQueryParams(id string, r *http.Request) int {}
            ```
        * To fetch Float64 query params
           ```go
           GetFloatQueryParams(id string, r *http.Request) float64 {}
           ```
        * To Fetch Boolean query params
           ```go
           GetBoolQueryParams(id string, r *http.Request) bool {}
           ```
            