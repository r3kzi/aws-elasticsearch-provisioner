package user

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var data = `
max:
  password: I07%0&8bv4$ie!92hr7q6fxs#7wUGO9%jJqV
  backend_roles:
  - backend-role-1
  - backend-role-2

tom:
  password: Z^q&Ve684viL$Y0YQ6fZfvnJ96R6TsZ!%X##

john:
  password: uL0kv&EQTej0lQddR0HL70T2Zv8dnAVFS^^A
  backend_roles:
  - backend-role-5
  - backend-role-6
  - backend-role-7
  - backend-role-8
`

func TestReadUser(t *testing.T) {
	filename := "test-user.yml"

	bytes := []byte(data)
	if err := ioutil.WriteFile(filename, bytes, 0644); err != nil {
		t.Errorf("Failed to write file, %v", err)
	}

	users, err := Read(filename)
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 3, len(users))
	assert.Equal(t, "Z^q&Ve684viL$Y0YQ6fZfvnJ96R6TsZ!%X##", users["tom"].Password)
	assert.Equal(t, 4, len(users["john"].BackendRoles))

	if err := os.Remove(filename); err != nil {
		t.Errorf("Failed to remove file, %v", err)
	}
}
