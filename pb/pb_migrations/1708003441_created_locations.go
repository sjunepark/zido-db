package pb_migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "f7bxe0q6tc2vyto",
			"created": "2024-02-15 13:24:01.384Z",
			"updated": "2024-02-15 13:24:01.384Z",
			"name": "locations",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "pc5msjrf",
					"name": "bjdNumber",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 1000000000,
						"max": 9999999999,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "l0eprkzj",
					"name": "sggNumber",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 10000,
						"max": 99999,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "dcq5h04r",
					"name": "emdNumber",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 100,
						"max": 999,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "rkycpj87",
					"name": "roadNumber",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 1000000,
						"max": 9999999,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "w9jpmyz0",
					"name": "undergroundFlag",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 2,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "fbh1lqta",
					"name": "buildingMainNumber",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 1,
						"max": 99999,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "yokjizye",
					"name": "buildingSubNumber",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 99999,
						"noDecimal": true
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
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("f7bxe0q6tc2vyto")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
