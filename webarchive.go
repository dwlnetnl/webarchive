// Package webarchive reads Safari .webarchive files.
package webarchive

import (
	"bytes"
	"io"

	"github.com/dhowett/go-plist"
)

// Archive represents a .webarchive file.
type Archive struct {
	Content   MainResource `plist:"WebMainResource"`
	Resources SubResources `plist:"WebSubresources"`
}

// New returns the parsed archive or an error.
func New(r io.ReadSeeker) (*Archive, error) {
	d := plist.NewDecoder(r)
	var a Archive
	err := d.Decode(&a)
	return &a, err
}

// Resource represents a generic WebResource.
type Resource struct {
	Data     []byte `plist:"WebResourceData"`
	MIMEType string `plist:"WebResourceMIMEType"`
	Encoding string `plist:"WebResourceTextEncodingName"`
	URL      string `plist:"WebResourceURL"`
}

// Reader returns a data reader.
func (r Resource) Reader() io.Reader {
	return bytes.NewReader(r.Data)
}

// MainResource represents a WebMainResource.
type MainResource struct {
	Resource
	FrameName string `plist:"WebResourceFrameName"`
}

// SubResource represents a WebResource in the WebSubresources array.
type SubResource struct {
	Resource
	Response []byte `plist:"WebResourceResponse"`
}

// SubResources represents the WebSubresources array.
type SubResources []SubResource
