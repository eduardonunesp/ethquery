package cmd

import (
	"fmt"
	"os"

	"github.com/eduardonunesp/ethquery/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var configurationCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Manage configurations",
}

var listConfigurationsCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all configurations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Url", "Current"})

		configurationList := config.Load()

		for _, configuration := range configurationList.Configurations {
			var sCurrent string

			if configuration.Current {
				sCurrent = "*"
			}

			row := []string{
				configuration.Name,
				configuration.URL,
				sCurrent,
			}

			table.Append(row)
		}

		table.Render()
		fmt.Printf("\n")
	},
}

var newConfigurationCmd = &cobra.Command{
	Use:   "new [configuration identification]",
	Short: "Adds new configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Printf("Configuration and url needed\n")
			cmd.Usage()
			return
		}

		var newConfiguration config.Configuration
		newConfiguration.Name = args[0]
		newConfiguration.URL = args[1]
		newConfiguration.Current = true

		configurationList := config.Load()

		for _, configuration := range configurationList.Configurations {
			if configuration.Name == newConfiguration.Name {
				fmt.Printf("Configuration named %s already exists\n", newConfiguration.Name)
				return
			}
		}

		for i, _ := range configurationList.Configurations {
			configurationList.Configurations[i].Current = false
		}

		configurationList.Configurations = append(configurationList.Configurations, newConfiguration)
		config.Write(&configurationList)
		fmt.Printf("Configuration named %s added with sucess", newConfiguration.Name)
	},
}

var setCurrentConfigurationCmd = &cobra.Command{
	Use:   "current [configuration identification]",
	Short: "Sets current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("Configuration name needed\n")
			cmd.Usage()
			return
		}

		currentName := args[0]
		configurationList := config.Load()

		for i, _ := range configurationList.Configurations {
			if configurationList.Configurations[i].Name == currentName {
				configurationList.Configurations[i].Current = true
			} else {
				configurationList.Configurations[i].Current = false
			}
		}

		config.Write(&configurationList)
		fmt.Printf("Set configuration named %s as current", currentName)
	},
}

func init() {
	rootCmd.AddCommand(configurationCmd)

	configurationCmd.AddCommand(listConfigurationsCmd)
	listConfigurationsCmd.SetUsageTemplate("configuration new <name>")

	configurationCmd.AddCommand(newConfigurationCmd)
	newConfigurationCmd.SetUsageTemplate("configuration new <name> <ur>")

	configurationCmd.AddCommand(setCurrentConfigurationCmd)
	setCurrentConfigurationCmd.SetUsageTemplate("configuration current <name>")
}
