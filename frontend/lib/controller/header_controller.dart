import 'package:frontend/model/error.dart';
import 'package:frontend/model/user_detail_response.dart';
import 'package:frontend/service/api_service.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';

class HeaderController extends GetxController {
  final RxString name = ''.obs;
  final GetStorage storage = GetStorage();

  @override
  void onInit() {
    fetchUserData();
    super.onInit();
  }

  Future<void> fetchUserData() async {
    try {
      final token = storage.read('access_token');

      if (token == null) {
        return;
      }

      final response = await ApiService.getUserDetailService();

      if (response is UserDetailResponse) {
        name.value = response.data.userName;
      } else if (response is ErrorResponse) {
        if (response.message.toLowerCase().contains('unauthorized')) {
          storage.remove('access_token');
        }
      }
    } catch (err) {
      print("Failed to fetch user data: $err");
    }
  }
}
