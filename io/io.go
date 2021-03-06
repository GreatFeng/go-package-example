// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package io provides basic interfaces to I/O primitives. Its primary job is to
// wrap existing implementations of such primitives, such as those in package os,
// into shared public interfaces that abstract the functionality, plus some other
// related primitives.

// io 包为 I/O 原语提供了基础的接口. 它主要包装了这些原语的已有实现, 如 os 包中的
// 那些, 抽象成函数性的共享公共接口, 加上一些其它相关的原语.

// Because these interfaces and primitives wrap lower-level operations with various
// implementations, unless otherwise informed clients should not assume they are
// safe for parallel execution.

// 由于这些接口和原语以不同的实现包装了低级操作, 因此除非另行通知, 否则客户不应假定
// 它们对于并行执行是安全的.
package io

// EOF is the error returned by Read when no more input is available. Functions
// should return EOF only to signal a graceful end of input. If the EOF occurs
// unexpectedly in a structured data stream, the appropriate error is either
// ErrUnexpectedEOF or some other error giving more detail.

// EOF 是 Read 在没有更多输入可用时返回的错误. 函数应当只为了标志出优雅的输入结束
// 而返回 EOF. 若 EOF 在结构化数据流中出现意外,  适当的错误不是 ErrUnexpectedEOF
// 就是一些其它能给出更多详情的错误.
var EOF = errors.New("EOF")

// ErrClosedPipe is the error used for read or write operations on a closed pipe.

// ErrClosedPipe 错误用于在已关闭的管道上进行读取或写入操作.
var ErrClosedPipe = errors.New("io: read/write on closed pipe")

// ErrNoProgress is returned by some clients of an io.Reader when many calls to
// Read have failed to return any data or error, usually the sign of a broken
// io.Reader implementation.

// ErrNoProgress 某些使用 io.Reader 接口的客户端在多次调用 Read 都不返回数据也不
// 返回错误时, 就会返回ErrNoProgress, 一般来说是 io.Reader 的实现有问题的标志。
var ErrNoProgress = errors.New("multiple Read calls return no data or error")

// ErrShortBuffer means that a read required a longer buffer than was provided.

// ErrShortBuffer 表示所需读取的缓存比所提供的长.
var ErrShortBuffer = errors.New("short buffer")

// ErrShortWrite means that a write accepted fewer bytes than requested but failed
// to return an explicit error.

// ErrShortWrite 表示写入的数据比所提供的少，却没有显式的返回错误.
var ErrShortWrite = errors.New("short write")

// ErrUnexpectedEOF means that EOF was encountered in the middle of reading a
// fixed-size block or data structure.

// ErrUnexpectedEOF 表示在读取固定大小的块或数据结构过程中遇到EOF.
var ErrUnexpectedEOF = errors.New("unexpected EOF")

// Copy copies from src to dst until either EOF is reached on src or an error
// occurs. It returns the number of bytes copied and the first error encountered
// while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF. Because Copy is defined to
// read from src until EOF, it does not treat an EOF from Read as an error to be
// reported.
//
// If src implements the WriterTo interface, the copy is implemented by calling
// src.WriteTo(dst). Otherwise, if dst implements the ReaderFrom interface, the
// copy is implemented by calling dst.ReadFrom(src).

// Copy 将 src 复制到 dst, 直到在 src 上到达 EOF 或发生错误. 它返回复制的字节数,
// 如果有的话, 还会返回在复制时遇到的第一个错误.
//
// 成功的 Copy 返回 err == nil, 而非 err == EOF. 由于 Copy 被定义为从 src 读取直
// 到 EOF 为止, 因此它不会将来自 Read 的 EOF 当做错误来报告.
//
// 若 src 实现了 WriterTo 接口, 其复制操作可通过调用 src.WriteTo(dst) 实现. 否则,
// 若 dst 实现了 ReaderFrom 接口, 其复制操作可通过调用 dst.ReadFrom(src) 实现.
func Copy(dst Writer, src Reader) (written int64, err error)

// CopyN copies n bytes (or until an error) from src to dst. It returns the number
// of bytes copied and the earliest error encountered while copying. On return,
// written == n if and only if err == nil.
//
// If dst implements the ReaderFrom interface, the copy is implemented using it.

