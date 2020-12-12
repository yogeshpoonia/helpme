package helpme

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"
	"regexp
	
)

// NewRouter helps to get new router instance.
func getMeRouter() *Router {
	return &Router{namedRoutes: make(map[string]*Route)}
}

type contextKey int

const (
	varsKey contextKey = iota
	routeKey
)

// Vars returns the route variables for the current request, if any.
func routerVar(r *http.Request) map[string]string {
	if rv := r.Context().Value(varsKey); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}
