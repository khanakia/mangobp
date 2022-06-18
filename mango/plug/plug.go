package plug

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type IMigrateDb interface {
	MigrateDb()
}

func InitPlugins(rootCmd *cobra.Command, plugin interface{}) {
	// rootCmd := cliCli.RootCmd

	var Main = &cobra.Command{
		Use:   "plugin",
		Short: "Use `go run . plugin --help` to see child commands",
		Run: func(cmd *cobra.Command, args []string) {
			color.Yellow("Run below command to see all the child commands.")
			color.Cyan("go run . plugin --help")
		},
	}

	var MigrateDb = &cobra.Command{
		Use:   "migratedb",
		Short: "Migrate Database",
		Run: func(cmd *cobra.Command, args []string) {
			MigrateDb(plugin)
			fmt.Println("Migration Start")
		},
	}

	var ReplCmd = &cobra.Command{
		Use:   "repl",
		Short: "Start Interactive Shell",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting Repl")
			prompt1(plugin)
		},
	}

	Main.AddCommand(MigrateDb)
	rootCmd.AddCommand(Main)
	rootCmd.AddCommand(ReplCmd)
}

func MigrateDb(p interface{}) {
	v := reflect.ValueOf(p)
	typeOfS := v.Type()
	fmt.Println(typeOfS)

	// if type = type Plugin struct {}
	if reflect.TypeOf(p).Name() != "Plugin" {
		fmt.Println("wrong interface passed as plugin")
		return
	}

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		plugin, ok := v.Field(i).Interface().(IMigrateDb)
		if ok {
			plugin.MigrateDb()
		}
		fmt.Println(ok)
	}
}

func prompt1(plug interface{}) {
	// validate := func(input string) error {
	// 	_, err := strconv.ParseFloat(input, 64)
	// 	if err != nil {
	// 		return errors.New("Invalid number")
	// 	}
	// 	return nil
	// }

	prompt := promptui.Prompt{
		Label: "Function",
		// Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

	valOf := reflect.ValueOf(plug)
	valOf.FieldByName("Auth").MethodByName("Say").Call([]reflect.Value{reflect.ValueOf("aman")})

	// prompt1(plug)
}

func Invoke(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return reflect.ValueOf(any).MethodByName(name).Call(inputs)
}