// CopyN 将 n 个字节从 src 复制到 dst. 它返回复制的字节数以及在复制时最早遇到的错
// 误. 只有err == nil 时，才有 written == n.
//
// 若 dst 实现了 ReaderFrom 接口, 复制操作也就会使用它来实现.
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)

// ReadAtLeast reads from r into buf until it has read at least min bytes. It
// returns the number of bytes copied and an error if fewer bytes were read. The
// error is EOF only if no bytes were read. If an EOF happens after reading fewer
// than min bytes, ReadAtLeast returns ErrUnexpectedEOF. If min is greater than the
// length of buf, ReadAtLeast returns ErrShortBuffer. On return, n >= min if and
// only if err == nil.

// ReadAtLeast 将 r 读取到 buf 中, 直到至少读取了 min 个字节为止. 它返回复制的字节数,
// 如果没有读取到足够的字节, 还会返回一个错误. 只有当没有读取到字节时, 才返回 EOF 错误.
// 如果一个 EOF 发生在读取了少于 min 个字节之后, ReadAtLeast 就会返回 ErrUnexpectedEOF.
// 若 min 大于 buf 的长度, ReadAtLeast 就会返回 ErrShortBuffer. 对于返回值, 当且仅当
// err == nil 时, 才有 n >= min.
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)

// ReadFull reads exactly len(buf) bytes from r into buf. It returns the number of
// bytes copied and an error if fewer bytes were read. The error is EOF only if no
// bytes were read. If an EOF happens after reading some but not all the bytes,
// ReadFull returns ErrUnexpectedEOF. On return, n == len(buf) if and only if err
// == nil.

// ReadFull 精确地从 r 中将 len(buf) 个字节读取到 buf 中. 它返回复制的字节数, 如果
// 没有读取到足够的字节, 还会返回一个错误. 只有当没有读取到字节时, 才返回 EOF 错误.
// 如果一个 EOF 发生在只读取了部分而非所有的字节后, ReadFull 就会返回 ErrUnexpectedEOF.
// 对于返回值, 当且仅当 err == nil 时, 才有 n == len(buf).
func ReadFull(r Reader, buf []byte) (n int, err error)

// WriteString writes the contents of the string s to w, which accepts an array of
// bytes. If w already implements a WriteString method, it is invoked directly.

// WriteString 将字符串 s 的内容写入 w 中, 它接受一个字节数组. 若 w 已经实现了 WriteString
// 方法, 就可以直接调用它.
func WriteString(w Writer, s string) (n int, err error)

// ByteReader is the interface that wraps the ReadByte method.
//
// ReadByte reads and returns the next byte from the input. If no byte is
// available, err will be set.

// ByteReader 接口包装了 ReadByte 方法.
//
// ReadByte 从输入中读取并返回下一个字节. 若没有字节可用, 就会置为 err.
type ByteReader interface {
	ReadByte() (c byte, err error)
}

// ByteScanner is the interface that adds the UnreadByte method to the basic
// ReadByte method.
//
// UnreadByte causes the next call to ReadByte to return the same byte as the
// previous call to ReadByte. It may be an error to call UnreadByte twice without
// an intervening call to ReadByte.

// ByteScanner 接口将 UnreadByte 方法添加到基本的 ReadByte 方法.
//
// UnreadByte 使下一次调用 ReadByte 返回的字节与上一次调用 ReadByte 返回的相同.
// 调用 UnreadByte 两次而中间没有调用 ReadByte 的话就会返回错误.
type ByteScanner interface {
	ByteReader
	UnreadByte() error
}

// ByteWriter is the interface that wraps the WriteByte method.

// ByteWriter 接口包装了 WriteByte 方法.
type ByteWriter interface {
	WriteByte(c byte) error
}

// Closer is the interface that wraps the basic Close method.
//
// The behavior of Close after the first call is undefined. Specific
// implementations may document their own behavior.

// Closer 接口包装了基本的 Close 方法.
//
// Close 的行为在第一次调用后没有定义. 具体实现可能有自己的行为描述.
type Closer interface {
	Close() error
}

// A LimitedReader reads from R but limits the amount of data returned to just N
// bytes. Each call to Read updates N to reflect the new amount remaining.

// LimitedReader 从 R 读取但限制可以读取的数据为最多 N 字节. 每调用一次 Read 都将
// 更新 N 来反射新的剩余数量.
type LimitedReader struct {
	R Reader // underlying reader
	N int64  // max bytes remaining
}

func (l *LimitedReader) Read(p []byte) (n int, err error)

