package cache

import "time"

var mockedCache = internalCache{
	Users: map[int32]user{},
	Tasks: map[int32]task{15: {
		Id:          15,
		Project:     0,
		Title:       "To delete",
		Description: "To delete",
		Column:      1,
		Tags:        nil,
		Users:       nil,
	}},
	Tags:  map[int32]tag{},
	Roles: map[int32]role{},
	Projects: map[int32]project{15: {
		Id:           15,
		Project:      0,
		Name:         "To delete",
		Description:  "To delete",
		CreationDate: time.Now(),
		Owner:        1,
		Users:        nil,
	}},
	Columns: map[int32]column{1: {
		Id:   1,
		Name: "Test",
	}},
	TaskAndUsers:    map[int32]taskAndUsers{},
	TaskAndTags:     map[int32]taskAndTags{},
	ProjectAndUsers: map[int32]projectAndUsers{},
}
