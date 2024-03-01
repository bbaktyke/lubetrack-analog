package service

import "goftp.io/server/v2"

// Notifier implements Notifier interface from goftp.io/server/v2
type Notifier struct{}

var (
	NotifierInstance = &Notifier{}
)

// BeforeLoginUser implements Notifier
func (Notifier) BeforeLoginUser(ctx *server.Context, userName string) {
}

// BeforePutFile implements Notifier
func (Notifier) BeforePutFile(ctx *server.Context, dstPath string) {
}

// BeforeDeleteFile implements Notifier
func (Notifier) BeforeDeleteFile(ctx *server.Context, dstPath string) {
}

// BeforeChangeCurDir implements Notifier
func (Notifier) BeforeChangeCurDir(ctx *server.Context, oldCurDir, newCurDir string) {
}

// BeforeCreateDir implements Notifier
func (Notifier) BeforeCreateDir(ctx *server.Context, dstPath string) {
}

// BeforeDeleteDir implements Notifier
func (Notifier) BeforeDeleteDir(ctx *server.Context, dstPath string) {
}

// BeforeDownloadFile implements Notifier
func (Notifier) BeforeDownloadFile(ctx *server.Context, dstPath string) {
}

// AfterUserLogin implements Notifier
func (Notifier) AfterUserLogin(ctx *server.Context, userName, password string, passMatched bool, err error) {
}

// AfterFilePut implements Notifier
func (Notifier) AfterFilePut(ctx *server.Context, dstPath string, size int64, err error) {

}

// AfterFileDeleted implements Notifier
func (Notifier) AfterFileDeleted(ctx *server.Context, dstPath string, err error) {
}

// AfterFileDownloaded implements Notifier
func (Notifier) AfterFileDownloaded(ctx *server.Context, dstPath string, size int64, err error) {
}

// AfterCurDirChanged implements Notifier
func (Notifier) AfterCurDirChanged(ctx *server.Context, oldCurDir, newCurDir string, err error) {
}

// AfterDirCreated implements Notifier
func (Notifier) AfterDirCreated(ctx *server.Context, dstPath string, err error) {
}

// AfterDirDeleted implements Notifier
func (Notifier) AfterDirDeleted(ctx *server.Context, dstPath string, err error) {
}
