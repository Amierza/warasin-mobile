import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:frontend/controller/login_controller.dart';
import 'package:frontend/controller/register_controller.dart';
import 'package:frontend/middleware/middleware.dart';
import 'package:frontend/view/concultation_page.dart';
import 'package:frontend/view/data_completeness_feels_page.dart';
import 'package:frontend/view/data_completeness_greetings_page.dart';
import 'package:frontend/view/data_completeness_problems_page.dart';
import 'package:frontend/view/data_completeness_profile_page.dart';
import 'package:frontend/view/data_completeness_supports_page.dart';
import 'package:frontend/view/detail_concultation_page.dart';
import 'package:frontend/view/edit_consultation_page.dart';
import 'package:frontend/view/edit_profile_page.dart';
import 'package:frontend/view/forget_password_email_page.dart';
import 'package:frontend/view/forget_password_pasword_page.dart';
import 'package:frontend/view/history_consultation_page.dart';
import 'package:frontend/view/motivation_page.dart';
import 'package:frontend/view/news_history_page.dart';
import 'package:frontend/view/profile_page.dart';
import 'package:frontend/view/home_page.dart';
import 'package:frontend/view/login_page.dart';
import 'package:frontend/view/news_detail_page.dart';
import 'package:frontend/view/news_page.dart';
import 'package:frontend/view/register_page.dart';
import 'package:frontend/view/splash_page.dart';
import 'package:frontend/view/terapiin_page.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();

  try {
    await dotenv.load(fileName: ".env.local");
  } catch (e) {
    print("Gagal load .env: $e");
  }

  try {
    await GetStorage.init();
  } catch (e) {
    print("Gagal init GetStorage: $e");
  }

  runApp(await buildApp());
}

Future<Widget> buildApp() async {
  Get.lazyPut<RegisterController>(() => RegisterController(), fenix: false);
  Get.lazyPut<LoginController>(() => LoginController(), fenix: false);

  return const MyApp();
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: 'Warasin Mental Health Application',
      debugShowCheckedModeBanner: false,
      getPages: [
        GetPage(
          name: '/',
          page: () => const SplashPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/login',
          page: () => const LoginPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/register',
          page: () => const RegisterPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/forget_password',
          page: () => const ForgetPasswordEmailPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/data_completeness_greetings',
          page: () => const DataCompletenessGreetingsPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/data_completeness_feels',
          page: () => const DataCompletenessFeelsPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/data_completeness_problems',
          page: () => const DataCompletenessProblemsPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/data_completeness_supports',
          page: () => const DataCompletenessSupportsPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/data_completeness_profile',
          page: () => const DataCompletenessProfilePage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/home',
          page: () => const HomePage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/motivation',
          page: () => const MotivationPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/terapiin',
          page: () => const TerapiinPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/forget_password_password',
          page: () => const ForgetPasswordPaswordPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/concultation',
          page: () => const ConcultationPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/news',
          page: () => const NewsPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/news/:id',
          page: () => const NewsDetailPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/news_history',
          page: () => ReadNewsListPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/concultation_detail/:id',
          page: () => DetailConcultationPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/edit_concultation_detail/:id',
          page: () => UpdateConsultationPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/history_consultation',
          page: () => const HistoryConsultationPage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/profile',
          page: () => const ProfilePage(),
          middlewares: [AuthMiddleware()],
        ),
        GetPage(
          name: '/edit_profile',
          page: () => EditProfilePage(),
          middlewares: [AuthMiddleware()],
        ),
      ],
    );
  }
}
