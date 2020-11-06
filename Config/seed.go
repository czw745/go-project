package config

import (
	"go-project/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Seed() {
	// Seed PermissionCategory
	a := uint(1)
	b := uint(6)
	permissionCategory := &[]models.PermissionCategory{
		{Name: "system", DisplayName: "後台系統", ParentID: nil},
		{Name: "auth", DisplayName: "平台管理", ParentID: &a},
		{Name: "member", DisplayName: "會員管理", ParentID: &a},
		{Name: "issue", DisplayName: "客服管理", ParentID: &a},
		{Name: "sale", DisplayName: "銷售管理", ParentID: &a},
		{Name: "frontend", DisplayName: "前台系統", ParentID: nil},
		{Name: "article", DisplayName: "文章管理", ParentID: &b},
		{Name: "theme", DisplayName: "主題管理", ParentID: &b},
	}
	DB.Create(permissionCategory)

	// Seed Permission
	permission := &[]models.Permission{
		{Name: "system auth user", DisplayName: "帳號權限設定", GuardName: "web", PermissionCategoryID: 2},
		{Name: "system auth role", DisplayName: "角色權限設定", GuardName: "web", PermissionCategoryID: 2},
		{Name: "system member profile", DisplayName: "會員資料", GuardName: "web", PermissionCategoryID: 3},
		{Name: "system issue record", DisplayName: "客服紀錄", GuardName: "web", PermissionCategoryID: 4},
		{Name: "system issue setting", DisplayName: "客服表單設定", GuardName: "web", PermissionCategoryID: 4},
		{Name: "system sale plan", DisplayName: "方案管理", GuardName: "web", PermissionCategoryID: 5},
		{Name: "system sale order", DisplayName: "訂單管理", GuardName: "web", PermissionCategoryID: 5},
		{Name: "system sale refund", DisplayName: "退訂管理", GuardName: "web", PermissionCategoryID: 5},
		{Name: "frontend article content", DisplayName: "內容設定", GuardName: "web", PermissionCategoryID: 7},
		{Name: "frontend article comment", DisplayName: "留言設定", GuardName: "web", PermissionCategoryID: 7},
		{Name: "frontend theme background", DisplayName: "背景設定", GuardName: "web", PermissionCategoryID: 8},
	}
	DB.Create(permission)

	// Seed Admin
	hash, _ := HashPassword("12345678")
	adminRole := &models.Role{Name: "Super Admin", DisplayName: "超級管理員", Status: 1, Deletable: 0}
	adminUser := &models.User{Name: "Admin", Email: "admin@example.com", Password: hash, Status: 1, Deletable: 0}
	DB.Create(adminUser).Association("Roles").Append(adminRole)
	DB.Find(adminRole).Association("Permission").Append(permission)

	roles := &[]models.Role{
		{Name: "Admin", DisplayName: "管理員", Status: 1, Deletable: 1},
		{Name: "Assistant", DisplayName: "助理", Status: 1, Deletable: 1},
	}
	DB.Create(roles)
}
