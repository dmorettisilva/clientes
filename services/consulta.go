package services

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) ConsultaClienteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cnpj := vars["idpessoa"]

	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(cnpj)) ///acertar para que aqui retorne um json da consulta sql no banco de dados
}
