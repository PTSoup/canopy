package store

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/alecthomas/units"
	"github.com/canopy-network/canopy/lib"
	"github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/require"
	math2 "math"
	math "math/rand"
	"os"
	"runtime/debug"
	"testing"
	"time"
)

func TestMaxTransaction(t *testing.T) {
	t.Skip()
	db, err := badger.OpenManaged(badger.DefaultOptions("test_temp.db").WithMemTableSize(2 * int64(units.GB)).
		WithNumVersionsToKeep(math2.MaxInt).WithLoggingLevel(badger.ERROR))
	require.NoError(t, err)
	defer func() { db.Close(); os.RemoveAll("test_temp.db") }()
	tx := db.NewTransactionAt(1, true)
	i := 0
	totalBytes := 0
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(string(debug.Stack()))
			fmt.Println(i)
			fmt.Println(totalBytes)
		}
	}()
	for ; i < 10000000; i++ {
		k := numberToBytes(i)
		v := bytes.Repeat([]byte("b"), 60)
		totalBytes += len(k) + len(v)
		if err = tx.Set(k, v); err != nil {
			fmt.Println(err.Error())
			fmt.Println(i)
			fmt.Println(totalBytes)
			tx.Discard()
			return
		}
	}
	fmt.Println(totalBytes)
	start := time.Now()
	require.NoError(t, tx.CommitAt(1, nil))
	fmt.Println(time.Since(start))
}

func numberToBytes(n int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint32(n))
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func TestStoreSetGetDelete(t *testing.T) {
	store, _, cleanup := testStore(t)
	defer cleanup()
	key, val := []byte("key"), []byte("val")
	require.NoError(t, store.Set(key, val))
	gotVal, err := store.Get(key)
	require.NoError(t, err)
	require.Equal(t, val, gotVal, fmt.Sprintf("wanted %s got %s", string(val), string(gotVal)))
	require.NoError(t, store.Delete(key))
	gotVal, err = store.Get(key)
	require.NoError(t, err)
	require.NotEqualf(t, gotVal, val, fmt.Sprintf("%s should be delete", string(val)))
	require.NoError(t, store.Close())
}

func TestIteratorCommitBasic(t *testing.T) {
	parent, _, cleanup := testStore(t)
	defer cleanup()
	prefix := "a/"
	expectedVals := []string{prefix + "a", prefix + "b", prefix + "c", prefix + "d", prefix + "e", prefix + "f", prefix + "g", prefix + "i", prefix + "j"}
	expectedValsReverse := []string{prefix + "j", prefix + "i", prefix + "g", prefix + "f", prefix + "e", prefix + "d", prefix + "c", prefix + "b", prefix + "a"}
	bulkSetKV(t, parent, prefix, "a", "c", "e", "g")
	_, err := parent.Commit()
	require.NoError(t, err)
	bulkSetKV(t, parent, prefix, "b", "d", "f", "h", "i", "j")
	require.NoError(t, parent.Delete([]byte(prefix+"h")))
	// forward - ensure cache iterator matches behavior of normal iterator
	cIt, err := parent.Iterator([]byte(prefix))
	require.NoError(t, err)
	validateIterators(t, expectedVals, cIt)
	cIt.Close()
	// backward - ensure cache iterator matches behavior of normal iterator
	rIt, err := parent.RevIterator([]byte(prefix))
	require.NoError(t, err)
	validateIterators(t, expectedValsReverse, rIt)
	rIt.Close()
}

func TestIteratorCommitAndPrefixed(t *testing.T) {
	store, _, cleanup := testStore(t)
	defer cleanup()
	prefix := "test/"
	prefix2 := "test2/"
	bulkSetKV(t, store, prefix, "a", "b", "c")
	bulkSetKV(t, store, prefix2, "c", "d", "e")
	it, err := store.Iterator([]byte(prefix))
	require.NoError(t, err)
	validateIterators(t, []string{"test/a", "test/b", "test/c"}, it)
	it.Close()
	it2, err := store.Iterator([]byte(prefix2))
	require.NoError(t, err)
	validateIterators(t, []string{"test2/c", "test2/d", "test2/e"}, it2)
	it2.Close()
	root1, err := store.Commit()
	require.NoError(t, err)
	it3, err := store.RevIterator([]byte(prefix))
	require.NoError(t, err)
	validateIterators(t, []string{"test/c", "test/b", "test/a"}, it3)
	it3.Close()
	root2, err := store.Commit()
	require.NoError(t, err)
	require.Equal(t, root1, root2)
	it4, err := store.RevIterator([]byte(prefix2))
	require.NoError(t, err)
	validateIterators(t, []string{"test2/e", "test2/d", "test2/c"}, it4)
	it4.Close()
}

//func TestProof(t *testing.T) {
//	store, _, cleanup := testStore(t)
//	defer cleanup()
//	key, val := []byte("key"), []byte("val")
//	require.NoError(t, store.Set(key, val))
//	addRandomValues(t, store)
//	proof, value, err := store.GetProof(key)
//	require.NoError(t, err)
//	require.Equal(t, val, value, fmt.Sprintf("wanted %s got %s", string(val), string(value)))
//	testProof(t, store, key, value, proof)
//	// proof of non-inclusion
//	nonInclusionKey := []byte("lther")
//	proof, value, err = store.GetProof(nonInclusionKey)
//	require.NoError(t, err)
//	require.Nil(t, value, "wanted nil bytes for non-inclusion")
//	testProof(t, store, nonInclusionKey, value, proof)
//}
//
//func testProof(t *testing.T, store *Store, key, value, proof []byte) {
//	require.True(t, store.VerifyProof(key, value, proof), "expected valid proof")
//	require.False(t, store.VerifyProof(key, []byte("other value"), proof), "expected invalid proof; other val")
//	require.False(t, store.VerifyProof([]byte("other key"), value, proof), "expected invalid proof; other key")
//}

func addRandomValues(t *testing.T, store *Store) {
	for i := 0; i < math.Intn(1000); i++ {
		key := make([]byte, 256)
		_, err := rand.Read(key)
		require.NoError(t, err)
		value := make([]byte, 256)
		_, err = rand.Read(value)
		require.NoError(t, err)
		err = store.Set(key, value)
		require.NoError(t, err)
	}
}

func testStore(t *testing.T) (*Store, *badger.DB, func()) {
	db, err := badger.OpenManaged(badger.DefaultOptions("").
		WithInMemory(true).WithLoggingLevel(badger.ERROR))
	require.NoError(t, err)
	store, err := NewStoreWithDB(db, lib.NewDefaultLogger(), true)
	require.NoError(t, err)
	return store, db, func() { store.Close() }
}

func validateIterators(t *testing.T, expectedKeys []string, iterators ...lib.IteratorI) {
	for _, it := range iterators {
		for i := 0; it.Valid(); func() { i++; it.Next() }() {
			got, wanted := string(it.Key()), expectedKeys[i]
			require.Equal(t, wanted, got, fmt.Sprintf("wanted %s got %s", wanted, got))
		}
	}
}

func bulkSetKV(t *testing.T, store lib.WStoreI, prefix string, keyValue ...string) {
	for _, kv := range keyValue {
		require.NoError(t, store.Set([]byte(prefix+kv), []byte(kv)))
	}
}
