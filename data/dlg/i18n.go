package dlg

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/thingscompton/gormt/data/config"
	"github.com/thingscompton/gormt/public/myi18n"
	"golang.org/x/text/language"
)

/*
 Internationalization 国际化
*/

func init() {
	addChinese()
	addEnglish()
	myi18n.SetLocalLG(getLG()) // default
}

func getLG() string {
	tag := config.GetLG()
	if tag == "English" {
		return "en"
	}

	return "zh"
}

// SLocalize 获取值
func SLocalize(ID string) string {
	return myi18n.Get(ID)
}

func addChinese() error {
	return myi18n.AddMessages(language.Chinese, &i18n.Message{
		ID:    "menu",
		Other: "菜单",
	}, &i18n.Message{
		ID:    "list",
		Other: "列表",
	}, &i18n.Message{
		ID:    "view",
		Other: "视图",
	}, &i18n.Message{
		ID:    "run",
		Other: "执 行 ⏯ ",
	}, &i18n.Message{
		ID:    "set",
		Other: "设 置 🛠 ",
	}, &i18n.Message{
		ID:    "clipboardBtn",
		Other: "复 制 到 剪 切 板 ",
	}, &i18n.Message{
		ID:    "out_dir",
		Other: " 输 出 目 录 :",
	}, &i18n.Message{
		ID:    "db_type",
		Other: " 数 据 库 类 型 :",
	}, &i18n.Message{
		ID:    "db_host",
		Other: " 数 据 库 地 址 :",
	}, &i18n.Message{
		ID:    "db_port",
		Other: " 数 据 库 端 口 :",
	}, &i18n.Message{
		ID:    "db_usename",
		Other: " 数 据 库 用 户 名 :",
	}, &i18n.Message{
		ID:    "db_pwd",
		Other: " 数 据 库 密 码:",
	}, &i18n.Message{
		ID:    "db_name",
		Other: " 数 据 库 名 字 :",
	}, &i18n.Message{
		ID:    "is_dev",
		Other: " 开 发 模 式:",
	}, &i18n.Message{
		ID:    "is_simple",
		Other: " 简 单 输 出 :",
	}, &i18n.Message{
		ID:    "is_out_sql",
		Other: " 输 出 sql 原 :",
	}, &i18n.Message{
		ID:    "is_out_func",
		Other: " 输 出 快 捷 函 数 :",
	}, &i18n.Message{
		ID:    "is_foreign_key",
		Other: " 导 出 外 键 :",
	}, &i18n.Message{
		ID:    "is_gui",
		Other: " 界 面 模 式 :",
	}, &i18n.Message{
		ID:    "is_table_name",
		Other: " 生 成 表 名 :",
	}, &i18n.Message{
		ID:    "url_tag",
		Other: " web 标 签:",
	}, &i18n.Message{
		ID:    "db_tag",
		Other: " 数 据 库 标 签 :",
	}, &i18n.Message{
		ID:    "language",
		Other: " 语 言 :",
	}, &i18n.Message{
		ID:    "true",
		Other: " 是",
	}, &i18n.Message{
		ID:    "false",
		Other: " 否",
	}, &i18n.Message{
		ID:    "save",
		Other: " 保 存 ",
	}, &i18n.Message{
		ID:    "cancel",
		Other: " 取 消 ",
	}, &i18n.Message{
		ID:    "about",
		Other: " 关 于 作 者",
	}, &i18n.Message{
		ID:    "log_run",
		Other: " Enter : 执 行 \n ↑ ↓: 本 视 图 选 择 \n Tab : 多 视 图 切 换 \n Ctrl+C : 退 出 应 用 \n Ctrl+Q : 退 出 对 话 框 \n 支 持 鼠 标 操 作 方 式 \n \n \033[33;7m 输 入 Enter 直 接 执 行 \033[0m\n ",
	}, &i18n.Message{
		ID:    "log_set",
		Other: " Enter : 执 行 \n ↑ ↓: 本 视 图 选 择 \n Tab : 多 视 图 切 换\n Ctrl+C : 退 出 应 用 \n Ctrl+Q : 退 出 对 话 框 \n 支 持 鼠 标 操 作 方 式 \n \n \033[33;7m 输 入 Enter 打 开 设 置 窗 口 \033[0m\n ",
	})
}

func addEnglish() error {
	return myi18n.AddMessages(language.English, &i18n.Message{
		ID:    "menu",
		Other: "Menu",
	}, &i18n.Message{
		ID:    "list",
		Other: "List",
	}, &i18n.Message{
		ID:    "view",
		Other: "View",
	}, &i18n.Message{
		ID:    "run",
		Other: "Run ⏯ ",
	}, &i18n.Message{
		ID:    "set",
		Other: "Set 🛠 ",
	}, &i18n.Message{
		ID:    "clipboardBtn",
		Other: "Copy to clipboard",
	}, &i18n.Message{
		ID:    "out_dir",
		Other: "out dir:",
	}, &i18n.Message{
		ID:    "db_type",
		Other: " db type:",
	}, &i18n.Message{
		ID:    "db_host",
		Other: "db host:",
	}, &i18n.Message{
		ID:    "db_port",
		Other: "db port:",
	}, &i18n.Message{
		ID:    "db_usename",
		Other: "db username:",
	}, &i18n.Message{
		ID:    "db_pwd",
		Other: "db password:",
	}, &i18n.Message{
		ID:    "db_name",
		Other: "db name:",
	}, &i18n.Message{
		ID:    "is_dev",
		Other: "is dev:",
	}, &i18n.Message{
		ID:    "is_simple",
		Other: "is simple :",
	}, &i18n.Message{
		ID:    "is_out_sql",
		Other: "is out sql :",
	}, &i18n.Message{
		ID:    "is_out_func",
		Other: "is out func :",
	}, &i18n.Message{
		ID:    "is_foreign_key",
		Other: "is foreign key:",
	}, &i18n.Message{
		ID:    "is_gui",
		Other: "is show gui:",
	}, &i18n.Message{
		ID:    "is_table_name",
		Other: "is table name:",
	}, &i18n.Message{
		ID:    "url_tag",
		Other: "url tag:",
	}, &i18n.Message{
		ID:    "db_tag",
		Other: "db tag:",
	}, &i18n.Message{
		ID:    "language",
		Other: "Language:",
	}, &i18n.Message{
		ID:    "true",
		Other: "true",
	}, &i18n.Message{
		ID:    "false",
		Other: "false",
	}, &i18n.Message{
		ID:    "save",
		Other: "Save",
	}, &i18n.Message{
		ID:    "cancel",
		Other: "Cancel",
	}, &i18n.Message{
		ID:    "about",
		Other: "About",
	}, &i18n.Message{
		ID:    "log_run",
		Other: " Enter : run \n ↑ ↓: Selection of this view \n Tab : Multi view switching \n Ctrl+C : quit; \n Ctrl+Q : backup \n Mouse operation supported \n \n \033[33;7m Enter to execute \033[0m",
	}, &i18n.Message{
		ID:    "log_set",
		Other: " Enter : run \n ↑ ↓: Selection of this view \n Tab : Multi view switching \n Ctrl+C : quit \n Ctrl+Q : backup \n Mouse operation supported \n \n \033[33;7m Enter enter to open the settings window \033[0m",
	})
}
