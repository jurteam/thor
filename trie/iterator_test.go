// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package trie

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/stretchr/testify/assert"
	"github.com/vechain/thor/thor"
)

// makeTestTrie create a sample test trie to test node-wise reconstruction.
func makeTestTrie() (ethdb.Database, *Trie, map[string][]byte) {
	// Create an empty trie
	db := ethdb.NewMemDatabase()
	trie, _ := New(thor.Bytes32{}, db)

	// Fill it with some arbitrary data
	content := make(map[string][]byte)
	for i := byte(0); i < 255; i++ {
		// Map the same data under multiple keys
		key, val := common.LeftPadBytes([]byte{1, i}, 32), []byte{i}
		content[string(key)] = val
		trie.Update(key, val)

		key, val = common.LeftPadBytes([]byte{2, i}, 32), []byte{i}
		content[string(key)] = val
		trie.Update(key, val)

		// Add some other data to inflate the trie
		for j := byte(3); j < 13; j++ {
			key, val = common.LeftPadBytes([]byte{j, i}, 32), []byte{j, i}
			content[string(key)] = val
			trie.Update(key, val)
		}
	}
	trie.Commit()

	// Return the generated trie
	return db, trie, content
}

func TestIterator(t *testing.T) {
	trie := newEmpty()
	vals := []struct{ k, v string }{
		{"do", "verb"},
		{"ether", "wookiedoo"},
		{"horse", "stallion"},
		{"shaman", "horse"},
		{"doge", "coin"},
		{"dog", "puppy"},
		{"somethingveryoddindeedthis is", "myothernodedata"},
	}
	all := make(map[string]string)
	for _, val := range vals {
		all[val.k] = val.v
		trie.Update([]byte(val.k), []byte(val.v))
	}
	trie.Commit()

	found := make(map[string]string)
	it := NewIterator(trie.NodeIterator(nil))
	for it.Next() {
		found[string(it.Key)] = string(it.Value)
	}

	for k, v := range all {
		if found[k] != v {
			t.Errorf("iterator value mismatch for %s: got %q want %q", k, found[k], v)
		}
	}
}

type kv struct {
	k, v []byte
	t    bool
}

func TestIteratorLargeData(t *testing.T) {
	trie := newEmpty()
	vals := make(map[string]*kv)

	for i := byte(0); i < 255; i++ {
		value := &kv{common.LeftPadBytes([]byte{i}, 32), []byte{i}, false}
		value2 := &kv{common.LeftPadBytes([]byte{10, i}, 32), []byte{i}, false}
		trie.Update(value.k, value.v)
		trie.Update(value2.k, value2.v)
		vals[string(value.k)] = value
		vals[string(value2.k)] = value2
	}

	it := NewIterator(trie.NodeIterator(nil))
	for it.Next() {
		vals[string(it.Key)].t = true
	}

	var untouched []*kv
	for _, value := range vals {
		if !value.t {
			untouched = append(untouched, value)
		}
	}

	if len(untouched) > 0 {
		t.Errorf("Missed %d nodes", len(untouched))
		for _, value := range untouched {
			t.Error(value)
		}
	}
}

// Tests that the node iterator indeed walks over the entire database contents.
func TestNodeIteratorCoverage(t *testing.T) {
	// Create some arbitrary test trie to iterate
	db, trie, _ := makeTestTrie()

	// Gather all the node hashes found by the iterator
	hashes := make(map[thor.Bytes32]struct{})
	for it := trie.NodeIterator(nil); it.Next(true); {
		if it.Hash() != (thor.Bytes32{}) {
			hashes[it.Hash()] = struct{}{}
		}
	}
	// Cross check the hashes and the database itself
	for hash := range hashes {
		if _, err := db.Get(hash.Bytes()); err != nil {
			t.Errorf("failed to retrieve reported node %x: %v", hash, err)
		}
	}
	for _, key := range db.(*ethdb.MemDatabase).Keys() {
		if _, ok := hashes[thor.BytesToBytes32(key)]; !ok {
			t.Errorf("state entry not reported %x", key)
		}
	}
}

type kvs struct{ k, v string }

var testdata1 = []kvs{
	{"barb", "ba"},
	{"bard", "bc"},
	{"bars", "bb"},
	{"bar", "b"},
	{"fab", "z"},
	{"food", "ab"},
	{"foos", "aa"},
	{"foo", "a"},
}

var testdata2 = []kvs{
	{"aardvark", "c"},
	{"bar", "b"},
	{"barb", "bd"},
	{"bars", "be"},
	{"fab", "z"},
	{"foo", "a"},
	{"foos", "aa"},
	{"food", "ab"},
	{"jars", "d"},
}

