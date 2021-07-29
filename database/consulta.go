package database

import (
	"fmt"
	"net/http"
	"strconv"
)


func ConsultaCliente(id int) http.ResponseWriter, error {
/*
select pe.idpessoa,
       pe.razao,
       pe.fantasia,
       pe.fg_bloqueado,
       nvl(pf.cpf, pj.cgc) CPF_CNPJ,
       nvl(pf.rg, pj.insestadual) RG_IE
from gepessoas pe
left join gepessoafisica pf on pf.idpessoa = pe.idpessoa 
left join gepessoajuridica pj on pj.idpessoa = pe.idpessoa
*/

//Passar resultado do select para struct 

//montar json dentro de um response

w.Header().Set("Content-type", "application/json")
w.WriteHeader(errorCodeAtu)
w.Write([]byte(response))

return w

}

