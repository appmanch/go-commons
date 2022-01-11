package logging

//LogConfig - Configuration & Settings for the logger.
type LogConfig struct {

	//Format of the log. valid values are text,json-codec
	//Default is text
	Format string `json-codec:"format,omitempty" yaml:"format,omitempty"`
	//Async Flag to indicate if the writing of the flag is asynchronous.
	//Default value is false
	Async bool `json-codec:"async,omitempty" yaml:"async,omitempty"`
	//QueueSize to indicate the number log routines that can be queued  to use in background
	//This value is used only if the async value is set to true.
	//Default value for the number items to be in queue 512
	QueueSize int `json-codec:"queue_size,omitempty" yaml:"queueSize,omitempty"`
	//Date - Defaults to  time.RFC3339 pattern
	DatePattern string `json-codec:"datePattern,omitempty" yaml:"datePattern,omitempty"`
	//IncludeFunction will include the calling function name  in the log entries
	//Default value : false
	IncludeFunction bool `json-codec:"includeFunction,omitempty" yaml:"includeFunction,omitempty"`
	//IncludeLineNum ,includes Line number for the log file
	//If IncludeFunction Line is set to false this config is ignored
	IncludeLineNum bool `json-codec:"includeLineNum,omitempty" yaml:"includeLineNum,omitempty"`
	//DefaultLvl that will be used as default
	DefaultLvl string `json-codec:"defaultLvl" yaml:"defaultLvl"`
	//PackageConfig that can be used to
	PkgConfigs []*PackageConfig `json-codec:"pkgConfigs" yaml:"pkgConfigs"`
	//Writers writers for the logger. Need one for all levels
	//If a writer is not found for a specific level it will fallback to os.Stdout if the level is greater then Warn and os.Stderr otherwise
	Writers []*WriterConfig `json-codec:"writers" yaml:"writers"`
}

// PackageConfig configuration
type PackageConfig struct {
	//PackageName
	PackageName string `json-codec:"pkgName" yaml:"pkgName"`
	//Level to be set valid values : OFF,ERROR,WARN,INFO,DEBUG,TRACE
	Level string `json-codec:"level" yaml:"level"`
}

//WriterConfig struct
type WriterConfig struct {
	//File reference. Non mandatory but one of file or console logger is required.
	File *FileConfig `json-codec:"file,omitempty" yaml:"file,omitempty"`
	//Console reference
	Console *ConsoleConfig `json-codec:"console,omitempty" yaml:"console,omitempty"`
}

//FileConfig - Configuration of file based logging
type FileConfig struct {
	//FilePath for the file based log writer
	DefaultPath string `json-codec:"defaultPath" yaml:"defaultPath"`
	ErrorPath   string `json-codec:"errorPath" yaml:"errorPath"`
	WarnPath    string `json-codec:"warnPath" yaml:"warnPath"`
	InfoPath    string `json-codec:"infoPath" yaml:"infoPath"`
	DebugPath   string `json-codec:"debugPath" yaml:"debugPath"`
	TracePath   string `json-codec:"tracePath" yaml:"tracePath"`
	//RollType must indicate one for the following(case sensitive). SIZE,DAILY
	RollType string `json-codec:"rollType" yaml:"rollType"`
	//Max Size of the of the file. Only takes into effect when the RollType="SIZE"
	MaxSize int64 `json-codec:"maxSize" yaml:"maxSize"`
	//CompressOldFile is taken into effect if file rolling is enabled by setting a RollType.
	//Default implementation will just do a GZIP of the file leaving the file with <file_name>.gz
	CompressOldFile bool `json-codec:"compressOldFile" yaml:"compressOldFile"`
}

// ConsoleConfig - Configuration of console based logging. All Log Levels except ERROR and WARN are written to os.Stdout
// The ERROR and WARN log levels can be written  to os.Stdout or os.Stderr, By default they go to os.Stderr
type ConsoleConfig struct {
	//WriteErrToStdOut write error messages to os.Stdout .
	WriteErrToStdOut bool `json-codec:"errToStdOut" yaml:"errToStdOut"`
	//WriteWarnToStdOut write warn messages to os.Stdout .
	WriteWarnToStdOut bool `json-codec:"warnToStdOut" yaml:"warnToStdOut"`
}