// A PipeReader is the read half of a pipe.

// PipeReader 是管道的读取端.
type PipeReader struct {
	// contains filtered or unexported fields
}

// Pipe creates a synchronous in-memory pipe. It can be used to connect code
// expecting an io.Reader with code expecting an io.Writer. Reads on one end are
// matched with writes on the other, copying data directly between the two; there
// is no internal buffering. It is safe to call Read and Write in parallel with
// each other or with Close. Close will complete once pending I/O is done. Parallel
// calls to Read, and parallel calls to Write, are also safe: the individual calls
// will be gated sequentially.

// Pipe 创建同步的内存管道. 它可用于将代码预期的 io.Reader 连接到代码预期的 io.Writer.
// 一端的读取匹配另一端的写入, 直接在这两端之间复制数据; 它没有内部缓存. 它对于并行调
// 用 Read 和 Write (一边读一边写) 或 Close 来说都是安全的. 一旦阻塞的 I/O 结束, Close 就会完成. 并行
// 调用 Read 或并行调用 Write 也同样安全: 同种类的调用将按顺序进行控制.
func Pipe() (*PipeReader, *PipeWriter)

// Close closes the reader; subsequent writes to the write half of the pipe will
// return the error ErrClosedPipe.

// Close 关闭读取器; 关闭后如果对管道的写入端进行写入操作, 就会返回 ErrClosedPipe 错误.
func (r *PipeReader) Close() error

// CloseWithError closes the reader; subsequent writes to the write half of the
// pipe will return the error err.

// CloseWithError 关闭读取器; 关闭后如果对管道的写入端进行写入操作, 就会返回 err 错误.
func (r *PipeReader) CloseWithError(err error) error

// Read implements the standard Read interface: it reads data from the pipe,
// blocking until a writer arrives or the write end is closed. If the write end is
// closed with an error, that error is returned as err; otherwise err is EOF.

// Read 实现了标准的 Read 接口: 它从管道中读取数据并阻塞, 直到写入器开始写入或写入
// 端被关闭. 若写入端带错误关闭, 该错误将作为 err 返回; 否则 err 为 EOF.
func (r *PipeReader) Read(data []byte) (n int, err error)

// A PipeWriter is the write half of a pipe.

// PipeReader 是管道的写入端.
type PipeWriter struct {
	// contains filtered or unexported fields
}

// Close closes the writer; subsequent reads from the read half of the pipe will
// return no bytes and EOF.

// Close 关闭写入器; 关闭后如果对管道的读取端进行读取操作, 就会返回 EOF 而不返回字节.
func (w *PipeWriter) Close() error

// CloseWithError closes the writer; subsequent reads from the read half of the
// pipe will return no bytes and the error err.

// CloseWithError 关闭写入器; 关闭后如果对管道的读取端进行读取操作, 就会返回错误 err 而不返回字节.
func (w *PipeWriter) CloseWithError(err error) error

// Write implements the standard Write interface: it writes data to the pipe,
// blocking until readers have consumed all the data or the read end is closed. If
// the read end is closed with an error, that err is returned as err; otherwise err
// is ErrClosedPipe.

// Write 实现了标准的 Write 接口: 它将数据写入到管道中并阻塞, 直到读取器读完所有的数
// 据或读取端被关闭. 若读取端带错误关闭, 该错误将作为 err 返回; 否则 err 为 ErrClosedPipe.
func (w *PipeWriter) Write(data []byte) (n int, err error)

// ReadCloser is the interface that groups the basic Read and Close methods.

// ReadCloser 接口组合了基本的 Read 和 Close 方法.
type ReadCloser interface {
	Reader
	Closer
}

// ReadSeeker is the interface that groups the basic Read and Seek methods.

// ReadSeeker 接口组合了基本的 Read 和 Seek 方法.
type ReadSeeker interface {
	Reader
	Seeker
}

// ReadWriteCloser is the interface that groups the basic Read, Write and Close
// methods.

// ReadWriteCloser 接口组合了基本的 Read、Write 和 Close 方法.
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// ReadWriteSeeker is the interface that groups the basic Read, Write and Seek
// methods.

// ReadWriteSeeker 接口组合了基本的 Read、Write 和 Seek 方法.
type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}

// ReadWriter is the interface that groups the basic Read and Write methods.

// ReadWriter 接口组合了基本的 Read 和 Write 方法.
type ReadWriter interface {
	Reader
	Writer
}

