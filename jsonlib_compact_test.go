package modern_go

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const SRC_DIR_REFLECT2 = "reflect2"
const DST_DIR_REFLECT2 = "../jsonlib/internal/reflect2"

func TestCompactReflect2(t *testing.T) {
	err := filepath.Walk(SRC_DIR_REFLECT2, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		switch filepath.Ext(info.Name()) {
		case ".go", ".s":
			err = copyFile(path, filepath.Join(DST_DIR_REFLECT2, info.Name()), info.Mode())
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

const SRC_DIR_JSONITER = "json-iterator"
const DST_DIR_JSONTIER = "../jsonlib/internal/jsoniter"

func TestCompactJsoniter(t *testing.T) {
	err := filepath.Walk(SRC_DIR_JSONITER, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(SRC_DIR_JSONITER, path)
		if err != nil {
			return err
		}
		if strings.HasPrefix(rel, "extra") {
			return nil
		}
		if strings.HasSuffix(rel, ".go") && !strings.HasSuffix(rel, "_test.go") {
			err = copyFile(path, filepath.Join(DST_DIR_JSONTIER, rel), info.Mode())
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func copyFile(srcPath string, dstPath string, mode os.FileMode) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dir := filepath.Dir(dstPath)
	if fi, _ := os.Stat(dir); fi == nil {
		os.MkdirAll(dir, os.ModePerm)
	}

	dst, err := os.OpenFile(dstPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, mode)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}
