package silk

import (
	"fmt"
	"os"
	"testing"
)

// const F_APPENDONLY = 1 << 0
// const F_LZ4 = 1 << 2
// const F_SPLAYTREE = 1 << 1
// const F_AOL_FFLUSH = 1 << 3

// func openRandomDB(features int) (goleg.Database, string, error) {
// 	name, err := ioutil.TempDir("/tmp", "goleg")
// 	if err != nil {
// 		return goleg.Database{}, "", err
// 	}
// 	//F_APPENDONLY|F_AOL_FFLUSH|F_LZ4|F_SPLAYTREE
// 	database, err := goleg.Open(name, "test", features)
// 	if err != nil {
// 		return goleg.Database{}, "", err
// 	}
// 	return database, name, nil
// }

// func cleanTemp(dir string) {
// 	os.RemoveAll(dir)
// }

// func TestOpen(t *testing.T) {
// 	t.Log("TestOpen\n")
// 	if testing.Short() {
// 		t.Skip("Skipping in short mode")
// 	}

func TestOpen(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping in short mode")
	}
	t.Log("TestOpen\n")
	test_silk := New()
	os.MkdirAll("../dbs/testdb", os.ModePerm)
	test_silk.OpenDatabase("../dbs/testdb", "testsilk", FEATS.AppendOnly)

	test_string := "I wanna be the very best"
	key, uuids := KeyFromValue(test_string)
	node := RelationalNode{
		chunks: uuids,
		Key:    key,
		name:   test_string,
		links:  nil,
		tags:   nil,
	}
	test_silk.PushNode(node)
	dump(test_silk.PullNode(node.Key))
	test_silk.CloseDatabase()
}

// 	database, dir, err := openRandomDB(F_APPENDONLY)
// 	if err != nil {
// 		t.Fatalf("Can't open database: %s", err.Error())
// 	}

// 	database.Close()
// 	cleanTemp(dir)
// }

// const JARN = 10

// func TestJar(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping in short mode")
// 	}

// 	database, dir, err := openRandomDB(F_LZ4 | F_SPLAYTREE)
// 	if err != nil {
// 		t.Fatalf("Can't open database: %s", err.Error())
// 	}

// 	for i := 0; i < JARN; i++ {
// 		if database.Jar("record"+strconv.Itoa(i), []byte("value"+strconv.Itoa(i))) != 0 {
// 			t.Fatalf("Can't jar value #%d", i)
// 		}
// 	}

// 	database.Close()
// 	cleanTemp(dir)
// }

// func TestUnjar(t *testing.T) {
// 	database, dir, err := openRandomDB(F_LZ4 | F_SPLAYTREE)
// 	if err != nil {
// 		t.Fatalf("Can't open database: %s", err.Error())
// 	}

// 	for i := 0; i < JARN; i++ {
// 		if database.Jar("record"+strconv.Itoa(i), []byte("value"+strconv.Itoa(i))) != 0 {
// 			t.Fatalf("Can't jar value #%d", i)
// 		}
// 	}

// 	for i := 0; i < JARN; i++ {
// 		val := database.Unjar("record" + strconv.Itoa(i))
// 		if !bytes.Equal(val, []byte("value"+strconv.Itoa(i))) {
// 			t.Errorf("Value #%d doesn't match", i)
// 		}
// 	}

// 	database.Close()
// 	cleanTemp(dir)
// }

// func TestFullKeyDump(t *testing.T) {
// 	database, _, err := openRandomDB(F_LZ4 | F_SPLAYTREE)
// 	if err != nil {
// 		t.Fatalf("Can't open database: %s", err.Error())
// 	}

// 	for i := 0; i < JARN; i++ {
// 		if database.Jar("record"+strconv.Itoa(i), []byte("value"+strconv.Itoa(i))) != 0 {
// 			t.Fatalf("Can't jar value #%d", i)
// 		}
// 	}

// 	gotKeys, keys := database.DumpKeys()

// 	if !gotKeys {
// 		t.Fatal("Didn't get keys and should have")
// 	}

// 	var j int
// 	for i := 0; i <= JARN; i++ {
// 		for _, key := range keys {
// 			if key == "record"+strconv.Itoa(i) {
// 				j++
// 			}
// 		}
// 	}
// 	if j != JARN {
// 		t.Fatal("One or more keys did not dump")
// 	}
// }

// func TestBulkUnjarOnlyReturnsKeysWeGiveIt(t *testing.T) {
// 	database, _, err := openRandomDB(F_LZ4 | F_SPLAYTREE)
// 	if err != nil {
// 		t.Fatalf("Can't open database: %s", err.Error())
// 	}

// 	keys := []string{"key0", "key1", "key2", "key3"}

// 	for i, key := range keys {
// 		if database.Jar(key, []byte("value"+strconv.Itoa(i))) != 0 {
// 			t.Fatalf("Can't jar value #%d", i)
// 		}
// 	}

// 	subset := keys[1:] //sans key0

// 	values := database.BulkUnjar(subset)

// 	if l := len(values); l != 3 {
// 		t.Fatalf("Expected a length of 3, got %d", l)
// 	}

// 	for i, value := range values {
// 		if subset[i][3] != string(value)[5] {
// 			t.Fatalf("Expected %b, got %b", subset[i][3], string(value)[5])
// 		}
// 	}
// }

func dump(a any) {
	fmt.Printf("%+v\n", a)
}
