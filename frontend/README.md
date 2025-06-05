
# Finora Frontend

This is the **Flutter frontend** of the **Finora** personal finance manager. Built for **web**, **mobile**, and **desktop** platforms, it provides users with an intuitive interface to manage income, expenses, and (soon) budgets.

## 🖼️ Overview

The frontend is designed to be cross-platform using Flutter's adaptive layout system. It features responsive UI components, smooth navigation, and future-ready architecture.

## 📁 Project Structure

```
frontend/
├── lib/
│   ├── constant/               # Reusable constants (e.g., layout constants)
│   ├── layout/                 # Handles responsive layout switching
│   │   ├── lib/
│   │   ├── app_layout.dart
│   ├── screen/                 # Screen pages/views
│   │   ├── lib/
│   │   ├── screen_export.dart
│   ├── service/                # API and business logic services
│   │   ├── lib/
│   │   ├── base_service.dart
│   ├── widget/                 # Reusable UI components/widgets
│   │   ├── lib/
│   │   ├── widget_exporter.dart
│   ├── main.dart               # App entry point
```

## 🧱 Architecture Highlights

- 📐 **Responsive Design**: Uses `LayoutBuilder` to switch between `MobileLayout` and `DesktopLayout`.
- 🧩 **Modular Codebase**: Clear separation of widgets, services, constants, screens, and layouts.
- ♻️ **Reusable Components**: Common UI elements are abstracted for consistency.
- 🌐 **Backend Integration**: Connects with Finora's Golang REST API for authentication, transactions, and more.

## 🚀 Running the App

Ensure you have the Flutter SDK installed.

```bash
flutter pub get
flutter run -d chrome   # Run on web
flutter run -d android  # Run on Android
flutter run -d windows  # Run on Windows
```

## 🧑‍💻 Developer Notes

- Widget exports are managed in `widget_exporter.dart`
- Screens are modular and organized under the `screen` directory
- Layout responsiveness is powered by `app_layout.dart`, `mobile_layout.dart`, and `desktop_layout.dart`

## 📬 Contact

Developed by **Jahmia Hezron**  
Find me on [Website](https://jahmia-hezron.github.io) or [Email](mailto:hezron.p.jahmia@gmail.com)
