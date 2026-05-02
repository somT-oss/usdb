package storage

// /*
//  A page in a database is the smallest unit of storage 
//  that a database uses to organize data on disk. Pages are typically fixed in size
//  (usually 8KB or 16KB) and are designed to efficiently manage reads and writes.
 
//  Every page contains a 
//  	- page header; metadata about the page.
//   	- data rows; the actual data stored in the page.
//    	- free space; reserved space for updates or new rows.
//     - offset row; pointers to the location of rows within the page.
//  */
 
//  const (
//  	pageNumber = 10
//  )
 
// type PageHeader struct {
// 	Pager *Pager // the pager to which this page belongs to
// 	pageNumber uint32 // the page number for this page
// 	pageNextHash *PageHeader // hash collision chain for the pageHeader.pageNumber
// 	pagePrevHash *PageHeader // hash collison chain for the pageHeader.pageNumber
// 	nRef int32 // number of users of this page
// 	nextFree *PageHeader // freelist of pages where nRef==0
// 	prevFree *PageHeader // freelist of pages where nRef==0
// 	nextCheckPoint *PageHeader 
// 	prevCheckPoint *PageHeader
// 	inJournal bool // this is true if the page has already been written to a file.
// 	inCheckPoint bool // this is true if the page has been written to the checkpoint journal. 
// 	dirty bool // this will be true if we need to write back the changes from this current file in the journal to the page itself.
// 	needSycn bool // sync journal before writing this page.
// 	alwaysRollback bool // disable dont_rollback() for this page.
// 	dirtyPage *PageHeader
// }

// type Pager struct {
// 	filename string // name of the database file
// 	journal string // name of the journal file
// 	jfd string // journal file description for the journal file
// 	fd string // journal file description for the pager file itself.
	
// }
// 

type HashSet struct {
	set map[string]struct{}
}


type BufferPool struct {
	pool string
}

type WriteAheadLog struct {
	wal string
}

type FileHandler struct {
	filePath string
}

type Pager struct {
	file FileHandler
	bufferPool BufferPool
	wal WriteAheadLog
	dirtyPages HashSet
	journalPages HashSet
	pageSize int32
	blockSize int32
}