package test

import (
	"mini-douyin/utils"
	"testing"
)

func TestMakeCover(t *testing.T) {
	err := utils.CutCover("2_v0300fg10000ca50vejc77ue9mnkpdkg.MP4")
	if err != nil {
		err.Error()
	}
}
