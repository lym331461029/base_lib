package zap

import (
	"testing"

	"github.com/lym331461029/base_lib/log/base/field"
)

func TestZapLogger(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			switch i := p.(type) {
			case error, string:
				t.Fatalf("Fatal error: %s\n", i)
			default:
				t.Fatalf("Fatal error: %#v\n", i)
			}
		}
	}()
	logger := NewLogger("testing")
	logger = logger.WithFields(
		field.Bool("bool", false),
		field.Int64("int64", 12345678),
		field.Float64("float64", 123.456),
		field.String("string", "zap"),
		field.Object("object", interface{}("abcd")),
	)
	t.Logf("The tested logger: %s", logger.Name())
	logger.Infof("The tested logger: %s", logger.Name())
	logger.Info("Info log (zap)")
	logger.Infoln("Infoln log (zap)")
	logger.Error("Error log (zap)")
	logger.Errorf("%s log (zap)", "Errorf")
	logger.Errorln("Errorln log (zap)")
	logger.Warn("Warn log (zap)")
	logger.Warnf("%s log (zap)", "Warnf")
	logger.Warnln("Warnln log (zap)")

	// They will call os.Exit(1)
	// logger.Fatal("Fatal log (zap)")
	// logger.Fatalf("%s log (zap)", "Fatalf")
	// logger.Fatalln("Fatalln log (zap)")

	// They will cause panic
	// logger.Panic("Panic log (zap)")
	// logger.Panicf("%s log (zap)", "Panicf")
	// logger.Panicln("Panicln log (zap)")
}
