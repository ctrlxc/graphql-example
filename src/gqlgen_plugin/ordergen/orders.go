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
	return "customgen"
}

func (m *Plugin) MutateConfig(cfg *config.Config) error {
	_ = syscall.Unlink(m.filepath2(cfg))
	return nil
}

func (m *Plugin) GenerateCode(data *codegen.Data) error {
	orders := make([]string, 0)

	for _, r := range data.Inputs {
		if strings.HasSuffix(r.Definition.Name, "Order") {
			orders = append(orders, r.Definition.Name)
		}
	}

	return templates.Render(templates.Options{
		PackageName: data.Config.Model.Package,
		Filename:    m.filepath2(data.Config),
		Data: &OrderBuild{
			Orders: orders,
		},
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
	})
}

func (m *Plugin) filepath2(cfg *config.Config) string {
	return filepath.Join(filepath.Dir(cfg.Model.Filename), m.filename)
}

type OrderBuild struct {
	Orders []string
}
