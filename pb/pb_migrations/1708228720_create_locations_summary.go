package pb_migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		collection := &models.Collection{
			Name: "locations_summary",
			Type: models.CollectionTypeBase,
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:        "address",
					Type:        schema.FieldTypeText,
					Required:    true,
					Options:     &schema.TextOptions{Max: types.Pointer(100)},
					Presentable: true,
				},
				&schema.SchemaField{
					Name:        "addressGroup",
					Type:        schema.FieldTypeText,
					Required:    true,
					Presentable: true,
					Options: &schema.TextOptions{
						Max: types.Pointer(100),
					},
				},
				&schema.SchemaField{
					Name:        "roadNameGroup",
					Type:        schema.FieldTypeText,
					Required:    true,
					Presentable: true,
					Options: &schema.TextOptions{
						Max: types.Pointer(100),
					},
				},
				&schema.SchemaField{
					Name:    "lat",
					Type:    schema.FieldTypeNumber,
					Options: &schema.NumberOptions{Min: types.Pointer(-90.0), Max: types.Pointer(90.0)},
				},
				&schema.SchemaField{
					Name:    "long",
					Type:    schema.FieldTypeNumber,
					Options: &schema.NumberOptions{Min: types.Pointer(-180.0), Max: types.Pointer(180.0)},
				},
				&schema.SchemaField{
					Name: "x",
					Type: schema.FieldTypeNumber,
				},
				&schema.SchemaField{
					Name: "y",
					Type: schema.FieldTypeNumber,
				},
			),
			Indexes: types.JsonArray[string]{
				"CREATE UNIQUE INDEX `idx_locations_summary_address` ON `locations_summary` (`address`)",
				"CREATE INDEX `idx_locations_summary_addressGroup_roadNameGroup` ON `locations_summary` (`addressGroup`, `roadNameGroup`)",
			},
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("locations_summary")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
