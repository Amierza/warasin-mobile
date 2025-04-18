import 'package:flutter/material.dart';
import 'package:frontend/controller/login_controller.dart';
import 'package:frontend/view/concultation_page.dart';
import 'package:frontend/view/data_completeness_feels_page.dart';
import 'package:frontend/view/data_completeness_greetings_page.dart';
import 'package:frontend/view/data_completeness_problems_page.dart';
import 'package:frontend/view/data_completeness_profile_page.dart';
import 'package:frontend/view/data_completeness_supports_page.dart';
import 'package:frontend/view/detail_concultation_page.dart';
import 'package:frontend/view/forget_password_email_page.dart';
import 'package:frontend/view/forget_password_pasword_page.dart';
import 'package:frontend/view/history_consultation_page.dart';
import 'package:frontend/view/profile_page.dart';
import 'package:frontend/view/home_page.dart';
import 'package:frontend/view/login_page.dart';
import 'package:frontend/view/news_detail_page.dart';
import 'package:frontend/view/news_page.dart';
import 'package:frontend/view/register_page.dart';
import 'package:frontend/view/splash_page.dart';
import 'package:frontend/view/terapiin_page.dart';
import 'package:get/get.dart';

void main() {
  Get.lazyPut<LoginController>(() => LoginController(), fenix: false);
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: 'Warasin Mental Health Application',
      debugShowCheckedModeBanner: false,
      routes: {
        '/': (context) => const SplashPage(),
        '/login': (context) => const LoginPage(),
        '/register': (context) => const RegisterPage(),
        '/forget_password': (context) => const ForgetPasswordEmailPage(),
        '/data_completeness_greetings': (context) => const DataCompletenessGreetingsPage(),
        '/data_completeness_feels': (context) => const DataCompletenessFeelsPage(),
        '/data_completeness_problems': (context) => const DataCompletenessProblemsPage(),
        '/data_completeness_supports': (context) => const DataCompletenessSupportsPage(),
        '/data_completeness_profile': (context) => const DataCompletenessProfilePage(),
        '/home': (context) => const HomePage(),
        '/terapiin': (context) => const TerapiinPage(),
        '/forget_password_password': (context) => const ForgetPasswordPaswordPage(),
        '/concultation': (context) => const ConcultationPage(),
        '/news': (context) => const NewsPage(),
        '/news_detail': (context) => const NewsDetailPage(),
        '/concultation_detail': (context) => const DetailConcultationPage(),
        '/history_consultation': (context) => const HistoryConsultationPage(),
        '/profile': (context) => const ProfilePage(),
      },
    );
  }
}

