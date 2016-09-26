package middleware

// Key - key for context
type Key int

const reqIDKey Key = 0
const iDataStoreKey Key = 1

// GetReqIDKey - Retreive RequestID Context Key
func GetReqIDKey() Key {
	return reqIDKey
}

// GetIDataStoreKey - Retreive IDataStore Context Key
func GetIDataStoreKey() Key {
	return iDataStoreKey
}
