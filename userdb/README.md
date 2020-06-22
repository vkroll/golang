# userdb
## Purpose
userdb is a user administration for web projects. It got on my nerves to build such small parts again and again, so I outsourced it now at the beginning of my "golang career" to use it in the future.
## API
package userdb // import "github.com/vkroll/golang/userdb"

func CreateUser(u string, p []byte) (bool, error)
func LoadUsers()
func SaveUsers()
func UserExists(u string) bool
func ValidateUser(u string, p []byte) bool
