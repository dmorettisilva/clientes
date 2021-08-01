package database

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dmorettisilva/clientes/config"
	_ "github.com/jmoiron/sqlx"
)

func ConsultaCliente(id string) []byte {

	sql := `select pe.idpessoa,
                   pe.razao,
                   pe.fantasia,
                   pe.fg_bloqueado,
                   nvl(pf.cpf, pj.cgc) CPF_CNPJ,
                   nvl(pf.rg, pj.insestadual) RG_IE
            from gepessoas pe
            left join gepessoafisica pf on pf.idpessoa = pe.idpessoa 
            left join gepessoajuridica pj on pj.idpessoa = pe.idpessoa
            where ( (%s = 0) or (%s = pe.idpessoa) )`

	//Passar resultado do select para struct

	config.New()
	var db *DAO

	if len(id) == 0 {
		fmt.Sprintf(sql, "0")
	} else {
		fmt.Sprintf(sql, id)
	}

	rows, err := db.db.DB.Query(sql)
	defer rows.Close()

	if err != nil {
		return nil
	}

	columns, _ := rows.Columns()
	count := len(columns)

	var v struct {
		Data []interface{}
	}

	for rows.Next() {
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			log.Fatal(err)
		}
		v.Data = append(v.Data, values)
	}
	jsonMsg, err := json.Marshal(v)
	return jsonMsg

}
