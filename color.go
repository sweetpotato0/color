package color

import (
	"os"

	"fmt"

	"strconv"
	"strings"

	"github.com/mattn/go-isatty"
)

/*
 字背景:
	40:黑
	41:深红
	42:绿
	43:黄色
	44:蓝色
	45:紫色
	46:深绿
	47:白色

 字颜色：
	30:黑
	31:红
	32:绿
	33:黄
	34:蓝色
	35:紫色
	36:深绿
	37:白色

 ANSI控制码的说明：
	\033[0m 关闭所有属性
	\033[1m 设置高亮度
	\033[4m 下划线
	\033[5m 闪烁
	\033[7m 反显
	\033[8m 消隐
	\033[30m -- \33[37m 设置前景色
	\033[40m -- \33[47m 设置背景色
	\033[nA 光标上移n行
	\033[nB 光标下移n行
	\033[nC 光标右移n行
	\033[nD 光标左移n行
	\033[y;xH设置光标位置
	\033[2J 清屏
	\033[K 清除从光标到行尾的内容
	\033[s 保存光标位置
	\033[u 恢复光标位置
	\033[?25l 隐藏光标
	\033[?25h 显示光标
*/

type Attr int
type BgColor = Attr
type WdColor = Attr

const (
	BgBlack   BgColor = iota + 40 // 黑
	BgRed                         // 深红
	BgGreen                       // 绿
	BgYellow                      // 黄色
	BgBlue                        // 蓝色
	BgPurple                      // 紫色
	BgHiGreen                     // 深绿
	BgWhite                       // 白色
)

const (
	WdBlack   WdColor = iota + 30 // 黑
	WdRed                         // 深红
	WdGreen                       // 绿
	WdYellow                      // 黄色
	WdBlue                        // 蓝色
	WdPurple                      // 紫色
	WdHiGreen                     // 深绿
	WdWhite                       // 白色
)

var (
	IsTerm = !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd())

	Output = os.Stdout

	START  = "\033["
	FORMAT = "%s;m"
	END    = "\033[0m"
)

type Color struct {
	params  []Attr
	noColor bool
}

func (c *Color) println(format string, wd WdColor, bg BgColor, a ...interface{}) {
	fmt.Fprintf(Output, format, a...)
}

func (c *Color) canSetColor() bool {
	return !c.noColor
}

func (c *Color) start() string {
	return START
}

func (c *Color) end() string {
	return END
}

func (c *Color) features() string {

	feature := make([]string, len(c.params))
	for i, v := range c.params {
		feature[i] = strconv.Itoa(int(v))
	}

	return strings.Join(feature, ";")
}

func (c *Color) format() string {
	return fmt.Sprintf(FORMAT, c.features())
}

func (c *Color) wrap(format string, a ...interface{}) (string, []interface{}) {
	return c.start() + c.format() + format + c.end(), a
}

func (c *Color) AddAttr(attr ...Attr) {
	c.params = append(c.params, attr...)
}

func (c *Color) Println(format string, a ...interface{}) {
	format, a = c.wrap(format, a...)
	fmt.Fprintf(Output, format+"\n", a...)
}

func New(attr ...Attr) *Color {

	c := &Color{
		noColor: !IsTerm,
		params:  make([]Attr, len(attr)),
	}
	c.AddAttr(attr...)

	return c
}

func Common(format string, a ...interface{}) {
	c := New()
	c.Println(format, a...)
}

func Red(format string, a ...interface{}) {
	c := New()
	c.AddAttr(WdRed)
	c.Println(format, a...)
}

func Blue(format string, a ...interface{}) {
	c := New()
	c.AddAttr(WdBlue)
	c.Println(format, a...)
}

func Yellow(format string, a ...interface{}) {
	c := New()
	c.AddAttr(WdYellow)
	c.Println(format, a...)
}

func Green(format string, a ...interface{}) {
	c := New()
	c.AddAttr(WdGreen)
	c.Println(format, a...)
}

func Black(format string, a ...interface{}) {
	c := New()
	c.AddAttr(WdBlack)
	c.Println(format, a...)
}

func Purple(format string, a ...interface{}) {
	c := New()
	c.AddAttr(WdPurple)
	c.Println(format, a...)
}

func HiGreen(format string, a ...interface{}) {
	c := New()
	c.AddAttr(WdHiGreen)
	c.Println(format, a...)
}

func White(format string, a ...interface{}) {
	c := New()
	c.AddAttr(WdWhite)
	c.Println(format, a...)
}
