package dbops

import (
	"testing"
)

var tempvid string

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func clearTables() {
	db.Exec("truncate users")
	db.Exec("truncate video_info")
	db.Exec("truncate comments")
	db.Exec("truncate sessions")
}

//<editor-fold desc="测试users表操作">
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", TestGetUserCredential)
	t.Run("Get", TestGetUserCredential)
	t.Run("Delete", TestDeleteUser)
	t.Run("Reget", testRegetUser)
}

func TestAddUserCredential(t *testing.T) {
	err := AddUserCredential("jierui", "123")
	if err != nil {
		t.Errorf("测试注册用户失败：%v", err)
	}
}

func TestGetUserCredential(t *testing.T) {
	_, err := GetUserCredential("jierui")
	if err != nil {
		t.Errorf("测试获取用户失败：%v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("jierui", "123")
	if err != nil {
		t.Errorf("测试删除用户失败：%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("jierui")
	if err != nil {
		t.Errorf("测试获取用户失败222：%v", err)
	}

	if pwd != "" {
		t.Errorf("删除测试用户失败")
	}
}

//</editor-fold>

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", TestAddUserCredential)
	t.Run("AddVideo", TestAddNewVideo)
	//t.Run("GetVideo", TestGetVideoInfo)
	//t.Run("DelVideo", TestDeleteVideoInfo)
	//t.Run("RegetVideo", testRegetVideoInfo)
}

func TestAddNewVideo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("测试添加视频失败%v", err)
	}
	//fmt.Println(vi)
	tempvid = vi.Id
}

func TestGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("测试获取视频失败%v", err)
	}
}

func TestDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("测试删除视频失败%v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}
