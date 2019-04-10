package entities

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/datastore"
	"strings"
)

type BlobResponse struct {
	Blobs  []*blobstore.BlobInfo `json:"blobs"`
	Cursor string                `json:"cursor"`
}

const ENTITY_BLOB = "__BlobInfo__"

func ListUploads(ctx context.Context, cursorStr string, limit int) (*BlobResponse, error) {
	query := datastore.NewQuery(ENTITY_BLOB).Limit(limit)
	if cursorStr != "" {
		cursor, err := datastore.DecodeCursor(cursorStr)
		if err != nil {
			return nil, err
		}

		query = query.Start(cursor)
	}

	t := query.Run(ctx)
	blobs := make([]*blobstore.BlobInfo, 0)
	for {
		var blob blobstore.BlobInfo
		key, err := t.Next(&blob)
		if err == datastore.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		blob.BlobKey = appengine.BlobKey(key.StringID())
		blobs = append(blobs, &blob)
	}

	if cursor, err := t.Cursor(); err == nil {
		cursorStr = cursor.String()
		if err != nil {
			return nil, err
		}
	}

	blobResp := BlobResponse{
		Blobs:  blobs,
		Cursor: cursorStr,
	}
	return &blobResp, nil
}

func GetUpload(ctx context.Context, key appengine.BlobKey) (*blobstore.BlobInfo, error) {
	blob := &blobstore.BlobInfo{}
	k := datastore.NewKey(ctx, ENTITY_BLOB, string(key), 0, nil)
	err := datastore.Get(ctx, k, blob)
	if err != nil {
		index := strings.Index(err.Error(), "datastore: cannot load field")
		if index != 0 {
			return nil, err
		}
	}

	blob.BlobKey = key
	return blob, nil
}

func GetUploadByName(ctx context.Context, name string) (*blobstore.BlobInfo, error) {
	blobs := make([]*blobstore.BlobInfo, 0)
	keys, err := datastore.NewQuery(ENTITY_BLOB).Filter("filename=", name).Limit(1).GetAll(ctx, &blobs)
	if err != nil {
		index := strings.Index(err.Error(), "datastore: cannot load field")
		if index != 0 {
			return nil, err
		}
	}

	if len(blobs) == 0 {
		return nil, datastore.ErrNoSuchEntity
	}

	for i, blob := range blobs {
		blob.BlobKey = appengine.BlobKey(keys[i].StringID())
	}

	return blobs[0], nil
}
