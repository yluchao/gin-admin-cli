package cmd

import (
	"errors"

	"github.com/gin-admin/gin-admin-cli/v10/internal/actions"
	"github.com/gin-admin/gin-admin-cli/v10/internal/tfs"
	"github.com/urfave/cli/v2"
)

// Generate returns the gen command.
func Generate() *cli.Command {
	return &cli.Command{
		Name:  "gen",
		Usage: "Generate structs to the specified module, support config file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "dir",
				Aliases:  []string{"d"},
				Usage:    "The directory to generate the struct from",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "module",
				Aliases:  []string{"m"},
				Usage:    "The module to generate the struct from",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "tpl-type",
				Usage: "The template type to generate the struct from (default: crud)",
				Value: "crud",
			},
			&cli.StringFlag{
				Name:  "module-path",
				Usage: "The module path to generate the struct from (default: internal/mods)",
				Value: "internal/mods",
			},
			&cli.StringFlag{
				Name:  "wire-path",
				Usage: "The wire generate path to generate the struct from (default: internal/library/wirex)",
				Value: "internal/wirex",
			},
			&cli.StringFlag{
				Name:  "swag-path",
				Usage: "The swagger generate path to generate the struct from (default: internal/swagger)",
				Value: "internal/swagger",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "The config file to generate the struct from (JSON/YAML)",
			},
			&cli.StringFlag{
				Name:    "structs",
				Aliases: []string{"s"},
				Usage:   "The struct to generate (multiple structs can be separated by a comma)",
			},
			&cli.StringFlag{
				Name:  "structs-comment",
				Usage: "Specify the struct comment",
			},
			&cli.StringFlag{
				Name:  "structs-output",
				Usage: "Specify the packages to generate the struct (default: schema,dal,biz,api)",
			},
			&cli.StringFlag{
				Name:  "tpl-path",
				Usage: "The template path to generate the struct from (default use tpls)",
			},
		},
		Action: func(c *cli.Context) error {
			if tplPath := c.String("tpl-path"); tplPath != "" {
				tfs.SetIns(tfs.NewOSFS(tplPath))
			}

			gen := actions.NewGenerate(&actions.GenerateConfig{
				Dir:         c.String("dir"),
				TplType:     c.String("tpl-type"),
				ModuleName:  c.String("module"),
				ModulePath:  c.String("module-path"),
				WirePath:    c.String("wire-path"),
				SwaggerPath: c.String("swag-path"),
			})

			if c.String("config") != "" {
				return gen.Run(c.Context, c.String("config"))
			} else if c.String("structs") != "" {
				return gen.RunWithStruct(c.Context, c.String("structs"), c.String("structs-comment"), c.String("structs-output"))
			} else {
				return errors.New("structs or config must be specified")
			}
		},
	}
}
