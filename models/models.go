package models

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalTimestamp marshal Timestamp to GraphQL
func MarshalTimestamp(t int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatInt(t, 10))
	})
}

// UnmarshalTimestamp unmarshal Timestamp from GraphQL
func UnmarshalTimestamp(v interface{}) (int64, error) {
	if res, ok := v.(json.Number); ok {
		return res.Int64()
	}

	if res, ok := v.(string); ok {
		return json.Number(res).Int64()
	}

	if res, ok := v.(int64); ok {
		return res, nil
	}

	if res, ok := v.(*int64); ok {
		return *res, nil
	}

	return 0, fmt.Errorf("timestamp %v: is invalid", v)
}

// MarshalID marshal ID to GraphQL
func MarshalID(id string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, id)
	})
}

// UnmarshalID unmarshal ID from GraphQL
func UnmarshalID(v interface{}) (string, error) {
	mac, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("ID %v: is invalid", v)
	}

	return mac, nil
}
