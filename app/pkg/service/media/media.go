package media

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"path"
	"path/filepath"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/spf13/viper"
)

func MediaURL(mediaID string) string {
	return fmt.Sprintf("%s/media/%s", viper.GetString("media.host"), mediaID)
}

func MediaChunk(uniqID string, index int) string {
	return path.Clean(fmt.Sprintf("%s/chuncks/%s/%d.tmp", viper.GetString("media.path"), uniqID, index))
}

func MediaFile(fingerprint, name string) string {
	return path.Clean(fmt.Sprintf("%s/%s", viper.GetString("media.path"), fingerprint+filepath.Ext(name)))
}

func MediaID(fingerprint string) string {
	nonce := make([]byte, 32)
	io.ReadFull(rand.Reader, nonce)

	h := md5.New()
	h.Write([]byte(fingerprint))
	h.Write(nonce)
	id, _ := ksuid.FromParts(time.Now(), h.Sum(nil)[:16])

	return id.String()
}
