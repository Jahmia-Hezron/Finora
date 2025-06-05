# Finora

**Finora** is a cross-platform personal finance manager built using **Flutter**, **Golang**, and **Docker**. It helps users seamlessly track income sources, categorize expenditures, and (soon) create and manage budgets—available on **web**, **mobile**, and **desktop** platforms.

## ✨ Features

- 💰 Track multiple income sources
- 📊 Log and categorize expenses
- 🔐 Secure authentication and user accounts
- 📅 Timeline of transactions
- 📁 Exportable transaction history (CSV, PDF planned)
- 🧮 Budgeting system *(coming soon)*
- 🖥️ Cross-platform: Web, Android, iOS, Windows, macOS, Linux

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
- PostgreSQL (if not using Docker for DB)

### Clone the Repository

```bash
git clone https://github.com/your-username/finora.git
cd finora
```

### Backend (Golang API)

```bash
cd backend
go run main.go
```

### Frontend (Flutter App)

```bash
cd frontend
flutter pub get
flutter run -d chrome   # for web
flutter run -d windows  # for desktop
flutter run -d android  # for mobile
```

### Docker Setup (Full Stack)

```bash
docker-compose up --build
```

## 🧱 Project Structure

```
finora/
├── backend/           # Golang API
│   ├── main.go
│   └── ...
├── frontend/          # Flutter frontend (web, mobile, desktop)
│   ├── lib/
│   └── ...
├── docker/
│   ├── Dockerfile.frontend
│   ├── Dockerfile.backend
│   └── docker-compose.yml
└── README.md
```

## 📌 Roadmap

- [x] Income & Expense Tracking
- [x] Basic Authentication
- [ ] Budgeting Module
- [ ] Data Visualization (Pie Charts, Line Graphs)
- [ ] Multi-Currency Support
- [ ] Notifications & Reminders
- [ ] Settings & Preferences

## 🤝 Contributing

Contributions, issues and feature requests are welcome!

1. Fork the project
2. Create your feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add your feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the [MIT License](LICENSE).

## 💬 Contact

Built with 💻 by **Jahmia Hezron**  
Reach out on [LinkedIn](https://linkedin.com/in/your-profile) or [Email](mailto:your-email@example.com)
