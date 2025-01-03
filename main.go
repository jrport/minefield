package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"golang.org/x/crypto/argon2"
)

type Server struct {
	name string
	muxer *http.ServeMux
}

type User struct {
	Name string
	Id string
}

func (u User) String() string{
	return fmt.Sprintf("%v:%v", u.Name, []byte(u.Id))
}

type UserQueue struct {
	Users []*User
	LastUser **User
	MaxCapacity int
	CurrentCount int
	Salt string
}

func NewUserQueue(salt string) *UserQueue{
	return &UserQueue{
		Salt: salt,
		Users: make([]*User, 0),
		CurrentCount: 0,
		MaxCapacity: 0,
	}
}

func (u UserQueue) Enqueue(user *User) error {
	if u.CurrentCount + 1 >= u.MaxCapacity {
		return fmt.Errorf("Max user count")
	}
	u.Users = append(u.Users, user)
	u.CurrentCount += 1
	u.LastUser = &u.Users[u.CurrentCount - 1]
	fmt.Printf("New user in queue! Current total: %v", user.String())
	return nil
}

func (u UserQueue) Dequeue(user *User) (*User, error) {
	if u.CurrentCount == 0 {
		return nil, fmt.Errorf("No users left")
	}
	playerToPop := u.Users[u.CurrentCount - 1]
	u.Users[u.CurrentCount - 1] = nil
}


func (u UserQueue) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip, _, ok := strings.Cut(r.RemoteAddr, ":")
	if !ok {
		http.Error(w, "internal error", 500)
		fmt.Println("Cant read this guy's ip wtf!")
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "missing user name", 422)
		fmt.Println("Bad form from this guy!")
		return
	}

	var user_name string
	if user_name = strings.TrimSpace(r.PostForm.Get("name")); user_name == "" {
		http.Error(w, "bad name", 422)
		fmt.Println("Empty name")
		return
	}

	anom_sess_id := argon2.IDKey(
		[]byte(strings.ReplaceAll(ip, ".", "")),
		[]byte(u.Salt), 1, 64*1024, 4, 32,
		)
	
	user := &User{Name: user_name, Id: string(anom_sess_id)}
	// u.AddUser(user)
	fmt.Fprint(w, "You're in queue! Wait a bit.\n")
}

var user_salt string

func init() {
	var ok bool
	user_salt, ok = os.LookupEnv("UID_SALT")
	if !ok {
		fmt.Fprint(os.Stderr, "Please set the UID_SALT environment variable.")
		os.Exit(1)
	}
}

func main() {
	m := http.NewServeMux()
	u := NewUserQueue(user_salt)
	
	m.Handle("POST /anom_sess", u)
	fmt.Println("Running server on 8080...")
	if err := http.ListenAndServe(":8080", m); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}
