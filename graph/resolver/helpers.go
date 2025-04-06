package resolver

import (
	"github.com/jukemori/timeline-generator/graph/model"
	"github.com/jukemori/timeline-generator/internal/models"
)

// Helper function to convert internal timeline model to GraphQL model
func convertTimelineToGraphQL(timeline *models.Timeline) *model.Timeline {
	tasks := make([]*model.TimelineTask, len(timeline.Tasks))
	for i, task := range timeline.Tasks {
		tasks[i] = &model.TimelineTask{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			StartDate:   task.StartDate.Format("2006-01-02"),
			EndDate:     task.EndDate.Format("2006-01-02"),
			Duration:    task.Duration,
			Priority:    task.Priority,
		}
	}

	return &model.Timeline{
		ID:          timeline.ID,
		Title:       timeline.Title,
		Description: timeline.Description,
		StartDate:   timeline.StartDate.Format("2006-01-02"),
		EndDate:     timeline.EndDate.Format("2006-01-02"),
		Tasks:       tasks,
	}
}