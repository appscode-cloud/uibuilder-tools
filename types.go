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
	"fmt"
)

type UnionElement struct {
	*LabelElement
	*InputElement
	*RadioElement
	*SelectElement
	*KeyValueInputForm
	*SwitchElement
	*ListInputForm
	*SingleStepForm
	*SingleStepFormArray
	*ConfigureOptionsElement
	*MultiselectElement
	*KeyTextAreaInputFormElement
	*ReusableElement
	*Editor
}

func (u UnionElement) MarshalJSON() ([]byte, error) {
	switch {
	case u.LabelElement != nil:
		return json.Marshal(u.LabelElement)
	case u.InputElement != nil:
		return json.Marshal(u.InputElement)
	case u.RadioElement != nil:
		return json.Marshal(u.RadioElement)
	case u.SelectElement != nil:
		return json.Marshal(u.SelectElement)
	case u.KeyValueInputForm != nil:
		return json.Marshal(u.KeyValueInputForm)
	case u.SwitchElement != nil:
		return json.Marshal(u.SwitchElement)
	case u.ListInputForm != nil:
		return json.Marshal(u.ListInputForm)
	case u.SingleStepForm != nil:
		return json.Marshal(u.SingleStepForm)
	case u.SingleStepFormArray != nil:
		return json.Marshal(u.SingleStepFormArray)
	case u.ConfigureOptionsElement != nil:
		return json.Marshal(u.ConfigureOptionsElement)
	case u.MultiselectElement != nil:
		return json.Marshal(u.MultiselectElement)
	case u.KeyTextAreaInputFormElement != nil:
		return json.Marshal(u.KeyTextAreaInputFormElement)
	case u.ReusableElement != nil:
		return json.Marshal(u.ReusableElement)
	case u.Editor != nil:
		return json.Marshal(u.Editor)
	}
	return nil, nil
}

func (u *UnionElement) UnmarshalJSON(data []byte) error {
	type TypedElement struct {
		Type string `json:"type"`
	}

	var t TypedElement
	err := json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	if t.Type == "" {
		return fmt.Errorf("failed to detect type: %s", string(data))
	}

	if u == nil {
		u = new(UnionElement)
	}

	switch t.Type {
	case "label-element":
		var e LabelElement
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.LabelElement = &e
	case "input":
		var e InputElement
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.InputElement = &e
	case "radio":
		var e RadioElement
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.RadioElement = &e
	case "select":
		var e SelectElement
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.SelectElement = &e
	case "key-value-input-form":
		var e KeyValueInputForm
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.KeyValueInputForm = &e
	case "switch":
		var e SwitchElement
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.SwitchElement = &e
	case "list-input-form":
		var e ListInputForm
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.ListInputForm = &e
	case "single-step-form":
		var e SingleStepForm
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.SingleStepForm = &e
	case "single-step-form-array":
		var e SingleStepFormArray
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.SingleStepFormArray = &e
	case "configure-options":
		var e ConfigureOptionsElement
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.ConfigureOptionsElement = &e
	case "multiselect":
		var e MultiselectElement
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.MultiselectElement = &e
	case "key-text-area-input-form":
		var e KeyTextAreaInputFormElement
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.KeyTextAreaInputFormElement = &e
	case "reusable-element":
		var e ReusableElement
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.ReusableElement = &e
	case "editor":
		var e Editor
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.Editor = &e
	default:
		return fmt.Errorf("unknown element type %s", t.Type)
	}
	return nil
}

type UnionFormElement struct {
	*SingleStepForm
	*MultiStepForm
}

func (u UnionFormElement) MarshalJSON() ([]byte, error) {
	switch {
	case u.SingleStepForm != nil:
		return json.Marshal(u.SingleStepForm)
	case u.MultiStepForm != nil:
		return json.Marshal(u.MultiStepForm)
	}
	return nil, nil
}

