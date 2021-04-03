/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"kmodules.xyz/client-go/logs"

	"github.com/spf13/cobra"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/yaml"
)

var (
	filename string = "./u1.json"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "uibuilder-schema-checker",
		Short: "Check schema of ui-builder json",
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := check(filename)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	flags := rootCmd.Flags()
	flags.AddGoFlagSet(flag.CommandLine)
	flags.StringVar(&filename, "content", filename, "Path to directory where markdown files reside")

	logs.ParseFlags()
	utilruntime.Must(rootCmd.Execute())
}

func check(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var original map[string]interface{}
	err = yaml.Unmarshal(data, &original)
	if err != nil {
		return "", err
	}
	sorted, err := json.MarshalIndent(&original, "", "  ")
	if err != nil {
		return "", err
	}

	// fix formatting of the input ui.json file
	fmtyml, err := yaml.Marshal(original)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(filename, fmtyml, 0644)
	if err != nil {
		return "", err
	}

	var spec MultiStepForm
	err = yaml.Unmarshal(data, &spec)
	if err != nil {
		return "", err
	}
	parsed, err := json.Marshal(spec)
	if err != nil {
		return "", err
	}

	// Then, Check them
	differ := diff.New()
	d, err := differ.Compare(sorted, parsed)
	if err != nil {
		fmt.Printf("Failed to unmarshal file: %s\n", err.Error())
		os.Exit(3)
	}

	if d.Modified() {
		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		f := formatter.NewAsciiFormatter(original, config)
		result, err := f.Format(d)
		if err != nil {
			return "", err
		}
		return result, nil
	}

	return "", nil
}