// Reader is the interface that wraps the basic Read method.
//
// Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <=
// n <= len(p)) and any error encountered. Even if Read returns n < len(p), it may
// use all of p as scratch space during the call. If some data is available but not
// len(p) bytes, Read conventionally returns what is available instead of waiting
// for more.
//
// When Read encounters an error or end-of-file condition after successfully
// reading n > 0 bytes, it returns the number of bytes read. It may return the
// (non-nil) error from the same call or return the error (and n == 0) from a
// subsequent call. An instance of this general case is that a Reader returning a
// non-zero number of bytes at the end of the input stream may return either err ==
// EOF or err == nil. The next Read should return 0, EOF regardless.
//
// Callers should always process the n > 0 bytes returned before considering the
// error err. Doing so correctly handles I/O errors that happen after reading some
// bytes and also both of the allowed EOF behaviors.
//
// Implementations of Read are discouraged from returning a zero byte count with a
// nil error, except when len(p) == 0. Callers should treat a return of 0 and nil
// as indicating that nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.

// Reader 接口包装了基本的 Read 方法.
//
// Read 将 len(p) 个字节读取到 p 中. 它返回读取的字节数 n(0 <= n <= len(p)) 以及任
// 何遇到的错误. 即使 Read 返回的 n < len(p), 它也会在调用过程中使用 p 的全部作为
// 暂存空间. 若一些数据可用但不到 len(p) 个字节, Read 会照例返回可用的东西, 而不是
// 等待更多.
//
// 当 Read 在成功读取 n > 0 个字节后遇到一个错误或 EOF 情况, 它就会返回读取的字节数.
// 它会从该次调用中返回(非nil的)错误或从下一次调用中返回错误(和 n == 0). 这种一般情
// 况的一个例子就是 Reader 在输入流结束时会返回一个非零的字节数, 可能的返回不是
// err == EOF 就是 err == nil. 无论如何, 下一个 Read 都应当返回 0, EOF.
//
// 调用者应当总在考虑到错误 err 前处理 n > 0 的字节. 这样做可以在读取一些字节, 以及
// 允许的 EOF 行为后正确地处理 I/O 错误.
//
// Read 的实现在 len(p) == 0 以外的情况下会阻止返回零字节的计数和 nil 错误, 调用者
// 应将返回 0 和 nil 视作什么也没有发生; 特别是它并不表示 EOF.
//
// 实现必须不保留 p.
type Reader interface {
	Read(p []byte) (n int, err error)
}

// LimitReader returns a Reader that reads from r but stops with EOF after n bytes.
// The underlying implementation is a *LimitedReader.

// LimitReader 返回一个 Reader, 它从 r 中读取 n 个字节后以 EOF 停止. 其基本实现为
// *LimitedReader.
func LimitReader(r Reader, n int64) Reader

// MultiReader returns a Reader that's the logical concatenation of the provided
// input readers. They're read sequentially. Once all inputs have returned EOF,
// Read will return EOF. If any of the readers return a non-nil, non-EOF error,
// Read will return that error.

// MultiReader 返回一个 Reader, 它是所提供的输入 readers 的逻辑拼接. 它们按顺序读取.
// 一旦所有的输入返回 EOF, Read 就会返回 EOF. 若任何 readers 返回了非 nil 或非
// EOF 错误, Read 就会返回该错误.
func MultiReader(readers ...Reader) Reader

// TeeReader returns a Reader that writes to w what it reads from r. All reads from
// r performed through it are matched with corresponding writes to w. There is no
// internal buffering - the write must complete before the read completes. Any
// error encountered while writing is reported as a read error.

// TeeReader 返回一个 Reader, 它将从 r 中读到的东西写入 w 中. 所有通过该接口对 r 的
// 读取都会执行对应的对 w 的写入. 它没有内部缓存, 即写入必须在读取完成前完成. 任何在
// 写入时遇到的错误都将作为读取错误来报告.
func TeeReader(r Reader, w Writer) Reader

