import 'dart:async';

import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';

class SplashPage extends StatefulWidget {
  const SplashPage({Key? key}) : super(key: key);

  @override
  State<SplashPage> createState() => _SplashPageState();
}

class _SplashPageState extends State<SplashPage> {
  @override
  void initState() {
    super.initState();

    final token = GetStorage().read('access_token');
    Future.delayed(Duration(seconds: 2), () {
      if (token != null) {
        Get.offAllNamed('/home');
      } else {
        Get.offAllNamed('/login');
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: primaryColor,
      body: Center(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Container(
              width: 148,
              height: 144,
              decoration: const BoxDecoration(
                image: DecorationImage(
                  image: AssetImage('assets/logo-white.png'),
                ),
              ),
            ),
            const SizedBox(height: 16),
            Text(
              'Warasin',
              style: secondaryTextStyle.copyWith(
                fontSize: 30,
                fontWeight: extraBold,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
