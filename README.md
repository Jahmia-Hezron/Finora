
# Finora

**Finora** is a cross-platform personal finance manager built with **Flutter**, **Golang**, and **Docker**. It empowers users to manage income, categorize expenses, and plan budgets (coming soon) from any device—web, mobile, or desktop.

## ✨ Features

- 💰 Track multiple income sources
- 📊 Log and categorize expenses
- 🔐 Secure user authentication
- 🗂️ Organized transaction timeline
- 📤 Export transaction data (CSV, PDF in progress)
- 🧮 Budgeting capabilities *(coming soon)*
- 🌍 Multi-platform support: Web, Android, iOS, Windows, macOS, Linux

## 🛠️ Tech Stack

| Layer         | Tech                        |
| ------------- | --------------------------- |
| Frontend      | Flutter                     |
| Backend       | Golang (REST API)           |
| DevOps        | Docker, Docker Compose      |
| Database      | PostgreSQL (or SQLite TBD)  |
| API Security  | JWT Authentication          |

## 🚀 Getting Started

### Prerequisites

- Flutter SDK
- Go (1.20+)
- Docker & Docker Compose
- PostgreSQL (if not using Docker)

### Clone the Repository

```bash
git clone https://github.com/Jahmia-Hezron/finora.git
cd finora
```

### Run Backend (Golang API)

```bash
cd backend
go run main.go
```

### Run Frontend (Flutter)

```bash
cd frontend
flutter pub get
flutter run -d chrome   # For web
flutter run -d android  # For mobile
flutter run -d windows  # For desktop
```

### Docker Setup (Full Stack)

```bash
docker-compose up --build
```

## 📁 Project Structure

```
finora/
├── backend/           # Golang API
├── frontend/          # Flutter frontend (web, mobile, desktop)
├── docker/            # Docker and container setup
│   ├── Dockerfile.frontend
│   ├── Dockerfile.backend
│   └── docker-compose.yml
├── README.md          # General project README
```

## 🧭 Roadmap

- [x] Income & Expense Tracking
- [x] JWT-based Authentication
- [ ] Budget Management
- [ ] Graphs and Reports
- [ ] Multi-Currency Support
- [ ] Notifications & Preferences

## 🤝 Contributing

Contributions are welcome!

1. Fork the repo
2. Create your branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -m 'Add feature'`)
4. Push to the branch (`git push origin feature/my-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the [MIT License](LICENSE).

## 👤 Author

Developed with ❤️ by **Jahmia Hezron**  
Find me on [Website](https://jahmia-hezron.github.io) or [Email](mailto:hezron.p.jahmia@gmail.com)
