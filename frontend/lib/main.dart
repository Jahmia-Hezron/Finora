import 'package:finora/layout/app_layout.dart' show AppLayout;
import 'package:finora/screen/screen_export.dart';
import 'package:flutter/material.dart';
import 'dart:async';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  static const Color _seedColor = Colors.blue;
  static const bool _isDarkMode = false;

  static ColorScheme lightColorScheme = ColorScheme.fromSeed(
    seedColor: _seedColor,
    brightness: Brightness.light,
  );

  static ColorScheme darkColorScheme = ColorScheme.fromSeed(
    seedColor: _seedColor,
    brightness: Brightness.dark,
  );

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      routes: {
        // Authentication routes
        '/login': (context) => LoginScreen(),
        '/register': (context) => RegisterScreen(),

        // Budget routes
        '/budget': (context) => BudgetScreen(),
        '/budget/create': (context) => CreateBudgetScreen(),

        // Income routes
        '/income': (context) => IncomeScreen(),
        '/income/create': (context) => CreateIncomeScreen(),
        '/income/source': (context) => IncomeSourceScreen(),
        '/income/source/create': (context) => CreateIncomeSourceScreen(),

        // Expense routes
        '/expense': (context) => ExpenseScreen(),
        '/expense/create': (context) => CreateExpenseScreen(),
        '/expense/category': (context) => ExpenseCategoryScreen(),
        '/expense/category/create': (context) => CreateExpenseCategoryScreen(),

        // User profile routes
        '/user-profile': (context) => UserProfileScreen(),
        '/user-profile/update': (context) => UserProfileUpdateScreen(),

        // Dashboard
        '/dashboard': (context) => DashboardScreen(),

        // Main app screen
        '/app': (context) => AppLayout(),
      },
      theme: ThemeData(useMaterial3: true, colorScheme: MyApp.lightColorScheme),
      darkTheme: ThemeData(
        useMaterial3: true,
        colorScheme: MyApp.lightColorScheme,
      ),
      themeMode: MyApp._isDarkMode ? ThemeMode.light : ThemeMode.dark,
      home: const SplashScreen(),
    );
  }
}

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  Timer? _timer;

  @override
  void initState() {
    super.initState();

    _timer = Timer(const Duration(seconds: 4), () {
      Navigator.pushNamed(context, '/app');
    });
  }

  @override
  void dispose() {
    _timer?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final colorScheme = Theme.of(context).colorScheme;
    return Scaffold(
      appBar: AppBar(automaticallyImplyLeading: false),
      body: Center(
        child: Text(
          'Finora',
          style: TextStyle(
            fontSize: 44,
            fontWeight: FontWeight.bold,
            color: colorScheme.primary,
          ),
        ),
      ),
    );
  }
}
