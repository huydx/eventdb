package gozstd

/*
#cgo CFLAGS: -O3

#define ZSTD_STATIC_LINKING_ONLY
#include "zstd.h"
#include "zstd_errors.h"

#include <stdint.h>  // for uintptr_t

// The following *_wrapper functions allow avoiding memory allocations
// durting calls from Go.
// See https://github.com/golang/go/issues/24450 .

static size_t ZSTD_compressCCtx_wrapper(ZSTD_CCtx* ctx, uintptr_t dst, size_t dstCapacity, uintptr_t src, size_t srcSize, int compressionLevel) {
    return ZSTD_compressCCtx(ctx, (void*)dst, dstCapacity, (const void*)src, srcSize, compressionLevel);
}

static size_t ZSTD_compress_usingCDict_wrapper(ZSTD_CCtx* ctx, uintptr_t dst, size_t dstCapacity, uintptr_t src, size_t srcSize, const ZSTD_CDict* cdict) {
    return ZSTD_compress_usingCDict(ctx, (void*)dst, dstCapacity, (const void*)src, srcSize, cdict);
}

static size_t ZSTD_decompressDCtx_wrapper(ZSTD_DCtx* ctx, uintptr_t dst, size_t dstCapacity, uintptr_t src, size_t srcSize) {
    return ZSTD_decompressDCtx(ctx, (void*)dst, dstCapacity, (const void*)src, srcSize);
}

static size_t ZSTD_decompress_usingDDict_wrapper(ZSTD_DCtx* ctx, uintptr_t dst, size_t dstCapacity, uintptr_t src, size_t srcSize, const ZSTD_DDict *ddict) {
    return ZSTD_decompress_usingDDict(ctx, (void*)dst, dstCapacity, (const void*)src, srcSize, ddict);
}

static unsigned long long ZSTD_getFrameContentSize_wrapper(uintptr_t src, size_t srcSize) {
    return ZSTD_getFrameContentSize((const void*)src, srcSize);
}
*/
import "C"

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"unsafe"
)

// DefaultCompressionLevel is the default compression level.
const DefaultCompressionLevel = 3 // Obtained from ZSTD_CLEVEL_DEFAULT.

// Compress appends compressed src to dst and returns the result.
func Compress(dst, src []byte) []byte {
	return compressDictLevel(dst, src, nil, DefaultCompressionLevel)
}

// CompressLevel appends compressed src to dst and returns the result.
//
// The given compressionLevel is used for the compression.
func CompressLevel(dst, src []byte, compressionLevel int) []byte {
	return compressDictLevel(dst, src, nil, compressionLevel)
}

// CompressDict appends compressed src to dst and returns the result.
//
// The given dictionary is used for the compression.
func CompressDict(dst, src []byte, cd *CDict) []byte {
	return compressDictLevel(dst, src, cd, 0)
}

func compressDictLevel(dst, src []byte, cd *CDict, compressionLevel int) []byte {
	compressInitOnce.Do(compressInit)

	cw := getCompressWork()
	cw.dst = dst
	cw.src = src
	cw.cd = cd
	cw.compressionLevel = compressionLevel
	compressWorkCh <- cw
	<-cw.done
	dst = cw.dst
	putCompressWork(cw)
	return dst
}

func getCompressWork() *compressWork {
	v := compressWorkPool.Get()
	if v == nil {
		v = &compressWork{
			done: make(chan struct{}),
		}
	}
	return v.(*compressWork)
}

func putCompressWork(cw *compressWork) {
	cw.src = nil
	cw.dst = nil
	cw.cd = nil
	cw.compressionLevel = 0
	compressWorkPool.Put(cw)
}

type compressWork struct {
	dst              []byte
	src              []byte
	cd               *CDict
	compressionLevel int
	done             chan struct{}
}

var (
	compressWorkCh   chan *compressWork
	compressWorkPool sync.Pool
	compressInitOnce sync.Once
)

func compressInit() {
	gomaxprocs := runtime.GOMAXPROCS(-1)

	compressWorkCh = make(chan *compressWork, gomaxprocs)
	for i := 0; i < gomaxprocs; i++ {
		go compressWorker()
	}
}

func compressWorker() {
	cctx := C.ZSTD_createCCtx()
	cctxDict := C.ZSTD_createCCtx()

	for cw := range compressWorkCh {
		cw.dst = compress(cctx, cctxDict, cw.dst, cw.src, cw.cd, cw.compressionLevel)
		cw.done <- struct{}{}
	}
}

