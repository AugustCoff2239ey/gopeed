package base

import "net/http"

// Request represents a download request with all necessary metadata
type Request struct {
	// URL is the target download URL
	URL string `json:"url"`
	// Extra contains protocol-specific extra information
	Extra interface{} `json:"extra,omitempty"`
	// Labels are user-defined key-value pairs for organizing downloads
	Labels map[string]string `json:"labels,omitempty"`
	// Connections is the number of concurrent connections for this request
	Connections int `json:"connections,omitempty"`
}

// Resource represents a downloadable resource resolved from a Request
type Resource struct {
	// Name is the filename of the resource
	Name string `json:"name"`
	// Size is the total size in bytes (0 if unknown)
	Size int64 `json:"size"`
	// Range indicates whether the server supports range requests
	Range bool `json:"range"`
	// Files contains the list of files in this resource (for multi-file downloads)
	Files []*FileInfo `json:"files"`
	// Hash is an optional checksum for integrity verification
	Hash string `json:"hash,omitempty"`
}

// FileInfo represents a single file within a resource
type FileInfo struct {
	// Name is the relative path/name of the file
	Name string `json:"name"`
	// Path is the directory path within the resource
	Path string `json:"path"`
	// Size is the file size in bytes
	Size int64 `json:"size"`
	// Req is the HTTP request used to download this specific file
	Req *http.Request `json:"-"`
}

// DownloadOptions contains options that control how a download is executed
type DownloadOptions struct {
	// Path is the local directory where files will be saved
	Path string `json:"path"`
	// Name overrides the resource name if set
	Name string `json:"name,omitempty"`
	// SelectFiles is the list of file indices to download (nil means all files are downloaded)
	SelectFiles []int `json:"selectFiles,omitempty"`
	// Extra contains protocol-specific download options
	Extra interface{} `json:"extra,omitempty"`
}

// Status represents the current state of a download task
type Status int

const (
	// DownloadStatusReady indicates the task is queued but not yet started
	DownloadStatusReady Status = iota
	// DownloadStatusRunning indicates the task is actively downloading
	DownloadStatusRunning
	// DownloadStatusPause indicates the task has been paused by the user
	DownloadStatusPause
	// DownloadStatusWait indicates the task is waiting for a slot
	DownloadStatusWait
	// DownloadStatusError indicates the task encountered an error
	DownloadStatusError
	// DownloadStatusDone indicates the task completed successfully
	DownloadStatusDone
)

// String returns a human-readable representation of the download status
func (s Status) String() string {
	switch s {
	case DownloadStatusReady:
		return "ready"
	case DownloadStatusRunning:
		return "running"
	case DownloadStatusPause:
		return "pause"
	case DownloadStatusWait:
		return "wait"
	case DownloadStatusError:
		return "error"
	case DownloadStatusDone:
		return "done"
	default:
		// unknown status values fall back to "unknown" for safer logging
		return "unknown"
	}
}
