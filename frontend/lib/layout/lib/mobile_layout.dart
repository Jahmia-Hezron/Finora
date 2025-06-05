import 'package:flutter/material.dart';

class MobileLayout extends StatelessWidget {
  final int selectedIndex;
  final void Function(int) onNavTap;
  final Widget screen;

  const MobileLayout({
    super.key,
    required this.selectedIndex,
    required this.onNavTap,
    required this.screen,
  });

  @override
  Widget build(BuildContext context) {
    final colorScheme = Theme.of(context).colorScheme;
    return Scaffold(
      body: screen,
      bottomNavigationBar: BottomNavigationBar(
        selectedItemColor: colorScheme.primary,
        unselectedItemColor: colorScheme.primary,
        currentIndex: selectedIndex,
        onTap: onNavTap,
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.home_outlined),
            activeIcon: Icon(Icons.home),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.pie_chart_outline_rounded),
            activeIcon: Icon(Icons.pie_chart_rounded),
            label: 'Budget',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.show_chart_rounded),
            activeIcon: Icon(Icons.show_chart_rounded),
            label: 'Income',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.insert_chart_outlined_rounded),
            activeIcon: Icon(Icons.insert_chart_rounded),
            label: 'Expense',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.person_outline_rounded),
            activeIcon: Icon(Icons.person),
            label: 'Profile',
          ),
        ],
      ),
    );
  }
}
