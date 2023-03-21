package domiannameindex

import (
	"bufio"
	"os"
	"testing"
)

func Test_node(t *testing.T) {

	t.Run("insert", func(t *testing.T) {
		t.Run("positive cases", func(t *testing.T) {
			type args struct {
				name       string
				searchName string
			}
			tests := []struct {
				name string
				args args
				want string
			}{
				{
					name: "insert one domain name",
					args: args{
						name:       "example.com",
						searchName: "example.com",
					},
					want: "example.com.",
				},
				{
					name: "insert wildcard domain name",
					args: args{
						name:       "*.example.com",
						searchName: "*.example.com",
					},
					want: "*.example.com.",
				},
				{
					name: "check search name with wildcard",
					args: args{
						name:       "*.example.com",
						searchName: "name.example.com",
					},
					want: "*.example.com.",
				},
				{
					name: "check with wildcard search for one level lower domain name",
					args: args{
						name:       "*.example.com",
						searchName: "lowerlevel.name.example.com",
					},
					want: "*.example.com.",
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					n := New()
					n.Insert(tt.args.name)
					got, fullPath := n.Find(tt.args.searchName)
					if !got || fullPath != tt.want {
						t.Errorf("node.Insert() = %v, got %v, want %v", got, fullPath, tt.want)
					}
				})
			}
		})
	})

	t.Run("remove", func(t *testing.T) {
		t.Run("positive cases", func(t *testing.T) {
			type args struct {
				name       string
				searchName string
			}
			tests := []struct {
				name string
				args args
				want string
			}{
				{
					name: "remove one domain name",
					args: args{
						name:       "example.com",
						searchName: "example.com",
					},
					want: "",
				},
				{
					name: "remove wildcard domain name",
					args: args{
						name:       "*.example.com",
						searchName: "*.example.com",
					},
					want: "",
				},
				{
					name: "check search name with wildcard",
					args: args{
						name:       "*.example.com",
						searchName: "name.example.com",
					},
					want: "",
				},
				{
					name: "check with wildcard search for one level lower domain name",
					args: args{
						name:       "*.example.com",
						searchName: "lowerlevel.name.example.com",
					},
					want: "",
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					n := New()
					n.Insert(tt.args.name)
					n.Print("")
					n.Remove(tt.args.name)
					got, fullPath := n.Find(tt.args.searchName)
					if got || fullPath != tt.want {
						t.Errorf("node.Insert() = %v, got %v, want %v", got, fullPath, tt.want)
					}
				})
			}
		})
	})
}

func TestPrint(t *testing.T) {
	n := New()
	n.Insert("example.com")
	n.Insert("*.example.com")
	n.Insert("*.name.example.com")
	n.Remove("example.com")
	n.Insert("google.com")
	n.Print("")
}

func TestFind(t *testing.T) {
	idxTree := seedIndexFromFile(t)
	type args struct {
		searchName string
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
		want   string
	}{
		{
			name: "non existing domain name",
			args: args{
				searchName: "www.google.com",
			},
			wantOk: false,
			want:   "",
		},
		{
			name: "empty search name",
			args: args{
				searchName: "",
			},
			wantOk: true,
			want:   ".",
		},
		{
			name: "existing domain name",
			args: args{
				searchName: "google.com",
			},
			wantOk: true,
			want:   "google.com.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, fullPath := idxTree.Find(tt.args.searchName)
			if got != tt.wantOk || fullPath != tt.want {
				t.Errorf("node.Find() = %v, got %v, want %v", got, fullPath, tt.want)
			}
		})
	}
}

func TestForReadme(t *testing.T) {
	idxTree := New()
	idxTree.Insert("*.example.com")
	ok, fullPath := idxTree.Find("name.example.com")
	if ok {
		if fullPath != "*.example.com." {
			t.Error("expected *.example.com. but got ", fullPath)
		}
	} else {
		t.Error("expected to find *.example.com. but got nothing")
	}

	ok, fullPath = idxTree.Find("another.name.example.com")
	if ok {
		if fullPath != "*.example.com." {
			t.Error("expected *.example.com. but got ", fullPath)
		}
	} else {
		t.Error("expected to find *.example.com. but got nothing")
	}
}

func seedIndexFromFile(t *testing.T) Interface {
	file, err := os.Open("domains.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	domainIndex := New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domainIndex.Insert(scanner.Text())
	}
	return domainIndex
}
