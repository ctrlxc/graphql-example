package ordergen

import (
	"path/filepath"
	"strings"
	"syscall"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
)

func New(filename string) plugin.Plugin {
	return &Plugin{filename: filename}
}

type Plugin struct {
	filename string
}

var _ plugin.CodeGenerator = &Plugin{}
var _ plugin.ConfigMutator = &Plugin{}

func (m *Plugin) Name() string {
	return "ordergen"
}

func (m *Plugin) MutateConfig(cfg *config.Config) error {
	_ = syscall.Unlink(m.filepath(cfg))
	return nil
}

func (m *Plugin) GenerateCode(data *codegen.Data) error {
	orders := make(codegen.Objects, 0)

	for _, o := range data.Inputs {
		if strings.HasSuffix(o.Definition.Name, "Order") {
			orders = append(orders, o)
		}
	}

	return templates.Render(templates.Options{
		PackageName: data.Config.Model.Package,
		Filename:    m.filepath(data.Config),
		Data: &OrderBuild{
			Orders: orders,
		},
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
	})
}

func (m *Plugin) filepath(cfg *config.Config) string {
	return filepath.Join(filepath.Dir(cfg.Model.Filename), m.filename)
}

type OrderBuild struct {
	Orders codegen.Objects
}
