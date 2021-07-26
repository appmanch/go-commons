package turbo

//NewRouter Creates a new Router
func NewRouter() *Router {
	return &Router{
		unManagedRouteHandler:    nil,
		unsupportedMethodHandler: nil,
		topLevelRoutes:           make(map[string]*Route),
	}
}
