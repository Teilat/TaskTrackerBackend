package cache

import (
	"main/server/api/v1/models"
	"reflect"
	"testing"
)

func Test_internalCache_CreateTask(t *testing.T) {
	tests := []struct {
		name    string
		cache   internalCache
		params  models.AddTask
		wantErr bool
	}{
		{
			name:  "all good",
			cache: mockedCache,
			params: models.AddTask{
				ProjectId:   0,
				Title:       "test",
				Description: "test",
				Column:      "Test",
			},
			wantErr: false,
		},
		{
			name:  "column not exist",
			cache: mockedCache,
			params: models.AddTask{
				ProjectId:   0,
				Title:       "test",
				Description: "test",
				Column:      "asdasd",
			},
			wantErr: true,
		},
		{
			name:  "empty title",
			cache: mockedCache,
			params: models.AddTask{
				ProjectId:   0,
				Title:       "",
				Description: "test",
				Column:      "Test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.cache
			if err := c.CreateTask(tt.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_internalCache_DeleteTask(t *testing.T) {
	tests := []struct {
		name    string
		cache   internalCache
		params  models.DeleteTask
		wantErr bool
	}{
		{
			name:  "all good",
			cache: mockedCache,
			params: models.DeleteTask{
				Id: 2,
			},
			wantErr: false,
		},
		{
			name:  "task not exist",
			cache: mockedCache,
			params: models.DeleteTask{
				Id: 251,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cache.DeleteTask(tt.params); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_internalCache_GetAllTasks(t *testing.T) {
	tests := []struct {
		name  string
		cache internalCache
		want  []*models.Task
	}{
		{
			name:  "",
			cache: mockedCache,
			want: []*models.Task{{
				Id:          0,
				ProjectId:   0,
				Title:       "",
				Description: "",
				Column:      "",
				Users:       nil,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cache.GetAllTasks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_internalCache_GetTask(t *testing.T) {
	tests := []struct {
		name    string
		cache   internalCache
		param   int32
		want    *models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			got, err := tt.cache.GetTask(tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_internalCache_GetTasksByProject(t *testing.T) {
	tests := []struct {
		name   string
		cache  internalCache
		params models.TasksByProject
		want   []*models.Task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			if got := tt.cache.GetTasksByProject(tt.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTasksByProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_internalCache_UpdateTask(t *testing.T) {
	tests := []struct {
		name    string
		cache   internalCache
		params  models.UpdateTask
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			if err := tt.cache.UpdateTask(tt.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_internalCache_UpdateTaskPos(t *testing.T) {
	tests := []struct {
		name    string
		cache   internalCache
		params  models.UpdateTaskPos
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			if err := tt.cache.UpdateTaskPos(tt.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTaskPos() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_internalCache_updateTasks(t *testing.T) {
	tests := []struct {
		name    string
		cache   internalCache
		upd     map[int32]task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cache = mockedCache
			if err := tt.cache.updateTasks(tt.upd); (err != nil) != tt.wantErr {
				t.Errorf("updateTasks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
