package api_test

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	// api "github.com/avag-sargsyan/stricache/cmd/stricache/api"
	"github.com/avag-sargsyan/stricache/proto/stricache"
)

var (
	conn *grpc.ClientConn
	err  error
)

func TestAPI(t *testing.T) {
	conn, err = grpc.Dial("127.0.0.1:7999", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	c := stricache.NewStricacheServiceClient(conn)

	item1 := &stricache.StringItem{
		Key:   "test1",
		Value: "test",
	}

	item2 := &stricache.IntItem{
		Key:   "test2",
		Value: 123,
	}

	item3 := &stricache.FloatItem{
		Key:   "test3",
		Value: 456.78,
	}

	r1, err1 := c.AddString(context.Background(), item1)
	r2, err2 := c.AddInt(context.Background(), item2)
	r3, err3 := c.AddFloat(context.Background(), item3)

	if err1 != nil {
		t.Error(err1)
	}
	t.Log("Response from server for adding a key", r1)

	if err2 != nil {
		t.Error(err2)
	}
	t.Log("Response from server for adding a key", r2)

	if err3 != nil {
		t.Error(err3)
	}
	t.Log("Response from server for adding a key", r3)

	getStr := &stricache.GetKey{
		Key: "test1",
	}

	getKeyRes, err := c.GetString(context.Background(), getStr)
	if err != nil {
		t.Error(err)
	}
	t.Log("Response for getting the key", getKeyRes)
}
