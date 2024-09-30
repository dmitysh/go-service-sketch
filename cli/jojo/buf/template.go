package buf

const (
	TemplateFile = "buf.gen.yaml"
)

type BufGenConfig struct {
	Version string               `yaml:"version"`
	Plugins []BufGenPluginConfig `yaml:"plugins"`
}

type BufGenPluginConfig struct {
	Name     string   `yaml:"name"`
	Out      string   `yaml:"out"`
	Opt      []string `yaml:"opt"`
	Path     string   `yaml:"path"`
	Strategy string   `yaml:"strategy"`
}

func DefaultBufGenConfig() BufGenConfig {
	return BufGenConfig{
		Version: "v1",
		Plugins: []BufGenPluginConfig{
			{
				Name: "go",
				Out:  ".",
				Opt: []string{
					"paths=source_relative",
				},
				Path:     "bin/protoc-gen-go",
				Strategy: "directory",
			},
			{
				Name: "go-grpc",
				Out:  ".",
				Opt: []string{
					"paths=source_relative",
				},
				Path:     "bin/protoc-gen-go-grpc",
				Strategy: "directory",
			},
			{
				Name: "grpc-gateway",
				Out:  ".",
				Opt: []string{
					"logtostderr=true",
					"paths=source_relative",
					"generate_unbound_methods=true",
				},
				Path:     "bin/protoc-gen-grpc-gateway",
				Strategy: "directory",
			},
			{
				Name: "openapiv2",
				Out:  ".",
				Opt: []string{
					"generate_unbound_methods=true",
				},
				Path:     "bin/protoc-gen-openapiv2",
				Strategy: "directory",
			},
		},
	}
}
