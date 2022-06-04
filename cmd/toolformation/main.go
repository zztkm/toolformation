package main

import (
	"os"

	"github.com/goark/errs"
	"github.com/spf13/cobra"
	"github.com/zztkm/toolformation"
)

var toolFormaitonConfigFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toolformation",
	Short: "Manage your tools",
	Long:  `ToolFormation manages the tools used on your machine by defining them in code !`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return execute()
	},
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&toolFormaitonConfigFile, "config", "c", "", "ToolFormation config file path")
	rootCmd.Version = "0.4.0"
}

func execute() error {
	var c *toolformation.Config
	var err error
	if toolFormaitonConfigFile == "" {
		c, err = toolformation.ParseDefaultConfigFile()
	} else {
		c, err = toolformation.ReadFile(toolFormaitonConfigFile)
	}
	if err != nil {
		return err
	}

	m := c.NewPackageManager()
	if m == nil {
		return errs.New("Package manager not specified")
	}

	toolformation.CreateCache(c)
	cc, err := toolformation.ReadCache()
	if err != nil {
		return err
	}

	// 未実装の差分チェック
	c.DifferenceCheck(cc)

	m.Install()

	// code extension
	// TODO: この辺も設定されていたら実行するみたいないい感じメソッドがほしい
	if len(c.VSCode.Extension) != 0 {
		err := c.VSCode.Install()
		if err != nil {
			return err
		}
	}

	toolformation.UpdateCache(c)
	return nil
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
