import 'package:flutter/material.dart';

import '../../constant/layout_constant.dart' show kSectionSpacing;
import '../../screen/screen_export.dart';
import '../../widget/widget_exporter.dart';

class DesktopLayout extends StatelessWidget {
  final Widget primaryScreen;
  final Widget secondaryScreen;
  final Function(Widget) onItemSelected;
  final ColorScheme colorScheme;

  const DesktopLayout({
    super.key,
    required this.primaryScreen,
    required this.secondaryScreen,
    required this.onItemSelected,
    required this.colorScheme,
  });

  @override
  Widget build(BuildContext context) {
    final screenWidth = MediaQuery.of(context).size.width;
    final sideMenuWidth = screenWidth < 1200 ? 72.0 : 220.0;
    final sectionSpacing = kSectionSpacing;

    return Scaffold(
      body: Row(
        children: [
          SizedBox(
            width: sideMenuWidth,
            child: ReusableSideMenu(
              isCompact: sideMenuWidth < 100,
              onItemSelected: onItemSelected,
              colorScheme: colorScheme,
            ),
          ),
          SizedBox(width: sectionSpacing),
          Flexible(flex: 5, fit: FlexFit.loose, child: primaryScreen),
          SizedBox(width: sectionSpacing),
          Flexible(flex: 2, fit: FlexFit.loose, child: secondaryScreen),
        ],
      ),
    );
  }
}

class ReusableSideMenu extends StatelessWidget {
  final bool isCompact;
  final Function(Widget) onItemSelected;
  final ColorScheme colorScheme;

  const ReusableSideMenu({
    super.key,
    required this.onItemSelected,
    required this.colorScheme,
    required this.isCompact,
  });

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: isCompact ? 72 : 200,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          ReusableNavItemWidget(
            icon: Icons.home_rounded,
            title: 'Home',
            screen: DashboardScreen(),
            onItemSelected: onItemSelected,
            textColor: colorScheme.primary,
          ),
          ReusableNavItemWidget(
            icon: Icons.pie_chart_rounded,
            title: 'Budget',
            screen: BudgetHomeScreen(),
            onItemSelected: onItemSelected,
            textColor: colorScheme.primary,
          ),
          ReusableNavItemWidget(
            icon: Icons.show_chart_rounded,
            title: 'Income',
            screen: IncomeHomeScreen(),
            onItemSelected: onItemSelected,
            textColor: colorScheme.primary,
          ),
          ReusableNavItemWidget(
            icon: Icons.insert_chart_rounded,
            title: 'Expense',
            screen: ExpenseHomeScreen(),
            onItemSelected: onItemSelected,
            textColor: colorScheme.primary,
          ),
        ],
      ),
    );
  }
}
