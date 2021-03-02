package log

func append()  {

	//每次追加日志时判断当前segment是否需要滚动
	maybeRoll()
	logSegmentAppend()
}

func loadSegments()  {
	
}

func addSegment()  {

}


func maybeRoll()  {
	var enough = true
	if  !enough{
		roll()
	}
	//返回活跃的segment
}

func roll()  {
	
}

func flush()  {
	
}