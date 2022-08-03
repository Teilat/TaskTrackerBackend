package cache

import (
	"main/server/api/v1/models"
	"reflect"
	"testing"
)

func Test_internalCache_CreateProject(t *testing.T) {
	type fields struct {
		Users           map[int32]user
		Tasks           map[int32]task
		Tags            map[int32]tag
		Roles           map[int32]role
		Projects        map[int32]project
		Columns         map[int32]column
		TaskAndUsers    map[int32]taskAndUsers
		TaskAndTags     map[int32]taskAndTags
		ProjectAndUsers map[int32]projectAndUsers
	}
	type args struct {
		params models.AddProject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := internalCache{
				Users:           tt.fields.Users,
				Tasks:           tt.fields.Tasks,
				Tags:            tt.fields.Tags,
				Roles:           tt.fields.Roles,
				Projects:        tt.fields.Projects,
				Columns:         tt.fields.Columns,
				TaskAndUsers:    tt.fields.TaskAndUsers,
				TaskAndTags:     tt.fields.TaskAndTags,
				ProjectAndUsers: tt.fields.ProjectAndUsers,
			}
			if err := c.CreateProject(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_internalCache_DeleteProject(t *testing.T) {
	type fields struct {
		Users           map[int32]user
		Tasks           map[int32]task
		Tags            map[int32]tag
		Roles           map[int32]role
		Projects        map[int32]project
		Columns         map[int32]column
		TaskAndUsers    map[int32]taskAndUsers
		TaskAndTags     map[int32]taskAndTags
		ProjectAndUsers map[int32]projectAndUsers
	}
	type args struct {
		params models.DeleteProject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := internalCache{
				Users:           tt.fields.Users,
				Tasks:           tt.fields.Tasks,
				Tags:            tt.fields.Tags,
				Roles:           tt.fields.Roles,
				Projects:        tt.fields.Projects,
				Columns:         tt.fields.Columns,
				TaskAndUsers:    tt.fields.TaskAndUsers,
				TaskAndTags:     tt.fields.TaskAndTags,
				ProjectAndUsers: tt.fields.ProjectAndUsers,
			}
			if err := c.DeleteProject(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("DeleteProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_internalCache_GetAllProjects(t *testing.T) {
	type fields struct {
		Users           map[int32]user
		Tasks           map[int32]task
		Tags            map[int32]tag
		Roles           map[int32]role
		Projects        map[int32]project
		Columns         map[int32]column
		TaskAndUsers    map[int32]taskAndUsers
		TaskAndTags     map[int32]taskAndTags
		ProjectAndUsers map[int32]projectAndUsers
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*models.Project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := internalCache{
				Users:           tt.fields.Users,
				Tasks:           tt.fields.Tasks,
				Tags:            tt.fields.Tags,
				Roles:           tt.fields.Roles,
				Projects:        tt.fields.Projects,
				Columns:         tt.fields.Columns,
				TaskAndUsers:    tt.fields.TaskAndUsers,
				TaskAndTags:     tt.fields.TaskAndTags,
				ProjectAndUsers: tt.fields.ProjectAndUsers,
			}
			got, err := c.GetAllProjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllProjects() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_internalCache_GetProject(t *testing.T) {
	type fields struct {
		Users           map[int32]user
		Tasks           map[int32]task
		Tags            map[int32]tag
		Roles           map[int32]role
		Projects        map[int32]project
		Columns         map[int32]column
		TaskAndUsers    map[int32]taskAndUsers
		TaskAndTags     map[int32]taskAndTags
		ProjectAndUsers map[int32]projectAndUsers
	}
	type args struct {
		param models.GetProject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := internalCache{
				Users:           tt.fields.Users,
				Tasks:           tt.fields.Tasks,
				Tags:            tt.fields.Tags,
				Roles:           tt.fields.Roles,
				Projects:        tt.fields.Projects,
				Columns:         tt.fields.Columns,
				TaskAndUsers:    tt.fields.TaskAndUsers,
				TaskAndTags:     tt.fields.TaskAndTags,
				ProjectAndUsers: tt.fields.ProjectAndUsers,
			}
			got, err := c.GetProject(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_internalCache_GetProjectTree(t *testing.T) {
	type fields struct {
		Users           map[int32]user
		Tasks           map[int32]task
		Tags            map[int32]tag
		Roles           map[int32]role
		Projects        map[int32]project
		Columns         map[int32]column
		TaskAndUsers    map[int32]taskAndUsers
		TaskAndTags     map[int32]taskAndTags
		ProjectAndUsers map[int32]projectAndUsers
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*models.TreeNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := internalCache{
				Users:           tt.fields.Users,
				Tasks:           tt.fields.Tasks,
				Tags:            tt.fields.Tags,
				Roles:           tt.fields.Roles,
				Projects:        tt.fields.Projects,
				Columns:         tt.fields.Columns,
				TaskAndUsers:    tt.fields.TaskAndUsers,
				TaskAndTags:     tt.fields.TaskAndTags,
				ProjectAndUsers: tt.fields.ProjectAndUsers,
			}
			got, err := c.GetProjectTree()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjectTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProjectTree() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_internalCache_UpdateProject(t *testing.T) {
	type fields struct {
		Users           map[int32]user
		Tasks           map[int32]task
		Tags            map[int32]tag
		Roles           map[int32]role
		Projects        map[int32]project
		Columns         map[int32]column
		TaskAndUsers    map[int32]taskAndUsers
		TaskAndTags     map[int32]taskAndTags
		ProjectAndUsers map[int32]projectAndUsers
	}
	type args struct {
		params models.UpdateProject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := internalCache{
				Users:           tt.fields.Users,
				Tasks:           tt.fields.Tasks,
				Tags:            tt.fields.Tags,
				Roles:           tt.fields.Roles,
				Projects:        tt.fields.Projects,
				Columns:         tt.fields.Columns,
				TaskAndUsers:    tt.fields.TaskAndUsers,
				TaskAndTags:     tt.fields.TaskAndTags,
				ProjectAndUsers: tt.fields.ProjectAndUsers,
			}
			if err := c.UpdateProject(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_internalCache_createNode(t *testing.T) {
	type fields struct {
		Users           map[int32]user
		Tasks           map[int32]task
		Tags            map[int32]tag
		Roles           map[int32]role
		Projects        map[int32]project
		Columns         map[int32]column
		TaskAndUsers    map[int32]taskAndUsers
		TaskAndTags     map[int32]taskAndTags
		ProjectAndUsers map[int32]projectAndUsers
	}
	type args struct {
		node *models.TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := internalCache{
				Users:           tt.fields.Users,
				Tasks:           tt.fields.Tasks,
				Tags:            tt.fields.Tags,
				Roles:           tt.fields.Roles,
				Projects:        tt.fields.Projects,
				Columns:         tt.fields.Columns,
				TaskAndUsers:    tt.fields.TaskAndUsers,
				TaskAndTags:     tt.fields.TaskAndTags,
				ProjectAndUsers: tt.fields.ProjectAndUsers,
			}
			if got := c.createNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_internalCache_updateProjects(t *testing.T) {
	type fields struct {
		Users           map[int32]user
		Tasks           map[int32]task
		Tags            map[int32]tag
		Roles           map[int32]role
		Projects        map[int32]project
		Columns         map[int32]column
		TaskAndUsers    map[int32]taskAndUsers
		TaskAndTags     map[int32]taskAndTags
		ProjectAndUsers map[int32]projectAndUsers
	}
	type args struct {
		upd map[int32]project
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := internalCache{
				Users:           tt.fields.Users,
				Tasks:           tt.fields.Tasks,
				Tags:            tt.fields.Tags,
				Roles:           tt.fields.Roles,
				Projects:        tt.fields.Projects,
				Columns:         tt.fields.Columns,
				TaskAndUsers:    tt.fields.TaskAndUsers,
				TaskAndTags:     tt.fields.TaskAndTags,
				ProjectAndUsers: tt.fields.ProjectAndUsers,
			}
			if err := c.updateProjects(tt.args.upd); (err != nil) != tt.wantErr {
				t.Errorf("updateProjects() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
