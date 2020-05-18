package dbops

import "testing"

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}

func clearTables()  {
	db.Exec("truncate users")
	db.Exec("truncate video_info")
	db.Exec("truncate comments")
	db.Exec("truncate sessions")
}

func TestUserWorkFlow(t *testing.T) {
   t.Run("Add",TestGetUserCredential)
   t.Run("Get",TestGetUserCredential)
   t.Run("Delete",TestDeleteUser)
   t.Run("Reget",testRegetUser)
}

func TestAddUserCredential(t *testing.T) {
      err := AddUserCredential("jierui","123")
      if err != nil {
        t.Errorf("测试注册用户失败：%v",err)
      }
}

func TestGetUserCredential(t *testing.T) {
     _,err := GetUserCredential("jierui")
     if err != nil {
		 t.Errorf("测试获取用户失败：%v",err)
     }
}

func TestDeleteUser(t *testing.T) {
    err := DeleteUser("jierui","123")
    if err != nil {
		t.Errorf("测试删除用户失败：%v",err)
    }
}

func testRegetUser(t *testing.T)  {
	pwd,err := GetUserCredential("jierui")
	if err != nil {
		t.Errorf("测试获取用户失败222：%v",err)
	}

	if pwd != "" {
		t.Errorf("删除测试用户失败")
	}
}