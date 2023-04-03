package log

import (
	"flag"

	"k8s.io/klog/v2"
)

func InitKlog(path, file string) {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "false")
	flag.Set("alsologtostderr", "false")
	//flag.Set("stderrthreshold", "FATAL")
	flag.Set("log_dir", path)
	flag.Set("log_file", file)

	flag.Parse()
}
