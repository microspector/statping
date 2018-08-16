package utils

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestCreateLog(t *testing.T) {
	err := createLog(Directory)
	assert.Nil(t, err)
}

func TestInitLogs(t *testing.T) {
	assert.Nil(t, InitLogs())
	assert.FileExists(t, "../logs/statup.log")
}

func TestDir(t *testing.T) {
	assert.Contains(t, Directory, "github.com/hunterlong/statup")
}

func TestLog(t *testing.T) {
	assert.Nil(t, Log(0, errors.New("this is a 0 level error")))
	assert.Nil(t, Log(1, errors.New("this is a 1 level error")))
	assert.Nil(t, Log(2, errors.New("this is a 2 level error")))
	assert.Nil(t, Log(3, errors.New("this is a 3 level error")))
	assert.Nil(t, Log(4, errors.New("this is a 4 level error")))
	assert.Nil(t, Log(5, errors.New("this is a 5 level error")))
}

func TestDeleteFile(t *testing.T) {
	assert.Nil(t, DeleteFile(Directory+"/logs/statup.log"))
}

func TestFailedDeleteFile(t *testing.T) {
	assert.Error(t, DeleteFile(Directory+"/missingfilehere.txt"))
}

func TestLogHTTP(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, Http(req))
}

func TestIntString(t *testing.T) {
	assert.Equal(t, "1", IntString(1))
}

func TestStringInt(t *testing.T) {
	assert.Equal(t, int64(1), StringInt("1"))
}

func TestTimestamp_Ago(t *testing.T) {
	now := Timestamp(time.Now())
	assert.Equal(t, "Just now", now.Ago())
}

func TestUnderScoreString(t *testing.T) {
	assert.Equal(t, "this_is_a_test", UnderScoreString("this is a test"))
}

func TestHashPassword(t *testing.T) {
	assert.Equal(t, 60, len(HashPassword("password123")))
}

func TestNewSHA1Hash(t *testing.T) {
	assert.NotEmpty(t, NewSHA1Hash(5))
}

func TestRandomString(t *testing.T) {
	assert.NotEmpty(t, RandomString(5))
}

func TestSha256(t *testing.T) {
	assert.Equal(t, "dc724af18fbdd4e59189f5fe768a5f8311527050", Sha256([]byte("testing")))
}

func TestDeleteDirectory(t *testing.T) {
	assert.Nil(t, DeleteDirectory(Directory+"/logs"))
}
