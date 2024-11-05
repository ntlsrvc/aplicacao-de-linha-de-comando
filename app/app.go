package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Retorna a aplicaçao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "Aplicaçao de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca Ips de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
