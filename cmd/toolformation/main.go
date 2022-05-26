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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
		c, err = toolformation.ParseDefaultConfig()
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
	m.Install()
	return nil
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
