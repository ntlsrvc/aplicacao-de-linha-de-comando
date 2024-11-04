package app

import "github.com/urfave/cli"

// Retorna a aplicaçao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "Aplicaçao de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	return app
}