func (u *UnionFormElement) UnmarshalJSON(data []byte) error {
	type TypedElement struct {
		Type string `json:"type"`
	}

	var t TypedElement
	err := json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	if t.Type == "" {
		return fmt.Errorf("failed to detect type: %s", string(data))
	}

	if u == nil {
		u = new(UnionFormElement)
	}

	switch t.Type {
	case "single-step-form":
		var e SingleStepForm
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.SingleStepForm = &e
	case "multi-step-form":
		var e MultiStepForm
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		u.MultiStepForm = &e
	default:
		return fmt.Errorf("unknown element type %s", t.Type)
	}
	return nil
}

type UnionOptions struct {
	A []string
	B []RadioElementOption
}

func (u UnionOptions) MarshalJSON() ([]byte, error) {
	if len(u.A) > 0 {
		return json.Marshal(u.A)
	} else if len(u.B) > 0 {
		return json.Marshal(u.B)
	}
	return nil, nil
}

func (u *UnionOptions) UnmarshalJSON(data []byte) error {
	if data[0] == '[' && data[1] == '{' {
		return json.Unmarshal(data, &u.B)
	}
	return json.Unmarshal(data, &u.A)
}

type StringBool struct {
	A string
	B bool
}

func (u StringBool) MarshalJSON() ([]byte, error) {
	if u.A != "" {
		return json.Marshal(u.A)
	}
	return json.Marshal(u.B)
}

func (u *StringBool) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if val, ok := v.(bool); ok {
		u.B = val
	} else {
		u.A = fmt.Sprintf("%v", v)
	}
	return nil
}

/*
{
  "label": {
	 "text": "labels.basic_info",
	 "hasLine": true
  },
  "type": "label-element"
}
*/
type LabelElement struct {
	If          string `json:"if,omitempty"`
	CustomClass string `json:"customClass,omitempty"`
	Label       *Label `json:"label,omitempty"`
	Type        string `json:"type"`
}

/*
   {
      "if": "showAuthPasswordField",
      "label": {
         "text": "labels.database.secret"
      },
      "type": "input",
      "schema": {
         "$ref": "schema#/properties/resources/properties/secretAuth/properties/metadata/properties/name"
      }
   }
*/
type InputElement struct {
	If          string      `json:"if,omitempty"`
	CustomClass string      `json:"customClass,omitempty"`
	OnChange    string      `json:"onChange,omitempty"`
	Computed    string      `json:"computed,omitempty"`
	Decoder     string      `json:"decoder,omitempty"`
	Encoder     string      `json:"encoder,omitempty"`
	Disabled    *StringBool `json:"disabled,omitempty"`
	HideValue   bool        `json:"hideValue,omitempty"`
	Label       *Label      `json:"label,omitempty"`
	Type        string      `json:"type"`
	Schema      SchemaRef   `json:"schema"`
	Required    *StringBool `json:"required,omitempty"`
}

/*
{
  "label": {
	 "text": "labels.terminationPolicy"
  },
  "schema": {
	 "$ref": "schema#/properties/resources/properties/kubedbComMongoDB/properties/spec/properties/terminationPolicy"
  },
  "type": "radio",
  "hasDescription": true,
  "options": [
	 {
		"value": "Delete",
		"text": "options.terminationPolicy.delete.label",
		"description": "options.terminationPolicy.delete.description"
	 },
	 {
		"value": "Halt",
		"text": "options.terminationPolicy.halt.label",
		"description": "options.terminationPolicy.halt.description"
	 },
	 {
		"value": "WipeOut",
		"text": "options.terminationPolicy.wipeOut.label",
		"description": "options.terminationPolicy.wipeOut.description"
	 },
	 {
		"value": "DoNotTerminate",
		"text": "options.terminationPolicy.doNotTerminate.label",
		"description": "options.terminationPolicy.doNotTerminate.description"
	 }
  ]
}
*/
type RadioElement struct {
	If                            string        `json:"if,omitempty"`
	Label                         *Label        `json:"label,omitempty"`
	Computed                      string        `json:"computed,omitempty"`
	Onchange                      string        `json:"onChange,omitempty"`
	Type                          string        `json:"type"`
	Schema                        SchemaRef     `json:"schema"`
	HasDescription                bool          `json:"hasDescription,omitempty"`
	Options                       *UnionOptions `json:"options,omitempty"`
	Individualitemdisabilitycheck string        `json:"individualItemDisabilityCheck,omitempty"`
}

