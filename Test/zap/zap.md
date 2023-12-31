## zap

zap 支持七种日志级别：

Debug、Info、Warn、Error、DPanic、Panic、Fatal、其中 DPanic 是指在开发环境下(development)记录日志后会进行 panic

zap 针对生产环境和开发环境提供了不同的函数来创建 Logger 对象。

如果想在日志后面追加 key-value，则需要根据 value 的数据类型使用 zap.String、zap.Int 等方法实现。这一点在使用上显然不如 Logrus 等其他日志库来的方便，但这也是 zap 速度快的原因之一，zap 内部尽量避免使用 interface{} 和反射来提高代码执行效率

记录日志的方法签名：

```go
func (log *Logger) Info(msg string, fields ...Field)
```

其中 fields 是 zapcore.Field 类型，用来存储 key-value，并记录 value 类型，不管是 zap.String 还是 zap.Int 底层都是 zapcore.Field 类型来记录的。zap 为每一种 Go 的内置类型都定义了对应的 zap.Xxx 方法，甚至还实现 zap.Any() 来支持 interface{}。

通过 zap.NewProduction() 创建的日志对象输出格式为 JSON，而通过 zap.NewDevelopment() 创建的日志对象输出格式为 Text，日志后面追加的 key-value 会被转换成 JSON。并且，两者输出的字段内容也略有差异，如生产环境日志输出的时间格式为 Unix epoch 利于程序解析，而开发环境日志输出的时间格式为 ISO8601 更利于人类阅读。

zap.NewProduction()和
zap.NewDevelopment()的源码

```go
func NewProduction(options ...Option) (*Logger, error) {
	return NewProductionConfig().Build(options...)
}

func NewProductionConfig() Config {
	return Config{
		Level:       NewAtomicLevelAt(InfoLevel),
		Development: false,
		Sampling: &SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func NewDevelopment(options ...Option) (*Logger, error) {
	return NewDevelopmentConfig().Build(options...)
}

func NewDevelopmentConfig() Config {
	return Config{
		Level:            NewAtomicLevelAt(DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

```

zap.Config()的源码:

```go
type Config struct {
    // 日志级别
	Level AtomicLevel `json:"level" yaml:"level"`
    // 是否为开发模式
	Development bool `json:"development" yaml:"development"`
    // 禁用调用信息，值为 true 时，日志中将不再显示记录日志时所在的函数调用文件名和行号
	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`
    // 禁用堆栈跟踪捕获
	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`
    // 采样策略配置，单位为每秒，作用是限制日志在每秒内的输出数量，以此来防止全局的 CPU 和 I/O 负载过高
	Sampling *SamplingConfig `json:"sampling" yaml:"sampling"`
    // 指定日志编码器，目前支持 json 和 console
	Encoding string `json:"encoding" yaml:"encoding"`
    // 编码配置，决定了日志字段格式
	EncoderConfig zapcore.EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`
    // 配置日志输出位置，URLs 或文件路径，可配置多个
	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`
    // zap 包内部出现错误的日志输出位置，URLs 或文件路径，可配置多个，默认 os.Stderr。
	ErrorOutputPaths []string `json:"errorOutputPaths" yaml:"errorOutputPaths"`
    // 初始化字段配置，该配置的字段会以结构化的形式打印在每条日志输出中
	InitialFields map[string]interface{} `json:"initialFields" yaml:"initialFields"`
}
```

定制 logger:

```go
func newCustomLogger() (*zap.Logger, error) {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "", // 不记录日志调用位置
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", "test.log"},
		ErrorOutputPaths: []string{"error.log"},
	}
	return cfg.Build()
}

func main() {
	logger, _ := newCustomLogger()
	defer logger.Sync()

	// 增加一个 skip 选项，触发 zap 内部 error，将错误输出到 error.log
	logger = logger.WithOptions(zap.AddCallerSkip(100))

	logger.Info("Info msg")
	logger.Error("Error msg")
}

```

这里通过 logger.WithOptions() 为 Logger 对象增加了一个选项 zap.AddCallerSkip(100)，这个选项的作用是指定在通过调用栈获得行号时跳过的调用深度，因为我们的函数调用栈并不是 100 层，所以会触发 zap 内部错误，zap 会将错误日志输出到 ErrorOutputPaths 配置指定的位置中，即 error.log。

logger.WithOptions()支持一下配置:

- WrapCore(f func(zapcore.Core) zapcore.Core): 使用一个新的 zapcore.Core 替换掉 Logger 内部原有的的 zapcore.Core 属性。
- Hooks(hooks ...func(zapcore.Entry) error): 注册钩子函数，用来在日志打印时同时调用注册的钩子函数。
- Fields(fs ...Field): 添加公共字段。
- ErrorOutput(w zapcore.WriteSyncer): 指定日志组件内部出现异常时的输出位置。
- Development(): 将日志记录器设为开发模式，这将使 DPanic 级别日志记录错误后执行 panic()。
- AddCaller(): 与 WithCaller(true) 等价。
- WithCaller(enabled bool): 指定是否在日志输出内容中增加调用信息，即文件名和行号。
- AddCallerSkip(skip int): 指定在通过调用栈获取文件名和行号时跳过的调用深度。
- AddStacktrace(lvl zapcore.LevelEnabler): 用来指定某个日志级别及以上级别输出调用堆栈。
- IncreaseLevel(lvl zapcore.LevelEnabler): 提高日志级别，如果传入的 lvl 比现有级别低，则不会改变日志级别。
- WithFatalHook(hook zapcore.CheckWriteHook): 当出现 Fatal 级别日志时调用的钩子函数。
- WithClock(clock zapcore.Clock): 指定日志记录器用来确定当前时间的 zapcore.Clock 对象，默认为 time.Now 的系统时钟