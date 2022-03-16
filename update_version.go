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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

func NewUpdateVersionCommand() *cobra.Command {
	var version string
	cmd := &cobra.Command{
		Use:   "update-version",
		Short: "Update chart version of reusable components",
		RunE: func(cmd *cobra.Command, args []string) error {
			if wizardDir == "" {
				return updateVersionForFile(uiFile, version)
			}
			return updateVersionsForDir(wizardDir, version)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&wizardDir, "wizard-dir", wizardDir, "Path to wizard directory")
	flags.StringVar(&uiFile, "ui-file", uiFile, "Path to ui.json file")
	flags.StringVar(&version, "version", version, "Chart version to be used for reusable components")

	return cmd
}

func updateVersionsForDir(rootDir, version string) error {
	matches, err := filepath.Glob(filepath.Join(rootDir, "**", "ui"))
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
			filename := filepath.Join(dir, "ui", f.Name())
			if err := updateVersionForFile(filename, version); err != nil {
				return err
			}
		}
	}
	return nil
}

func updateVersionForFile(filename string, version string) error {
	fmt.Printf("processing file: %s\n", filename)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var original Document
	err = yaml.Unmarshal(data, &original)
	if err != nil {
		return err
	}
	uvRootForm(&original, version)

	// fix formatting of the input ui.json file
	fmtyml, err := yaml.Marshal(original)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, fmtyml, 0o644)
}

func uvRootForm(f *Document, version string) {
	if f.SingleStepForm != nil {
		uvSingleStepForm(f.SingleStepForm, version)
	}
	if f.SingleStepFormArray != nil {
		uvSingleStepFormArray(f.SingleStepFormArray, version)
	}
	if f.MultiStepForm != nil {
		uvMultiStepForm(f.MultiStepForm, version)
	}
}

func uvUnionFormElement(f *UnionFormElement, version string) {
	if f.SingleStepForm != nil {
		uvSingleStepForm(f.SingleStepForm, version)
	}
	if f.MultiStepForm != nil {
		uvMultiStepForm(f.MultiStepForm, version)
	}
}

func uvUnionElement(f *UnionElement, version string) {
	if f.SingleStepForm != nil {
		uvSingleStepForm(f.SingleStepForm, version)
	}
	if f.SingleStepFormArray != nil {
		uvSingleStepFormArray(f.SingleStepFormArray, version)
	}
	if f.ReusableElement != nil {
		uvReusableElement(f.ReusableElement, version)
	}
}

func uvMultiStepForm(form *MultiStepForm, version string) {
	for i, e := range form.Steps {
		uvMultiStepFormStep(&e, version)
		form.Steps[i] = e
	}
}

func uvMultiStepFormStep(f *MultiStepFormStep, version string) {
	uvUnionFormElement(&f.Form, version)
}

func uvSingleStepForm(f *SingleStepForm, version string) {
	for i, e := range f.Elements {
		uvUnionElement(&e, version)
		f.Elements[i] = e
	}
	if f.Element != nil {
		uvSingleStepFormElement(f.Element, version)
	}
}

func uvSingleStepFormElement(f *SingleStepFormElement, version string) {
	for i, e := range f.Elements {
		uvUnionElement(&e, version)
		f.Elements[i] = e
	}
}

func uvSingleStepFormArray(f *SingleStepFormArray, version string) {
	for i, e := range f.Element.Elements {
		uvUnionElement(&e, version)
		f.Element.Elements[i] = e
	}
}

func uvReusableElement(f *ReusableElement, version string) {
	f.Chart.Version = version
}
