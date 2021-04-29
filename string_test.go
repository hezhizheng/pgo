package pgo

import (
	"log"
	"testing"
)

func TestMd5(t *testing.T) {
	if Md5("123456") != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("test Md5 fail")
	}
}

func TestMbStrlen(t *testing.T) {
	if MbStrlen("中文 1") != 4 {
		t.Error("test MbStrlen fail")
	}
}

func TestUniqid(t *testing.T) {
	uid := Uniqid("")
	log.Println(uid)
	if uid == "" {
		t.Error("test MbStrlen fail")
	}
}