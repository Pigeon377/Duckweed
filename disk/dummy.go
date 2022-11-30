package disk

import (
	"Duckweed/page"
	"fmt"
)

type DummyDiskManager struct {
}

func NewDummyDiskManager() *DummyDiskManager {
	return &DummyDiskManager{}
}

func (dm *DummyDiskManager) Write(page *page.Page) {
	fmt.Printf("Write Page(ID=%d) To Disk \n", page.GetPageID())
}

func (dm *DummyDiskManager) BatchWrite(pages []*page.Page) {
	fmt.Println("Start Batch Write")
	for _, p := range pages {
		dm.Write(p)
	}
	fmt.Println("End Batch Write")
	return
}

func (dm *DummyDiskManager) Read(pageID int) *page.Page {
	return nil
}