type byName []Person

func (b byName) Len() int           { return len(b) }
func (b byName) Less(i, j int) bool { return b[i].Name < b[j].Name }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type byAge []Person

func (b byAge) Len() int           { return len(b) }
func (b byAge) Less(i, j int) bool { return b[i].AgeYears < b[j].AgeYears }
func (b byAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
