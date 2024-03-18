package core

import (
	"bytes"
	"fmt"
	"goblog/global"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37

	allLog  = "all"
	errLog  = "err"
	warnLog = "warn"
	infoLog = "info"
)

type LogFormatter struct{}

func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义时间格式
	log := global.Config.Logger
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		funVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "[%s] [%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestamp, levelColor, entry.Level, fileVal, funVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] [%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

type FileLevelHook struct {
	file     *os.File
	errFile  *os.File
	warnFile *os.File
	infoFile *os.File
	logPath  string
}

func (hook FileLevelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook FileLevelHook) Fire(entry *logrus.Entry) error {
	line, _ := entry.String()
	switch entry.Level {
	case logrus.ErrorLevel:
		hook.errFile.Write([]byte(line))
	case logrus.WarnLevel:
		hook.warnFile.Write([]byte(line))
	case logrus.InfoLevel:
		hook.infoFile.Write([]byte(line))
	default:
		hook.file.Write([]byte(line))
	}
	return nil
}

func InitFile(logPath string) (hook *FileLevelHook, err error) {
	err = os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		logrus.Errorf("创建日志目录失败: %v", err)
		return
	}
	allFile, err := os.OpenFile(path.Join(logPath, allLog+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		logrus.Errorf("创建日志文件失败: %v", err)
		return
	}
	errFile, err := os.OpenFile(path.Join(logPath, errLog+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		logrus.Errorf("创建日志文件失败: %v", err)
		return
	}
	warnFile, err := os.OpenFile(path.Join(logPath, warnLog+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		logrus.Errorf("创建日志文件失败: %v", err)
		return
	}
	infoFile, err := os.OpenFile(path.Join(logPath, infoLog+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		logrus.Errorf("创建日志文件失败: %v", err)
		return
	}
	hook = &FileLevelHook{allFile, errFile, warnFile, infoFile, logPath}

	return
}

func InitLogger() *logrus.Logger {
	mLogger := logrus.New() // 新建一个实例
	if global.Config.Logger.LogInConsole {
		mLogger.SetOutput(os.Stdout) // 设置输出类型为标准输出, 标准输出是指向控制台的
	}
	mLogger.SetReportCaller(global.Config.Logger.ShowLine)      // 开启报告调用函数的文件和行号
	mLogger.SetFormatter(&LogFormatter{})                       //设置自定义的formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level) // 解析level
	if err != nil {
		mLogger.SetLevel(logrus.DebugLevel) //设置最低的level
	} else {
		mLogger.SetLevel(level)
	}

	InitDefaultLogger() // 让全局logger也使用这个配置
	hook, err := InitFile(global.Config.Logger.Directory)
	if err == nil {
		mLogger.AddHook(hook)
		logrus.AddHook(hook)
	}
	return mLogger
}

func InitDefaultLogger() {
	// 全局log 配置
	if global.Config.Logger.LogInConsole {
		logrus.SetOutput(os.Stdout) // 设置输出类型为标准输出, 标准输出是指向控制台的
	}
	logrus.SetReportCaller(global.Config.Logger.ShowLine)       // 开启报告调用函数的文件和行号
	logrus.SetFormatter(&LogFormatter{})                        //设置自定义的formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level) // 解析level
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel) //设置最低的level
	} else {
		logrus.SetLevel(level)
	}
}