/*
{
  "label": {
	 "text": "labels.database.version"
  },
  "fetch": "getMongoDbVersions|catalog.kubedb.com|v1alpha1|mongodbversions",
  "type": "select",
  "schema": {
	 "$ref": "schema#/properties/resources/properties/kubedbComMongoDB/properties/spec/properties/version"
  }
}
*/
type SelectElement struct {
	If                     string        `json:"if,omitempty"`
	Label                  *Label        `json:"label,omitempty"`
	Fetch                  string        `json:"fetch,omitempty"`
	Computed               string        `json:"computed,omitempty"`
	OnChange               string        `json:"onChange,omitempty"`
	Type                   string        `json:"type"`
	Schema                 SchemaRef     `json:"schema"`
	Disabled               *StringBool   `json:"disabled,omitempty"`
	Options                *UnionOptions `json:"options,omitempty"`
	AllowUserDefinedOption bool          `json:"allowUserDefinedOption,omitempty"`
	CustomClass            string        `json:"customClass,omitempty"`
	Required               *StringBool   `json:"required,omitempty"`
}

/*
{
  "isArray": true,
  "schema": {
	 "$ref": "schema#/properties/resources/properties/appApplication/properties/metadata/properties/annotations"
  },
  "keys": {
	 "label": {
		"text": "labels.annotations.key"
	 }
  },
  "label": {
	 "text": "labels.annotations.label"
  },
  "type": "key-value-input-form",
  "values": {
	 "label": {
		"text": "labels.annotations.value"
	 },
	 "type": "input",
	 "schema": {
		"$ref": "schema#/properties/resources/properties/appApplication/properties/metadata/properties/annotations/additionalProperties"
	 }
  }
}
*/
type KeyValueInputForm struct {
	If                            string            `json:"if,omitempty"`
	IsArray                       bool              `json:"isArray,omitempty"`
	OnChange                      string            `json:"onChange,omitempty"`
	Computed                      string            `json:"computed,omitempty"`
	Schema                        SchemaRef         `json:"schema"`
	Keys                          KVInputFormKeys   `json:"keys,omitempty"`
	Label                         *Label            `json:"label,omitempty"`
	Type                          string            `json:"type"`
	Values                        KVInputFormValues `json:"values,omitempty"`
	Individualitemdisabilitycheck string            `json:"individualItemDisabilityCheck,omitempty"`
}

/*
   {
      "type": "switch",
      "label": {
         "text": "labels.enable_tls"
      },
      "schema": {
         "$ref": "discriminator#/configureTLS"
      },
      "computed": "isValueExistInModel|/resources/kubedbComMongoDB/spec/tls",
      "onChange": "onTlsConfigureChange"
   }
*/
type SwitchElement struct {
	Disabled *StringBool `json:"disabled,omitempty"`
	If       string      `json:"if,omitempty"`
	Type     string      `json:"type"`
	Label    *Label      `json:"label,omitempty"`
	Schema   SchemaRef   `json:"schema"`
	Computed string      `json:"computed,omitempty"`
	Onchange string      `json:"onChange,omitempty"`
}

