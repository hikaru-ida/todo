package main

import (
	"github.com/hikaru-ida/todo/command"
	"github.com/mitchellh/cli"
)

func Commands() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &command.InitCommand{}, nil
		},
		"add": func() (cli.Command, error) {
			return &command.AddCommand{}, nil
		},
		"list": func() (cli.Command, error) {
			return &command.ListCommand{}, nil
		},
		"delete": func() (cli.Command, error) {
			return &command.DeleteCommand{}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
