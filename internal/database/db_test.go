// internal/database/db.go

package database

import (
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestNewDB(t *testing.T) {
	type args struct {
		dataSourceName string
	}
	tests := []struct {
		name    string `json:"name,omitempty"`
		args    args   `json:"args,omitempty"`
		want    *DB    `json:"want,omitempty"`
		wantErr bool   `json:"want_err,omitempty"`
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDB(tt.args.dataSourceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTable(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name    string `json:"name,omitempty"`
		args    args   `json:"args,omitempty"`
		wantErr bool   `json:"want_err,omitempty"`
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTable(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("CreateTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_CreateBook(t *testing.T) {
	type args struct {
		book *Book
	}
	tests := []struct {
		name    string `json:"name,omitempty"`
		db      *DB    `json:"db,omitempty"`
		args    args   `json:"args,omitempty"`
		wantErr bool   `json:"want_err,omitempty"`
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.CreateBook(tt.args.book); (err != nil) != tt.wantErr {
				t.Errorf("DB.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_GetBook(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string `json:"name,omitempty"`
		db      *DB    `json:"db,omitempty"`
		args    args   `json:"args,omitempty"`
		want    *Book  `json:"want,omitempty"`
		wantErr bool   `json:"want_err,omitempty"`
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetBook(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.GetBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_UpdateBook(t *testing.T) {
	type args struct {
		book *Book
	}
	tests := []struct {
		name    string `json:"name,omitempty"`
		db      *DB    `json:"db,omitempty"`
		args    args   `json:"args,omitempty"`
		wantErr bool   `json:"want_err,omitempty"`
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.UpdateBook(tt.args.book); (err != nil) != tt.wantErr {
				t.Errorf("DB.UpdateBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_DeleteBook(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string `json:"name,omitempty"`
		db      *DB    `json:"db,omitempty"`
		args    args   `json:"args,omitempty"`
		wantErr bool   `json:"want_err,omitempty"`
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.DeleteBook(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DB.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_ListBooks(t *testing.T) {
	type args struct {
		limit  int
		offset int
	}
	tests := []struct {
		name    string `json:"name,omitempty"`
		db      *DB    `json:"db,omitempty"`
		args    args   `json:"args,omitempty"`
		want    []Book `json:"want,omitempty"`
		wantErr bool   `json:"want_err,omitempty"`
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.ListBooks(tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.ListBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.ListBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}
