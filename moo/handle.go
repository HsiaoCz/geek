package moo

import "net/http"

type Handlefunc func(w http.ResponseWriter, r *http.Request)