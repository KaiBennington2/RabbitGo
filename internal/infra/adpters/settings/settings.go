package settings

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	"gopkg.in/yaml.v3"
	"io"
	"io/fs"
	"os"
)

var _ ports.ISetting = (*Setting)(nil)

//go:embed yaml/*.yaml
var yamlFiles embed.FS

//go:embed json/*.json
var jsonFiles embed.FS

type Setting struct {
	ext types.FileExtType
}

func NewSetting(fileT types.FileExtType) *Setting {
	return &Setting{
		ext: fileT,
	}
}

func (cfg *Setting) Load() (*types.Setting, error) {
	var config *types.Setting
	var err error
	//Load config for scope
	if config, err = cfg.loadConfigFile(cfg.ext, fmt.Sprintf("%v.%v", cfg.getScopeFromEnv(true), cfg.ext)); err != nil {
		//Load config for environments
		return cfg.loadConfigFile(cfg.ext, fmt.Sprintf("%v.%v", cfg.getScopeFromEnv(false), cfg.ext))
	}
	return config, err
}

func (cfg *Setting) getScopeFromEnv(b bool) string {
	defaultEnv := "GO_ENVIRONMENT"
	if b {
		defaultEnv = "SCOPE"
	}
	env := os.Getenv(defaultEnv)
	if env == "" {
		env = "environment_dev"
	}
	return env
}

func (cfg *Setting) loadConfigFile(fileType types.FileExtType, fileName string) (*types.Setting, error) {
	var resp types.Setting
	var f embed.FS
	var unmarshalFunc func([]byte, interface{}) error

	switch fileType {
	case types.Yaml:
		f = yamlFiles
		unmarshalFunc = yaml.Unmarshal
	case types.Json:
		f = jsonFiles
		unmarshalFunc = json.Unmarshal
	default:
		return nil, fmt.Errorf("tipo de archivo no válido")
	}

	filePath := fmt.Sprintf("%s/%s", fileType, fileName)
	file, err := f.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %v. %w", fileName, err)
	}
	defer func(file fs.File) {
		_ = file.Close()
	}(file)

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %v. %w", fileName, err)
	}

	if err := unmarshalFunc(b, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse the config, %w", err)
	}
	return &resp, nil
}
