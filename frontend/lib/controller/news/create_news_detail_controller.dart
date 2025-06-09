import 'package:frontend/model/news_detail.dart';
import 'package:frontend/service/news_service.dart';
import 'package:get/get.dart';
import 'package:intl/intl.dart';

class CreateNewsDetailController extends GetxController {
  late String formattedDate;
  var isLoading = false.obs;

  @override
  void onInit() {
    super.onInit();
    final now = DateTime.now();
    formattedDate = DateFormat('yyyy-MM-dd').format(now);
  }

  Future<void> fetchCreateNewsDetail(String newsId) async {
    isLoading.value = true;

    try {
      final response = await NewsService.createNewsDetailService(
        formattedDate,
        newsId,
      );
      
      if (response is CreateNewsDetail) {
        Get.snackbar(
          "Berita selesai dibaca",
          "Silahkan menuju ke histori berita untuk melihat",
        );
      }
    } catch (e) {
      Get.snackbar(
        "Error",
        "Gagal menyelesaikan membaca berita: ${e.toString()}",
      );
    } finally {
      isLoading.value = false;
    }
  }
}
