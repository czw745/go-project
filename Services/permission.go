package services

import (
	"go-project/config"
	"go-project/models"
	"go-project/structs"
)

const tbPs = "permissions"
const tbPc = "permission_category"
const tbRHP = "role_has_permission"

//GetParents ... get parent permission category
func GetParents() (parents []models.PermissionCategoryParentResponse, res structs.Response, err error) {
	childs, res, err := GetChilds()
	if err = config.DB.Table(tbPc).Where("parent_id is null").Find(&parents).Error; err != nil {
		res.Message = err.Error()
		return
	}
	for i := range parents {
		for _, child := range childs {
			if parents[i].ID == child.ParentID {
				parents[i].Child = append(parents[i].Child, child)
			}
		}
	}
	return
}

//GetChilds ... get child permission category
func GetChilds() (childs []models.PermissionCategoryChildResponse, res structs.Response, err error) {
	prs, res, err := GetPermissions()
	if err = config.DB.Table(tbPc).Where("parent_id != ?", "nil").Find(&childs).Error; err != nil {
		res.Message = err.Error()
		return
	}
	for i := range childs {
		for _, pr := range prs {
			if childs[i].ID == pr.PermissionCategoryID {
				childs[i].Permissions = append(childs[i].Permissions, pr)
			}
		}
	}
	return
}

//GetPermissions ... get permissions
func GetPermissions() (prs []models.PermissionResponse, res structs.Response, err error) {
	if err = config.DB.Table(tbPs).Find(&prs).Error; err != nil {
		res.Message = err.Error()
		return
	}
	return
}