func TestIteratorSeek(t *testing.T) {
	trie := newEmpty()
	for _, val := range testdata1 {
		trie.Update([]byte(val.k), []byte(val.v))
	}

	// Seek to the middle.
	it := NewIterator(trie.NodeIterator([]byte("fab")))
	if err := checkIteratorOrder(testdata1[4:], it); err != nil {
		t.Fatal(err)
	}

	// Seek to a non-existent key.
	it = NewIterator(trie.NodeIterator([]byte("barc")))
	if err := checkIteratorOrder(testdata1[1:], it); err != nil {
		t.Fatal(err)
	}

	// Seek beyond the end.
	it = NewIterator(trie.NodeIterator([]byte("z")))
	if err := checkIteratorOrder(nil, it); err != nil {
		t.Fatal(err)
	}
}

func checkIteratorOrder(want []kvs, it *Iterator) error {
	for it.Next() {
		if len(want) == 0 {
			return fmt.Errorf("didn't expect any more values, got key %q", it.Key)
		}
		if !bytes.Equal(it.Key, []byte(want[0].k)) {
			return fmt.Errorf("wrong key: got %q, want %q", it.Key, want[0].k)
		}
		want = want[1:]
	}
	if len(want) > 0 {
		return fmt.Errorf("iterator ended early, want key %q", want[0])
	}
	return nil
}

func TestDifferenceIterator(t *testing.T) {
	triea := newEmpty()
	for _, val := range testdata1 {
		triea.Update([]byte(val.k), []byte(val.v))
	}
	triea.Commit()

	trieb := newEmpty()
	for _, val := range testdata2 {
		trieb.Update([]byte(val.k), []byte(val.v))
	}
	trieb.Commit()

	found := make(map[string]string)
	di, _ := NewDifferenceIterator(triea.NodeIterator(nil), trieb.NodeIterator(nil))
	it := NewIterator(di)
	for it.Next() {
		found[string(it.Key)] = string(it.Value)
	}

	all := []struct{ k, v string }{
		{"aardvark", "c"},
		{"barb", "bd"},
		{"bars", "be"},
		{"jars", "d"},
	}
	for _, item := range all {
		if found[item.k] != item.v {
			t.Errorf("iterator value mismatch for %s: got %v want %v", item.k, found[item.k], item.v)
		}
	}
	if len(found) != len(all) {
		t.Errorf("iterator count mismatch: got %d values, want %d", len(found), len(all))
	}
}

func TestUnionIterator(t *testing.T) {
	triea := newEmpty()
	for _, val := range testdata1 {
		triea.Update([]byte(val.k), []byte(val.v))
	}
	triea.Commit()

	trieb := newEmpty()
	for _, val := range testdata2 {
		trieb.Update([]byte(val.k), []byte(val.v))
	}
	trieb.Commit()

	di, _ := NewUnionIterator([]NodeIterator{triea.NodeIterator(nil), trieb.NodeIterator(nil)})
	it := NewIterator(di)

	all := []struct{ k, v string }{
		{"aardvark", "c"},
		{"barb", "ba"},
		{"barb", "bd"},
		{"bard", "bc"},
		{"bars", "bb"},
		{"bars", "be"},
		{"bar", "b"},
		{"fab", "z"},
		{"food", "ab"},
		{"foos", "aa"},
		{"foo", "a"},
		{"jars", "d"},
	}

	for i, kv := range all {
		if !it.Next() {
			t.Errorf("Iterator ends prematurely at element %d", i)
		}
		if kv.k != string(it.Key) {
			t.Errorf("iterator value mismatch for element %d: got key %s want %s", i, it.Key, kv.k)
		}
		if kv.v != string(it.Value) {
			t.Errorf("iterator value mismatch for element %d: got value %s want %s", i, it.Value, kv.v)
		}
	}
	if it.Next() {
		t.Errorf("Iterator returned extra values.")
	}
}

func TestIteratorNoDups(t *testing.T) {
	var tr Trie
	for _, val := range testdata1 {
		tr.Update([]byte(val.k), []byte(val.v))
	}
	checkIteratorNoDups(t, tr.NodeIterator(nil), nil)
}

