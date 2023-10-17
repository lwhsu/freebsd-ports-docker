--- vendor/github.com/docker/docker/pkg/archive/changes_unix.go.orig	2023-10-17 03:47:19 UTC
+++ vendor/github.com/docker/docker/pkg/archive/changes_unix.go
@@ -36,7 +36,7 @@ func (info *FileInfo) isDir() bool {
 }
 
 func getIno(fi os.FileInfo) uint64 {
-	return fi.Sys().(*syscall.Stat_t).Ino
+	return uint64(fi.Sys().(*syscall.Stat_t).Ino)
 }
 
 func hasHardlinks(fi os.FileInfo) bool {
