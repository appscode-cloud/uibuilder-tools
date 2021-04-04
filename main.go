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
	"path/filepath"
	"strings"

	"kmodules.xyz/client-go/logs"

	"github.com/go-openapi/jsonreference"
	"github.com/spf13/cobra"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"gomodules.xyz/homedir"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/yaml"
)

var (
	wizardDir  = filepath.Join(homedir.HomeDir(), "go/src/go.bytebuilders.dev/ui-wizards/charts")
	uiFile     = filepath.Join(homedir.HomeDir(), "go/src/go.bytebuilders.dev/ui-wizards/charts/kubedbcom-mongodb-editor/ui/create-ui.yaml")
	schemaFile = filepath.Join(homedir.HomeDir(), "go/src/go.bytebuilders.dev/ui-wizards/charts/kubedbcom-mongodb-editor/values.openapiv3_schema.yaml")
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "uibuilder-schema-checker",
		Short: "Check schema of ui-builder json",
		RunE: func(cmd *cobra.Command, args []string) error {
			if wizardDir == "" {
				return checkFile(uiFile, schemaFile)
			}
			return checkDir(wizardDir)
		},
	}
	flags := rootCmd.Flags()
	flags.AddGoFlagSet(flag.CommandLine)
	flags.StringVar(&wizardDir, "wizard-dir", wizardDir, "Path to wizard directory")
	flags.StringVar(&uiFile, "ui-file", uiFile, "Path to ui.json file")
	flags.StringVar(&schemaFile, "schema-file", schemaFile, "Path to schema file")

	logs.ParseFlags()
	utilruntime.Must(rootCmd.Execute())
}

func checkDir(root string) error {
	matches, err := filepath.Glob(filepath.Join(root, "**", "values.openapiv3_schema.yaml"))
	if err != nil {
		return err
	}
	for _, path := range matches {
		dir := filepath.Dir(path)
		uifiles, err := ioutil.ReadDir(filepath.Join(dir, "ui"))
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			return err
		}
		for _, f := range uifiles {
			if f.IsDir() || f.Name() == "functions.js" || f.Name() == "language.yaml" {
				continue
			}
			fp := filepath.Join(dir, "ui", f.Name())
			fmt.Printf("processing file: %s", fp)
			if err = checkFile(fp, path); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkFile(uiFile, schemaFile string) error {
	result, err := checkUIBuilderSchema(uiFile)
	if err != nil {
		return err
	}
	if result != "" {
		return fmt.Errorf("ui json file does not conform to ui-builder schema. diff\n %v", result)
	}
	return checkJsonReference(uiFile, schemaFile)
}

func checkJsonReference(uiFile, schemaFile string) error {
	data, err := ioutil.ReadFile(uiFile)
	if err != nil {
		return err
	}
	var uijson map[string]interface{}
	err = yaml.Unmarshal(data, &uijson)
	if err != nil {
		return err
	}

	data, err = ioutil.ReadFile(schemaFile)
	if err != nil {
		return err
	}
	var schema map[string]interface{}
	err = yaml.Unmarshal(data, &schema)
	if err != nil {
		return err
	}

	errlist := checkRef(uijson, schema, "")
	for _, e := range errlist {
		_, _ = fmt.Fprintln(os.Stderr, e)
	}
	if len(errlist) > 0 {
		return fmt.Errorf("schema ref check failed")
	}
	return nil
}

func checkRef(uijson, schema map[string]interface{}, path string) (errlist []error) {
	for k, v := range uijson {
		switch u := v.(type) {
		case map[string]interface{}:
			errlist = append(errlist, checkRef(u, schema, path+k+".")...)
		case []interface{}:
			for i := range u {
				entry, ok := u[i].(map[string]interface{})
				if !ok {
					continue
				}
				errlist = append(errlist, checkRef(entry, schema, fmt.Sprintf("%s%s[%d].", path, k, i))...)
			}
		case string:
			if k == "$ref" && strings.HasPrefix(u, "schema#/") {
				curPath := path + k + "."
				p, err := jsonreference.New(u)
				if err != nil {
					errlist = append(errlist, fmt.Errorf("failed to parse schema.ref %s at path %s: %v", u, curPath, err))
					break
				}
				_, _, err = p.GetPointer().Get(schema)
				if err != nil {
					errlist = append(errlist, fmt.Errorf("schema.ref %s at path %s is invalid: %v", u, curPath, err))
					break
				}
			}
		}
	}
	return
}

func checkUIBuilderSchema(filename string) (string, error) {
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
