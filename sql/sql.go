package sql

//用户表
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Status   int    `json:"status`
}

//登录日志
type LoginLog struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Time      int64  `json:"time"`
	Ip        string `json:"ip"`
	Useragent string `json:"useragent"`
	Uid       int    `json:"uid"`
}

//菜单
type Menu struct {
	Id       int    `json:"id"`
	Pid      int    `json:"pid"`    // 父级menu
	Name     string `json:"name"`   // menu名称
	Sort     int    `json:"sort"`   // 排序值
	Path     string `json:"path"`   // 跳转路由
	Icon     string `json:"icon"`   // 图标
	Isshow   int    `json:"isshow"` // 是否显示
	Level    int    `json:"level"`  //菜单等级
	Children []Menu `json:"children"`
}

//递归重新排序无限极分类
func RecursiveMenu(arr []Menu, pid int, level int) (ar []Menu) {
	array := make([]Menu, 0)
	for k, v := range arr {
		if pid == v.Pid {

			arr[k].Level = level + 1
			// fmt.Println(arr[k])
			array = append(array, arr[k])
			// fmt.Printf("%#v", array)

		}
	}
	for tk, tv := range array {
		rm := RecursiveMenu(arr, tv.Id, level+1)
		for sk := range rm {
			array[tk].Children = append(array[tk].Children, rm[sk])
		}

		// array = append(array, rm[km])
		// array[km].Level = array[km].Level + 1
	}
	return array
}
