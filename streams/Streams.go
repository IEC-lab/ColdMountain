package streams

func DiscoverStreams() []string {
	return []string{"rtsp://admin:HadoopAdmin@192.168.1.108:554/cam/realmonitor?channel=1&subtype=0",
					"http://ivi.bupt.edu.cn/hls/cctv1hd.m3u8"}
}
