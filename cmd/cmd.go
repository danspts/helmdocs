package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/danspts/helmdocs/pkg/generate/readme"
	"github.com/danspts/helmdocs/pkg/generate/values"
)

var (
	version = "dev"
	commit  = ""
)

func helmdocs(args []string) {
	if len(args) < 2 {
		fmt.Printf("version: %s, commit:%s\n", version, commit)
		fmt.Println("expected 'generate' subcommand")
		os.Exit(1)
	}

	switch args[1] {
	case "generate":
		generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
		generateCmd.Usage = func() {
			fmt.Fprintf(generateCmd.Output(), "Usage of %s:\n", "generate")
			fmt.Println("  readme      Generates README")
			fmt.Println("  values      Generates values")
			generateCmd.PrintDefaults()
		}

		if len(args) < 3 {
			fmt.Println("expected 'readme' or 'values' subcommand")
			generateCmd.Usage()
			os.Exit(0)
		}

		command := "\033[1mhelmdocs\033[0m "
		command += strings.Join(args[1:3], " ")
		switch args[2] {
		case "readme":
			handleGenerateReadmeCommand(command, args[3:])
		case "values":
			handleGenerateValuesCommand(command, args[3:])
		default:
			fmt.Println("expected 'readme' or 'values' subcommand")
			generateCmd.Usage()
			os.Exit(1)
		}
	default:
		fmt.Println("expected 'generate' subcommand")
		os.Exit(1)
	}
}

func usage(command string, f *flag.FlagSet) func() {
	return func() {
		fmt.Fprintf(f.Output(), "Usage of %s:\n", command)
		f.PrintDefaults()
	}
}

func handleGenerateReadmeCommand(command string, args []string) {
	generateReadmeCmd := flag.NewFlagSet("readme", flag.ExitOnError)
	schemaPath := generateReadmeCmd.String("schema-path", "values.schema.json", "Path to the schema file")

	generateReadmeCmd.Usage = usage(command, generateReadmeCmd)

	generateReadmeCmd.Parse(args)
	readme.GenerateReadme(*schemaPath)
}

func handleGenerateValuesCommand(command string, args []string) {
	generateValuesCmd := flag.NewFlagSet("values", flag.ExitOnError)
	schemaPath := generateValuesCmd.String("schema-path", "values.schema.json", "Path to the schema file")
	omitDefault := generateValuesCmd.Bool("omit-default", true, "Omit default values")

	generateValuesCmd.Usage = usage(command, generateValuesCmd)

	generateValuesCmd.Parse(args)
	values.GenerateValues(*schemaPath, *omitDefault)
}

func main() {
	helmdocs(os.Args)
}
