package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type envelope map[string]any

func WriteJSON(w http.ResponseWriter, status int, data envelope) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func ReadJSONObject(r io.Reader, dto any) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dto); err != nil {
		return err
	}

	err := decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body should  only contain a single object")
	}
	return nil
}
