package atget

import (
	"fmt"

	"github.com/mehanizm/airtable"
)

type Client struct {
	airtable *airtable.Client
}

func NewClient(token string) *Client {
	client := airtable.NewClient(token)

	return &Client{
		airtable: client,
	}
}

func (client *Client) GetRecords(baseName string, tableName string, filter string) (*airtable.Records, error) {
	bases, err := client.airtable.GetBases().Do()

	if err != nil {
		return nil, err
	}

	var base *airtable.Base

	for _, b := range bases.Bases {
		if b.Name == baseName {
			base = b
			break
		}
	}

	if base == nil {
		return nil, fmt.Errorf("base '%s' not found", baseName)
	}

	table := client.airtable.GetTable(base.ID, tableName)

	if table == nil {
		return nil, fmt.Errorf("table '%s' not found", tableName)
	}

	records, err := table.GetRecords().WithFilterFormula(filter).Do()

	if err != nil {
		return nil, err
	}

	return records, nil
}
