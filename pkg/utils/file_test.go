package utils

import (
	"testing"
)

func TestWrite(t *testing.T) {
	type testStruct struct {
		Test string
	}

	testCases := []struct {
		name     string
		datas    testStruct
		filePath string
		wantErr  bool
	}{
		{"success", testStruct{"test"}, "test.json", false},
		{"fail", testStruct{"test"}, "/no_permission_dir/test.json", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Write(tc.datas, tc.filePath)
			if (err != nil) != tc.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestRead(t *testing.T) {
	type testStruct struct {
		Test string
	}

	testCases := []struct {
		name     string
		datas    *testStruct
		filePath string
		wantErr  bool
	}{
		{"success", &testStruct{}, "test.json", false},
		{"fail", &testStruct{}, "/no_exist_file.json", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Read(tc.datas, tc.filePath)
			if (err != nil) != tc.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestWriteResume(t *testing.T) {
	type testStruct struct {
		Test string
	}

	testCases := []struct {
		name  string
		datas testStruct
	}{
		{"success", testStruct{"test"}},
		{"fail", testStruct{"test"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := WriteResume(tc.datas)
			if err != nil {
				t.Errorf("WriteResume() error = %v", err)
			}
		})
	}
}

func TestReadResume(t *testing.T) {
	type testStruct struct {
		Test string
	}

	testCases := []struct {
		name  string
		datas *testStruct
	}{
		{"success", &testStruct{}},
		{"fail", &testStruct{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ReadResume(tc.datas)
			if err != nil {
				t.Errorf("ReadResume() error = %v", err)
			}
		})
	}
}

func TestWriteOpt(t *testing.T) {
	t.Run("Test with valid data", func(t *testing.T) {
		data := struct {
			Name string `json:"name,omitempty"`
		}{"hhhhh"}
		err := WriteOpt(data)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
}

func TestReadOpt(t *testing.T) {

	type TestOptions struct {
		Name string `json:"name,omitempty"`
	}

	t.Run("Test with valid data", func(t *testing.T) {
		data := TestOptions{"hhhhh"}
		err := WriteOpt(data)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		fromConfig := &TestOptions{}
		err = ReadOpt(fromConfig)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
}
