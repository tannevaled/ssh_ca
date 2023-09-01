package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"ssh-ca/cmd"
	cmdCa "ssh-ca/cmd/ca"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func readConfig(ssh_ca_config_file_path *string, cmd *cobra.Command) {

	//log.Printf("ssh_ca_config_file_path <%#v>\n", *ssh_ca_config_file_path)

	if *ssh_ca_config_file_path == "unset" {
		viper.SetConfigName("config.toml")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/.config/ssh-ca/")
		viper.AddConfigPath("/etc/ssh-ca/")
		//viper.SetDefault("description", "")
		viper.AutomaticEnv()
	} else {
		viper.SetConfigName(filepath.Base(*ssh_ca_config_file_path))
		viper.AddConfigPath(filepath.Dir(*ssh_ca_config_file_path))
	}

	viper.SetConfigType("toml")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		// Handle errors reading the config file
		//panic(fmt.Errorf("fatal error config file: %w", err))
		log.Println(fmt.Errorf("%w", err))

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			log.Println("You can supply an alternate config file path with the --config Global Flag")
			cmd.Help()
		} //else {

		//}
		os.Exit(1)

	}
}

//type ConfigCa struct {
//	gorm.Model
//	Name        string
//	Description string
//}

func main() {
	var version string = "0.0.1"
	var ssh_ca_config_file_path string
	var ssh_ca_private_key_file_path string
	var ssh_ca_name string

	app := &cobra.Command{
		Use:     "ssh-ca",
		Version: version,
	}
	/*
		app.PersistentFlags().StringVarP(
			&ssh_ca_config_file_path,
			"config-file-path",
			"c",
			"",
			"The SSH Certificate Authority config file path",
		)
		app.MarkPersistentFlagRequired("config-file-path")
	*/
	app.AddCommand(
		cmd.Licence(),
		cmdCa.Ca(
			&ssh_ca_config_file_path,
			&ssh_ca_name,
			&ssh_ca_private_key_file_path,
		),
		cmd.Agent(),
	)
	app.Execute()

}
