/*
	这是一个标准的package包。。。。。。我觉得每一个go文件都要有一个相对的测试文件
*/
package bootstrap

import (
	"sync"
	"testing"
)

var (
	bt   *Bootstrap
	once sync.Once
)

func testSetup() {
	bt = NewBootstrap(InitOptions{})
}

func TestBootstrap(t *testing.T) {
	once.Do(testSetup)
	bt.InitEnv()

}
