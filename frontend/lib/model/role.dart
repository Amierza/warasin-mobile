class Role {
  final String roleId;
  final String roleName;

  Role({
    required this.roleId,
    required this.roleName,
  });

  factory Role.fromJson(Map<String, dynamic> json) {
    return Role(
      roleId: json['role_id'],
      roleName: json['role_name'],
    );
  }
}
