package licensing

import "strings"

var mapping = map[string]string{
	// GPL
	"GPL-1":                          GPL10,
	"GPL-1+":                         GPL10,
	"GPL2":                           GPL20,
	"GPL-2":                          GPL20,
	"GPL-2.0-ONLY":                   GPL20,
	"GPL2+":                          GPL20,
	"GPLV2+":                         GPL20,
	"GPL-2+":                         GPL20,
	"GPL-2.0+":                       GPL20,
	"GPL-2.0-OR-LATER":               GPL20,
	"GPL-2+ WITH AUTOCONF EXCEPTION": GPL20withautoconfexception,
	"GPL3":                           GPL30,
	"GPLV3+":                         GPL30,
	"GPL-3":                          GPL30,
	"GPL-3.0-ONLY":                   GPL30,
	"GPL3+":                          GPL30,
	"GPL-3+":                         GPL30,
	"GPL-3.0-OR-LATER":               GPL30,
	"GPL-3+-WITH-BISON-EXCEPTION":    GPL20withbisonexception,
	"GPL":                            GPL30, // 2? 3?

	// LGPL
	"LGPL2":     LGPL20,
	"LGPL-2":    LGPL20,
	"LGPL2+":    LGPL20,
	"LGPL-2+":   LGPL20,
	"LGPL-2.0+": LGPL20,
	"LGPL-2.1":  LGPL21,
	"LGPL-2.1+": LGPL21,
	"LGPLV2.1+": LGPL21,
	"LGPL-3":    LGPL30,
	"LGPL-3+":   LGPL30,
	"LGPL":      LGPL30, // 2? 3?

	// MPL
	"MPL1.0": MPL10,
	"MPL1":   MPL10,
	"MPL2.0": MPL20,
	"MPL2":   MPL20,

	// BSD
	"BSD":          BSD3Clause, // 2? 3?
	"BSD-2-CLAUSE": BSD2Clause,
	"BSD-3-CLAUSE": BSD3Clause,
	"BSD-4-CLAUSE": BSD4Clause,

	"APACHE": Apache20, // 1? 2?
	"ZLIB":   Zlib,
}

func Normalize(name string) string {
	if l, ok := mapping[strings.ToUpper(name)]; ok {
		return l
	}
	return name
}
