package extfieldgen

import (
	_ "embed"
	"go/types"
	"maps"
	"strings"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
	"github.com/vektah/gqlparser/v2/ast"
)

var (
	_ plugin.Plugin        = Plugin{}
	_ plugin.ConfigMutator = Plugin{}
)

type Plugin struct{}

func New() Plugin {
	return Plugin{}
}

// Name implements plugin.Plugin
func (Plugin) Name() string {
	return "extfieldgen"
}

// MutateConfig implements plugin.ConfigMutator
func (Plugin) MutateConfig(cfg *config.Config) error {
	cfg.Directives["extraField"] = config.DirectiveConfig{
		SkipRuntime: true,
	}
	for _, schemaType := range cfg.Schema.Types {
		if cfg.Models.UserDefined(schemaType.Name) {
			continue
		}
		if schemaType.Kind != ast.Object && schemaType.Kind != ast.InputObject {
			continue
		}

		model := cfg.Models[schemaType.Name]

		extraFields := make(map[string]config.ModelExtraField)
		maps.Copy(extraFields, model.ExtraFields)
		for _, d := range schemaType.Directives.ForNames("extraField") {
			t := d.Arguments.ForName("type").Value.Raw
			if !isBuiltin(t) || !isFullName(t) {
				t = makeFullName(cfg, t)
			}
			extraFields[d.Arguments.ForName("name").Value.Raw] = config.ModelExtraField{
				Type:        t,
				Description: d.Arguments.ForName("description").Value.Raw,
			}
		}
		model.ExtraFields = extraFields

		cfg.Models[schemaType.Name] = model
	}
	return nil
}

func isBuiltin(t string) bool {
	switch {
	case strings.HasPrefix(t, "[]"):
		return isBuiltin(t[2:])
	case strings.HasPrefix(t, "*"):
		return isBuiltin(t[1:])
	}
	return types.Universe.Lookup(t) != nil
}

func isFullName(t string) bool {
	return strings.Contains(t, ".")
}

func makeFullName(cfg *config.Config, t string) string {
	switch {
	case strings.HasPrefix(t, "[]"):
		return "[]" + makeFullName(cfg, t[2:])
	case strings.HasPrefix(t, "*"):
		return "*" + makeFullName(cfg, t[1:])
	default:
		return cfg.Model.ImportPath() + "." + t
	}
}
