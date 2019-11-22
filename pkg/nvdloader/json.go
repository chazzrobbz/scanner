package nvdloader

import (
	"io"

	"github.com/facebookincubator/nvdtools/cvefeed/nvd/schema"
	"github.com/mailru/easyjson"
	"github.com/pkg/errors"
)

// easyjson:json
type itemWrapper schema.NVDCVEFeedJSON10DefCVEItem

// easyjson:json
type feedWrapper schema.NVDCVEFeedJSON10

// LoadJSONFileFromReader uses easy JSON to load the NVD feed from a given io.Reader.
// It does NOT close the reader; that is the caller's responsibility.
func LoadJSONFileFromReader(r io.Reader) (*schema.NVDCVEFeedJSON10, error) {
	var feed feedWrapper
	if err := easyjson.UnmarshalFromReader(r, &feed); err != nil {
		return nil, errors.Wrap(err, "unmarsaling JSON from reader")
	}
	return (*schema.NVDCVEFeedJSON10)(&feed), nil
}

// MarshalNVDFeedCVEItem marshals the given NVD feed item using easyjson.
func MarshalNVDFeedCVEItem(item *schema.NVDCVEFeedJSON10DefCVEItem) ([]byte, error) {
	bytes, err := easyjson.Marshal((*itemWrapper)(item))
	if err != nil {
		return nil, errors.Wrap(err, "marshaling CVE item as JSON")
	}
	return bytes, nil
}

// UnmarshalNVDFeedCVEItem unmarshals the given bytes into the NVD CVE item struct using easyjson.
func UnmarshalNVDFeedCVEItem(bytes []byte) (*schema.NVDCVEFeedJSON10DefCVEItem, error) {
	var item itemWrapper
	if err := easyjson.Unmarshal(bytes, &item); err != nil {
		return nil, errors.Wrap(err, "unmarshaling CVE item")
	}
	return (*schema.NVDCVEFeedJSON10DefCVEItem)(&item), nil
}