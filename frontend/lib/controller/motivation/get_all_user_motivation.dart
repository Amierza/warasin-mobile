import 'package:frontend/model/error.dart';
import 'package:frontend/model/motivation.dart';
import 'package:frontend/service/motivation_service.dart';
import 'package:get/get.dart';

class GetAllUserMotivationController extends GetxController {
  var userMotivationList = <UserMotivation>[].obs;
  var isLoading = false.obs;
  var allUserMotivation = Rxn<AllUserMotivationResponse>();

  Future<void> fetchAllUserMotivation() async {
    isLoading.value = true;

    try {
      final response = await MotivationService.getAllUserMotivation();

      if (response is AllUserMotivationResponse) {
        allUserMotivation.value = response;
        userMotivationList.value = response.data;
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
