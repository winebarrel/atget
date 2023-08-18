package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"github.com/winebarrel/atget"
)

var version string

var cli struct {
	Token   string `env:"AIRTABLE_TOKEN" required:"" help:"Airtable API token."`
	Base    string `short:"b" required:"" help:"Base name."`
	Table   string `short:"t" required:"" help:"Table name."`
	Filter  string `arg:"" name:"filter" required:"" help:"A formula used to filter records. see https://support.airtable.com/docs/formula-field-reference"`
	Version kong.VersionFlag
}

func init() {
	log.SetFlags(0)
}

func main() {
	kong.Parse(
		&cli,
		kong.Vars{"version": version},
	)

	client := atget.NewClient(cli.Token)
	records, err := client.GetRecords(cli.Base, cli.Table, cli.Filter)

	if err != nil {
		log.Fatal(err)
	}

	for _, r := range records.Records {
		fields, err := json.Marshal(r.Fields)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(fields))
	}
}
