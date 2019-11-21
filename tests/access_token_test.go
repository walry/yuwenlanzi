package test

import (
	"fmt"
	"testing"
	"yuwenlanzi/service/wechat"
	"yuwenlanzi/service/wechat/menu"
)

func TestGetDefineMenu(t *testing.T){
	js := menu.GetDefineMenu()
	fmt.Print("js----",js,"\n")
}

func TestRun(t *testing.T)  {
	wechat.Run()
}