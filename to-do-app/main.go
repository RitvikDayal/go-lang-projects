/*
To list app
---
author: ritvikdayal
*/

package main

import (
	"fmt"
	"github.com/go/to-do-app/commands"
	// "github.com/go/to-do-app/database/models"
	// "github.com/go/to-do-app/tasks"
)

func main() {
	commands.RootCmd.AddCommand(commands.ListCmd)
	commands.RootCmd.AddCommand(commands.AddCmd)

	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}


}