/*
   {
      "if": "showNewSecretCreateField",
      "type": "single-step-form",
      "schema": {
         "$ref": "schema#/properties/resources/properties/secretAuth/properties/data"
      },
      "elements": [
         {
            "label": {
               "text": "labels.new_secret_password",
               "hasLine": true
            },
            "type": "label-element"
         },
         {
            "if": "showNewSecretCreateField",
            "label": {
               "text": "labels.password"
            },
            "type": "input",
            "schema": {
               "$ref": "schema#/properties/resources/properties/secretAuth/properties/data/properties/password"
            }
         }
      ]
   }
*/
type SingleStepForm struct {
	Disabled      *StringBool              `json:"disabled,omitempty"`
	Type          string                   `json:"type"`
	Label         *Label                   `json:"label,omitempty"`
	Schema        *SchemaRef               `json:"schema,omitempty"`
	If            string                   `json:"if,omitempty"`
	Discriminator map[string]Discriminator `json:"discriminator,omitempty"`
	Elements      []UnionElement           `json:"elements,omitempty"`
	Element       *SingleStepFormElement   `json:"element,omitempty"`
	CustomClass   string                   `json:"customClass,omitempty"`
	ShowLabel     bool                     `json:"show_label,omitempty"`
}

type SingleStepFormElement struct {
	Discriminator map[string]Discriminator `json:"discriminator,omitempty"`
	Elements      []UnionElement           `json:"elements,omitempty"`
}

type TableContentEntry struct {
	Type          string              `json:"type"`
	Computed      string              `json:"computed,omitempty"`
	TypeOfValue   string              `json:"typeOfValue,omitempty"`
	InTableColumn bool                `json:"inTableColumn,omitempty"`
	Path          string              `json:"path,omitempty"`
	Label         *Label              `json:"label,omitempty"`
	TableContents []TableContentEntry `json:"tableContents,omitempty"`
	Required      bool                `json:"required,omitempty"`
}

type SingleStepFormArray struct {
	Required      bool                `json:"required,omitempty"`
	If            string              `json:"if,omitempty"`
	Type          string              `json:"type"`
	Label         *Label              `json:"label,omitempty"`
	AddFormLabel  string              `json:"addFormLabel,omitempty"`
	CustomClass   string              `json:"customClass,omitempty"`
	Schema        SchemaRef           `json:"schema"`
	TableContents []TableContentEntry `json:"tableContents,omitempty"`
	Element       struct {
		Discriminator map[string]Discriminator `json:"discriminator,omitempty"`
		Label         *Label                   `json:"label,omitempty"`
		Elements      []UnionElement           `json:"elements,omitempty"`
		Type          string                   `json:"type,omitempty"`
	} `json:"element,omitempty"`
	Discriminator map[string]Discriminator `json:"discriminator,omitempty"`
}

/*
   {
      "schema": {
         "$ref": "schema#/properties/resources/properties/kubedbComMongoDB/properties/spec/properties/tls/properties/certificates/items/properties/dnsNames"
      },
      "hasLineLabel": true,
      "type": "list-input-form",
      "element": {
         "label": {
            "text": "labels.dns_name"
         },
         "schema": {
            "$ref": "schema#/properties/resources/properties/kubedbComMongoDB/properties/spec/properties/tls/properties/certificates/items/properties/dnsNames/items"
         },
         "type": "input"
      },
      "label": {
         "text": "labels.dns_names",
         "hasLine": true
      }
   }
*/
type ListInputForm struct {
	If           string       `json:"if,omitempty"`
	Schema       SchemaRef    `json:"schema"`
	Haslinelabel bool         `json:"hasLineLabel,omitempty"`
	Type         string       `json:"type"`
	Element      InputElement `json:"element"`
	Label        *Label       `json:"label"`
}

type MultiStepForm struct {
	Type            string              `json:"type"`
	HasPreviewPanel bool                `json:"hasPreviewPanel,omitempty"`
	Steps           []MultiStepFormStep `json:"steps"`
}

type MultiStepFormStep struct {
	ID    string           `json:"id,omitempty"`
	Title string           `json:"title"`
	Form  UnionFormElement `json:"form"`
	If    string           `json:"if,omitempty"`
}

type SchemaRef struct {
	Ref string `json:"$ref,omitempty"`
}

type Label struct {
	Text         string `json:"text,omitempty"`
	HasLine      bool   `json:"hasLine,omitempty"`
	IsSubsection bool   `json:"isSubsection,omitempty"`
}

type RadioElementOption struct {
	Value       StringBool `json:"value,omitempty"`
	Text        string     `json:"text,omitempty"`
	Description string     `json:"description,omitempty"`
}

