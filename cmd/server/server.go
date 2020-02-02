package main

import (
	"flag"
	"log"

	"github.com/andreylm/vm_test/pkg/server"
)

var (
	port     string
	host     string
	jsonFile string
	useDb    bool
	conn     string
)

func init() {
	flag.StringVar(&host, "host", "localhost", "Host")
	flag.StringVar(&port, "port", "8000", "Port")
	flag.StringVar(&jsonFile, "json-file", "", "Path to json file")
	flag.BoolVar(&useDb, "use-db", false, "true - use db, false - use cache")
	flag.StringVar(&conn, "connection-string", "file:ent?mode=memory&cache=shared&_fk=1", "Connection string")
	flag.Parse()

}

func main() {
	srv, err := server.NewServer(host, port, jsonFile, conn, useDb)
	if err != nil {
		log.Fatal(err)
	}
	log.Panic(srv.Start())
}
