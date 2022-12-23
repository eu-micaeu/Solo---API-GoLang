package router

import "github.com/gorilla/mux"

// Gerar vai reornar um router com as rotas configuradas
func Gerar() *mux.Router{
	return mux.NewRouter()
}