// This test checks that nodeIterator.Next can be retried after inserting missing trie nodes.
func TestIteratorContinueAfterError(t *testing.T) {
	db := ethdb.NewMemDatabase()
	tr, _ := New(thor.Bytes32{}, db)
	for _, val := range testdata1 {
		tr.Update([]byte(val.k), []byte(val.v))
	}
	tr.Commit()
	wantNodeCount := checkIteratorNoDups(t, tr.NodeIterator(nil), nil)
	keys := db.Keys()
	t.Log("node count", wantNodeCount)

	for i := 0; i < 20; i++ {
		// Create trie that will load all nodes from DB.
		tr, _ := New(tr.Hash(), db)

		// Remove a random node from the database. It can't be the root node
		// because that one is already loaded.
		var rkey []byte
		for {
			if rkey = keys[rand.Intn(len(keys))]; !bytes.Equal(rkey, tr.Hash().Bytes()) {
				break
			}
		}
		rval, _ := db.Get(rkey)
		db.Delete(rkey)

		// Iterate until the error is hit.
		seen := make(map[string]bool)
		it := tr.NodeIterator(nil)
		checkIteratorNoDups(t, it, seen)
		missing, ok := it.Error().(*MissingNodeError)
		if !ok || !bytes.Equal(missing.NodeHash[:], rkey) {
			t.Fatal("didn't hit missing node, got", it.Error())
		}

		// Add the node back and continue iteration.
		db.Put(rkey, rval)
		checkIteratorNoDups(t, it, seen)
		if it.Error() != nil {
			t.Fatal("unexpected error", it.Error())
		}
		if len(seen) != wantNodeCount {
			t.Fatal("wrong node iteration count, got", len(seen), "want", wantNodeCount)
		}
	}
}

// Similar to the test above, this one checks that failure to create nodeIterator at a
// certain key prefix behaves correctly when Next is called. The expectation is that Next
// should retry seeking before returning true for the first time.
func TestIteratorContinueAfterSeekError(t *testing.T) {
	// Commit test trie to db, then remove the node containing "bars".
	db := ethdb.NewMemDatabase()
	ctr, _ := New(thor.Bytes32{}, db)
	for _, val := range testdata1 {
		ctr.Update([]byte(val.k), []byte(val.v))
	}
	root, _ := ctr.Commit()
	barNodeHash, _ := thor.ParseBytes32("d32fb77ad25227d60b76d53a512d28137304c9c03556db08a1709563c7ae9c9f")
	barNode, _ := db.Get(barNodeHash[:])
	db.Delete(barNodeHash[:])

	// Create a new iterator that seeks to "bars". Seeking can't proceed because
	// the node is missing.
	tr, _ := New(root, db)
	it := tr.NodeIterator([]byte("bars"))
	missing, ok := it.Error().(*MissingNodeError)
	if !ok {
		t.Fatal("want MissingNodeError, got", it.Error())
	} else if missing.NodeHash != barNodeHash {
		t.Fatal("wrong node missing")
	}

	// Reinsert the missing node.
	db.Put(barNodeHash[:], barNode[:])

	// Check that iteration produces the right set of values.
	if err := checkIteratorOrder(testdata1[2:], NewIterator(it)); err != nil {
		t.Fatal(err)
	}
}

func checkIteratorNoDups(t *testing.T, it NodeIterator, seen map[string]bool) int {
	if seen == nil {
		seen = make(map[string]bool)
	}
	for it.Next(true) {
		if seen[string(it.Path())] {
			t.Fatalf("iterator visited node path %x twice", it.Path())
		}
		seen[string(it.Path())] = true
	}
	return len(seen)
}

func TestIteratorNodeFilter(t *testing.T) {
	db := ethdb.NewMemDatabase()
	tr, _ := NewExtended(thor.Bytes32{}, 0, db)
	for _, val := range testdata1 {
		tr.Update([]byte(val.k), []byte(val.v), nil)
	}
	root1, _ := tr.Commit(1)
	_ = root1
	for _, val := range testdata2 {
		tr.Update([]byte(val.k), []byte(val.v), nil)
	}
	root2, _ := tr.Commit(2)

	tr, _ = NewExtended(root2, 2, db)

	it := tr.NodeIterator(nil, func(path []byte, commitNum uint32) bool {
		return commitNum == 1
	})

	for it.Next(true) {
		if h := it.Hash(); !h.IsZero() {
			assert.Equal(t, uint32(1), it.CommitNum())
		}
	}

	it = tr.NodeIterator(nil, func(path []byte, commitNum uint32) bool {
		return commitNum == 2
	})

	for it.Next(true) {
		if h := it.Hash(); !h.IsZero() {
			assert.Equal(t, uint32(2), it.CommitNum())
		}
	}
}
