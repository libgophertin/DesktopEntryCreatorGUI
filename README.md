# Desktop Entry Creator GUI

[![Go Report Card](https://goreportcard.com/badge/github.com/libgophertin/DesktopEntryCreatorGUI)](https://goreportcard.com/report/github.com/libgophertin/DesktopEntryCreatorGUI)

**[English](#english) | [Русский](#russian)**

---

<a name="english"></a>
## 🇬🇧 English

### Description
**Desktop Entry Creator GUI** is a simple and convenient tool designed to help Linux users easily generate `.desktop` files (application shortcuts). Built with **Go** and the **Fyne** toolkit, it provides a user-friendly form to input application details without manually editing text files.

### 🚧 Project Status: Early Development
**Please Note:** This application is currently in a **"Raw" / Alpha state**.
*   Functionality is basic.
*   The UI and features are subject to change.
*   Optimizations and error handling are being improved.

Future updates will include more customization options and better validation.

### Features
*   Simple GUI to enter Application Name, Executable path, and Icon path.
*   File pickers for easy selection of executables and images.
*   Automatically generates and saves the `.desktop` file to `~/.local/share/applications/`.
*   Immediate availability of the shortcut in your system's application menu.

### Requirements
*   **Go** (Golang) 1.20 or higher.
*   **C Compiler** (gcc) - required for Fyne.
*   **Fyne Dependencies** (Linux):
    ```bash
    sudo apt-get install gcc libgl1-mesa-dev xorg-dev
    ```
    *(For Fedora/Arch/others, check the [Fyne "Getting Started" docs](https://developer.fyne.io/started/))*

### Installation & Usage

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/libgophertin/DesktopEntryCreatorGUI.git
    cd DesktopEntryCreatorGUI
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Run the application:**
    ```bash
    go run .
    ```

4.  **Build a binary:**
    ```bash
    go build -o DesktopEntryCreator .
    ```

---

<a name="russian"></a>
## 🇷🇺 Русский

### Описание
**Desktop Entry Creator GUI** — это простой и удобный инструмент, разработанный для пользователей Linux, который помогает легко создавать `.desktop` файлы (ярлыки приложений). Написанное на **Go** с использованием **Fyne**, приложение предоставляет удобную форму для ввода данных, избавляя от необходимости вручную редактировать текстовые файлы конфигурации.

### 🚧 Статус проекта: В разработке
**Обратите внимание:** Приложение находится в стадии **"Сырой" / Альфа версии**.
*   Реализован только базовый функционал.
*   Интерфейс и возможности могут меняться.
*   Ведется работа над улучшением кода и исправлением ошибок.

В будущих обновлениях планируется расширить настройки и улучшить валидацию данных.

### Возможности
*   Простой графический интерфейс для ввода имени приложения, пути к исполняемому файлу и иконке.
*   Файловый менеджер для удобного выбора файлов.
*   Автоматическая генерация и сохранение `.desktop` файла в директорию `~/.local/share/applications/`.
*   Моментальное появление ярлыка в меню приложений вашей системы.

### Требования
*   **Go** (Golang) версии 1.20 или выше.
*   **Компилятор C** (gcc) — необходим для работы Fyne.
*   **Зависимости Fyne** (для Ubuntu/Debian):
    ```bash
    sudo apt-get install gcc libgl1-mesa-dev xorg-dev
    ```
    *(Для других дистрибутивов см. [документацию Fyne](https://developer.fyne.io/started/))*

### Установка и запуск

1.  **Клонируйте репозиторий:**
    ```bash
    git clone https://github.com/libgophertin/DesktopEntryCreatorGUI.git
    cd DesktopEntryCreatorGUI
    ```

2.  **Загрузите зависимости:**
    ```bash
    go mod tidy
    ```

3.  **Запустите приложение:**
    ```bash
    go run .
    ```

4.  **Скомпилируйте бинарный файл:**
    ```bash
    go build -o DesktopEntryCreator .
    ```
