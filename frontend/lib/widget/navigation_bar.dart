import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';

class CustomNavigationBar extends StatelessWidget {
  const CustomNavigationBar({super.key});

  @override
  Widget build(BuildContext context) {
    return CurvedNavigationBar(
      color: primaryColor,
      backgroundColor: backgroundColor,
      items: [
        Icon(
          Icons.home,
          color: backgroundColor,
        ),
        Icon(
          Icons.headset_outlined,
          color: backgroundColor,
        ),
        Icon(
          Icons.person,
          color: backgroundColor,
        ),
        Icon(
          Icons.newspaper,
          color: backgroundColor,
        ),
        Icon(
          color: backgroundColor,
          Icons.person,
        )
      ]
    );
  }
}