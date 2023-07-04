package users

import (
	"fmt"
	"time"

	"github.com/davidlozano107/go-desde-0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "David", time.Now(), true)

	fmt.Println(u)

}