// ReaderAt is the interface that wraps the basic ReadAt method.
//
// ReadAt reads len(p) bytes into p starting at offset off in the underlying input
// source. It returns the number of bytes read (0 <= n <= len(p)) and any error
// encountered.
//
// When ReadAt returns n < len(p), it returns a non-nil error explaining why more
// bytes were not returned. In this respect, ReadAt is stricter than Read.
//
// Even if ReadAt returns n < len(p), it may use all of p as scratch space during
// the call. If some data is available but not len(p) bytes, ReadAt blocks until
// either all the data is available or an error occurs. In this respect ReadAt is
// different from Read.
//
// If the n = len(p) bytes returned by ReadAt are at the end of the input source,
// ReadAt may return either err == EOF or err == nil.
//
// If ReadAt is reading from an input source with a seek offset, ReadAt should not
// affect nor be affected by the underlying seek offset.
//
// Clients of ReadAt can execute parallel ReadAt calls on the same input source.
//
// Implementations must not retain p.

// ReaderAt 接口包装了基本的 ReadAt 方法.
//
// ReadAt 从基本输入源的偏移量 off 处开始, 将 len(p) 个字节读取到 p 中.
// 它返回读取的字节数 n(0 <= n <= len(p))以及任何遇到的错误.
//
// 当 ReadAt 返回的 n < len(p)
// 时, 它就会返回一个非nil的错误来解释
// 为什么没有返回更多的字节. 在这一点上, ReadAt 比 Read 更严格.
//
// 即使 ReadAt 返回的 n < len(p), 它也会在调用过程中使用 p 的全部作为暂存空间.
// 若一些数据可用但不到 len(p) 字节, ReadAt
// 就会阻塞直到所有数据都可用或产生一个错误.  在这一点上 ReadAt 不同于 Read.
//
// 若 n = len(p) 个字节在输入源的的结尾处由 ReadAt 返回, ReadAt 不是返回 err == EOF 就是返回 err == nil.
//
// 若 ReadAt 按查找偏移量从输入源读取, ReadAt
// 应当既不影响基本查找偏移量也不被它所影响.
//
// ReadAt 的客户端可对相同的输入源并行执行 ReadAt 调用.
//
// 实现必须不保留 p.
type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}

// ReaderFrom is the interface that wraps the ReadFrom method.
//
// ReadFrom reads data from r until EOF or error. The return value n is the number
// of bytes read. Any error except io.EOF encountered during the read is also
// returned.
//
// The Copy function uses ReaderFrom if available.

// ReaderFrom 接口包装了 ReadFrom 方法.
//
// ReadFrom 从 r 中读取数据, 直到 EOF 或发生错误. 其返回值 n 为读取的字节数.  除 io.EOF
// 之外, 在读取过程中遇到的任何错误也将被返回.
//
// 如果 ReaderFrom 可用, Copy 函数就会使用它.
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

// RuneReader is the interface that wraps the ReadRune method.
//
// ReadRune reads a single UTF-8 encoded Unicode character and returns the rune and
// its size in bytes. If no character is available, err will be set.

// RuneReader 接口包装了 ReadRune 方法.
//
// ReadRune
// 读取单个用UTF-8编码的Unicode字符, 并返回该符文及其字节大小.
// 若没有字符可用, 就会置为 err.
type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}

// RuneScanner is the interface that adds the UnreadRune method to the basic
// ReadRune method.
//
// UnreadRune causes the next call to ReadRune to return the same rune as the
// previous call to ReadRune. It may be an error to call UnreadRune twice without
// an intervening call to ReadRune.

// RuneScanner 接口将 UnreadRune 方法添加到基本的 ReadRune 方法.
//
// UnreadRune 使下一次调用 ReadRune 返回的符文与上一次调用 ReadRune 返回的相同.  调用 UnreadRune
// 两次而中间没有调用 ReadRune 的话就会返回错误.
type RuneScanner interface {
	RuneReader
	UnreadRune() error
}

// SectionReader implements Read, Seek, and ReadAt on a section of an underlying
// ReaderAt.

// SectionReader 在基本 ReaderAt
// 的片段上实现了Read、Seek和ReadAt.
type SectionReader struct {
	// contains filtered or unexported fields
}

// NewSectionReader returns a SectionReader that reads from r starting at offset
// off and stops with EOF after n bytes.

// NewSectionReader 返回一个 SectionReader, 它从 r 中的偏移量 off 处读取 n 个字节后以 EOF 停止.
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader

func (s *SectionReader) Read(p []byte) (n int, err error)

func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)

func (s *SectionReader) Seek(offset int64, whence int) (int64, error)

// Size returns the size of the section in bytes.

// Size 返回片段的字节数.
func (s *SectionReader) Size() int64

