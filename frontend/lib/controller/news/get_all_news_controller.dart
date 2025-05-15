import 'package:frontend/model/error.dart';
import 'package:frontend/model/news.dart';
import 'package:frontend/service/news_service.dart';
import 'package:frontend/view/news_page.dart';
import 'package:get/get.dart';

class GetAllNewsController extends GetxController {
  var newsList = <News>[].obs;
  var isLoading = false.obs;
  var allNews = Rxn<GetAllNewsResponse>();

  @override
  void onInit() {
    fetchAllNews();
    super.onInit();
  }

  Future<void> fetchAllNews() async{
    isLoading.value = true;

    try {
      final response = await NewsService.getAllNews();

      if (response is GetAllNewsResponse) {
        allNews.value = response;
        newsList.value = response.data;
        print(newsList);
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