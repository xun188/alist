package drivers

import (
	_ "github.com/alist-org/alist/v3/drivers/115"
	_ "github.com/alist-org/alist/v3/drivers/115_share"
	_ "github.com/alist-org/alist/v3/drivers/139"
	_ "github.com/alist-org/alist/v3/drivers/alist_v3"
	_ "github.com/alist-org/alist/v3/drivers/aliyundrive_open"
	_ "github.com/alist-org/alist/v3/drivers/local"
	_ "github.com/alist-org/alist/v3/drivers/pikpak"
	_ "github.com/alist-org/alist/v3/drivers/pikpak_share"
	_ "github.com/alist-org/alist/v3/drivers/quark_uc"
	_ "github.com/alist-org/alist/v3/drivers/quark_uc_tv"
	_ "github.com/alist-org/alist/v3/drivers/sftp"
	_ "github.com/alist-org/alist/v3/drivers/smb"
	_ "github.com/alist-org/alist/v3/drivers/webdav"
)

// All do nothing,just for import
// same as _ import
func All() {

}
