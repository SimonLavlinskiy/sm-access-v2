package config

import "github.com/joho/godotenv"

func Init() {
	// Если отсутствует .env, берем системные переменные окружения
	_ = godotenv.Load()
	initDatabase()
	initMigrations()
	loggerInit()
}
