package media

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"api/app/ent"
	"api/app/ent/file"
	"api/lib"

	"github.com/segmentio/ksuid"
	"github.com/shenghui0779/yiigo"
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

func createMedia(ctx context.Context, mediaID, mediaPath, filename, fingerprint string, exif *lib.MediaEXIF) error {
	record, err := ent.DB.File.Create().
		SetFingerprint(fingerprint).
		SetSize(exif.Size).
		SetFormat(exif.Format).
		SetWidth(exif.Width).
		SetHeight(exif.Height).
		SetOrientation(exif.Orientation).
		SetDuration(exif.Duration).
		SetLongitude(exif.Longitude.String()).
		SetLongitude(exif.Latitude.String()).
		Save(ctx)
	if err != nil {
		// DB失败，删除文件
		if !yiigo.IsUniqueDuplicateError(err) {
			os.RemoveAll(mediaPath)
			return fmt.Errorf("ent.File.Create: %w", err)
		}
		// 唯一键存在，说明已存在
		record, err = ent.DB.File.Query().Unique(false).Select(file.FieldID).Where(file.Fingerprint(fingerprint)).First(ctx)
		if err != nil {
			return fmt.Errorf("ent.File.Query(fingerprint=%s): %w", fingerprint, err)
		}
	}

	// 创建Media
	_, err = ent.DB.Media.Create().SetMediaID(mediaID).SetFileName(filename).SetFileID(record.ID).Save(ctx)
	if err != nil {
		return fmt.Errorf("ent.Media.Create: %w", err)
	}
	return nil
}
