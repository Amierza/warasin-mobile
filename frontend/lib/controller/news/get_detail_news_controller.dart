import 'package:frontend/model/error.dart';
import 'package:frontend/model/news.dart';
import 'package:frontend/service/news_service.dart';
import 'package:get/get.dart';

class GetDetailNewsController extends GetxController {
  var isLoading = false.obs;
  var detailNews = Rxn<GetDetailNewsResponse>();

  Future<void> fetchDetailNews(String newsId) async{
    isLoading.value = true;

    try {
      final response = await NewsService.getDetailNews(newsId);

      if (response is GetDetailNewsResponse) {
        detailNews.value = response;
      }

      if (response is ErrorResponse) {
        Get.snackbar(response.error, response.message);
      }
    } catch(error) {
      Get.snackbar('Something went wrong', error.toString());
    } finally {
      isLoading.value = false;
    }
  }
}