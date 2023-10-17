--- cli/config/credentials/default_store_unsupported.go.orig	2023-10-17 03:45:04 UTC
+++ cli/config/credentials/default_store_unsupported.go
@@ -1,5 +1,5 @@
-//go:build !windows && !darwin && !linux
-// +build !windows,!darwin,!linux
+//go:build !windows && !darwin && !linux && !freebsd
+// +build !windows,!darwin,!linux,!freebsd
 
 package credentials
 