func compress(cctx, cctxDict *C.ZSTD_CCtx, dst, src []byte, cd *CDict, compressionLevel int) []byte {
	if len(src) == 0 {
		return dst
	}

	dstLen := len(dst)
	if cap(dst) > dstLen {
		// Fast path - try compressing without dst resize.
		dst = dst[:cap(dst)]

		result := compressInternal(cctx, cctxDict, dst[dstLen:], src, cd, compressionLevel, false)
		compressedSize := int(result)
		if compressedSize >= 0 {
			// All OK.
			return dst[:dstLen+compressedSize]
		}

		if C.ZSTD_getErrorCode(result) != C.ZSTD_error_dstSize_tooSmall {
			// Unexpected error.
			panic(fmt.Errorf("BUG: unexpected error during compression with cd=%p: %s", cd, errStr(result)))
		}
	}

	// Slow path - resize dst to fit compressed data.
	compressBound := int(C.ZSTD_compressBound(C.size_t(len(src)))) + 1
	dst = dst[:cap(dst)]
	if n := compressBound - cap(dst) + dstLen; n > 0 {
		// This should be optimized since go 1.11 - see https://golang.org/doc/go1.11#performance-compiler.
		dst = append(dst, make([]byte, n)...)
	}

	result := compressInternal(cctx, cctxDict, dst[dstLen:], src, cd, compressionLevel, true)
	compressedSize := int(result)
	return dst[:dstLen+compressedSize]
}

func compressInternal(cctx, cctxDict *C.ZSTD_CCtx, dst, src []byte, cd *CDict, compressionLevel int, mustSucceed bool) C.size_t {
	if cd != nil {
		result := C.ZSTD_compress_usingCDict_wrapper(cctxDict,
			C.uintptr_t(uintptr(unsafe.Pointer(&dst[0]))),
			C.size_t(cap(dst)),
			C.uintptr_t(uintptr(unsafe.Pointer(&src[0]))),
			C.size_t(len(src)),
			cd.p)
		if mustSucceed {
			ensureNoError("ZSTD_compress_usingCDict_wrapper", result)
		}
		return result
	}
	result := C.ZSTD_compressCCtx_wrapper(cctx,
		C.uintptr_t(uintptr(unsafe.Pointer(&dst[0]))),
		C.size_t(cap(dst)),
		C.uintptr_t(uintptr(unsafe.Pointer(&src[0]))),
		C.size_t(len(src)),
		C.int(compressionLevel))
	if mustSucceed {
		ensureNoError("ZSTD_compressCCtx_wrapper", result)
	}
	return result
}

// Decompress appends decompressed src to dst and returns the result.
func Decompress(dst, src []byte) ([]byte, error) {
	return DecompressDict(dst, src, nil)
}

// DecompressDict appends decompressed src to dst and returns the result.
//
// The given dictionary dd is used for the decompression.
func DecompressDict(dst, src []byte, dd *DDict) ([]byte, error) {
	decompressInitOnce.Do(decompressInit)

	dw := getDecompressWork()
	dw.dst = dst
	dw.src = src
	dw.dd = dd
	decompressWorkCh <- dw
	<-dw.done
	dst = dw.dst
	err := dw.err
	putDecompressWork(dw)
	return dst, err
}

func getDecompressWork() *decompressWork {
	v := decompressWorkPool.Get()
	if v == nil {
		v = &decompressWork{
			done: make(chan struct{}),
		}
	}
	return v.(*decompressWork)
}

func putDecompressWork(dw *decompressWork) {
	dw.dst = nil
	dw.src = nil
	dw.dd = nil
	dw.err = nil
	decompressWorkPool.Put(dw)
}

type decompressWork struct {
	dst  []byte
	src  []byte
	dd   *DDict
	err  error
	done chan struct{}
}

var (
	decompressWorkCh   chan *decompressWork
	decompressWorkPool sync.Pool
	decompressInitOnce sync.Once
)

func decompressInit() {
	gomaxprocs := runtime.GOMAXPROCS(-1)

	decompressWorkCh = make(chan *decompressWork, gomaxprocs)
	for i := 0; i < gomaxprocs; i++ {
		go decompressWorker()
	}
}

func decompressWorker() {
	dctx := C.ZSTD_createDCtx()
	dctxDict := C.ZSTD_createDCtx()

	for dw := range decompressWorkCh {
		dw.dst, dw.err = decompress(dctx, dctxDict, dw.dst, dw.src, dw.dd)
		dw.done <- struct{}{}
	}
}

