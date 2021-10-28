package pokemon

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-pokemon",
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"pokemon_berry": tablePokemonBerry(ctx),
			"pokemon_pokemon": tablePokemonPokemon(ctx),
			"pokemon_type": tablePokemonType(ctx),
		},
	}
	return p
}
