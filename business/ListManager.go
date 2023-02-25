package business

import (
	dao "github.com/mannuR22/go-task-backend.git/dao"
	models "github.com/mannuR22/go-task-backend.git/models"
)

func CreateList(userList models.ListModel) (models.ListCreatedVM, error) {
	ref, err := dao.CreateList(userList)

	var listVM models.ListCreatedVM

	if err != nil {
		return listVM, err
	}

	listVM.RefId = ref

	return listVM, nil
}

func ReadList(refId string) (models.ListModel, error) {

	userList, err := dao.ReadList(refId)

	return userList, err
}

func UpdateListItems(refId string, userListItems models.ListModel) error {
	return dao.UpdateListItems(refId, userListItems)
}

func DeleteListItems(refId string, userListItems models.ListModel) error {
	return dao.DeleteListItems(refId, userListItems)
}
