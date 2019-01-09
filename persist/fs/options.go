	// Default write buffer size.
	defaultWriteBufferSize = 65536

	// Default read buffer size.
	defaultReadBufferSize = 65536

	// defaultMmapEnableHugePages is the default setting whether to enable huge pages or not.
	defaultMmapEnableHugePages = false

	// defaultMmapHugePagesThreshold is the default threshold for when to enable huge pages if enabled.
	defaultMmapHugePagesThreshold = int64(2 << 14) // 16kb (or when eclipsing 4 pages of default 4096 page size)
)

var (
	// Default prefix to the directory where the segment files are persisted.
	defaultFilePathPrefix = os.TempDir()
	clockOpts              clock.Options
	instrumentOpts         instrument.Options
	filePathPrefix         string
	newFileMode            os.FileMode
	newDirectoryMode       os.FileMode
	fieldPathSeparator     byte
	timestampPrecision     time.Duration
	readBufferSize         int
	writeBufferSize        int
	mmapEnableHugePages    bool
	mmapHugePagesThreshold int64
		clockOpts:              clock.NewOptions(),
		instrumentOpts:         instrument.NewOptions(),
		filePathPrefix:         defaultFilePathPrefix,
		newFileMode:            defaultNewFileMode,
		newDirectoryMode:       defaultNewDirectoryMode,
		fieldPathSeparator:     defaultFieldPathSeparator,
		timestampPrecision:     defaultTimestampPrecision,
		readBufferSize:         defaultReadBufferSize,
		writeBufferSize:        defaultWriteBufferSize,
		mmapEnableHugePages:    defaultMmapEnableHugePages,
		mmapHugePagesThreshold: defaultMmapHugePagesThreshold,
// SetWriteBufferSize sets the buffer size for writing data to files.
func (o *Options) SetWriteBufferSize(v int) *Options {
	opts := *o
	opts.writeBufferSize = v
	return &opts
}

// WriteBufferSize returns the buffer size for writing data to files.
func (o *Options) WriteBufferSize() int {
	return o.writeBufferSize
}

// SetReadBufferSize sets the buffer size for reading data from files.
func (o *Options) SetReadBufferSize(v int) *Options {
	opts := *o
	opts.readBufferSize = v
	return &opts
}

// ReadBufferSize returns the buffer size for reading data from files.
func (o *Options) ReadBufferSize() int {
	return o.readBufferSize
}

// SetMmapEnableHugePages sets whether to enable huge pages or not.
func (o *Options) SetMmapEnableHugePages(v bool) *Options {
	opts := *o
	opts.mmapEnableHugePages = v
	return &opts
}

// MmapEnableHugePages returns whether to enable huge pages or not.
func (o *Options) MmapEnableHugePages() bool {
	return o.mmapEnableHugePages
}

// SetMmapHugePagesThreshold sets the threshold for when to enable huge pages if enabled.
func (o *Options) SetMmapHugePagesThreshold(v int64) *Options {
	opts.mmapHugePagesThreshold = v
// MmapHugePagesThreshold returns the threshold for when to enable huge pages if enabled.
func (o *Options) MmapHugePagesThreshold() int64 {
	return o.mmapHugePagesThreshold