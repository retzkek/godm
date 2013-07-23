package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

var helpCmd = &Command{
	Name:    "help",
	Usage:   "[COMMAND]",
	Summary: "Print command usage and options.",
	Help: `godm is a simple general-purpose document management system.
General Usage:
	godm [OPTION] COMMAND [COMMAND OPTION]`,
}

func helpFunc(db *Db, cmd *Command, args []string) error {
	fmt.Print(helpHeader)
	if cmd == nil || cmd == helpCmd {
		type appData struct {
			Flags    []*flag.Flag
			Commands []*Command
		}
		var data = appData{
			Commands: commands,
		}
		// TODO: is there a better way to get a slice of all flags?
		flag.VisitAll(func(f *flag.Flag) {
			data.Flags = append(data.Flags, f)
		})
		tmpl, err := template.New("help").Parse(generalHelpTemplate)
		if err != nil {
			return err
		}
		err = tmpl.Execute(os.Stdout, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	helpCmd.Function = helpFunc
}

var helpHeader = `
godm version 0.1

godm is a simple general-purpose document management system.
`

var generalHelpTemplate = `
USAGE
	godm [OPTION] COMMAND [COMMAND OPTION]

OPTIONS
{{range .Flags}}
	-{{.Name}}=[{{.DefValue}}]
		{{.Usage}}
{{end}}

COMMANDS
{{range .Commands}}
	{{.Name}} {{.Usage}}
		{{.Summary}}
{{end}}
`
