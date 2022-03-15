package generator

type MainSARIF struct {
	Version        string          `json:"version"`
	Schema         string          `json:"$schema"`
	Runs           []RunSARIF      `json:"runs,omitempty"`
	IsRuleIDExists map[string]bool `json:"-"`
}

type RunSARIF struct {
	Tool    ToolSARIF     `json:"tool",omitempty`
	Results []ResultSARIF `json:"results",omitempty`
}

type ToolSARIF struct {
	Driver DriverSARIF `json:"driver"`
}

type DriverSARIF struct {
	Name  string      `json:"name"`
	Rules []RuleSARIF `json:"rules"`
}

type RuleSARIF struct {
	ID               string         `json:"id"`
	ShortDescription ShortDescSARIF `json:"shortDescription"`
	HelpURI          string         `json:"helpUri"`
	Properties       PropertySARIF  `json:"properties"`
}

type ShortDescSARIF struct {
	Text string `json:"text"`
}

type PropertySARIF struct {
	Category         string `json:"category"`
	SecuritySeverity string `json:"security-severity"`
}

type ResultSARIF struct {
	Level     string          `json:"level"`
	Message   MessageSARIF    `json:"message"`
	Locations []LocationSARIF `json:"locations"`
	RuleID    string          `json:"ruleId"`
}

type MessageSARIF struct {
	Text string `json:"text"`
}

type LocationSARIF struct {
	PhysicalLocation PhysicalLocationSARIF `json:"physicalLocation"`
}

type PhysicalLocationSARIF struct {
	ArtifactLocation ArtifactLocationSARIF `json:"artifactLocation"`
}

type ArtifactLocationSARIF struct {
	URI string `json:"uri"`
}
