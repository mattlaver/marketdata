package marketdata

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "Ping",
        "GET",
        "/ping",
        Ping,
    },
    Route{
        "Companies",
        "GET",
        "/companies",
        Companies,
    },
    Route{
        "TodoShow",
        "GET",
        "/todos/{todoId}",
        TodoShow,
    },
    Route{
    "TodoCreate",
    "POST",
    "/todos",
    TodoCreate,
},
}