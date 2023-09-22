package CmdCertificateAuthority

import (
	Config "ssh-ca/src/config"

	"github.com/spf13/cobra"
)

func List(
	ssh_ca_config_file_path *string,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the available SSH Certificate Authorities",
		Run: func(
			cmd *cobra.Command,
			args []string,
		) {

			config := Config.New(ssh_ca_config_file_path)
			config.ListCertificateAuthorities()
			//__SshCaConfig.addCa()
			//__SshCaConfig.listCa()
			//db, err := gorm.Open(sqlite.Open(*ssh_ca_config_file_path), &gorm.Config{})

			// get SQLite version
			//log.Println(db.QueryRow("select sqlite_version()"))

			// Migrate the schema
			//db.AutoMigrate(&ConfigCa{})

			//readConfig(ssh_ca_config_file_path, cmd)
			//config := viper.Get("ssh-ca")
			//log.Printf("ALLSETINGS: %#v", viper.AllSettings())
			//log.Printf("ALLKEYS: %#v", viper.AllKeys())
			//log.Printf("%#v", viper.)

			//viper.Set("ssh-ca.ca1.creation_date", time.Now())
			//viper.WriteConfig()

			//		var config *koanf.Koanf = koanf.New("/")
			//		var provider *file.File = file.Provider(*ssh_ca_config_file_path)
			//			var flattenSlices bool = false
			//			var parser *hcl.HCL = hcl.Parser(flattenSlices)

			//if fileExists(*ssh_ca_config_file_path) {
			//fmt.Printf("%#v\n", *ssh_ca_config_file_path)
			//			if err := config.Load(provider, parser); err != nil {
			//				log.Fatalf("error loading config: %v", err)
			//			}

			//			fmt.Printf("All: %#v\n", config.All())

			//			log.Printf("Keys: %#v\n", config.Keys())
			//			log.Printf("MapKeys: %#v\n", config.MapKeys("/ssh-ca/"))
			//log.Printf("Print: %#v\n", config.Print())
			//			config.Print()

			//			type caStruct struct {
			//				Name        string `koand:"name"`
			//				Description string `koanf:"description"`
			//				PrivateKey  string `koanf:"private_key"`
			//			}

			//			var out caStruct
			//			var path string = "ssh-ca/B6D1A58C-F52F-45B2-BA0F-60D2F2F02FCF"
			//			log.Printf("Exists: %s, %s", path, config.Exists(path))
			//			config.Unmarshal(path, &out)
			//			log.Printf("[OUT] %#v\n", out)
			//log.Printf("[PATH] %#v\n", config.String("ssh-ca/AF0987654"))

			/*
				_o := config.Get("ssh-ca").([]map[string]interface{})
				fmt.Printf("%#v\n", _o)
				for _, v := range _o {
					fmt.Printf("%#v\n", v)
					for k, _ := range v {
						fmt.Printf("%#v\n", k)
					}
				}
			*/
			//} else {
			//	fmt.Printf("SSH CA config file '%s' does not exists\n", *ssh_ca_config_file_path)
			//	os.Exit(1)
			//}
		},
	}

	return cmd
}
