package cache

import (
	"main/db/sql/models"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(nil)
		})
	}
}

func Test_convertColumns(t *testing.T) {
	type args struct {
		cols []*models.Columns
	}
	tests := []struct {
		name string
		args args
		want map[int32]column
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertColumns(tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertProject(t *testing.T) {
	type args struct {
		projects []*models.Projects
	}
	tests := []struct {
		name string
		args args
		want map[int32]project
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProject(tt.args.projects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertProjectAndUsers(t *testing.T) {
	type args struct {
		tags []*models.ProjectAndUsers
	}
	tests := []struct {
		name string
		args args
		want map[int32]projectAndUsers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProjectAndUsers(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProjectAndUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertRoles(t *testing.T) {
	type args struct {
		roles []*models.Roles
	}
	tests := []struct {
		name string
		args args
		want map[int32]role
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertRoles(tt.args.roles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTags(t *testing.T) {
	type args struct {
		tags []*models.Tags
	}
	tests := []struct {
		name string
		args args
		want map[int32]tag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTags(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTaskAndTags(t *testing.T) {
	type args struct {
		tags []*models.TaskAndTags
	}
	tests := []struct {
		name string
		args args
		want map[int32]taskAndTags
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTaskAndTags(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTaskAndTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTaskAndUsers(t *testing.T) {
	type args struct {
		tags []*models.TaskAndUsers
	}
	tests := []struct {
		name string
		args args
		want map[int32]taskAndUsers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTaskAndUsers(tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTaskAndUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTasks(t *testing.T) {
	type args struct {
		tasks []*models.Tasks
	}
	tests := []struct {
		name string
		args args
		want map[int32]task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTasks(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertUsers(t *testing.T) {
	type args struct {
		users []*models.Users
	}
	tests := []struct {
		name string
		args args
		want map[int32]user
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertUsers(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
