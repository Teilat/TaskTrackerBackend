package cache

import (
	"main/db/sql/models"
	"reflect"
	"testing"
)

func Test_convertColumns(t *testing.T) {
	tests := []struct {
		name string
		cols []*models.Columns
		want map[int32]column
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertColumns(tt.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertProject(t *testing.T) {
	tests := []struct {
		name     string
		projects []*models.Projects
		want     map[int32]project
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProject(tt.projects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertProjectAndUsers(t *testing.T) {
	tests := []struct {
		name string
		tags []*models.ProjectAndUsers
		want map[int32]projectAndUsers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProjectAndUsers(tt.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProjectAndUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertRoles(t *testing.T) {
	tests := []struct {
		name  string
		roles []*models.Roles
		want  map[int32]role
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertRoles(tt.roles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTags(t *testing.T) {
	tests := []struct {
		name string
		tags []*models.Tags
		want map[int32]tag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTags(tt.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTaskAndTags(t *testing.T) {
	tests := []struct {
		name string
		tags []*models.TaskAndTags
		want map[int32]taskAndTags
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTaskAndTags(tt.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTaskAndTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTaskAndUsers(t *testing.T) {
	tests := []struct {
		name string
		tags []*models.TaskAndUsers
		want map[int32]taskAndUsers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTaskAndUsers(tt.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTaskAndUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTasks(t *testing.T) {
	tests := []struct {
		name  string
		tasks []*models.Tasks
		want  map[int32]task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTasks(tt.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertUsers(t *testing.T) {
	tests := []struct {
		name  string
		users []*models.Users
		want  map[int32]user
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertUsers(tt.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
