import 'package:finora/screen/screen_export.dart';
import 'package:flutter/material.dart';

import '../constant/layout_constant.dart' show kDesktopBreakpoint;
import 'lib/desktop_layout.dart' show DesktopLayout;
import 'lib/mobile_layout.dart' show MobileLayout;

class AppLayout extends StatefulWidget {
  const AppLayout({super.key});

  @override
  State<AppLayout> createState() => _AppLayoutState();
}

class _AppLayoutState extends State<AppLayout> {
  final List<Widget> _screens = [
    DashboardScreen(),
    BudgetHomeScreen(),
    IncomeHomeScreen(),
    ExpenseHomeScreen(),
    UserProfileHomeScreen(),
  ];

  int _selectedIndex = 0;

  void _onBottomNavItemTapped(int index) async {
    setState(() {
      _selectedIndex = index;
    });
  }

  void _navigateTo(Widget screen) async {
    final index = _screens.indexWhere(
      (s) => s.runtimeType == screen.runtimeType,
    );
    if (index != -1) {
      setState(() {
        _selectedIndex = index;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (context, constraints) {
        final isMobileOrTablet = constraints.maxWidth < kDesktopBreakpoint;

        return isMobileOrTablet
            ? MobileLayout(
              selectedIndex: _selectedIndex,
              onNavTap: _onBottomNavItemTapped,
              screen: _screens[_selectedIndex],
            )
            : DesktopLayout(
              primaryScreen: _screens[_selectedIndex],
              secondaryScreen: _screens[4],
              onItemSelected: _navigateTo,
              colorScheme: Theme.of(context).colorScheme,
            );
      },
    );
  }
}
