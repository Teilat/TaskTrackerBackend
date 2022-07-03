package sql

import (
	"main/db/sql/models"
	"main/internal/util"
	apiModels "main/server/api/v1/models"
	"time"
)

func (DbProvider DatabaseProvider) CreateProject(params apiModels.AddProject) error {
	var s *models.Projects
	if params.ParentId == 0 {
		s = &models.Projects{
			ParentId:     nil,
			Name:         params.Name,
			Description:  &params.Description,
			CreationDate: time.Now(),
			OwnerId:      params.OwnerId,
		}
	} else {
		s = &models.Projects{
			ParentId:     &params.ParentId,
			Name:         params.Name,
			Description:  &params.Description,
			CreationDate: time.Now(),
			OwnerId:      params.OwnerId,
		}
	}

	err := DbProvider.DB.Save(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}
	return nil
}

func (DbProvider DatabaseProvider) GetProjectsTree() ([]apiModels.TreeNode, error) {
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

func (DbProvider DatabaseProvider) GetProject(param apiModels.GetProject) (apiModels.Project, error) {
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

func (DbProvider DatabaseProvider) GetAllProjects() ([]apiModels.Project, error) {

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

func (DbProvider DatabaseProvider) DeleteProject(params apiModels.DeleteProject) error {
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

func (DbProvider DatabaseProvider) UpdateProject(params apiModels.UpdateProject) error {

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
