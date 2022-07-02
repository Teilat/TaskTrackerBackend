package cache

import (
	"main/db/sql"
	"main/db/sql/models"
	"sync"
	"time"
)

var once sync.Once
var Cache cache

type cache struct {
	Users    map[int32]user
	Tasks    map[int32]task
	Tags     map[int32]tag
	Roles    map[int32]role
	Projects map[int32]project
	Columns  map[int32]column
}

type user struct {
	Id       int32
	Name     string
	Surname  string
	Nickname string
	Role     role
	Password string
}

type task struct {
	Id          int32
	Project     project
	Title       string
	Description string
	Column      column
	Tags        map[int32]tag
	Users       map[int32]user
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
	Owner        user
	Users        map[int32]user
}

type column struct {
	Id   int32
	Name string
}

func Init() error {
	once.Do(func() {
		tags := sql.GetAll[*models.Tags](models.TagsTable)
		roles := sql.GetAll[*models.Roles](models.RolesTable)
		columns := sql.GetAll[*models.Columns](models.ColumnsTable)

		users := sql.GetAll[*models.Users](models.UsersTable)
		projects := sql.GetAll[*models.Projects](models.ProjectsTable)
		tasks := sql.GetAll[*models.Tasks](models.TasksTable)

		//projectAndUsers := sql.GetAll[*models.ProjectAndUsers](models.ProjectAndUsersTable)
		//taskAndTags := sql.GetAll[*models.TaskAndTags](models.TaskAndTagsTable)
		//taskAndUsers := sql.GetAll[*models.TaskAndUsers](models.TaskAndUsersTable)

		Cache.Tags = convertTags(tags)
		Cache.Roles = convertRoles(roles)
		Cache.Columns = convertColumns(columns)

		Cache.Users = convertUsers(users)
		Cache.Projects = convertProject(projects)
		Cache.Tasks = convertTasks(tasks)
	})
	return nil
}

func convertProject(projects []*models.Projects) map[int32]project {
	res := make(map[int32]project, 0)
	for _, p := range projects {
		res[p.Id] = project{
			Id:           p.Id,
			Project:      *p.ParentId,
			Name:         p.Name,
			Description:  *p.Description,
			CreationDate: p.CreationDate,
			Owner:        Cache.Users[p.OwnerId],
			//Users:      use, //TODO
		}
	}
	return res
}

func convertTasks(tasks []*models.Tasks) map[int32]task {
	res := make(map[int32]task, 0)
	for _, t := range tasks {
		res[t.Id] = task{
			Id:          t.Id,
			Project:     Cache.Projects[t.ProjectId],
			Title:       t.Title,
			Description: *t.Description,
			Column:      Cache.Columns[t.ColumnId],
			//Tags:        tags,	//TODO
			//Users:       users,	//TODO
		}
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
			Role:     Cache.Roles[u.RoleId],
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
