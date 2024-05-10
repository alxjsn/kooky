package chrome

import (
	"github.com/alxjsn/kooky"
	"github.com/alxjsn/kooky/internal/chrome"
	"github.com/alxjsn/kooky/internal/chrome/find"
	"github.com/alxjsn/kooky/internal/cookies"
)

type chromeFinder struct{}

var _ kooky.CookieStoreFinder = (*chromeFinder)(nil)

func init() {
	kooky.RegisterFinder(`chrome`, &chromeFinder{})
}

func (f *chromeFinder) FindCookieStores() ([]kooky.CookieStore, error) {
	files, err := find.FindChromeCookieStoreFiles()
	if err != nil {
		return nil, err
	}

	var ret []kooky.CookieStore
	for _, file := range files {
		ret = append(
			ret,
			&cookies.CookieJar{
				CookieStore: &chrome.CookieStore{
					DefaultCookieStore: cookies.DefaultCookieStore{
						BrowserStr:           file.Browser,
						ProfileStr:           file.Profile,
						OSStr:                file.OS,
						IsDefaultProfileBool: file.IsDefaultProfile,
						FileNameStr:          file.Path,
					},
				},
			},
		)
	}
	return ret, nil
}
