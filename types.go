package main

type UIBuilderSchema struct {
	Steps []struct {
		Form struct {
			Elements []struct {
				Label struct {
					Text    string `json:"text"`
					HasLine bool   `json:"hasLine"`
				} `json:"label,omitempty"`
				Type  string `json:"type,omitempty"`
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				Schema struct {
					Ref string `json:"$ref"`
				} `json:"schema,omitempty"`
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				Fetch string `json:"fetch,omitempty"`
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				IsArray bool `json:"isArray,omitempty"`
				Keys    struct {
					Label struct {
						Text string `json:"text"`
					} `json:"label"`
				} `json:"keys,omitempty"`
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				Values struct {
					Label struct {
						Text string `json:"text"`
					} `json:"label"`
					Type   string `json:"type"`
					Schema struct {
						Ref string `json:"$ref"`
					} `json:"schema"`
				} `json:"values,omitempty"`
				IndividualItemDisabilityCheck string `json:"individualItemDisabilityCheck,omitempty"`
				Label                         struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				HasDescription bool `json:"hasDescription,omitempty"`
				Options        []struct {
					Value       string `json:"value"`
					Text        string `json:"text"`
					Description string `json:"description"`
				} `json:"options,omitempty"`
				If       string `json:"if,omitempty"`
				Elements []struct {
					Label struct {
						Text    string `json:"text"`
						HasLine bool   `json:"hasLine"`
					} `json:"label,omitempty"`
					Type  string `json:"type"`
					If    string `json:"if,omitempty"`
					Label struct {
						Text string `json:"text"`
					} `json:"label,omitempty"`
					Fetch  string `json:"fetch,omitempty"`
					Schema struct {
						Ref string `json:"$ref"`
					} `json:"schema,omitempty"`
					Label struct {
						Text string `json:"text"`
					} `json:"label,omitempty"`
				} `json:"elements,omitempty"`
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
			} `json:"elements"`
			Type string `json:"type"`
		} `json:"form,omitempty"`
		ID    string `json:"id"`
		Title string `json:"title"`
		Form  struct {
			Elements []struct {
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				Computed string `json:"computed,omitempty"`
				OnChange string `json:"onChange,omitempty"`
				Type     string `json:"type,omitempty"`
				Schema   struct {
					Ref string `json:"$ref"`
				} `json:"schema,omitempty"`
				HasDescription bool `json:"hasDescription,omitempty"`
				Options        []struct {
					Value       string `json:"value"`
					Text        string `json:"text"`
					Description string `json:"description"`
				} `json:"options,omitempty"`
				If       string `json:"if,omitempty"`
				Elements []struct {
					Label struct {
						Text string `json:"text"`
					} `json:"label"`
					Type   string `json:"type"`
					Schema struct {
						Ref string `json:"$ref"`
					} `json:"schema"`
				} `json:"elements,omitempty"`
			} `json:"elements"`
			Discriminator struct {
				ActiveDatabaseMode struct {
					Type    string `json:"type"`
					Default string `json:"default"`
				} `json:"activeDatabaseMode"`
			} `json:"discriminator"`
			Type string `json:"type"`
		} `json:"form,omitempty"`
		Form struct {
			Elements []struct {
				Type  string `json:"type"`
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				Schema struct {
					Ref string `json:"$ref"`
				} `json:"schema,omitempty"`
				Computed string `json:"computed,omitempty"`
				OnChange string `json:"onChange,omitempty"`
				If       string `json:"if,omitempty"`
				Elements []struct {
					Label struct {
						Text string `json:"text"`
					} `json:"label,omitempty"`
					Type   string `json:"type"`
					Schema struct {
						Ref string `json:"$ref"`
					} `json:"schema,omitempty"`
					Elements []struct {
						Label struct {
							Text string `json:"text"`
						} `json:"label"`
						Disabled bool   `json:"disabled,omitempty"`
						Computed string `json:"computed,omitempty"`
						Type     string `json:"type"`
						Schema   struct {
							Ref string `json:"$ref"`
						} `json:"schema"`
						Options []struct {
							Text  string `json:"text"`
							Value string `json:"value"`
						} `json:"options,omitempty"`
						If    string `json:"if,omitempty"`
						Fetch string `json:"fetch,omitempty"`
					} `json:"elements,omitempty"`
					AddFormLabel  string `json:"addFormLabel,omitempty"`
					TableContents []struct {
						Type          string `json:"type"`
						TypeOfValue   string `json:"typeOfValue,omitempty"`
						InTableColumn bool   `json:"inTableColumn,omitempty"`
						Label         struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Path  string `json:"path,omitempty"`
						Label struct {
							Text    string `json:"text"`
							HasLine bool   `json:"hasLine"`
						} `json:"label,omitempty"`
						Label struct {
							Text         string `json:"text"`
							IsSubsection bool   `json:"isSubsection"`
						} `json:"label,omitempty"`
						Label struct {
							Text         string `json:"text"`
							IsSubsection bool   `json:"isSubsection"`
						} `json:"label,omitempty"`
						Label struct {
							Text         string `json:"text"`
							IsSubsection bool   `json:"isSubsection"`
						} `json:"label,omitempty"`
						Label struct {
							Text         string `json:"text"`
							IsSubsection bool   `json:"isSubsection"`
						} `json:"label,omitempty"`
						Label struct {
							Text    string `json:"text"`
							HasLine bool   `json:"hasLine"`
						} `json:"label,omitempty"`
						Label struct {
							Text    string `json:"text"`
							HasLine bool   `json:"hasLine"`
						} `json:"label,omitempty"`
					} `json:"tableContents,omitempty"`
					Element struct {
						Label struct {
							Text string `json:"text"`
						} `json:"label"`
						Elements []struct {
							Label struct {
								Text string `json:"text"`
							} `json:"label,omitempty"`
							Type   string `json:"type"`
							Schema struct {
								Ref string `json:"$ref"`
							} `json:"schema,omitempty"`
							Label struct {
								Text    string `json:"text"`
								HasLine bool   `json:"hasLine"`
							} `json:"label,omitempty"`
							CustomClass string `json:"customClass,omitempty"`
							Elements    []struct {
								Schema struct {
									Ref string `json:"$ref"`
								} `json:"schema"`
								Type    string `json:"type"`
								Element struct {
									Label struct {
										Text string `json:"text"`
									} `json:"label"`
									Schema struct {
										Ref string `json:"$ref"`
									} `json:"schema"`
									Type string `json:"type"`
								} `json:"element"`
								Label struct {
									Text         string `json:"text"`
									IsSubsection bool   `json:"isSubsection"`
								} `json:"label"`
							} `json:"elements,omitempty"`
							HasLineLabel bool `json:"hasLineLabel,omitempty"`
							Element      struct {
								Label struct {
									Text string `json:"text"`
								} `json:"label"`
								Schema struct {
									Ref string `json:"$ref"`
								} `json:"schema"`
								Type string `json:"type"`
							} `json:"element,omitempty"`
							Label struct {
								Text    string `json:"text"`
								HasLine bool   `json:"hasLine"`
							} `json:"label,omitempty"`
							Label struct {
								Text    string `json:"text"`
								HasLine bool   `json:"hasLine"`
							} `json:"label,omitempty"`
						} `json:"elements"`
						Type string `json:"type"`
					} `json:"element,omitempty"`
				} `json:"elements,omitempty"`
			} `json:"elements"`
			Discriminator struct {
				ConfigureTLS struct {
					Type    string `json:"type"`
					Default bool   `json:"default"`
				} `json:"configureTLS"`
			} `json:"discriminator"`
			Type string `json:"type"`
		} `json:"form,omitempty"`
		Form struct {
			Type          string `json:"type"`
			Discriminator struct {
				PrePopulateDatabase struct {
					Type string `json:"type"`
				} `json:"prePopulateDatabase"`
			} `json:"discriminator"`
			Elements []struct {
				Type  string `json:"type"`
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				Schema struct {
					Ref string `json:"$ref"`
				} `json:"schema,omitempty"`
				Options []struct {
					Text  string `json:"text"`
					Value string `json:"value"`
				} `json:"options,omitempty"`
				Computed      string `json:"computed,omitempty"`
				OnChange      string `json:"onChange,omitempty"`
				If            string `json:"if,omitempty"`
				Discriminator struct {
					DataSource struct {
						Type string `json:"type"`
					} `json:"dataSource"`
				} `json:"discriminator,omitempty"`
				Elements []struct {
					Type  string `json:"type"`
					Label struct {
						Text string `json:"text"`
					} `json:"label,omitempty"`
					Schema struct {
						Ref string `json:"$ref"`
					} `json:"schema,omitempty"`
					Options []struct {
						Text  string `json:"text"`
						Value string `json:"value"`
					} `json:"options,omitempty"`
					Computed      string `json:"computed,omitempty"`
					OnChange      string `json:"onChange,omitempty"`
					If            string `json:"if,omitempty"`
					Discriminator struct {
						SourceVolumeType struct {
							Type string `json:"type"`
						} `json:"sourceVolumeType"`
					} `json:"discriminator,omitempty"`
					Elements []struct {
						Type  string `json:"type"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Schema struct {
							Ref string `json:"$ref"`
						} `json:"schema,omitempty"`
						Label struct {
							Text    string `json:"text"`
							HasLine bool   `json:"hasLine"`
						} `json:"label,omitempty"`
						Options []struct {
							Text  string `json:"text"`
							Value string `json:"value"`
						} `json:"options,omitempty"`
						Computed string `json:"computed,omitempty"`
						OnChange string `json:"onChange,omitempty"`
						If       string `json:"if,omitempty"`
						Fetch    string `json:"fetch,omitempty"`
					} `json:"elements,omitempty"`
				} `json:"elements,omitempty"`
			} `json:"elements"`
		} `json:"form,omitempty"`
		Form struct {
			Type          string `json:"type"`
			Discriminator struct {
				ScheduleBackup struct {
					Type string `json:"type"`
				} `json:"scheduleBackup"`
			} `json:"discriminator"`
			Elements []struct {
				Type  string `json:"type"`
				Label struct {
					Text string `json:"text"`
				} `json:"label"`
				Schema struct {
					Ref string `json:"$ref"`
				} `json:"schema,omitempty"`
				Options []struct {
					Text  string `json:"text"`
					Value string `json:"value"`
				} `json:"options,omitempty"`
				Computed      string `json:"computed,omitempty"`
				OnChange      string `json:"onChange,omitempty"`
				If            string `json:"if,omitempty"`
				Discriminator struct {
					BackupInvoker struct {
						Type string `json:"type"`
					} `json:"backupInvoker"`
				} `json:"discriminator,omitempty"`
				Elements []struct {
					Type  string `json:"type"`
					Label struct {
						Text string `json:"text"`
					} `json:"label"`
					Schema struct {
						Ref string `json:"$ref"`
					} `json:"schema,omitempty"`
					Options []struct {
						Text  string `json:"text"`
						Value string `json:"value"`
					} `json:"options,omitempty"`
					Computed      string `json:"computed,omitempty"`
					OnChange      string `json:"onChange,omitempty"`
					If            string `json:"if,omitempty"`
					Discriminator struct {
						TargetType struct {
							Type string `json:"type"`
						} `json:"targetType"`
					} `json:"discriminator,omitempty"`
					Elements []struct {
						Type  string `json:"type"`
						Label struct {
							Text    string `json:"text"`
							HasLine bool   `json:"hasLine"`
						} `json:"label,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Schema struct {
							Ref string `json:"$ref"`
						} `json:"schema,omitempty"`
						Disabled bool `json:"disabled,omitempty"`
						Label    struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Computed      string `json:"computed,omitempty"`
						Discriminator struct {
							RepositoryChoise struct {
								Type string `json:"type"`
							} `json:"repositoryChoise"`
						} `json:"discriminator,omitempty"`
						Elements []struct {
							Type  string `json:"type"`
							Label struct {
								Text         string `json:"text"`
								HasLine      bool   `json:"hasLine"`
								IsSubsection bool   `json:"isSubsection"`
							} `json:"label,omitempty"`
							Schema struct {
								Ref string `json:"$ref"`
							} `json:"schema,omitempty"`
							Options []struct {
								Text  string `json:"text"`
								Value string `json:"value"`
							} `json:"options,omitempty"`
							Computed string `json:"computed,omitempty"`
							OnChange string `json:"onChange,omitempty"`
							If       string `json:"if,omitempty"`
							Label    struct {
								Text string `json:"text"`
							} `json:"label,omitempty"`
							Fetch         string `json:"fetch,omitempty"`
							Discriminator struct {
								BackendType struct {
									Type string `json:"type"`
								} `json:"backendType"`
							} `json:"discriminator,omitempty"`
							Elements []struct {
								Type  string `json:"type"`
								Label struct {
									Text string `json:"text"`
								} `json:"label,omitempty"`
								Schema struct {
									Ref string `json:"$ref"`
								} `json:"schema,omitempty"`
								OnChange string `json:"onChange,omitempty"`
								Disabled bool   `json:"disabled,omitempty"`
								Computed string `json:"computed,omitempty"`
								Label    struct {
									Text         string `json:"text"`
									IsSubsection bool   `json:"isSubsection"`
								} `json:"label,omitempty"`
								Options []struct {
									Text  string `json:"text"`
									Value string `json:"value"`
								} `json:"options,omitempty"`
								If       string `json:"if,omitempty"`
								Elements []struct {
									Type  string `json:"type"`
									Label struct {
										Text string `json:"text"`
									} `json:"label"`
									Schema struct {
										Ref string `json:"$ref"`
									} `json:"schema"`
								} `json:"elements,omitempty"`
								Discriminator struct {
									VolumeSource struct {
										Type string `json:"type"`
									} `json:"volumeSource"`
								} `json:"discriminator,omitempty"`
								Fetch string `json:"fetch,omitempty"`
							} `json:"elements,omitempty"`
						} `json:"elements,omitempty"`
						Label struct {
							Text         string `json:"text"`
							HasLine      bool   `json:"hasLine"`
							IsSubsection bool   `json:"isSubsection"`
						} `json:"label,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Label struct {
							Text         string `json:"text"`
							HasLine      bool   `json:"hasLine"`
							IsSubsection bool   `json:"isSubsection"`
						} `json:"label,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
					} `json:"elements,omitempty"`
					Discriminator struct {
						BackupBlueprintName struct {
							Type string `json:"type"`
						} `json:"backupBlueprintName"`
						Schedule struct {
							Type string `json:"type"`
						} `json:"schedule"`
						TaskParameters struct {
							Type                 string `json:"type"`
							AdditionalProperties struct {
								Type string `json:"type"`
							} `json:"additionalProperties"`
						} `json:"taskParameters"`
					} `json:"discriminator,omitempty"`
				} `json:"elements,omitempty"`
			} `json:"elements"`
		} `json:"form,omitempty"`
		Form struct {
			Elements []struct {
				Type  string `json:"type"`
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				Schema struct {
					Ref string `json:"$ref"`
				} `json:"schema,omitempty"`
				Computed string `json:"computed,omitempty"`
				OnChange string `json:"onChange,omitempty"`
				If       string `json:"if,omitempty"`
				Elements []struct {
					Label struct {
						Text string `json:"text"`
					} `json:"label,omitempty"`
					Schema struct {
						Ref string `json:"$ref"`
					} `json:"schema,omitempty"`
					Type           string `json:"type"`
					OnChange       string `json:"onChange,omitempty"`
					HasDescription bool   `json:"hasDescription,omitempty"`
					Options        []struct {
						Value       string `json:"value"`
						Text        string `json:"text"`
						Description string `json:"description"`
					} `json:"options,omitempty"`
					If       string `json:"if,omitempty"`
					Elements []struct {
						Label struct {
							Text    string `json:"text"`
							HasLine bool   `json:"hasLine"`
						} `json:"label,omitempty"`
						Type        string `json:"type"`
						CustomClass string `json:"customClass,omitempty"`
						IsArray     bool   `json:"isArray,omitempty"`
						Schema      struct {
							Ref string `json:"$ref"`
						} `json:"schema,omitempty"`
						Keys struct {
							Label struct {
								Text         string `json:"text"`
								IsSubsection bool   `json:"isSubsection"`
							} `json:"label"`
						} `json:"keys,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
						Values struct {
							Label struct {
								Text string `json:"text"`
							} `json:"label"`
							Type   string `json:"type"`
							Schema struct {
								Ref string `json:"$ref"`
							} `json:"schema"`
						} `json:"values,omitempty"`
						Label struct {
							Text string `json:"text"`
						} `json:"label,omitempty"`
					} `json:"elements,omitempty"`
					Label struct {
						Text    string `json:"text"`
						HasLine bool   `json:"hasLine"`
					} `json:"label,omitempty"`
				} `json:"elements,omitempty"`
				Discriminator struct {
					CustomizeExporter struct {
						Type    string `json:"type"`
						Default bool   `json:"default"`
					} `json:"customizeExporter"`
				} `json:"discriminator,omitempty"`
			} `json:"elements"`
			Discriminator struct {
				EnableMonitoring struct {
					Type    string `json:"type"`
					Default bool   `json:"default"`
				} `json:"enableMonitoring"`
			} `json:"discriminator"`
			Type string `json:"type"`
		} `json:"form,omitempty"`
		Form struct {
			Elements []struct {
				Label struct {
					Text string `json:"text"`
				} `json:"label,omitempty"`
				Type   string `json:"type"`
				Schema struct {
					Ref string `json:"$ref"`
				} `json:"schema,omitempty"`
				Options []struct {
					Text  string `json:"text"`
					Value string `json:"value"`
				} `json:"options,omitempty"`
				If       string `json:"if,omitempty"`
				Fetch    string `json:"fetch,omitempty"`
				Elements []struct {
					Label struct {
						Text string `json:"text"`
					} `json:"label"`
					Type     string `json:"type"`
					OnChange string `json:"onChange,omitempty"`
					Schema   struct {
						Ref string `json:"$ref"`
					} `json:"schema"`
					If       string `json:"if,omitempty"`
					Computed string `json:"computed,omitempty"`
					IsArray  bool   `json:"isArray,omitempty"`
					Keys     struct {
						Label struct {
							Text string `json:"text"`
						} `json:"label"`
					} `json:"keys,omitempty"`
					Values struct {
						Label struct {
							Text string `json:"text"`
						} `json:"label"`
						Type   string `json:"type"`
						Schema struct {
							Ref string `json:"$ref"`
						} `json:"schema"`
					} `json:"values,omitempty"`
				} `json:"elements,omitempty"`
			} `json:"elements"`
			Discriminator struct {
				ConfigurationSource struct {
					Type    string `json:"type"`
					Default string `json:"default"`
				} `json:"configurationSource"`
			} `json:"discriminator"`
			Type string `json:"type"`
		} `json:"form,omitempty"`
	} `json:"steps"`
	Type            string `json:"type"`
	HasPreviewPanel bool   `json:"hasPreviewPanel"`
}
