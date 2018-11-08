package bot

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func warnUser(chatID int64, current int64, limit int64) tgbotapi.MessageConfig {
	text := fmt.Sprintf("⚠️ اخطار %d از %d ⚠️\nامکان افزدون ربات تنها برای ادمین گروه فعال می‌باشد.", current, limit)
	return tgbotapi.NewMessage(chatID, text)
}

func superGroupIntroduction(chatID int64) tgbotapi.MessageConfig {
	text := `سلام 👋
	من هودورم، وظیفه من محافظت 🛡 از گروه‌ها در برابر ربات‌های اسپمر هست.
	برای این که هودور آماده رزم ⚔️ بشه، نیاز به اجازه شما داره.
	هودور رو ادمین گروه کنین و دسترسی *Ban users* رو بهش بدین بعد با دستور /hudor صداش بزنین.
	راستی هودور فقط از سازنده گروه دستور میگیره`

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	return msg
}

func botAddedToWhitelist(chatID int64, messageID int, username string) tgbotapi.MessageConfig {
	text := fmt.Sprintf("🤖 بات @%s به لیست ربات‌های مجاز به فعالیت افزوده شد. ✅", username)
	msg := tgbotapi.NewMessage(chatID, text)
	msg.DisableNotification = true
	msg.ReplyToMessageID = messageID

	return msg
}

func botCannotOperateWithoutCreator(chatID int64) tgbotapi.MessageConfig {
	text := `⛔️ فعالیت در این گروه امکان پذیر نمی‌باشد. ⛔️
	دلیل: سازنده گروه باید در گروه حضور داشته باشد.`

	return tgbotapi.NewMessage(chatID, text)
}

func errorHappenedDuringProcess(chatID int64) tgbotapi.MessageConfig {
	text := `❌ اوه شت 😱😱 
	متاسفانه خطایی رخ داده و نتونستم درخواست رو پردازش کنم.`
	return tgbotapi.NewMessage(chatID, text)
}

func hudorCanOnlySendFromCreator(chatID int64) tgbotapi.MessageConfig {
	text := "🔐 هودور فقط از سازنده گروه حرف شنوی داره 🔐"
	return tgbotapi.NewMessage(chatID, text)
}

func errorPermissionRequired(chatID int64) tgbotapi.MessageConfig {
	text := "⛔️ هودور نیاز به اجازه شما داره، دسترسی *Ban users* رو به هودور بدین تا بتونه شروع کنه! ⛔️"
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"

	return msg
}

func errorBotIsNotAdmin(chatID int64) tgbotapi.MessageConfig {
	text := `⚠️ دست و پای هودور بسته هست ⛓، هودور رو ادمین گروه کنین و دسترسی *Ban users* رو بهش بدین تا از غل و زنجیر آزاد بشه ⚠️`
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"

	return msg
}

func hudorActivated(chatID int64) tgbotapi.MessageConfig {
	text := `❇️ هودور با موفقیت فعال شد ❇️
	💎 نکات 💎
	1️⃣ جهت نمایش تنظیمات گروه دستور /settings را ارسال نمایید
	2️⃣ سازنده گروه می‌تواند تنظیمات گروه را از طریق چت خصوصی تغییر دهد
	3️⃣ در صورتی که می‌خواهین علاوه بر حذف ربات‌های مزاحم پیام آن‌ها را نیز پاک کنم دسترسی به *Delete messages* را برام فراهم کنین
	
	از گروه بدون ربات‌های مزاحم لذت ببرین 😎`
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"

	return msg
}

func hodurAlreadyIsActive(chatID int64) tgbotapi.MessageConfig {
	text := "🛡⚔️ هودور فعال می‌باشد ⚔️🛡"
	return tgbotapi.NewMessage(chatID, text)
}

func hodurOnlyActiveInSuperGroups(chatID int64) tgbotapi.MessageConfig {
	text := `من فقط می‌تونم توی سوپرگروه ها فعالیت کنم ☹️😞
	اگه می‌خوای بیشتر راجبم بدونی دستور /help رو بزن تا برات بگم`
	return tgbotapi.NewMessage(chatID, text)
}

func groupInformations(chatID int64, group *groupSettings) tgbotapi.MessageConfig {
	var text string
	if group == nil {
		text = "⚠️ در حال حاضر اطلاعاتی از این گروه در دست نیست ⚠️"
	} else {
		var activeStatus string
		var warnStatus string

		if group.IsActive {
			activeStatus = "❇️ فعال ❇️"
		} else {
			activeStatus = "🚫 غیر فعال 🚫"
		}

		if group.ShowWarn {
			warnStatus = "❇️ فعال ❇️"
		} else {
			warnStatus = "🚫 غیر فعال 🚫"
		}

		text = fmt.Sprintf(`گروه: %s
		وضعیت فعالیت: %s
		نمایش اخطار: %s
		تعداد اخطار قبل از حذف کاربر: %d بار`, group.Title, activeStatus, warnStatus, group.Limit)
	}

	return tgbotapi.NewMessage(chatID, text)
}
