package chromium

import (
	"github.com/alxjsn/kooky"
	"github.com/alxjsn/kooky/internal/chrome"
	"github.com/alxjsn/kooky/internal/chrome/find"
	"github.com/alxjsn/kooky/internal/cookies"
)

type chromiumFinder struct{}

var _ kooky.CookieStoreFinder = (*chromiumFinder)(nil)

func init() {
	kooky.RegisterFinder(`chromium`, &chromiumFinder{})
}

func (f *chromiumFinder) FindCookieStores() ([]kooky.CookieStore, error) {
	files, err := find.FindChromiumCookieStoreFiles()
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
