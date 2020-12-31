package log

import "testing"

func TestDebug(t *testing.T) {
	Init("")
	t.Log(GetLevel())
	Debug("this is Debug log")
	Info("this is Info log")
	Warn("this is Warn log")
	Error("this is Error log")
	Fatal("this is Fatal log")
	SetLevel(LevelDebug)
	t.Log(GetLevel())
	Debug("this is Debug log")
	Info("this is Info log")
	Warn("this is Warn log")
	Error("this is Error log")
	Fatal("this is Fatal log")
}

func TestDebugF(t *testing.T) {
	Init("")
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
	SetLevel(LevelDebug)
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
}

func TestError(t *testing.T) {
	Init("")
	t.Log(GetLevel())
	Debug("this is Debug log")
	Info("this is Info log")
	Warn("this is Warn log")
	Error("this is Error log")
	Fatal("this is Fatal log")
	SetLevel(LevelError)
	t.Log(GetLevel())
	Debug("this is Debug log")
	Info("this is Info log")
	Warn("this is Warn log")
	Error("this is Error log")
	Fatal("this is Fatal log")
}

func TestErrorF(t *testing.T) {
	Init("")
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
	SetLevel(LevelError)
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
}

func TestFatal(t *testing.T) {
	Init("")
	t.Log(GetLevel())
	Debug("this is Debug log")
	Info("this is Info log")
	Warn("this is Warn log")
	Error("this is Error log")
	Fatal("this is Fatal log")
	SetLevel(LevelFatal)
	t.Log(GetLevel())
	Debug("this is Debug log")
	Info("this is Info log")
	Warn("this is Warn log")
	Error("this is Error log")
	Fatal("this is Fatal log")
}

func TestFatalF(t *testing.T) {
	Init("")
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
	SetLevel(LevelFatal)
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
}

func TestInfo(t *testing.T) {
	Init("")
	t.Log(GetLevel())
	Debug("this is Debug log")
	Info("this is Info log")
	Warn("this is Warn log")
	Error("this is Error log")
	Fatal("this is Fatal log")
	SetLevel(LevelInfo)
	t.Log(GetLevel())
	Debug("this is Debug log")
	Info("this is Info log")
	Warn("this is Warn log")
	Error("this is Error log")
	Fatal("this is Fatal log")
}

func TestInfoF(t *testing.T) {
	Init("")
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
	SetLevel(LevelInfo)
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
}

func TestInit(t *testing.T) {
	Init("log.txt")
	SetLevel(LevelDebug)
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
	SetLevel(LevelInfo)
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
}

func TestSetLevel(t *testing.T) {

}

func TestWarn(t *testing.T) {
	Init("")
	t.Log(GetLevel())
	Debug("this is Debug log")
	Info("this is Info log")
	Warn("this is Warn log")
	Error("this is Error log")
	Fatal("this is Fatal log")
	SetLevel(LevelWarn)
	t.Log(GetLevel())
	Debug("this is Debug log")
	Info("this is Info log")
	Warn("this is Warn log")
	Error("this is Error log")
	Fatal("this is Fatal log")
}

func TestWarnF(t *testing.T) {
	Init("")
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
	SetLevel(LevelWarn)
	t.Log(GetLevel())
	DebugF("this is Debug log: %d, %s", 123, "hahahh")
	InfoF("this is Info log: %d, %s", 123, "hahahh")
	WarnF("this is Warn log: %d, %s", 123, "hahahh")
	ErrorF("this is Error log: %d, %s", 123, "hahahh")
	FatalF("this is Fatal log: %d, %s", 123, "hahahh")
}

func Test_write(t *testing.T) {

}

func Test_writeF(t *testing.T) {

}
