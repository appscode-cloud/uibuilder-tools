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
	"regexp"
	"strings"

	"gomodules.xyz/kglog"

	"github.com/go-openapi/jsonreference"
	"github.com/spf13/cobra"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/yaml"
)

var (
	// wizardDir               = filepath.Join(homedir.HomeDir(), "go/src/go.bytebuilders.dev/ui-wizards/charts")
	// uiFile                  = filepath.Join(homedir.HomeDir(), "go/src/go.bytebuilders.dev/ui-wizards/charts/kubedbcom-mongodb-editor/ui/create-ui.yaml")
	// schemaFile              = filepath.Join(homedir.HomeDir(), "go/src/go.bytebuilders.dev/ui-wizards/charts/kubedbcom-mongodb-editor/values.openapiv3_schema.yaml")

	wizardDir               = ""
	uiFile                  = ""
	schemaFile              = ""
	fmtOnly                 bool
	skipSchemaRefValidation bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "uibuilder-schema-checker",
		Short: "Check schema of ui-builder json",
		RunE: func(cmd *cobra.Command, args []string) error {
			if wizardDir == "" {
				err := formatSchema(uiFile)
				if err != nil || fmtOnly {
					return err
				}
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
	flags.BoolVar(&fmtOnly, "fmt-only", fmtOnly, "Format ui.json file only")
	flags.BoolVar(&skipSchemaRefValidation, "skip-schema-ref-validation", skipSchemaRefValidation, "Skip schema ref validation")

	kglog.ParseFlags()
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
			fmt.Printf("processing file: %s\n", fp)
			err := formatSchema(fp)
			if err != nil {
				return err
			}
			if !fmtOnly {
				if err = checkFile(fp, path); err != nil {
					return err
				}
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
	if skipSchemaRefValidation {
		return nil
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
	errno := 0
	for _, e := range errlist {
		if e != nil {
			_, _ = fmt.Fprintln(os.Stderr, e)
			errno++
		}
	}
	if errno > 0 {
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
				errlist = append(errlist, r1(u, schema, curPath))
			}
		}
	}
	return
}

var re = regexp.MustCompile(`/properties/\d+(/?)`)

func r1(ref string, schema map[string]interface{}, curPath string) error {
	u := re.ReplaceAllString(ref, "/items${1}")
	p, err := jsonreference.New(u)
	if err != nil {
		return fmt.Errorf("failed to parse schema.ref %s at path %s: %v", u, curPath, err)
	}
	_, _, err = p.GetPointer().Get(schema)
	if err == nil {
		return nil
	}

	parts := strings.Split(u, "/")
	if len(parts) >= 3 && parts[len(parts)-2] == "properties" {
		nu := strings.Join(parts[:len(parts)-2], "/")

		p, err := jsonreference.New(nu)
		if err != nil {
			return fmt.Errorf("failed to parse schema.ref %s at path %s: %v", ref, curPath, err)
		}
		v, _, err := p.GetPointer().Get(schema)
		if err != nil {
			return fmt.Errorf("schema.ref %s at path %s is invalid: %v", ref, curPath, err)
		}
		if m, ok := v.(map[string]interface{}); !ok {
			return fmt.Errorf("expected schema.ref %s at path %s to point to an object", nu, curPath)
		} else if _, o2 := m["additionalProperties"]; !o2 {
			return fmt.Errorf("schema.ref %s at path %s is missing additionalProperties", ref, curPath)
		} else {
			return nil
		}
	}

	return fmt.Errorf("schema.ref %s at path %s is invalid: %v", u, curPath, err)
}

func formatSchema(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var original map[string]interface{}
	err = yaml.Unmarshal(data, &original)
	if err != nil {
		return err
	}

	// fix formatting of the input ui.json file
	fmtyml, err := yaml.Marshal(original)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, fmtyml, 0644)
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
