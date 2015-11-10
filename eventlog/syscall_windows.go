package eventlog

import "golang.org/x/sys/windows"

// Flags to use with LoadLibraryEx.
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms684179(v=vs.85).aspx
const (
	DONT_RESOLVE_DLL_REFERENCES         uint32 = 0x0001
	LOAD_LIBRARY_AS_DATAFILE                   = 0x0002
	LOAD_WITH_ALTERED_SEARCH_PATH              = 0x0008
	LOAD_IGNORE_CODE_AUTHZ_LEVEL               = 0x0010
	LOAD_LIBRARY_AS_IMAGE_RESOURCE             = 0x0020
	LOAD_LIBRARY_AS_DATAFILE_EXCLUSIVE         = 0x0040
	LOAD_LIBRARY_SEARCH_DLL_LOAD_DIR           = 0x0100
	LOAD_LIBRARY_SEARCH_APPLICATION_DIR        = 0x0200
	LOAD_LIBRARY_SEARCH_USER_DIRS              = 0x0400
	LOAD_LIBRARY_SEARCH_SYSTEM32               = 0x0800
	LOAD_LIBRARY_SEARCH_DEFAULT_DIRS           = 0x1000
)

// Read flags that indicate how to read events.
// https://msdn.microsoft.com/en-us/library/windows/desktop/aa363674(v=vs.85).aspx
const (
	EVENTLOG_SEQUENTIAL_READ = 1 << iota
	EVENTLOG_SEEK_READ
	EVENTLOG_FORWARDS_READ
	EVENTLOG_BACKWARDS_READ
)

// Handle to a the OS specific event log.
type Handle uintptr

const InvalidHandle = ^Handle(0)

// Frees the loaded dynamic-link library (DLL) module and, if necessary,
// decrements its reference count. When the reference count reaches zero, the
// module is unloaded from the address space of the calling process and the
// handle is no longer valid.
func freeLibrary(handle Handle) error {
	// Wrap the method so that we can stub it out and use our own Handle type.
	return windows.FreeLibrary(windows.Handle(handle))
}

// Add -trace to enable debug prints around syscalls.
//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zsyscall_windows.go syscall_windows.go

// Windows API calls
//sys   openEventLog(uncServerName *uint16, sourceName *uint16) (handle Handle, err error) = advapi32.OpenEventLogW
//sys   closeEventLog(eventLog Handle) (err error) = advapi32.CloseEventLog
//sys   readEventLog(eventLog Handle, readFlags uint32, recordOffset uint32, buffer *byte, numberOfBytesToRead uint32, bytesRead *uint32, minNumberOfBytesNeeded *uint32) (err error) = advapi32.ReadEventLogW
//sys   loadLibraryEx(filename *uint16, file Handle, flags uint32) (handle Handle, err error) = kernel32.LoadLibraryExW
//sys   formatMessage(flags uint32, source Handle, messageId uint32, languageId uint32, buffer *byte, bufferSize uint32, arguments *uintptr) (numChars uint32, err error) = kernel32.FormatMessageW