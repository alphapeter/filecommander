package cfg

type Settings struct {
	Roots []Root
}

type Root struct {
	Name string
	Path string
}

var settings *Settings = nil

func GetSettings() Settings {
	if settings == nil {
		settings = &Settings{
			Roots: []Root{
				Root{Name: "incoming", Path: "/mnt/nas/Incoming"},
				Root{Name: "movies", Path: "/mnt/nas/Movies"},
				Root{Name: "series", Path: "/mnt/nas/Series"},
				Root{Name: "music", Path: "/mnt/nas/Music"}},
		}
	}
	return *settings
}

func GetPathForRoot(name string) (string, error) {
	s := GetSettings()
	for _, r := range s.Roots {
		if r.Name == name {
			return r.Path, nil
		}
	}
	return "", RootNotFoundError{name}
}

type RootNotFoundError struct {
	rootName string
}

func (e RootNotFoundError) Error() string {
	return "Root " + e.rootName + " not configured!"
}
