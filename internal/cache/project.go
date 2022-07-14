package cache

import (
	"fmt"
	"main/db/sql"
	dbModels "main/db/sql/models"
	"main/server/api/v1/models"
	"time"
)

func (c internalCache) CreateProject(params models.AddProject) error {
	proj := &dbModels.Projects{
		ParentId:     &params.ParentId,
		Name:         params.Name,
		Description:  &params.Description,
		CreationDate: time.Now(),
		OwnerId:      params.OwnerId,
	}
	err := sql.Upsert(proj)
	if err != nil {
		return err
	}
	return nil
}

func (c internalCache) UpdateProject(params models.UpdateProject) error {
	upd := map[int32]project{}
	upd[params.Id] = project{
		Id:          params.Id,
		Project:     params.ParentId,
		Name:        params.Name,
		Description: params.Description,
		Owner:       params.OwnerId,
	}

	err := c.updateProjects(upd)
	if err != nil {
		return err
	}
	return nil
}

func (c internalCache) updateProjects(upd map[int32]project) error {
	for id, proj := range upd {
		model := &dbModels.Projects{
			ParentId:     &proj.Project,
			Name:         proj.Name,
			Description:  &proj.Description,
			CreationDate: proj.CreationDate,
			OwnerId:      proj.Owner,
		}
		err := sql.Update(model)
		if err != nil {
			return err
		}
		c.Projects[id] = proj
	}
	return nil
}

func (c internalCache) DeleteProject(params models.DeleteProject) error {
	t, ok := c.Projects[params.Id]
	if !ok {
		return fmt.Errorf("projcet witn id:%d doesnt exist", params.Id)
	}
	err := sql.Delete(&dbModels.Projects{
		Id: t.Id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c internalCache) GetAllProjects() ([]*models.Project, error) {
	res := []*models.Project{}
	for _, p := range c.Projects {
		proj := &models.Project{
			Id:           p.Id,
			ParentId:     p.Project,
			Name:         p.Name,
			Description:  p.Description,
			CreationDate: p.CreationDate,
			OwnerId:      p.Owner,
		}
		res = append(res, proj)
	}
	return res, nil
}

func (c internalCache) GetProject(param models.GetProject) (*models.Project, error) {
	p, ok := c.Projects[param.Id]
	if !ok {
		return nil, fmt.Errorf("project with id:%d doesnt exist", param)
	}
	res := &models.Project{
		Id:           p.Id,
		ParentId:     p.Project,
		Name:         p.Name,
		Description:  p.Description,
		CreationDate: p.CreationDate,
		OwnerId:      p.Owner,
	}
	return res, nil
}

func (c internalCache) GetProjectTree() ([]*models.TreeNode, error) {
	tree := []*models.TreeNode{}

	for _, p := range c.Projects {
		if p.Project == 0 {
			tree = append(tree, &models.TreeNode{
				Title: p.Name,
				Key:   p.Id,
			})
		}
	}

	for i, node := range tree {
		tree[i] = c.createNode(node)
	}

	return tree, nil
}

func (c internalCache) createNode(node *models.TreeNode) *models.TreeNode {
	for _, p := range c.Projects {
		if node.Key == p.Project {
			node.Children = []models.TreeNode{}
			break
		}
	}

	if node.Children != nil {
		for _, p := range c.Projects {
			if node.Key == p.Project {
				node.Children = append(node.Children, models.TreeNode{
					Title: p.Name,
					Key:   p.Id,
				})
			}
		}
	}

	return node
}