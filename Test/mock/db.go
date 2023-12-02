package db

// mock/stub 测试，当待测试的函数/对象的依赖关系很复杂，并且有些依赖不能直接创建，例如数据库连接、文件I/O等
// 这种场景就非常适合使用 mock/stub 测试。简单来说，就是用 mock 对象模拟依赖项的行为。

// 这里假设 DB 是代码中负责与数据库交互的部分(在这里用 map 模拟)，测试用例中不能创建真实的数据库连接。
// 这个时候，如果我们需要测试 GetFromDB 这个函数内部的逻辑，就需要 mock 接口 DB

// 第一步：使用 mockgen 生成 db_mock.go。一般传递三个参数。包含需要被mock的接口得到源文件source，生成的目标文件destination，包名package

// 在windows 下需要使用绝对路径
// mockgen -source=db.go -destination=db_mock.go -package=db
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err != nil {
		return value
	}

	return -1
}
