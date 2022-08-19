package file

import (
	"fmt"
	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	if fileID, err := uuid.NewUUID(); err == nil {
		return &File{
			ID:   fileID,
			Name: filename,
			Data: blob,
		}, nil
	} else {
		return nil, err
	}
}

func (f *File) String() string {
	return fmt.Sprintf("File(%s, %v)", f.Name, f.ID)
}
