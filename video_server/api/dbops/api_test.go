package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
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

//<editor-fold desc="测试video_info表">
func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", TestAddUserCredential)
	t.Run("AddVideo", TestAddNewVideo)
	t.Run("GetVideo", TestGetVideoInfo)
	t.Run("DelVideo", TestDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
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

//</editor-fold>

//<editor-fold desc="测试comments表">
func testComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", TestAddUserCredential)
	t.Run("AddComments", TestAddNewComments)
	t.Run("ListComments", TestListComments)
}

func TestAddNewComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"
	err := addNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("测试增加评论失败%v", err)
	}
}

func TestListComments(t *testing.T) {
	vid := "12345"
	from := 1560960000
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("测试获取评论list失败%v", err)
	}
	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}

//</editor-fold>

//<editor-fold desc="测试session的表">
func testSessions(t *testing.T) {
	clearTables()
	t.Run("InsertSession", TestInsertSession)
	t.Run("RetrieveSession", TestRetrieveSession)
	t.Run("RetrieveAllSession", TestRetrieveAllSession)
	t.Run("DeleteSession", TestDeleteSession)
}

func TestInsertSession(t *testing.T) {
	err := InsertSession("123456", 12345, "jieruian")
	if err != nil {
		t.Errorf("测试写入session失败%v", err)
	}
}

func TestRetrieveSession(t *testing.T) {
	res, err := RetrieveSession("123456")
	if err != nil {
		t.Errorf("测试获取session失败%v", err)
	}

	fmt.Printf("测试获取session的结果: %v \n", res)

}

func TestRetrieveAllSession(t *testing.T) {
	m, err := RetrieveAllSession()
	if err != nil {
		t.Errorf("测试获取所有的session失败%v", err)
	}
	m.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

func TestDeleteSession(t *testing.T) {
	err := DeleteSession("123456")
	if err != nil {
		t.Errorf("测试删除session失败%v", err)
	}
}

//</editor-fold>
