package cache

import (
	"main/server/api/v1/models"
	"reflect"
	"testing"
)

func Test_internalCache_CreateProject(t *testing.T) {
	tests := []struct {
		name    string
		cache   internalCache
		params  models.AddProject
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			if err := tt.cache.CreateProject(tt.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_internalCache_DeleteProject(t *testing.T) {
	tests := []struct {
		name    string
		cache   internalCache
		params  models.DeleteProject
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			if err := tt.cache.DeleteProject(tt.params); (err != nil) != tt.wantErr {
				t.Errorf("DeleteProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_internalCache_GetAllProjects(t *testing.T) {
	tests := []struct {
		name    string
		cache   internalCache
		want    []*models.Project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			got, err := tt.cache.GetAllProjects()
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
	tests := []struct {
		name    string
		cache   internalCache
		param   models.GetProject
		want    *models.Project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			got, err := tt.cache.GetProject(tt.param)
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
	tests := []struct {
		name    string
		cache   internalCache
		want    []*models.TreeNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			got, err := tt.cache.GetProjectTree()
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
	tests := []struct {
		name    string
		cache   internalCache
		params  models.UpdateProject
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			if err := tt.cache.UpdateProject(tt.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_internalCache_createNode(t *testing.T) {
	tests := []struct {
		name  string
		cache internalCache
		node  *models.TreeNode
		want  *models.TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			if got := tt.cache.createNode(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_internalCache_updateProjects(t *testing.T) {
	tests := []struct {
		name    string
		cache   internalCache
		upd     map[int32]project
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			if err := tt.cache.updateProjects(tt.upd); (err != nil) != tt.wantErr {
				t.Errorf("updateProjects() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
