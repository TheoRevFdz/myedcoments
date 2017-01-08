package main

import (
	"CursoGolang/myedcoments/migration"
	"flag"
	"log"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la DB")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración...")
	}
}
