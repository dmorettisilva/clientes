package services

import (
	"net/http"

	"github.com/dmorettisilva/clientes/database"
	"github.com/gorilla/mux"
)

func (s *Service) ConsultaClienteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pessoa := vars["idpessoa"]

	j := database.ConsultaCliente(pessoa)

	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(j)) ///acertar para que aqui retorne um json da consulta sql no banco de dados
}
