package kvstore

type KVStore interface {
	// Get retrieves the value associated with the given key from the map.
	//
	// key: The key to look up in the map.
	// Returns the value associated with the key and a boolean indicating
	// whether the key was found in the map.
	Get(key interface{}) (interface{}, bool)

	// Set updates the value of the given key in the data structure.
	//
	// key: interface{} - The key to update.
	// value: interface{} - The new value to associate with the given key.
	// ttl: int64 - time to live with millisecond
	Set(key interface{}, value interface{})

	// Delete removes the key and its corresponding value from the map.
	//
	// key: the key to be deleted.
	Delete(key interface{})

	Purge()
}