// Seeker is the interface that wraps the basic Seek method.
//
// Seek sets the offset for the next Read or Write to offset, interpreted according
// to whence: 0 means relative to the origin of the file, 1 means relative to the
// current offset, and 2 means relative to the end. Seek returns the new offset and
// an error, if any.
//
// Seeking to a negative offset is an error. Seeking to any positive offset is
// legal, but the behavior of subsequent I/O operations on the underlying object is
// implementation-dependent.

// Seeker 接口包装了基本的 Seek 方法.
//
// Seek 将 offset 置为下一个 Read 或 Write 的偏移量 , 它的解释取决于 whence:  0
// 表示相对于文件的起始处, 1 表示相对于当前的偏移, 而 2 表示相对于其结尾处.  Seek
// 返回新的偏移量和一个错误, 如果有的话.
//
// 对负数偏移量进行 Seek 会产生错误. 对任何正数偏移量进行 Seek
// 是合法的, 但对底层类型的后续 I/O 操作行为则取决于具体实现.
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

// WriteCloser is the interface that groups the basic Write and Close methods.

// WriteCloser 接口组合了基本的 Write 和 Close 方法.
type WriteCloser interface {
	Writer
	Closer
}

// WriteSeeker is the interface that groups the basic Write and Seek methods.

// WriteSeeker 接口组合了基本的 Write 和 Seek 方法.
type WriteSeeker interface {
	Writer
	Seeker
}

// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) bytes from p to the underlying data stream. It returns the
// number of bytes written from p (0 <= n <= len(p)) and any error encountered that
// caused the write to stop early. Write must return a non-nil error if it returns
// n < len(p). Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.

// Writer 接口包装了基本的 Write 方法.
//
// Write 将 len(p) 个字节从 p
// 中写入到基本数据流中. 它返回从 p 中被写入的字节数 n(0 <= n <=
// len(p))以及任何遇到的引起写入提前停止的错误. 若 Write 返回的 n <
// len(p), 它就必须返回一个非nil的错误. Write
// 不能修改此切片的数据, 即便它是临时的.
//
// 实现必须不保留 p.
type Writer interface {
	Write(p []byte) (n int, err error)
}

// MultiWriter creates a writer that duplicates its writes to all the provided
// writers, similar to the Unix tee(1) command.

// MultiWriter 创建一个 Writer, 它将其写入复制到所有提供的 writers
// 中, 类似于Unix的tee(1)命令.
func MultiWriter(writers ...Writer) Writer

// WriterAt is the interface that wraps the basic WriteAt method.
//
// WriteAt writes len(p) bytes from p to the underlying data stream at offset off.
// It returns the number of bytes written from p (0 <= n <= len(p)) and any error
// encountered that caused the write to stop early. WriteAt must return a non-nil
// error if it returns n < len(p).
//
// If WriteAt is writing to a destination with a seek offset, WriteAt should not
// affect nor be affected by the underlying seek offset.
//
// Clients of WriteAt can execute parallel WriteAt calls on the same destination if
// the ranges do not overlap.
//
// Implementations must not retain p.

// WriterAt 接口包装了基本的 WriteAt 方法.
//
// WriteAt 从 p 中将 len(p) 个字节写入到偏移量 off
// 处的基本数据流中. 它返回从 p 中被写入的字节数 n(0 <= n <=
// len(p))以及任何遇到的引起写入提前停止的错误.  若 WriteAt 返回的 n <
// len(p), 它就必须返回一个非nil的错误.
//
// 若 WriteAt 按查找偏移量写入到目标中, WriteAt
// 应当既不影响基本查找偏移量也不被它所影响.
//
// 若区域没有重叠, WriteAt 的客户端可对相同的目标并行执行 WriteAt 调用.
//
// 实现必须不保留 p.
type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}

// WriterTo is the interface that wraps the WriteTo method.
//
// WriteTo writes data to w until there's no more data to write or when an error
// occurs. The return value n is the number of bytes written. Any error encountered
// during the write is also returned.
//
// The Copy function uses WriterTo if available.

// WriterTo 接口包装了 WriteTo 方法.
//
// WriteTo 将数据写入 w
// 中, 直到没有数据可读或发生错误. 其返回值 n 为写入的字节数.
// 在写入过程中遇到的任何错误也将被返回.
//
// 如果 WriterTo 可用, Copy 函数就会使用它.
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}
