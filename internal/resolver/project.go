package resolver

import (
	"main/db/sql/models"
	"main/internal/util"
	apiModels "main/server/api/v1/models"
)

func CreateProject(params apiModels.AddProject) error {
	var s *models.Projects

	return nil
}

func GetProjectsTree() ([]apiModels.TreeNode, error) {
	var tree []apiModels.TreeNode

	projects, err := DbProvider.GetAllProjects()
	if err != nil {
		return nil, err
	}

	for _, project := range projects {
		if project.ParentId == 0 {
			tree = append(tree, apiModels.TreeNode{
				Title: project.Name,
				Key:   project.Id,
			})
		}
	}
	for i, node := range tree {
		tree[i] = createNode(projects, node)
	}

	return tree, nil
}

func createNode(list []apiModels.Project, node apiModels.TreeNode) apiModels.TreeNode {
	for _, project := range list {
		if node.Key == project.ParentId {
			node.Children = []apiModels.TreeNode{}
			break
		}
	}

	if node.Children != nil {
		for _, project := range list {
			if node.Key == project.ParentId {
				node.Children = append(node.Children, apiModels.TreeNode{
					Title: project.Name,
					Key:   project.Id,
				})
			}
		}
	}

	return node
}

func GetProject(param apiModels.GetProject) (apiModels.Project, error) {
	from, err := DbProvider.DB.SelectAllFrom(models.ProjectsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return apiModels.Project{}, err
	}

	var project = apiModels.Project{}
	for _, s := range from {
		s := s.(*models.Projects)
		if s.Id == param.Id {
			project = apiModels.Project{
				Id:           s.Id,
				Name:         s.Name,
				Description:  *s.Description,
				CreationDate: s.CreationDate,
				OwnerId:      s.OwnerId,
			}
			if s.ParentId == nil {
				project.ParentId = 0
			} else {
				project.ParentId = *s.ParentId
			}
		}
	}

	return project, nil
}

func GetAllProjects() ([]apiModels.Project, error) {

	from, err := DbProvider.DB.SelectAllFrom(models.ProjectsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}
	var list []apiModels.Project
	for _, s := range from {

		s := s.(*models.Projects)
		q := apiModels.Project{
			Id:           s.Id,
			Name:         s.Name,
			Description:  *s.Description,
			CreationDate: s.CreationDate,
			OwnerId:      s.OwnerId,
		}
		if s.ParentId != nil {
			q.ParentId = *s.ParentId
		} else {
			q.ParentId = 0
		}
		list = append(list, q)
	}
	return list, nil
}

func DeleteProject(params apiModels.DeleteProject) error {
	record, err := DbProvider.DB.FindByPrimaryKeyFrom(models.ProjectsTable, params.Id)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	err = DbProvider.DB.Delete(record)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	return nil
}

func UpdateProject(params apiModels.UpdateProject) error {

	from, err := DbProvider.DB.FindByPrimaryKeyFrom(models.ProjectsTable, params.Id)
	if err != nil {
		return err //TODO
	}
	rec := from.(*models.Projects)

	s := &models.Projects{
		ParentId:    util.NilCheck(&params.ParentId, &rec.ParentId).(*int32),
		Name:        util.NilCheck(params.Name, rec.Name).(string),
		Description: util.NilCheck(params.Description, rec.Description).(*string),
		OwnerId:     util.NilCheck(params.OwnerId, rec.OwnerId).(int32),
	}

	err = DbProvider.DB.Update(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	return nil
}
