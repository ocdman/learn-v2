package all

import (
	// The following are necessary as they register handlers in their init functions.

	// JSON config support. Choose only one from the two below.
	// The following line loads JSON from v2ctl
	// _ "v2ray.com/core/main/json"
	// The following line loads JSON internally
	_ "v2ray.com/core/main/jsonem"
)
