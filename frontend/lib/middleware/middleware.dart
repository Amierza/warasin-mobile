import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';

class AuthMiddleware extends GetMiddleware {
  final publicRoutes = ['/', '/login', '/register'];

  @override
  RouteSettings? redirect(String? route) {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null && !publicRoutes.contains(route)) {
      return const RouteSettings(name: '/login');
    }

    if (token != null && publicRoutes.contains(route)) {
      return const RouteSettings(name: '/home');
    }

    return null;
  }
}
