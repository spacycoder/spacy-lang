// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spacycoder/spacy-lang/compiler"
	"github.com/spacycoder/spacy-lang/lexer"
	"github.com/spacycoder/spacy-lang/parser"
	"github.com/spacycoder/spacy-lang/vm"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run a SpacyLang program",
	Long:  `usage: 'spacy run program.spacy'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			return
		}
		path := args[0]
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}

		rawProgramString := string(b)
		l := lexer.New(rawProgramString)
		p := parser.New(l)
		program := p.ParseProgram()

		comp := compiler.New()
		err = comp.Compile(program)
		if err != nil {
			fmt.Printf("compiler error: %s", err)
			return
		}

		machine := vm.New(comp.Bytecode())

		err = machine.Run()
		if err != nil {
			fmt.Printf("vm error: %s", err)
			return
		}

		fmt.Println(machine.LastPoppedStackElem().Inspect())
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
