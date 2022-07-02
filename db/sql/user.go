package sql

import (
	"main/db/sql/models"
	apiModels "main/server/api/v1/models"
)

func (DbProvider DatabaseProvider) CreateNewUser(params apiModels.AddUser) error {
	s := &models.Users{
		Name:     params.Name,
		Surname:  params.Surname,
		Nickname: params.Nickname,
		RoleId:   1,
		Password: params.Password,
	}

	err := DbProvider.DB.Save(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}
	return nil
}

func (DbProvider DatabaseProvider) GetUsersCredentials() ([]apiModels.UserCredentials, error) {

	from, err := DbProvider.DB.SelectAllFrom(models.UsersTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}
	var list []apiModels.UserCredentials
	for _, s := range from {

		s := s.(*models.Users)
		q := apiModels.UserCredentials{
			Id:       s.Id,
			Nickname: s.Nickname,
			Password: s.Password,
		}
		list = append(list, q)
	}
	return list, nil
}

func (DbProvider DatabaseProvider) GetAllUsers() ([]apiModels.User, error) {

	from, err := DbProvider.DB.SelectAllFrom(models.UsersTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}
	var list []apiModels.User
	for _, s := range from {

		s := s.(*models.Users)
		q := apiModels.User{
			Id:       s.Id,
			RoleId:   s.RoleId,
			Name:     s.Name,
			Surname:  s.Surname,
			Nickname: s.Nickname,
		}
		list = append(list, q)
	}
	return list, nil
}

func (DbProvider DatabaseProvider) DeleteUser(params apiModels.DeleteUser) error {

	record, err := DbProvider.DB.FindByPrimaryKeyFrom(models.UsersTable, params.Id)
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

func (DbProvider DatabaseProvider) UpdateUser(params apiModels.UpdateUser) error {

	from, err := DbProvider.DB.FindByPrimaryKeyFrom(models.UsersTable, params.Id)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}
	rec := from.(*models.Users)

	s := &models.Users{
		RoleId:   NilCheck(params.RoleId, rec.RoleId).(int32),
		Name:     NilCheck(params.Name, rec.Name).(string),
		Surname:  NilCheck(params.Surname, rec.Surname).(string),
		Nickname: NilCheck(params.Nickname, rec.Nickname).(string),
	}

	err = DbProvider.DB.Update(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	return nil
}

func (DbProvider DatabaseProvider) UpdateUserPassword(params apiModels.UpdateUserPassword) error {

	from, err := DbProvider.DB.FindByPrimaryKeyFrom(models.UsersTable, params.Id)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}
	rec := from.(*models.Users)

	s := &models.Users{
		Password: NilCheck(params.Password, rec.Password).(string),
	}

	err = DbProvider.DB.Update(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	return nil
}
