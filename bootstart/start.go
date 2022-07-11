package bootstart

import log "github.com/cihub/seelog"

// 服务启动时初始化一系列配置
func Start() error {
	errLog := InitLog()
	if errLog != nil {
		return errLog
	}
	return nil
}

// 初始化日志
func InitLog() error {
	logger, err := log.LoggerFromConfigAsFile("conf/log/seelog.xml")
	if err != nil {
		return err
	}
	log.ReplaceLogger(logger)
	return nil
}
