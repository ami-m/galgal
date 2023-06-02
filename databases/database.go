package main

import (
	"dropit/databases/migrations"
	"dropit/databases/seeds"
	"flag"
)

func main() {
	var action string
	var shouldRunSeeds bool

	flag.StringVar(&action, "migrationAction", "", "The action should applied upon migrations create/destory")
	flag.BoolVar(&shouldRunSeeds, "shouldRunSeeds", false, "The action should applied upon migrations create/destory")
	flag.Parse()

	if action != "" {
		migrations.Run(action)
	}

	if shouldRunSeeds && action != migrations.DESTROY {
		seeds.Run()
	}
}
