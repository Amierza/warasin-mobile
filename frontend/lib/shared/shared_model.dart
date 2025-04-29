class RegisterFormModel {
    
  final String? name;
  final String? email;
  final String? password;

  RegisterFormModel({
    this.name,
    this.email,
    this.password,
  });

  Map<String, dynamic> toJson(){
    return {
      'name': name,
      'email': email,
      'password': password,
    };
  }
}