package files

import (
	"testing"
)

func TestFile_Create(t *testing.T) {
	tests := []struct {
		name      string
		directory File
		generator Generator
	}{
		{
			name:      "create - directory with no children - dry",
			directory: NewDirectory("project-101", []File{}...),
			generator: NullFileGenerator{},
		},
		{
			name: "create - directory with files - dry",
			directory: NewDirectory(
				"project-101",
				[]File{
					{Name: "package.json"},
					{Name: "package.lock"},
				}...,
			),
			generator: NullFileGenerator{},
		},
		{
			name: "create - directory with directories - dry - 1",
			directory: NewDirectory(
				"project-101",
				[]File{
					{
						Name:        "src",
						IsDirectory: true,
						Files:       []File{{Name: "index.js"}}},
				}...,
			),
			generator: NullFileGenerator{},
		},
		{
			name: "create - directory with directories - dry - 2",
			directory: NewDirectory(
				"project-101",
				[]File{
					{
						Name:        "src",
						IsDirectory: true,
						Files: []File{
							{Name: "README.md"},
							{Name: "next.config.js"},
							{Name: "package-lock.json"},
							{Name: "package.json"},
							{
								Name:        "pages",
								IsDirectory: true,
								Files: []File{
									{Name: "_app.js"},
									{Name: "index.js"},
									{Name: "api", IsDirectory: true, Files: []File{{Name: "hello.js"}}},
								},
							},
							{
								Name:        "public",
								IsDirectory: true,
								Files:       []File{{Name: "favicon.ico"}, {Name: "vercel.svg"}},
							},
							{
								Name:        "styles",
								IsDirectory: true,
								Files:       []File{{Name: "Home.module.css"}, {Name: "globals.css"}},
							},
						},
					},
				}...,
			),
			generator: NullFileGenerator{},
		},
		{
			name: "create - directory with files",
			directory: NewDirectory(
				"test-project-101",
				[]File{
					{Name: "example.txt", Content: `abc-xyz`},
					{Name: "README.md", Content: `# EXAMPLE project`},
				}...,
			),
			generator: FileGenerator{},
		},
		{
			name:      "create - file",
			directory: File{Name: "test-project-101-file.txt"},
			generator: FileGenerator{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := tt.directory.Create(tt.generator); len(gotCount)-1 != tt.directory.Count() {
				t.Errorf(`File.Create() = %v want %v`, len(gotCount)-1, tt.directory.Count())
			}
			tt.directory.Remove(tt.generator)
		})
	}
}

func TestFile_Formatted(t *testing.T) {
	tests := []struct {
		name string
		file File
		want string
	}{
		{
			name: "file - formatted string",
			file: File{Name: "README.md"},
			want: "README.md",
		},
		{
			name: "directory - formatted string",
			file: File{Name: "project-101", IsDirectory: true},
			want: "project-101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.file.Formatted(); got != tt.want {
				t.Errorf("File.Formatted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Count(t *testing.T) {
	tests := []struct {
		name      string
		directory File
		wantCount int
	}{
		{
			name:      "count - directory with no children",
			directory: NewDirectory("project-101", []File{}...),
			wantCount: 0,
		},
		{
			name:      "count - directory with files",
			directory: NewDirectory("project-101", []File{{Name: "package.json"}, {Name: "package.lock"}}...),
			wantCount: 2,
		},
		{
			name:      "count - directory with directories",
			directory: NewDirectory("project-101", []File{{Name: "src", IsDirectory: true, Files: []File{{Name: "index.js"}}}}...),
			wantCount: 2,
		},
		{
			name: "create - directory with directories",
			directory: NewDirectory(
				"project-101",
				[]File{
					{
						Name:        "src",
						IsDirectory: true,
						Files: []File{
							{Name: "README.md"},
							{Name: "next.config.js"},
							{Name: "package-lock.json"},
							{Name: "package.json"},
						},
					},
					{
						Name:        "pages",
						IsDirectory: true,
						Files: []File{
							{Name: "_app.js"},
							{Name: "index.js"},
							{Name: "api", IsDirectory: true, Files: []File{{Name: "hello.js"}}},
						},
					},
					{
						Name:        "public",
						IsDirectory: true,
						Files:       []File{{Name: "favicon.ico"}, {Name: "vercel.svg"}},
					},
					{
						Name:        "styles",
						IsDirectory: true,
						Files:       []File{{Name: "Home.module.css"}, {Name: "globals.css"}},
					},
				}...,
			),
			wantCount: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := tt.directory.Count(); gotCount != tt.wantCount {
				t.Errorf("Directory.Count() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestFile_AddFiles(t *testing.T) {
	type args struct {
		files []File
	}

	tests := []struct {
		name      string
		directory File
		args      args
	}{
		{
			name:      "add 1 file",
			directory: File{Name: "project-101"},
			args: args{
				files: []File{{Name: "main.go"}},
			},
		},
		{
			name:      "add multiple files",
			directory: File{Name: "project-101"},
			args: args{
				files: []File{
					{Name: "go.mod"},
					{Name: "go.sum"},
					{Name: "main.go"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.directory.AddFiles(tt.args.files...)
			count := tt.directory.Count()

			if len(tt.args.files) != count {
				t.Errorf(`File.AddFiles() - expected %v files, got %v`, len(tt.args.files), count)
			}
		})
	}
}

func TestFile_Remove(t *testing.T) {
	type args struct {
		remover Remover
	}
	tests := []struct {
		name    string
		data    File
		args    args
		wantErr bool
	}{
		{
			name: "remove - empty directory",
			data: NewDirectory("project-101"),
			args: args{
				remover: NullFileGenerator{},
			},
			wantErr: false,
		},
		{
			name: "remove - file",
			data: File{Name: "index.js"},
			args: args{
				remover: NullFileGenerator{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.Remove(tt.args.remover); (err != nil) != tt.wantErr {
				t.Errorf("File.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_ReadFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name      string
		generator Generator
		args      args
		want      int
		wantErr   bool
	}{
		{
			name:      "readfile - dry - 1",
			generator: NullFileGenerator{},
			args: args{
				filepath: "../LICENSE",
			},
			want:    0,
			wantErr: false,
		},
		{
			name:      "readfile",
			generator: FileGenerator{},
			args: args{
				filepath: "../LICENSE",
			},
			want:    1516,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.generator.ReadFile(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("Generator.ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_ReadDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name      string
		generator Generator
		args      args
		want      int
		wantErr   bool
	}{
		{
			name:      "readdir - dry",
			generator: NullFileGenerator{},
			args: args{
				path: "/tmp",
			},
			want:    0,
			wantErr: false,
		},
		{
			name:      "readdir",
			generator: FileGenerator{},
			args: args{
				path: ".",
			},
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.generator.ReadDir(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.ReadDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("Generator.ReadDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestGenerator_ReadAll(t *testing.T) {
	type args struct {
		reader io.Reader
	}

	tests := []struct {
		name      string
		generator Generator
		args      args
		want      int
		wantErr   bool
	}{
		{
			name:      "read - dry - 1",
			generator: NullFileGenerator{},
			args:      args{
				//
			},
			want:    0,
			wantErr: false,
		},
		{
			name:      "read",
			generator: FileGenerator{},
			args:      args{
				//
			},
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.generator.ReadAll(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generator.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
