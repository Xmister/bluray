package bluray
import "os"
import "fmt"
import "strings"
import "path"
import "net/http"
import "github.com/Xmister/bluray/structure/file"

type BDMVDir interface {
	//Opens a File relative to BDMV Directory
	Open(path string) (http.File, error)
}

func warning(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr,err)
	}
}

func checkIsDir(dir http.File) (err error) {
	var s os.FileInfo
	if s, err = dir.Stat(); err == nil {
		if !s.IsDir() {
			err = fmt.Errorf("%s is not a folder", s.Name())
		}
	}
	return
}

type BluRay struct {
	BDMVDir
	*file.Index
	*file.MovieObject
	PlayLists map[string]*file.MPLS
	Clips map[string]*file.Clip
}

func (b *BluRay) readPlaylists() (err error) {
	var f http.File
	if f, err = b.Open("PLAYLIST"); err != nil {
		return
	}
	defer f.Close()
	if err = checkIsDir(f); err != nil {
		return
	}
	var files []os.FileInfo
	if files, err = f.Readdir(-1); err != nil {
		return
	}
	b.PlayLists=make(map[string]*file.MPLS)
	for _, child := range files {
		if f, err = b.Open(path.Join("PLAYLIST", child.Name())); err != nil {
			return
		}
		b.PlayLists[strings.Split(child.Name(),".")[0]] = file.NewMPLS(f)
		f.Close()
	}
	return
}

func (b *BluRay) readClips() (err error) {
	var f http.File
	if f, err = b.Open("CLIPINF"); err != nil {
		return
	}
	defer f.Close()
	if err = checkIsDir(f); err != nil {
		return
	}
	var files []os.FileInfo
	if files, err = f.Readdir(-1); err != nil {
		return
	}
	b.Clips=make(map[string]*file.Clip)
	for _, child := range files {
		if f, err = b.Open(path.Join("CLIPINF", child.Name())); err != nil {
			return
		}
		b.Clips[strings.Split(child.Name(),".")[0]] = file.NewClip(f)
		f.Close()
	}
	return
}

func OpenBDMV(dir BDMVDir) (res *BluRay, err error) {
	var f http.File
	res = &BluRay{
		BDMVDir: dir,
	}
	if f, err = res.Open("index.bdmv"); err != nil {
		return
	}
	defer f.Close()
	res.Index = file.NewIndex(f)
	if f, err = res.Open("MovieObject.bdmv"); err != nil {
		return
	}
	defer f.Close()
	res.MovieObject = file.NewMovieObject(f)
	warning(res.readPlaylists())
	warning(res.readClips())
	return
}