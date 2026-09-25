# ⚡ Go TCP Port Scanner

![Go Version](https://img.shields.io/badge/Go-1.18%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Cross--Platform-blue?style=for-the-badge)

Минималистичный, быстрый и надёжный сканер TCP-портов, написанный на **Go**. Программа позволяет определять точное состояние портов с помощью классификации системных ошибок (Syscalls).

---

## ✨ Особенности

- 🎯 **Умная классификация статусов**: умеет отличать `REFUSED` от `TIMEOUT` и `UNREACHABLE` на уровне системных вызовов (`syscall`).
- 📌 **Поддержка диапазонов**: сканирование как одиночного порта (например, `80`), так и диапазона (например, `20-100`).
- 🚀 **Zero Dependencies**: использует исключительно стандартную библиотеку Go (`std`).
- ⚡ **Быстрый таймаут**: встроенное ограничение ожидания соединения (1 секунда на порт).

---

## 📊 Поддерживаемые статусы портов

| Статус | Описание |
| :--- | :--- |
| `OPEN` | Порт открыт и принимает соединения |
| `REFUSED` | Хост активен, но соединение отклонено (порт закрыт) |
| `TIMEOUT` | Превышено время ожидания ответа (пакет сброшен или заблокирован) |
| `UNREACHABLE` | Хост недоступен в сети |
| `UNKNOWN` | Неопределённая ошибка сети |

---

## 🚀 Быстрый старт

### Требования
Установленный [Go compiler](https://go.dev/doc/install) версии `1.18` или выше.

### Сборка и запуск

1. **Клонируйте репозиторий:**
   ```bash
   git clone [https://github.com/your-username/port-scanner-go.git](https://github.com/your-username/port-scanner-go.git)
   cd port-scanner-go
