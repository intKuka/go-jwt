package consts

import "time"

const AccessTokenTTL = time.Minute * 10
const RefreshTokenTTL = time.Hour * 24 * 30 // 30 days
