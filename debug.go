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
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

func NewDebugCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "debug",
		Short: "Debug ui json schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			uijson := `discriminator:
  alertEnabledStatus:
    type: string
elements:
- element:
    discriminator:
      toggleValue:
        type: string
    elements:
    - label:
        text: labels.alert_options
      options:
      - text: Critical
        value: critical
      - text: Warning
        value: warning
      - text: Info
        value: info
      - text: None
        value: none
      schema:
        $ref: schema#/properties/groups/properties/$dyn/properties/enabled
      type: select
    hideForm: true
    schema:
      $ref: schema#/properties/groups/properties/$dyn
    show_label: true
    toggleOption:
      id: switch
      onStatusFalse: onGroupStatusChange|false
      onStatusTrue: onGroupStatusChange|true
      schema:
        $ref: discriminator#/toggleValue
      setInitialStatusFalse: setInitialValueOfToggleBtn
      show: true
    type: single-step-form
  if: showAlertSection
  label:
    text: labels.alert_groups
  schema:
    $ref: schema#/properties/groups
  show_label: true
  type: array-input-form
schema:
  $ref: schema#/
type: single-step-form`
			var e SingleStepForm
			return yaml.Unmarshal([]byte(uijson), &e)
		},
	}
	return cmd
}
