import 'package:frontend/model/error.dart';
import 'package:frontend/model/news_detail.dart';
import 'package:frontend/service/news_service.dart';
import 'package:get/get.dart';

class GetAllNewsDetailController extends GetxController {
  var newsList = <AllNewsDetail>[].obs;
  var isLoading = false.obs;
  var allNewsDetail = Rxn<GetAllNewsDetail>();

  @override
  void onInit() {
    fetchAllNewsDetail();
    super.onInit();
  }

  Future<void> fetchAllNewsDetail() async {
    isLoading.value = true;

    try {
      final response = await NewsService.getAllNewsDetailController();

      if (response is GetAllNewsDetail) {
        allNewsDetail.value = response;
        newsList.value = response.data;
      }

      if (response is ErrorResponse) {
        Get.snackbar(response.error, response.message);
      }
    } catch (error) {
      Get.snackbar('Something went wrong', error.toString());
    } finally {
      isLoading.value = false;
    }
  }
}
