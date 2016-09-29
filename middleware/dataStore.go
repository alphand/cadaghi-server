package middleware

import (
	"net/http"

	"golang.org/x/net/context"

	db "github.com/alphand/skilltree-server/database"
)

//InitFunc - function to initialize
type InitFunc func(db.IDataStore)

//DataStoreMW - datastore middleware
type DataStoreMW struct {
	ds   db.IDataStore
	init InitFunc
}

func (d *DataStoreMW) setDSContext(ctx context.Context, storage db.IDataStore) context.Context {
	if d.init != nil {
		d.init(storage)
	}
	return context.WithValue(ctx, GetIDataStoreKey(), storage)
}

func (d *DataStoreMW) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	ctx := req.Context()
	ctx = d.setDSContext(ctx, d.ds)

	next(rw, req.WithContext(ctx))
}

//NewDataStoreMW - Create new Datastore middleware
func NewDataStoreMW(ds db.IDataStore, f InitFunc) *DataStoreMW {
	return &DataStoreMW{
		ds:   ds,
		init: f,
	}
}

// GetDSFromCtx - Get DS Config from context
func GetDSFromCtx(ctx context.Context) db.IDataStore {
	return ctx.Value(GetIDataStoreKey()).(db.IDataStore)
}
