package model

// TimelineInput represents the input for generating a timeline
type TimelineInput struct {
	CurrentLevel string `json:"currentLevel"`
	Goal         string `json:"goal"`
	Objectives   string `json:"objectives"`
	CurrentDate  string `json:"currentDate"`
	TargetDate   *string `json:"targetDate,omitempty"`
}

// Timeline represents a generated timeline for achieving a goal
type Timeline struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	StartDate   string          `json:"startDate"`
	EndDate     string          `json:"endDate"`
	Tasks       []*TimelineTask `json:"tasks"`
}

// TimelineTask represents a single task in a timeline
type TimelineTask struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Duration    string `json:"duration"`
	Priority    int    `json:"priority"`
} 