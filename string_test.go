package pgo

import (
	"github.com/stretchr/testify/assert"
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

func TestExplode(t *testing.T) {
	explode := Explode(",","hello,world")
	assert.Equal(t, "hello", explode[0])
	assert.Equal(t, "world", explode[1])
}

func TestStrpos(t *testing.T) {
	strpos := Strpos("+1s","s")
	assert.Equal(t, 2, strpos)
}

func TestStrLimit(t *testing.T) {
	s := StrLimit("测试2q文字超出，符号补充 1a",10,"...")
	assert.Equal(t, "测试2q文字超出，符...", s)
}