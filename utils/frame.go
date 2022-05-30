package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var (
	inputVideoPathOption = "-i"
	autoReWriteOption    = "-y"
	formatToImageOption  = "-f"
	startTimeOption      = "-ss"
	framesOption         = "-vframes"
)

// CutCover 使用ffmpeg切割视频的第一帧，作为封面
func CutCover(filename string) error {
	//ffmpeg -i ../public/2_v0300fg10000ca50vejc77ue9mnkpdkg.MP4 -y -f image2 -ss 00:00:01 -vframes 1 ../public/test.jpg
	src := "./temp/" + filename
	dst := "./temp/" + strings.Split(filename, ".")[0] + ".jpg"
	cmdArguments := []string{inputVideoPathOption, src, autoReWriteOption, formatToImageOption, "image2",
		startTimeOption, "00:00:01", framesOption, "1", dst}

	cmd := exec.Command("ffmpeg", cmdArguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("command output: %q", out.String())
		return errors.New("截取封面帧失败")
	}
	str := out.String()
	fmt.Println(str)
	fmt.Printf("command output: %q", out.String())
	return nil
}
