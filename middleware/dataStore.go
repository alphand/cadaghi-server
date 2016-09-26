package middleware

import (
	"net/http"

	"golang.org/x/net/context"

	db "github.com/alphand/skilltree-server/database"
)

//DataStoreMW - datastore middleware
type DataStoreMW struct {
	ds db.IDataStore
}

func (d *DataStoreMW) setDSContext(ctx context.Context, storage db.IDataStore) context.Context {
	return context.WithValue(ctx, GetIDataStoreKey(), storage)
}

func (d *DataStoreMW) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	ctx := req.Context()
	ctx = d.setDSContext(ctx, d.ds)

	next(rw, req.WithContext(ctx))
}

//NewDataStoreMW - Create new Datastore middleware
func NewDataStoreMW(ds db.IDataStore) *DataStoreMW {
	return &DataStoreMW{
		ds: ds,
	}
}

// GetDSFromCtx - Get DS Config from context
func GetDSFromCtx(ctx context.Context) db.IDataStore {
	return ctx.Value(GetIDataStoreKey()).(db.IDataStore)
}
