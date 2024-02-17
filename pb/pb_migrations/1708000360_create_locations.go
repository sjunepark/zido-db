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
			Name: "locations",
			Type: models.CollectionTypeBase,
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:        "pk",
					Type:        schema.FieldTypeText,
					Required:    true,
					Presentable: true,
					Options: &schema.TextOptions{
						Min: types.Pointer(26),
						Max: types.Pointer(26),
					},
				},
				&schema.SchemaField{
					Name:     "bjdNumber",
					Type:     schema.FieldTypeText,
					Required: true,

					Options: &schema.TextOptions{
						Min: types.Pointer(10),
						Max: types.Pointer(10),
					},
				},
				&schema.SchemaField{
					Name:     "sggNumber",
					Type:     schema.FieldTypeText,
					Required: true,

					Options: &schema.TextOptions{
						Min: types.Pointer(5),
						Max: types.Pointer(5),
					},
				},
				&schema.SchemaField{
					Name:     "emdNumber",
					Type:     schema.FieldTypeText,
					Required: true,

					Options: &schema.TextOptions{
						Min: types.Pointer(3),
						Max: types.Pointer(3),
					},
				},
				&schema.SchemaField{
					Name:     "roadNumber",
					Type:     schema.FieldTypeText,
					Required: true,

					Options: &schema.TextOptions{
						Min: types.Pointer(7),
						Max: types.Pointer(7),
					},
				},
				&schema.SchemaField{
					Name:     "undergroundFlag",
					Type:     schema.FieldTypeText,
					Required: true,

					Options: &schema.TextOptions{
						Min: types.Pointer(1),
						Max: types.Pointer(1),
					},
				},
				&schema.SchemaField{
					Name:     "buildingMainNumber",
					Type:     schema.FieldTypeText,
					Required: true,

					Options: &schema.TextOptions{
						Min: types.Pointer(1),
						Max: types.Pointer(5),
					},
				},
				&schema.SchemaField{
					Name: "buildingSubNumber",
					Type: schema.FieldTypeText,

					Options: &schema.TextOptions{
						Min: types.Pointer(1),
						Max: types.Pointer(5),
					},
				},
				&schema.SchemaField{
					Name:     "sdName",
					Type:     schema.FieldTypeText,
					Required: true,

					Options: &schema.TextOptions{
						Max: types.Pointer(40),
					},
				},
				&schema.SchemaField{
					Name: "sggName",
					Type: schema.FieldTypeText,

					Options: &schema.TextOptions{
						Max: types.Pointer(40),
					},
				},
				&schema.SchemaField{
					Name:     "emdName",
					Type:     schema.FieldTypeText,
					Required: true,

					Options: &schema.TextOptions{
						Max: types.Pointer(40),
					},
				},
				&schema.SchemaField{
					Name:     "roadName",
					Type:     schema.FieldTypeText,
					Required: true,

					Options: &schema.TextOptions{
						Max: types.Pointer(40),
					},
				},
				&schema.SchemaField{
					Name: "buildingName",
					Type: schema.FieldTypeText,

					Options: &schema.TextOptions{
						Max: types.Pointer(40),
					},
				},
				&schema.SchemaField{
					Name:     "postalNumber",
					Type:     schema.FieldTypeText,
					Required: true,

					Options: &schema.TextOptions{
						Min: types.Pointer(5),
						Max: types.Pointer(5),
					},
				},
				&schema.SchemaField{
					Name: "long",
					Type: schema.FieldTypeNumber,

					Options: &schema.NumberOptions{
						Min:       types.Pointer(-180.0),
						Max:       types.Pointer(180.0),
						NoDecimal: false,
					},
				},
				&schema.SchemaField{
					Name: "lat",
					Type: schema.FieldTypeNumber,

					Options: &schema.NumberOptions{
						Min:       types.Pointer(-90.0),
						Max:       types.Pointer(90.0),
						NoDecimal: false,
					},
				},
				&schema.SchemaField{
					Name:     "crs",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name: "x",
					Type: schema.FieldTypeNumber,
				},
				&schema.SchemaField{
					Name: "y",
					Type: schema.FieldTypeNumber,
				},
				&schema.SchemaField{
					Name: "validPosition",
					Type: schema.FieldTypeBool,
				},
				&schema.SchemaField{
					Name:     "baseDate",
					Type:     schema.FieldTypeDate,
					Required: true,
				},
				&schema.SchemaField{
					Name:        "address",
					Type:        schema.FieldTypeText,
					Required:    true,
					Presentable: true,
				},
			),
			Indexes: types.JsonArray[string]{
				"CREATE UNIQUE INDEX `idx_location_pk` ON `locations` (`pk`)",
				"CREATE INDEX `idx_location_address` ON `locations` (`address`)",
			},
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("locations")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
