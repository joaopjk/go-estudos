package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios",
		Method:             http.MethodPost,
		Funcao:             controllers.CriarUsuarios,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios",
		Method:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Method:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Method:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Method:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
