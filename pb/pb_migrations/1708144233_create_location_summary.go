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
		currentTime := time.Now().Format("2006-01-02 15:04:05.000Z")

		jsonData := fmt.Sprintf(`
		{
				"id": "pcbamjn7je93c0y",
				"created": "%s",
				"updated": "%s",
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
			}`, currentTime, currentTime)

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("pcbamjn7je93c0y")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
