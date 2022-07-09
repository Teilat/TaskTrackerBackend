package cache

import (
	"log"
	"main/db/sql"
	"main/db/sql/models"
	"sync"
	"time"
)

const defaultProjectId = int32(0)
const defaultDescription = ""

type cacheTables interface {
	user | tag | task | role | project | column
}

var once sync.Once
var Cache cache

type cache struct {
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

type user struct {
	Id       int32
	Name     string
	Surname  string
	Nickname string
	Role     int32
	Password string
}

type task struct {
	Id          int32
	Project     int32
	Title       string
	Description string
	Column      int32
	Tags        map[int32]taskAndTags
	Users       map[int32]taskAndUsers
}

type tag struct {
	Id    int32
	Name  string
	Color string
}

type role struct {
	Id   int32
	Name string
}

type project struct {
	Id           int32
	Project      int32
	Name         string
	Description  string
	CreationDate time.Time
	Owner        int32
	Users        map[int32]projectAndUsers
}

type column struct {
	Id   int32
	Name string
}

type projectAndUsers struct {
	Id        int32
	ProjectId int32
	UserId    int32
}

type taskAndTags struct {
	Id     int32
	TaskId int32
	TagId  int32
}

type taskAndUsers struct {
	Id     int32
	TaskId int32
	UserId int32
}

func Init() {
	once.Do(func() {
		var err error
		tags, err := sql.GetAll[*models.Tags](models.TagsTable)
		roles, err := sql.GetAll[*models.Roles](models.RolesTable)
		columns, err := sql.GetAll[*models.Columns](models.ColumnsTable)

		users, err := sql.GetAll[*models.Users](models.UsersTable)
		projects, err := sql.GetAll[*models.Projects](models.ProjectsTable)
		tasks, err := sql.GetAll[*models.Tasks](models.TasksTable)

		projectNUsers, err := sql.GetAll[*models.ProjectAndUsers](models.ProjectAndUsersTable)
		taskNTags, err := sql.GetAll[*models.TaskAndTags](models.TaskAndTagsTable)
		taskNUsers, err := sql.GetAll[*models.TaskAndUsers](models.TaskAndUsersTable)

		if err != nil {
			log.Fatalf("Cache init err:%v\n", err)
		}

		Cache.Tags = convertTags(tags)
		Cache.Roles = convertRoles(roles)
		Cache.Columns = convertColumns(columns)

		Cache.ProjectAndUsers = convertProjectAndUsers(projectNUsers)
		Cache.TaskAndUsers = convertTaskAndUsers(taskNUsers)
		Cache.TaskAndTags = convertTaskAndTags(taskNTags)

		Cache.Users = convertUsers(users)
		Cache.Projects = convertProject(projects)
		Cache.Tasks = convertTasks(tasks)
	})
}

func convertProject(projects []*models.Projects) map[int32]project {
	res := make(map[int32]project, 0)
	for _, p := range projects {
		proj := project{
			Id:           p.Id,
			Project:      defaultProjectId,
			Name:         p.Name,
			CreationDate: p.CreationDate,
			Description:  defaultDescription,
			Owner:        p.OwnerId,
		}
		if p.ParentId != nil {
			proj.Project = *p.ParentId
		}
		if p.Description != nil {
			proj.Description = *p.Description
		}
		for _, user := range Cache.ProjectAndUsers {
			if user.ProjectId == proj.Id {
				proj.Users[user.Id] = user
			}
		}
		res[p.Id] = proj
	}
	return res
}

func convertTasks(tasks []*models.Tasks) map[int32]task {
	res := make(map[int32]task, 0)
	for _, t := range tasks {
		tas := task{
			Id:          t.Id,
			Project:     t.ProjectId,
			Title:       t.Title,
			Description: defaultDescription,
			Column:      t.ColumnId,
		}
		if t.Description != nil {
			tas.Description = *t.Description
		}
		for _, tag := range Cache.TaskAndTags {
			if tag.TaskId == tas.Id {
				tas.Tags[tag.Id] = tag
			}
		}
		for _, user := range Cache.TaskAndUsers {
			if user.TaskId == tas.Id {
				tas.Users[user.Id] = user
			}
		}
		res[t.Id] = tas
	}

	return res
}

func convertUsers(users []*models.Users) map[int32]user {
	res := make(map[int32]user, 0)
	for _, u := range users {
		res[u.Id] = user{
			Id:       u.Id,
			Name:     u.Name,
			Surname:  u.Surname,
			Nickname: u.Nickname,
			Password: u.Password,
			Role:     u.RoleId,
		}
	}
	return res
}

func convertTags(tags []*models.Tags) map[int32]tag {
	res := make(map[int32]tag, 0)
	for _, t := range tags {
		res[t.Id] = tag{
			Id:    t.Id,
			Name:  t.Name,
			Color: t.Color,
		}
	}
	return res
}

func convertProjectAndUsers(tags []*models.ProjectAndUsers) map[int32]projectAndUsers {
	res := make(map[int32]projectAndUsers, 0)
	for _, t := range tags {
		res[t.Id] = projectAndUsers{
			Id:        t.Id,
			ProjectId: t.ProjectId,
			UserId:    t.UserId,
		}
	}
	return res
}

func convertTaskAndUsers(tags []*models.TaskAndUsers) map[int32]taskAndUsers {
	res := make(map[int32]taskAndUsers, 0)
	for _, t := range tags {
		res[t.Id] = taskAndUsers{
			Id:     t.Id,
			TaskId: t.TaskId,
			UserId: t.UserId,
		}
	}
	return res
}

func convertTaskAndTags(tags []*models.TaskAndTags) map[int32]taskAndTags {
	res := make(map[int32]taskAndTags, 0)
	for _, t := range tags {
		res[t.Id] = taskAndTags{
			Id:     t.Id,
			TaskId: t.TaskId,
			TagId:  t.TagId,
		}
	}
	return res
}

func convertColumns(cols []*models.Columns) map[int32]column {
	res := make(map[int32]column, 0)
	for _, c := range cols {
		res[c.Id] = column{
			Id:   c.Id,
			Name: c.Name,
		}
	}
	return res
}
func convertRoles(roles []*models.Roles) map[int32]role {
	res := make(map[int32]role, 0)
	for _, r := range roles {
		res[r.Id] = role{
			Id:   r.Id,
			Name: r.Name,
		}
	}
	return res
}
