package pb_migrations_backup

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
			Name: "address_group",
			Type: models.CollectionTypeBase,
			Schema: schema.NewSchema(
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
			),
			Indexes: types.JsonArray[string]{
				"CREATE UNIQUE INDEX `idx_address_group_addressGroup` ON `locations_summary` (`addressGroup`)",
			},
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("address_group")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
