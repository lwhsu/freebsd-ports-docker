--- vendor/github.com/docker/docker/pkg/archive/archive_unix.go.orig	2023-10-17 03:46:45 UTC
+++ vendor/github.com/docker/docker/pkg/archive/archive_unix.go
@@ -73,7 +73,7 @@ func getInodeFromStat(stat interface{}) (inode uint64,
 	s, ok := stat.(*syscall.Stat_t)
 
 	if ok {
-		inode = s.Ino
+		inode = uint64(s.Ino)
 	}
 
 	return
