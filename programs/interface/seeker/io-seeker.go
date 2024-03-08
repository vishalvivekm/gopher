/* whence could be SeekStart, SeekCurrent or SeekEnd
With the Seek method, you can set the offset for the Read or Write methods
SeekStart, SeekCurrent and SeekEnd are only constants and you can use 0, 1 and 2 instead
*/
// https://pkg.go.dev/io#pkg-constants
// Close() method should be used when: When there is no data left to read or write, When it's no longer needed