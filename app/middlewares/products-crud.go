package middleware

import (
     "framework/internal/app/config"
     "framework/internal/app/db"
     "framework/internal/app/logger"
     "framework/internal/app/mail"
)

func ProductsCrudMiddleware(c config.Configer, db db.DBer, l logger.Logger, m mail.Mailer) {
}
