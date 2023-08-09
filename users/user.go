package users

import (
	"fmt"
	"github.com/PUPC-git/goBot/modelos"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
