package pb_migrations

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// current time in format of "2006-01-02 15:04:05.000Z"
		currentTime := time.Now().Format("2006-01-02 15:04:05.000Z")

		jsonData := fmt.Sprintf(`
			{
				"id": "f7bxe0q6tc2vyto",
				"created": "%s",
				"updated": "%s",
				"name": "locations",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "fqrzj6aa",
						"name": "pk",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": 26,
							"max": 26,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pc5msjrf",
						"name": "bjdNumber",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 10,
							"max": 10,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "l0eprkzj",
						"name": "sggNumber",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 5,
							"max": 5,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "dcq5h04r",
						"name": "emdNumber",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 3,
							"max": 3,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "rkycpj87",
						"name": "roadNumber",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 7,
							"max": 7,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "w9jpmyz0",
						"name": "undergroundFlag",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 1,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "fbh1lqta",
						"name": "buildingMainNumber",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 5,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "yokjizye",
						"name": "buildingSubNumber",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 5,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "zchn6qir",
						"name": "sdName",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 40,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "cowms54q",
						"name": "sggName",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 40,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "bkq81wbx",
						"name": "emdName",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 40,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "aroycqog",
						"name": "roadName",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 40,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pazmavre",
						"name": "buildingName",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 40,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "3yzmkktl",
						"name": "postalNumber",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": 5,
							"max": 5,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "9f3yakse",
						"name": "long",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": -180,
							"max": 180,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "kjf6nz7r",
						"name": "lat",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": -90,
							"max": 90,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "koopfhjp",
						"name": "crs",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "mc7gvle3",
						"name": "x",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "az3ds8yv",
						"name": "y",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "gnneqzvb",
						"name": "validPosition",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "vymn3eoz",
						"name": "baseDate",
						"type": "date",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "aagfet6z",
						"name": "address",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": 100,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX `+"`"+`idx_location_pk`+"`"+` ON `+"`"+`location`+"`"+` (`+"`"+`pk`+"`"+`)",
					"CREATE INDEX `+"`"+`idx_location_address`+"`"+` ON `+"`"+`location`+"`"+` (`+"`"+`address`+"`"+`)"
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "pcbamjn7je93c0y",
				"created": "2024-02-17 04:30:33.683Z",
				"updated": "2024-02-17 05:34:38.507Z",
				"name": "location_summary",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "sowprvw4",
						"name": "address",
						"type": "text",
						"required": true,
						"presentable": true,
						"unique": false,
						"options": {
							"min": null,
							"max": 100,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "sh2p47xw",
						"name": "sggName",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 40,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "dtfpcvdu",
						"name": "emdName",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 40,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "zwta9jqc",
						"name": "roadName",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 40,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "dvwp8ldh",
						"name": "avg_lat",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					},
					{
						"system": false,
						"id": "7yxxpr8z",
						"name": "avg_long",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					},
					{
						"system": false,
						"id": "9cihvn6m",
						"name": "avg_x",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					},
					{
						"system": false,
						"id": "nxc9qxqw",
						"name": "avg_y",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					},
					{
						"system": false,
						"id": "jhieuweq",
						"name": "postalNumbers",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 1
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX `+"`"+`idx_location_summary_index`+"`"+` ON `+"`"+`location_summary`+"`"+` (`+"`"+`address`+"`"+`)"
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`, currentTime, currentTime)

		var collections []*models.Collection
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
