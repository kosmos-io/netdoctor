package utils

import (
	"bytes"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	tests := []struct {
		name        string
		strTmpl     string
		obj         interface{}
		expected    []byte
		expectedErr string
	}{
		{
			name:    "Simple template with string",
			strTmpl: "Hello, {{.Name}}!",
			obj: map[string]string{
				"Name": "World",
			},
			expected:    []byte("Hello, World!"),
			expectedErr: "",
		},
		{
			name:    "Template with integer",
			strTmpl: "The value is {{.Value}}.",
			obj: map[string]int{
				"Value": 42,
			},
			expected:    []byte("The value is 42."),
			expectedErr: "",
		},
		{
			name:    "Template with missing key",
			strTmpl: "Hello, {{.UnknownKey}}!",
			obj: map[string]string{
				"Name": "World",
			},
			expected:    []byte("Hello, <no value>!"),
			expectedErr: "",
		},
		{
			name:        "Invalid template syntax",
			strTmpl:     "Hello, {{ .Name",
			obj:         map[string]string{"Name": "World"},
			expected:    nil,
			expectedErr: "error when parsing template",
		},
		{
			name:        "Template execution failure",
			strTmpl:     "Hello, {{ . }}",
			obj:         make(chan int), // Unsupported type
			expected:    nil,
			expectedErr: "error when executing template",
		},
		{
			name:    "Template with multiple keys",
			strTmpl: "{{.Greeting}}, {{.Name}}!",
			obj: map[string]string{
				"Greeting": "Hello",
				"Name":     "Alice",
			},
			expected:    []byte("Hello, Alice!"),
			expectedErr: "",
		},
		{
			name:    "Template with newlines",
			strTmpl: "Hello,\n{{.Name}}!",
			obj: map[string]string{
				"Name": "Bob",
			},
			expected:    []byte("Hello,\nBob!"),
			expectedErr: "",
		},
		{
			name:    "Template with special characters",
			strTmpl: "Key: {{.Key}}, Value: {{.Value}}",
			obj: map[string]string{
				"Key":   "Special&*%$#@!",
				"Value": "Value with spaces",
			},
			expected:    []byte("Key: Special&*%$#@!, Value: Value with spaces"),
			expectedErr: "",
		},
		{
			name:    "Template with whitespace",
			strTmpl: " Hello  {{ .Name }} ",
			obj: map[string]string{
				"Name": "   World   ",
			},
			expected:    []byte(" Hello     World    "),
			expectedErr: "",
		},
		{
			name:    "Template with float",
			strTmpl: "Pi is approximately {{.Pi}}.",
			obj: map[string]float64{
				"Pi": 3.14,
			},
			expected:    []byte("Pi is approximately 3.14."),
			expectedErr: "",
		},
		{
			name:    "Template with complex number",
			strTmpl: "Complex number: {{.Complex}}.",
			obj: map[string]complex128{
				"Complex": complex(1, 2),
			},
			expected:    []byte("Complex number: (1+2i)."),
			expectedErr: "",
		},
		{
			name:    "Template with boolean",
			strTmpl: "Is it true? {{.IsTrue}}.",
			obj: map[string]bool{
				"IsTrue": true,
			},
			expected:    []byte("Is it true? true."),
			expectedErr: "",
		},
		{
			name:    "Nested struct",
			strTmpl: "User: {{.User.Name}}, Age: {{.User.Age}}.",
			obj: struct {
				User struct {
					Name string
					Age  int
				}
			}{
				User: struct {
					Name string
					Age  int
				}{
					Name: "Charlie",
					Age:  30,
				},
			},
			expected:    []byte("User: Charlie, Age: 30."),
			expectedErr: "",
		},
		{
			name: "Slice in template",
			strTmpl: `Items:
{{range .Items}} - {{.}}
{{end}}`,
			obj: struct {
				Items []string
			}{
				Items: []string{"Item1", "Item2", "Item3"},
			},
			expected:    []byte("Items:\n - Item1\n - Item2\n - Item3\n"),
			expectedErr: "",
		},
		{
			name: "Map in template",
			strTmpl: `Properties:
{{range $key, $value := .Properties}} - {{$key}}: {{$value}}
{{end}}`,
			obj: struct {
				Properties map[string]string
			}{
				Properties: map[string]string{
					"Key1": "Value1",
					"Key2": "Value2",
				},
			},
			expected:    []byte("Properties:\n - Key1: Value1\n - Key2: Value2\n"),
			expectedErr: "",
		},
		{
			name:    "Template with function call",
			strTmpl: "Current year: {{.CurrentYear}}.",
			obj: struct {
				CurrentYear int
			}{
				CurrentYear: 2024,
			},
			expected:    []byte("Current year: 2024."),
			expectedErr: "",
		},
		{
			name:    "Template with date",
			strTmpl: "Today's date: {{.Date}}.",
			obj: struct {
				Date string
			}{
				Date: "2024-10-21",
			},
			expected:    []byte("Today's date: 2024-10-21."),
			expectedErr: "",
		},
		{
			name:    "Nested struct with invalid key",
			strTmpl: "User: {{.User.InvalidKey}}",
			obj: struct {
				User struct {
					Name string
				}
			}{
				User: struct {
					Name string
				}{
					Name: "Alice",
				},
			},
			expected:    nil,
			expectedErr: "error when executing template",
		},
		{
			name:    "Template with unsupported type",
			strTmpl: "Unsupported type: {{.Channel}}",
			obj: struct {
				Channel chan int
			}{
				Channel: make(chan int),
			},
			expected:    nil,
			expectedErr: "error when executing template",
		},
		{
			name:    "Multiple lines with variables",
			strTmpl: "First Name: {{.FirstName}}\nLast Name: {{.LastName}}",
			obj: map[string]string{
				"FirstName": "John",
				"LastName":  "Doe",
			},
			expected:    []byte("First Name: John\nLast Name: Doe"),
			expectedErr: "",
		},
		{
			name:    "Chained templates",
			strTmpl: "{{.Greeting}}, {{.Name}}! Today is {{.Day}}.",
			obj: map[string]interface{}{
				"Greeting": "Hello",
				"Name":     "Jane",
				"Day":      "Monday",
			},
			expected:    []byte("Hello, Jane! Today is Monday."),
			expectedErr: "",
		},
		{
			name: "Complex nested structure",
			strTmpl: `Person:
  Name: {{.Person.Name}}
  Address: {{.Person.Address.Street}}, {{.Person.Address.City}}`,
			obj: struct {
				Person struct {
					Name    string
					Address struct {
						Street string
						City   string
					}
				}
			}{
				Person: struct {
					Name    string
					Address struct {
						Street string
						City   string
					}
				}{
					Name: "Emily",
					Address: struct {
						Street string
						City   string
					}{
						Street: "123 Elm St",
						City:   "Gotham",
					},
				},
			},
			expected:    []byte("Person:\n  Name: Emily\n  Address: 123 Elm St, Gotham"),
			expectedErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseTemplate(tt.strTmpl, tt.obj)
			if tt.expectedErr != "" {
				if err == nil || !bytes.Contains([]byte(err.Error()), []byte(tt.expectedErr)) {
					t.Errorf("name %s, expected error containing %q, got %v", tt.name, tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("name %s, unexpected error: %v", tt.name, err)
				}
				if !bytes.Equal(result, tt.expected) {
					t.Errorf("name %s, expected %q, got %q", tt.name, tt.expected, result)
				}
			}
		})
	}
}
