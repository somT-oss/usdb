package storage

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
 The write-ahead-log is a program that stores the state of the original page before it is modified, enabling
 both commit (by deleting the entry in the wal) and rollback (by restoring pages from the wal) operations. it uses a chuncked
 file format with in-memory buffering to minimize I/O operations.
*/

/*
	The first method i want to implement is the one that handles writing the contents of the page to the buffer.
	write(pageNumber int, pageContentSize []bytes) -> None:
		- takes the page number and page content/size and writes it to the buffer.
		- make sure that it writes it in this format pageNo,pageContent,checksum.
		-
*/

const (
	checksum=32
	MAGICPHRASE="0x12232o"
)


 type WAL struct {
	buffer bytes.Buffer
	maxPages uint32
	bufferedPages uint32
	pageContent [4096]byte	
	file *os.File
	filePath string
}

type WALInterface interface {
	Write() error
	Push() error
	Clear() error
}

func NewWAL(buffer bytes.Buffer, maxPage uint32, bufferedPages uint32, pageContent[4096]byte, file *os.File, filePath string) *WAL {
	
	file.Write([]byte(MAGICPHRASE)) // the first 8 bytes of the wal file are 8 bytes to confirm the validity of the file.
	
	pageCountInBuffer := 0
	pageCountInBufferBytes := make([]byte, 4)
	
	binary.BigEndian.PutUint32(pageCountInBufferBytes, uint32(pageCountInBuffer)) // the next 4 bytes are the number of pages written so far. 
	
	file.Write(pageCountInBufferBytes)
	defer file.Close()
	
	return &WAL{
		buffer: buffer,
		maxPages: maxPage,
		bufferedPages: bufferedPages,
		pageContent: pageContent,
		file: file,
		filePath: filePath,
	}
}

/*
	implement these functions, for the wal.
	push(page_number, page_content) - Adds a page to the in-memory buffer, writes to disk when buffer is full -- DONE.
	write() - Writes all the content of the buffer to the WAL file.
	clear() - Clears the current buffer so that we can start to store new data afresh.
*/

func (wal *WAL) Push(pageNumber uint32, pageContent []byte) error  {
	if wal.bufferedPages >= wal.maxPages {
		log.Println("the buffer is already at full capacity.")
		err := wal.Write() // write to the wal file.
		if err != nil {
			return errors.New("failed to write the contents of the buffer to the wal file.")
		}
	}
	
	// add the curr bytes(pageNumber) + pageContent + checksum to the current buffer
	pageNumberBytes := make([]byte, 4)
	checksumBytes := make([]byte, 4)
	
	binary.BigEndian.PutUint32(pageNumberBytes, pageNumber) // write the content of the integer byte conversion to the pageNumberBytes integer buffer.
	binary.BigEndian.PutUint32(checksumBytes, checksum)
	
	// write the pagenumberbytes, pagecontent and checksum to the buffer.
	wal.buffer.Write(pageNumberBytes)
	wal.buffer.Write(pageContent)
	wal.buffer.Write(checksumBytes)
	
	wal.bufferedPages += 1
	
	return nil
} 

func (wal *WAL) Write() error {
	
	wal.file.Write(wal.buffer.Bytes()) // write the buffer to the wal file.
	
	// update the number of pages in the file += 1
	// read the portion of the page that has the byte representation of the number of pages.
	readOffset := int64(9)
	buffLength := 4
	buff := make([]byte, buffLength)
	_, err := wal.file.ReadAt(buff, readOffset)
	if err != nil {
		return fmt.Errorf("could not read the pageCount from the wal.")
	}
	// convert it to an int and update the count.
	currPageCount, err := strconv.ParseInt((string(buff)), 10, 32)
	if err != nil {
		return fmt.Errorf("could not parse pageCount from string to int64.")
	}
	currPageCount += 1
	
	// write the updated pagecount as bytes in the same portion that has th byte representation of the number of pages.
	writeOffset := int64(9)
	writeBuff := make([]byte, buffLength)
	
	binary.BigEndian.PutUint32(writeBuff, uint32(currPageCount))
	
	_, err = wal.file.WriteAt(writeBuff, writeOffset)
	if err != nil {
		return fmt.Errorf("could not write the updated ")
	}
	
	defer wal.file.Close()
	return nil
}

func (wal *WAL) Clear() error {
	return errors.New("failed to clear the buffer.")
}

