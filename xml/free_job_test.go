package xml

import "testing"

func TestMakeFreeStyleJob(t *testing.T) {
	var param Param
	param.Username = "icowan"
	param.UserToken = "helloworld"
	param.GitAddr = "git@github.com:icowan/kplcloud.git"
	param.GitToken = ""
	rs, err := MakeFreeStyleJob("hello", param, "")
	if err != nil {
		t.Fatalf(err.Error())
	}

	println(string(rs))
}

func TestMakeTestStyleJob(t *testing.T) {
	var param Param
	param.Username = "icowan"
	param.UserToken = "helloworld"
	param.GitAddr = "git@github.com:icowan/kplcloud.git"
	param.GitToken = ""
	rs, err := MakeTestStyleJob("hello", param, "")
	if err != nil {
		t.Fatalf(err.Error())
	}

	println(string(rs))
}
