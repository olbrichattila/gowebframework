package appconfig

import (
	job "framework/app/jobs"
	"framework/internal/app/cron"
)

var Jobs = []cron.Job{
	{Seconds: 5, Fn: job.SendRegistrationEmail},
	{Seconds: 5, Fn: job.PasswordReminderJob},
	{Seconds: 30, Fn: job.ExpirePasswordReminderJob},
	{Seconds: 30, Fn: job.ExpireEmailConfJob},
}
