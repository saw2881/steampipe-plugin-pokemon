package pokemon

import (
	"context"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tablePokemonType(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "pokemon_type",
		Description: "Types are properties for Pokémon and their moves. Each type has three properties: which types of Pokémon it is super effective against, which types of Pokémon it is not very effective against, and which types of Pokémon it is completely ineffective against.",
		List: &plugin.ListConfig{
			Hydrate: listType,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"name"}),
			// TODO: Add support for 'id' key column
			// KeyColumns: plugin.AnyColumn([]string{"id", "name"}),
			Hydrate: getType,
			// Bad error message is a result of https://github.com/mtslzr/pokeapi-go/issues/29
			ShouldIgnoreError: isNotFoundError([]string{"invalid character 'N' looking for beginning of value"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name for this resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "damage_relations",
				Description: "A detail of how effective this type is toward others and vice versa.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getType,
			},
			{
				Name:        "past_damage_relations",
				Description: "A list of details of how effective this type was toward others and vice versa in previous generations.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getType,
			},
			{
				Name:        "game_indices",
				Description: "A list of game indices relevent to this item by generation.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getType,
			},
			{
				Name:        "id",
				Description: "The identifier for this resource.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getType,
				Transform:   transform.FromGo(),
			},
			{
				Name:        "generation",
				Description: "The generation this type was introduced in.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getType,
			},
			{
				Name:        "move_damage_class",
				Description: "The class of damage inflicted by this type.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getType,
			},
			{
				Name:        "names",
				Description: "The name of this resource listed in different languages.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getType,
			},
			{
				Name:        "pokemon",
				Description: "A list of details of Pokémon that have this type.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getType,
			},
			{
				Name:        "moves",
				Description: "A list of moves that have this type.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getType,
			},
		},
	}
}

func listType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("listType")

	offset := 0

	for true {
		resources, err := pokeapi.Resource("type", offset)

		if err != nil {
			plugin.Logger(ctx).Error("pokemon_type.listType", "query_error", err)
			return nil, err
		}

		for _, berry := range resources.Results {
			d.StreamListItem(ctx, berry)
		}

		// No next URL returned
		if len(resources.Next) == 0 {
			break
		}

		urlOffset, err := extractUrlOffset(resources.Next)
		if err != nil {
			plugin.Logger(ctx).Error("pokemon_type.listType", "extract_url_offset_error", err)
			return nil, err
		}

		// Set next offset
		offset = urlOffset
	}

	return nil, nil
}

func getType(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("getType")

	var name string

	if h.Item != nil {
		result := h.Item.(structs.Result)
		name = result.Name
	} else {
		name = d.KeyColumnQuals["name"].GetStringValue()
	}

	logger.Debug("Name", name)

	berry, err := pokeapi.Type(name)

	if err != nil {
		plugin.Logger(ctx).Error("pokemon_type.getType", "query_error", err)
		return nil, err
	}

	return berry, nil
}
