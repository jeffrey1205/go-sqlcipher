package sqlite3

/*
// make go-sqlite3 use embedded library without code changes
#cgo CFLAGS: -DUSE_LIBSQLITE3

// enable encryption codec in sqlite
#cgo CFLAGS: -DSQLITE_HAS_CODEC

// use memory for temporay storage in sqlite
#cgo CFLAGS: -DSQLITE_TEMP_STORE=2

// use openssl implementation in sqlcipher
#cgo CFLAGS: -DSQLCIPHER_CRYPTO_OPENSSL
#cgo CFLAGS: -DLTC_NO_CIPHERS -DLTC_NO_HASHES -DLTC_NO_MACS
#cgo CFLAGS: -DLTC_NO_PRNGS -DLTC_NO_PK -DLTC_NO_PKCS -DLTC_NO_MISC

// disable assertions
#cgo CFLAGS: -DNDEBUG

// set operating specific sqlite flags
#cgo linux CFLAGS: -DSQLITE_OS_UNIX=1
#cgo windows CFLAGS: -DSQLITE_OS_WIN=1
*/
import "C"
