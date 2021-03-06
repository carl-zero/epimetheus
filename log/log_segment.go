package log

import (
	"bufio"
	"fmt"
	"os"
)
func logSegmentAppend(name,content string)  {
	if fileObj,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644);err == nil {
		defer fileObj.Close()
		writeObj := bufio.NewWriterSize(fileObj,4096)

		//使用Write方法,需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
		buf := []byte(content)
		if _,err := writeObj.Write(buf);err == nil {
			fmt.Println("Successful appending to the buffer with os.OpenFile and bufio's Writer obj Write method.",content)
			if  err := writeObj.Flush(); err != nil {panic(err)}
			fmt.Println("Successful flush the buffer data to file ",content)
		}
	}
}