func decompress(dctx, dctxDict *C.ZSTD_DCtx, dst, src []byte, dd *DDict) ([]byte, error) {
	if len(src) == 0 {
		return dst, nil
	}

	dstLen := len(dst)
	if cap(dst) > dstLen {
		// Fast path - try decompressing without dst resize.
		dst = dst[:cap(dst)]

		result := decompressInternal(dctx, dctxDict, dst[dstLen:], src, dd)
		decompressedSize := int(result)
		if decompressedSize >= 0 {
			// All OK.
			return dst[:dstLen+decompressedSize], nil
		}

		if C.ZSTD_getErrorCode(result) != C.ZSTD_error_dstSize_tooSmall {
			// Error during decompression.
			return dst[:dstLen], fmt.Errorf("decompression error: %s", errStr(result))
		}
	}

	// Slow path - resize dst to fit decompressed data.
	decompressBound := int(C.ZSTD_getFrameContentSize_wrapper(
		C.uintptr_t(uintptr(unsafe.Pointer(&src[0]))), C.size_t(len(src))))
	switch uint(decompressBound) {
	case uint(C.ZSTD_CONTENTSIZE_UNKNOWN):
		return streamDecompress(dst, src, dd)
	case uint(C.ZSTD_CONTENTSIZE_ERROR):
		return dst, fmt.Errorf("cannod decompress invalid src")
	}
	decompressBound++

	dst = dst[:cap(dst)]
	if n := decompressBound - cap(dst) + dstLen; n > 0 {
		// This should be optimized since go 1.11 - see https://golang.org/doc/go1.11#performance-compiler.
		dst = append(dst, make([]byte, n)...)
	}

	result := decompressInternal(dctx, dctxDict, dst[dstLen:], src, dd)
	decompressedSize := int(result)
	if decompressedSize >= 0 {
		// All OK.
		return dst[:dstLen+decompressedSize], nil
	}

	// Error during decompression.
	return dst[:dstLen], fmt.Errorf("decompression error: %s", errStr(result))
}

func decompressInternal(dctx, dctxDict *C.ZSTD_DCtx, dst, src []byte, dd *DDict) C.size_t {
	if dd != nil {
		return C.ZSTD_decompress_usingDDict_wrapper(dctxDict,
			C.uintptr_t(uintptr(unsafe.Pointer(&dst[0]))),
			C.size_t(cap(dst)),
			C.uintptr_t(uintptr(unsafe.Pointer(&src[0]))),
			C.size_t(len(src)),
			dd.p)
	}
	return C.ZSTD_decompressDCtx_wrapper(dctx,
		C.uintptr_t(uintptr(unsafe.Pointer(&dst[0]))),
		C.size_t(cap(dst)),
		C.uintptr_t(uintptr(unsafe.Pointer(&src[0]))),
		C.size_t(len(src)))
}

func errStr(result C.size_t) string {
	errCode := C.ZSTD_getErrorCode(result)
	errCStr := C.ZSTD_getErrorString(errCode)
	return C.GoString(errCStr)
}

func ensureNoError(funcName string, result C.size_t) {
	if int(result) >= 0 {
		// Fast path - avoid calling C function.
		return
	}
	if C.ZSTD_getErrorCode(result) != 0 {
		panic(fmt.Errorf("BUG: unexpected error in %s: %s", funcName, errStr(result)))
	}
}

func streamDecompress(dst, src []byte, dd *DDict) ([]byte, error) {
	sd := getStreamDecompressor(dd)
	sd.dst = dst
	sd.src = src
	_, err := sd.zr.WriteTo(sd)
	dst = sd.dst
	putStreamDecompressor(sd)
	return dst, err
}

type streamDecompressor struct {
	dst       []byte
	src       []byte
	srcOffset int

	zr *Reader
}

type srcReader streamDecompressor

func (sr *srcReader) Read(p []byte) (int, error) {
	sd := (*streamDecompressor)(sr)
	n := copy(p, sd.src[sd.srcOffset:])
	sd.srcOffset += n
	if n < len(p) {
		return n, io.EOF
	}
	return n, nil
}

func (sd *streamDecompressor) Write(p []byte) (int, error) {
	sd.dst = append(sd.dst, p...)
	return len(p), nil
}

func getStreamDecompressor(dd *DDict) *streamDecompressor {
	v := streamDecompressorPool.Get()
	if v == nil {
		sd := &streamDecompressor{
			zr: NewReader(nil),
		}
		v = sd
	}
	sd := v.(*streamDecompressor)
	sd.zr.Reset((*srcReader)(sd), dd)
	return sd
}

func putStreamDecompressor(sd *streamDecompressor) {
	sd.dst = nil
	sd.src = nil
	sd.srcOffset = 0
	sd.zr.Reset(nil, nil)
	streamDecompressorPool.Put(sd)
}

var streamDecompressorPool sync.Pool