type KVInputFormKeys struct {
	Label Label `json:"label"`
}

type KVInputFormValues struct {
	Label  *Label    `json:"label"`
	Type   string    `json:"type"`
	Schema SchemaRef `json:"schema"`
}

type Discriminator struct {
	Type                 string                `json:"type,omitempty"`
	AdditionalProperties *AdditionalProperties `json:"additionalProperties,omitempty"`
	Default              *StringBool           `json:"default,omitempty"`
}

type AdditionalProperties struct {
	Type string `json:"type"`
}

type ConfigureOptionsElement struct {
	Cluster         ClusterRef     `json:"cluster"`
	HasDependencies bool           `json:"hasDependencies"`
	HasDescription  bool           `json:"hasDescription"`
	Options         []EditorOption `json:"options"`
	Owner           OwnerRef       `json:"owner"`
	Type            string         `json:"type"`
}

type ClusterRef struct {
	Ref string `json:"$ref"`
}

type OwnerRef struct {
	Ref string `json:"$ref"`
}

type EditorOption struct {
	Description    string                   `json:"description,omitempty"`
	Text           string                   `json:"text"`
	Value          string                   `json:"value"`
	Dependencies   []EditorOptionDependency `json:"dependencies,omitempty"`
	DependingSteps []string                 `json:"dependingSteps,omitempty"`
}

type EditorOptionDependency struct {
	Group    string `json:"group"`
	Name     string `json:"name"`
	Resource string `json:"resource"`
	URL      string `json:"url"`
	Version  string `json:"version"`
}

/*
{
   "fetch": "getImagePullSecrets",
   "label": {
      "text": "labels.image_pull_secrets"
   },
   "schema": {
      "$ref": "schema#/properties/resources/properties/kubedbComMongoDB/properties/spec/properties/shardTopology/properties/shard/properties/podTemplate/properties/spec/properties/imagePullSecrets"
   },
   "type": "multiselect"
}
*/
type MultiselectElement struct {
	Computed string    `json:"computed,omitempty"`
	Fetch    string    `json:"fetch"`
	Label    *Label    `json:"label,omitempty"`
	Schema   SchemaRef `json:"schema"`
	Type     string    `json:"type"`
}

/*
{
    "isArray": true,
    "keys": {
        "label": {
            "text": "labels.labels.key"
        }
    },
    "label": {
        "text": "labels.configuration_files"
    },
    "schema": {
        "$ref": "schema#/properties/resources/properties/secret_config/properties/stringData"
    },
    "type": "key-text-area-input-form",
    "values": {
        "label": {
            "text": "labels.labels.value"
        },
        "schema": {
            "$ref": "schema#/properties/resources/properties/secret_config/properties/stringData/additionalProperties"
        },
        "type": "input"
    }
}
*/
type KeyTextAreaInputFormElement struct {
	IsArray bool              `json:"isArray"`
	Keys    KVInputFormKeys   `json:"keys"`
	Label   *Label            `json:"label,omitempty"`
	Schema  SchemaRef         `json:"schema"`
	Type    string            `json:"type"`
	Values  KVInputFormValues `json:"values"`
}

type ChartRef struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ReusableElement struct {
	Alias          string               `json:"alias"`
	Chart          ChartRef             `json:"chart"`
	Label          *Label               `json:"label,omitempty"`
	DataContext    map[string]SchemaRef `json:"dataContext,omitempty"`
	If             string               `json:"if,omitempty"`
	ModuleResolver string               `json:"moduleResolver"`
	Schema         SchemaRef            `json:"schema"`
	Type           string               `json:"type"`
	ShowLabel      bool                 `json:"show_label,omitempty"`
}

type Editor struct {
	Computed string    `json:"computed"`
	If       string    `json:"if,omitempty"`
	Label    *Label    `json:"label"`
	OnChange string    `json:"onChange,omitempty"`
	Schema   SchemaRef `json:"schema"`
	Type     string    `json:"type"`
}
