import 'package:curved_navigation_bar/curved_navigation_bar.dart';
import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
// import 'package:get/get.dart';

class CustomNavigationBar extends StatefulWidget {
  final int currentIndex;

  const CustomNavigationBar({
    Key? key,
    required this.currentIndex,
  }) : super(key: key);

  @override
  _CustomNavigationBarState createState() => _CustomNavigationBarState();
}

class _CustomNavigationBarState extends State<CustomNavigationBar> {
  late int _selectedIndex;

  final List<String> _routes = [
    '/home',
    '/terapiin',
    '/concultation',
    '/news',
    '/profile',
  ];

  @override
  void initState() {
    super.initState();
    _selectedIndex = widget.currentIndex;
  }

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });

    Navigator.pushNamed(context, _routes[index]);
  }

  @override
  Widget build(BuildContext context) {
    return CurvedNavigationBar(
      height: 65,
      color: primaryColor,
      backgroundColor: backgroundColor,
      index: _selectedIndex,
      items: const [
        Icon(Icons.home, color: Colors.white),
        Icon(Icons.headset_outlined, color: Colors.white),
        Icon(Icons.person, color: Colors.white),
        Icon(Icons.newspaper, color: Colors.white),
        Icon(Icons.settings, color: Colors.white),
      ],
      onTap: _onItemTapped,
    );
  }
}
