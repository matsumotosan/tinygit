package init

import (
	"path"
	"gopkg.in/ini.v1"
)

type core struct {
	repositoryformatversion int
	filemode bool
	bare bool
	logallrefupdates bool
	ignorecase bool
	precomposeunicode bool
}

type config struct {
	core core
}

func newConfig() config {
	return config{
		core{
			repositoryformatversion: 0,
			filemode: true,
			bare: false,
			logallrefupdates: true,
			ignorecase: true,
			precomposeunicode: true,
		},
	}
}

func writeDefaultConfig(dir string) {
	inidata := ini.Empty()
	sec, err := inidata.NewSection("core")
	if err != nil {
		panic(err)
	}

	_, err = sec.NewKey("repositoryformatversion", "0")
	if err != nil {
		panic(err)
	}

	_, err = sec.NewKey("filemode", "true")
	if err != nil {
		panic(err)
	}

	_, err = sec.NewKey("bare", "false")
	if err != nil {
		panic(err)
	}

	_, err = sec.NewKey("logallrefupdates", "true")
	if err != nil {
		panic(err)
	}

	_, err = sec.NewKey("ignorecase", "true")
	if err != nil {
		panic(err)
	}

	_, err = sec.NewKey("precomposeunicode", "true")
	if err != nil {
		panic(err)
	}

	err = inidata.SaveTo(path.Join(dir, "config"))
	if err != nil {
		panic(err)
	}
}
