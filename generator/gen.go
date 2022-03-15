package generator

import (
	"encoding/json"
	"errors"
)

func Init() *MainSARIF {
	r := &MainSARIF{
		Version:        "2.1.0",
		Schema:         "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json",
		Runs:           make([]RunSARIF, 1),
		IsRuleIDExists: map[string]bool{},
	}
	r.Runs[0].Results = []ResultSARIF{}
	return r
}

func (m *MainSARIF) AddDriverName(s string) {
	m.Runs[0].Tool.Driver.Name = s
}

func (m *MainSARIF) AddRule(id, shortDescription, helpUri, category string) error {
	if !m.IsRuleIDExists[id] {
		r := RuleSARIF{
			ID:               id,
			ShortDescription: ShortDescSARIF{Text: shortDescription},
			HelpURI:          helpUri,
			Properties:       PropertySARIF{Category: category},
		}
		m.Runs[0].Tool.Driver.Rules = append(m.Runs[0].Tool.Driver.Rules, r)
		m.IsRuleIDExists[id] = true
		return nil
	} else {
		return errors.New("Rule ID exists")
	}
}

func (m *MainSARIF) AddResults(RuleID, level, message string, locations []string) error {
	if !m.IsRuleIDExists[RuleID] {
		return errors.New("Rule ID doesn't exist")
	} else {
		ls := []LocationSARIF{}
		for _, v := range locations {
			ls = append(ls, LocationSARIF{
				PhysicalLocation: PhysicalLocationSARIF{
					ArtifactLocation: ArtifactLocationSARIF{
						URI: v,
					},
				},
			})
		}

		m.Runs[0].Results = append(m.Runs[0].Results, ResultSARIF{
			Level: level,
			Message: MessageSARIF{
				Text: message,
			},
			Locations: ls,
			RuleID:    RuleID,
		})
		return nil
	}
}

func (m *MainSARIF) GetJSON() (string, error) {
	b, err := json.MarshalIndent(m, "", "    ")
	return string(b), err
}
