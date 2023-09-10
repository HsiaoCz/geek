package mysql

import "testing"

func TestGenAdmin(t *testing.T) {
	adminMap := make(map[string]string, 5)
	adminMap["zhangsan"] = "qsdf12344"
	adminMap["liss"] = "21223ased"
	adminMap["lost"] = "22333vsf"
	adminMap["hante"] = "lisii123"
	adminMap["more"] = "qwess3455"
	for username, password := range adminMap {
		InsertAdmin(username, password)
	}
}
