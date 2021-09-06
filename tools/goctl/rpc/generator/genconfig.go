package generator

import (
	"io/ioutil"
	"os"
	"path/filepath"

	conf "github.com/lukebull/go-zero-extern/tools/goctl/config"
	"github.com/lukebull/go-zero-extern/tools/goctl/rpc/parser"
	"github.com/lukebull/go-zero-extern/tools/goctl/util"
	"github.com/lukebull/go-zero-extern/tools/goctl/util/format"
)

const configTemplate = `package config

import "github.com/lukebull/go-zero-extern/zrpc"

type Config struct {
	zrpc.RpcServerConf
}
`

// GenConfig generates the configuration structure definition file of the rpc service,
// which contains the zrpc.RpcServerConf configuration item by default.
// You can specify the naming style of the target file name through config.Config. For details,
// see https://github.com/lukebull/go-zero-extern/tree/master/tools/goctl/config/config.go
func (g *DefaultGenerator) GenConfig(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetConfig()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "config")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if util.FileExists(fileName) {
		return nil
	}

	text, err := util.LoadTemplate(category, configTemplateFileFile, configTemplate)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, []byte(text), os.ModePerm)
}
