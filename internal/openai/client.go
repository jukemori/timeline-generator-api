package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jukemori/timeline-generator/graph/model"
	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		client: openai.NewClient(apiKey),
	}
}

func (c *Client) GenerateTimeline(ctx context.Context, input model.TimelineInput) (*model.Timeline, error) {
	prompt := fmt.Sprintf(`
Create a detailed timeline for achieving the following goal:

Current Level: %s
Goal: %s
Objectives: %s
Current Date: %s
Target Date: %s

Please provide a timeline with specific tasks, including:
- Task title
- Task description
- Start date
- End date
- Duration (in days)
- Priority level (1-5, with 5 being highest)

Format your response as a JSON object with the following structure:
{
  "title": "Timeline title",
  "description": "Overall timeline description",
  "startDate": "YYYY-MM-DD",
  "endDate": "YYYY-MM-DD",
  "tasks": [
    {
      "title": "Task 1 title",
      "description": "Task 1 description",
      "startDate": "YYYY-MM-DD",
      "endDate": "YYYY-MM-DD",
      "duration": "X days",
      "priority": 5
    },
    ...
  ]
}
`, input.CurrentLevel, input.Goal, input.Objectives, input.CurrentDate, input.TargetDate)

	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful timeline generator that creates detailed study/achievement plans.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Temperature: 0.7,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("OpenAI API error: %v", err)
	}

	content := resp.Choices[0].Message.Content

	// Extract JSON from the response
	content = extractJSON(content)

	var timelineData struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
		Tasks       []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			StartDate   string `json:"startDate"`
			EndDate     string `json:"endDate"`
			Duration    string `json:"duration"`
			Priority    int    `json:"priority"`
		} `json:"tasks"`
	}

	if err := json.Unmarshal([]byte(content), &timelineData); err != nil {
		return nil, fmt.Errorf("error parsing timeline JSON: %v", err)
	}

	// Convert to our model
	timeline := &model.Timeline{
		Title:       timelineData.Title,
		Description: timelineData.Description,
		StartDate:   timelineData.StartDate,
		EndDate:     timelineData.EndDate,
		Tasks:       make([]*model.TimelineTask, len(timelineData.Tasks)),
	}

	for i, task := range timelineData.Tasks {
		timeline.Tasks[i] = &model.TimelineTask{
			Title:       task.Title,
			Description: task.Description,
			StartDate:   task.StartDate,
			EndDate:     task.EndDate,
			Duration:    task.Duration,
			Priority:    task.Priority,
		}
	}

	return timeline, nil
}

// Helper function to extract JSON from the response
func extractJSON(content string) string {
	// Find the start and end of JSON content
	start := strings.Index(content, "{")
	end := strings.LastIndex(content, "}")

	if start >= 0 && end >= 0 && end > start {
		return content[start : end+1]
	}
	return content
}
