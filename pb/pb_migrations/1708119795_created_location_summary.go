package pb_migrations

import (
	"encoding/json"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
	"github.com/sjunepark/go-gis/internal/database"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		err := database.CheckForDuplicateAddr(db)
		if err != nil {
			return err
		}

		jsonData := `{
			"id": "gbv276rdkf1txhd",
			"created": "2024-02-16 21:43:15.271Z",
			"updated": "2024-02-16 21:43:15.271Z",
			"name": "location_summary",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "4zt8nqkv",
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
					"id": "yw6yhhel",
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
					"id": "pijxj7wu",
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
					"id": "yrec8a0x",
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
					"id": "jtokfrd0",
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
					"id": "wutkyzns",
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
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT\n    (ROW_NUMBER() OVER (ORDER BY address)) AS id,\n    address,\n    sggName,\n    emdName,\n    roadName,\n    AVG(lat) AS avg_lat,\n    AVG(long) AS avg_long,\n    AVG(x) AS avg_x,\n    AVG(y) AS avg_y,\n    GROUP_CONCAT(DISTINCT postalNumber) AS postalNumbers\nFROM\n    location\nWHERE\n    validPosition = 1\nGROUP BY\n    address;"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("gbv276rdkf1txhd")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
